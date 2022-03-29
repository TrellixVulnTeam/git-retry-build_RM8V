// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package gerrit contains functions for interacting with gerrit/gitiles.
package gerrit

import (
	"context"
	"fmt"
	"infra/cros/internal/shared"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.chromium.org/luci/common/api/gerrit"
	"go.chromium.org/luci/common/errors"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
)

// This files contains legacy functions that do not use a client.
// You should not add functions here.

// ChangeRevKey is the necessary set of data for looking up a single Gerrit revision.
type ChangeRevKey struct {
	Host      string
	ChangeNum int64
	Revision  int32
}

func (cik ChangeRevKey) String() string {
	return fmt.Sprintf("%s:%d:%d", cik.Host, cik.ChangeNum, cik.Revision)
}

// ChangeRev contains data about a Gerrit change,revision pair.
type ChangeRev struct {
	ChangeRevKey
	Project string
	// Ref that this change targets, e.g.: "refs/heads/main"
	Branch string
	// The Git reference for the patch set, e.g. "refs/changes/23/123/5"
	Ref   string
	Files []string
}

var (
	// Override this to use a mock GerritClient rather than the real one.
	mockGerrit gerritpb.GerritClient
)

// Regex matching the path of a CL URL.
//
// The path must be of the form:
// "/c/<project>/+/<change number>/<patchset>".
var clURLPathRegex = regexp.MustCompile(`^\/c\/[^\+]+\/\+\/(\d+)\/(\d+)$`)

// ParseCLURL parses a url for a CL into a ChangeRevKey.
//
// The url must be of the form:
// "https://<host>-review.googlesource.com/c/<project>/+/<change number>/<patchset>".
func ParseCLURL(rawurl string) (*ChangeRevKey, error) {
	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	if parsedURL.Host == "" {
		return nil, fmt.Errorf("could not parse host from URL (scheme must be specified): %s", rawurl)
	}

	if !strings.HasSuffix(parsedURL.Host, "-review.googlesource.com") {
		return nil, fmt.Errorf(`host must end with "-review.googlesource.com: %s`, rawurl)
	}

	pathMatches := clURLPathRegex.FindStringSubmatch(parsedURL.Path)
	if pathMatches == nil {
		return nil, fmt.Errorf(
			`could not parse CL information from URL path, expected path format "/c/<project>/+/<change number>/<patchset>": %s`,
			rawurl,
		)
	}

	changeNum, err := strconv.ParseInt(pathMatches[1], 10, 64)
	if err != nil {
		return nil, errors.Annotate(err, "could not parse change num from URL: %s", rawurl).Err()
	}

	patchset, err := strconv.ParseInt(pathMatches[2], 10, 32)
	if err != nil {
		return nil, errors.Annotate(err, "could not parse patchset from URL: %s", rawurl).Err()
	}

	key := &ChangeRevKey{
		Host:      parsedURL.Host,
		ChangeNum: changeNum,
		Revision:  int32(patchset),
	}

	return key, nil
}

// GetChangeRev gets details from Gerrit about a change,revision pair.
func GetChangeRev(ctx context.Context, authedClient *http.Client, changeNum int64, revision int32, host string, retryOpts shared.Options) (*ChangeRev, error) {
	var g gerritpb.GerritClient
	var err error
	if mockGerrit != nil {
		g = mockGerrit
	} else {
		if g, err = gerrit.NewRESTClient(authedClient, host, true); err != nil {
			return nil, err
		}
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	ch := make(chan *gerritpb.ChangeInfo, 1)
	err = shared.DoWithRetry(ctx, retryOpts, func() error {
		// This sets the deadline for the individual API call, while the outer context sets
		// an overall timeout for all attempts.
		innerCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		change, err := g.GetChange(innerCtx, &gerritpb.GetChangeRequest{
			Number: changeNum,
			Options: []gerritpb.QueryOption{
				gerritpb.QueryOption_ALL_REVISIONS,
				gerritpb.QueryOption_ALL_FILES,
			}})
		if err != nil {
			return err
		}
		ch <- change
		return nil
	})
	if err != nil {
		return nil, err
	}
	change := <-ch
	for _, v := range change.GetRevisions() {
		if v.Number == revision {
			files := v.Files
			// In some cases (e.g. merge commit), GetChange doesn't return a file list.
			// We thus call into the ListFiles endpoint instead.
			if len(files) == 0 {
				listFiles, err := fetchFileList(ctx, host, change.Number, v.Number, authedClient)
				if err != nil {
					return nil, err
				}
				files = listFiles.Files
			}
			return &ChangeRev{
				ChangeRevKey: ChangeRevKey{
					Host:      host,
					ChangeNum: change.Number,
					Revision:  v.Number,
				},
				Branch:  change.Ref,
				Ref:     v.Ref,
				Project: change.Project,
				Files:   getKeys(files),
			}, nil
		}
	}
	return nil, fmt.Errorf("found no revision %d for change %d on host %s", revision, changeNum, host)
}

func fetchFileList(ctx context.Context, host string, change int64, revision int32, httpClient *http.Client) (*gerritpb.ListFilesResponse, error) {
	rest, err := gerrit.NewRESTClient(httpClient, host, true)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	ch := make(chan *gerritpb.ListFilesResponse, 1)
	shared.DoWithRetry(ctx, shared.DefaultOpts, func() error {
		// The "Parent: 1" is what makes ListFiles able to get file lists for merge commits.
		// It's a 1-indexed way to reference parent commits, and we always want a value of 1
		// in order to get the target branch ref.
		resp, err := rest.ListFiles(ctx, &gerritpb.ListFilesRequest{Number: int64(change), RevisionId: strconv.Itoa(int(revision)), Parent: 1})
		if err != nil {
			return err
		}
		ch <- resp
		return nil
	})
	results := <-ch
	return results, nil
}

// ChangeRevData encapsulates a bunch of Gerrit change revisions.
type ChangeRevData struct {
	m map[string]*ChangeRev
}

// GetChangeRev extracts a single Gerrit change revision from the ChangeRevData.
func (crv ChangeRevData) GetChangeRev(host string, changeNum int64, revision int32) (*ChangeRev, error) {
	key := ChangeRevKey{Host: host, ChangeNum: changeNum, Revision: revision}.String()
	val, found := crv.m[key]
	if !found {
		return nil, fmt.Errorf("No ChangeRev found for key %s", key)
	}
	return val, nil
}

// GetChangeRevData fetches the Gerrit changes for the provided ChangeIdKeys, and bundles the result
// into a ChangeRevData.
func GetChangeRevData(ctx context.Context, authedClient *http.Client, changeIds []ChangeRevKey) (*ChangeRevData, error) {
	output := &ChangeRevData{m: make(map[string]*ChangeRev)}

	// If there are a large number of changes, there is a change of exceeding
	// Gerrit quota. Use longer retry opts in this case.
	retryOpts := shared.DefaultOpts
	if len(changeIds) > 20 {
		retryOpts = shared.LongerOpts
	}

	for _, c := range changeIds {
		if _, exists := output.m[c.String()]; !exists {
			cr, err := GetChangeRev(ctx, authedClient, c.ChangeNum, c.Revision, c.Host, retryOpts)
			if err != nil {
				return output, err
			}
			output.m[c.String()] = cr
		}
	}
	return output, nil
}

// GetChangeRevsForTest is intended for testing only, and it allows creation of a ChangeRevData
// through the supplied ChangeRevs.
func GetChangeRevsForTest(cr []*ChangeRev) *ChangeRevData {
	output := &ChangeRevData{m: make(map[string]*ChangeRev)}
	for _, c := range cr {
		output.m[c.String()] = c
	}
	return output
}

// getKeys extracts the keyset from the provided map.
func getKeys(m map[string]*gerritpb.FileInfo) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
