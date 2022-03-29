// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package gerrit contains functions for interacting with gerrit/gitiles.
package gerrit

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	gerrs "errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"infra/cros/internal/shared"

	"go.chromium.org/luci/common/api/gitiles"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gitilespb "go.chromium.org/luci/common/proto/gitiles"
)

type Client interface {
	FetchFilesFromGitiles(ctx context.Context, host, project, ref string, paths []string) (*map[string]string, error)
	DownloadFileFromGitiles(ctx context.Context, host, project, ref, path string) (string, error)
	DownloadFileFromGitilesToPath(ctx context.Context, host, project, ref, path, saveToPath string) error
	Branches(ctx context.Context, host, project string) (map[string]string, error)
	Projects(ctx context.Context, host string) ([]string, error)
	ListFiles(ctx context.Context, host, project, ref, path string) ([]string, error)
}

// Client is a client for interacting with gerrit.
type ProdClient struct {
	isTestClient bool
	authedClient *http.Client
	// gitilesClient maps individual gerrit host to gitiles client.
	gitilesClient map[string]gitilespb.GitilesClient
}

// NewClient returns a new Client object.
func NewClient(authedClient *http.Client) (*ProdClient, error) {
	return &ProdClient{
		isTestClient:  false,
		authedClient:  authedClient,
		gitilesClient: map[string]gitilespb.GitilesClient{},
	}, nil
}

// getHostClient retrieves the inner gitilespb.GitilesClient for the specific
// host if it exists and creates a new one if it does not.
func (c *ProdClient) getHostClient(host string) (gitilespb.GitilesClient, error) {
	if client, ok := c.gitilesClient[host]; ok {
		return client, nil
	}
	if c.isTestClient {
		return nil, fmt.Errorf("test clients must have all inner clients set at initialization.")
	}
	var err error
	c.gitilesClient[host], err = gitiles.NewRESTClient(c.authedClient, host, true)
	if err != nil {
		return nil, err
	}
	return c.gitilesClient[host], err
}

// NewTestClient returns a new Client that uses the provided GitilesClient
// objects.
func NewTestClient(gcs map[string]gitilespb.GitilesClient) *ProdClient {
	return &ProdClient{
		isTestClient:  true,
		gitilesClient: gcs,
	}
}

// FetchFilesFromGitiles fetches file contents from gitiles.
//
// project is the git project to fetch from.
// ref is the git-ref to fetch from.
// paths lists the paths inside the git project to fetch contents for.
//
// fetchFilesFromGitiles returns a map from path in the git project to the
// contents of the file at that path for each requested path.
func (c *ProdClient) FetchFilesFromGitiles(ctx context.Context, host, project, ref string, paths []string) (*map[string]string, error) {
	gc, err := c.getHostClient(host)
	if err != nil {
		return nil, err
	}
	contents, err := obtainGitilesBytes(ctx, gc, project, ref)
	if err != nil {
		return nil, err
	}
	return extractGitilesArchive(ctx, contents, paths)
}

// DownloadFileFromGitiles downloads a file from Gitiles.
func (c *ProdClient) DownloadFileFromGitiles(ctx context.Context, host, project, ref, path string) (string, error) {
	gc, err := c.getHostClient(host)
	if err != nil {
		return "", err
	}
	req := &gitilespb.DownloadFileRequest{
		Project:    project,
		Path:       path,
		Committish: ref,
	}
	contents, err := gc.DownloadFile(ctx, req)
	if err != nil {
		return "", err
	}
	return contents.Contents, err
}

// DownloadFileFromGitilesToPath downloads a file from Gitiles to a specified path.
func (c *ProdClient) DownloadFileFromGitilesToPath(ctx context.Context, host, project, ref, path, saveToPath string) error {
	contents, err := c.DownloadFileFromGitiles(ctx, host, project, ref, path)
	if err != nil {
		return err
	}

	// Use existing file mode if the file already exists.
	fileMode := os.FileMode(int(0644))
	if fileData, err := os.Stat(saveToPath); err != nil && !gerrs.Is(err, os.ErrNotExist) {
		return err
	} else if fileData != nil {
		fileMode = fileData.Mode()
	}

	return os.WriteFile(saveToPath, []byte(contents), fileMode)
}

func obtainGitilesBytes(ctx context.Context, gc gitilespb.GitilesClient, project string, ref string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 8*time.Minute)
	defer cancel()
	ch := make(chan *gitilespb.ArchiveResponse, 1)

	err := shared.DoWithRetry(ctx, shared.LongerOpts, func() error {
		// This sets the deadline for the individual API call, while the outer context sets
		// an overall timeout for all attempts.
		innerCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		req := &gitilespb.ArchiveRequest{
			Project: project,
			Ref:     ref,
			Format:  gitilespb.ArchiveRequest_GZIP,
		}
		a, err := gc.Archive(innerCtx, req)
		if err != nil {
			return errors.Annotate(err, "obtain gitiles archive").Err()
		}
		ch <- a
		return nil
	})
	if err != nil {
		return nil, err
	}
	a := <-ch
	return a.Contents, nil
}

// extractGitilesArchive extracts file at each path in paths from the given
// gunzipped tarfile.
//
// extractGitilesArchive returns a map from path to the content of the file at
// that path in the archives for each requested path found in the archive.
//
// This function takes ownership of data. Caller should not use the byte array
// concurrent to / after this call. See io.Reader interface for more details.
func extractGitilesArchive(ctx context.Context, data []byte, paths []string) (*map[string]string, error) {
	// pmap maps files to the requested filename.
	// e.g. if "foo" is a symlink to "bar", then the entry "bar":"foo" exists.
	// if a file is not a symlink, it will be mapped to itself.
	pmap := make(map[string]string)
	for _, p := range paths {
		pmap[p] = p
	}

	res := make(map[string]string)
	// Do two passes to resolve links.
	for i := 0; i < 2; i++ {
		abuf := bytes.NewBuffer(data)
		gr, err := gzip.NewReader(abuf)
		if err != nil {
			return nil, errors.Annotate(err, "extract gitiles archive").Err()
		}
		defer gr.Close()

		tr := tar.NewReader(gr)
		for {
			h, err := tr.Next()
			eof := false
			switch {
			case err == io.EOF:
				// Scanned all files.
				eof = true
			case err != nil:
				return nil, errors.Annotate(err, "extract gitiles archive").Err()
			default:
				// good case.
			}
			if eof {
				break
			}
			requestedFile, found := pmap[h.Name]
			if !found {
				// not a requested file.
				continue
			}
			if _, ok := res[requestedFile]; ok {
				// already read this file.
				continue
			}
			if h.Typeflag == tar.TypeSymlink {
				if i == 0 {
					// if symlink, mark link in pmap so it gets picked up on the second pass.
					linkPath := path.Join(path.Dir(h.Name), h.Linkname)
					pmap[linkPath] = h.Name
				}
				continue
			}

			logging.Debugf(ctx, "Inventory data file %s size %d", h.Name, h.Size)
			data := make([]byte, h.Size)
			if _, err := io.ReadFull(tr, data); err != nil {
				return nil, errors.Annotate(err, "extract gitiles archive").Err()
			}
			res[requestedFile] = string(data)
		}
	}
	return &res, nil
}

// Branches returns a map of branches (to revisions) for a given repo.
func (c *ProdClient) Branches(ctx context.Context, host, project string) (map[string]string, error) {
	gc, err := c.getHostClient(host)
	if err != nil {
		return nil, err
	}
	req := &gitilespb.RefsRequest{
		Project:  project,
		RefsPath: "refs/heads",
	}
	refs, err := gc.Refs(ctx, req)
	if err != nil {
		return nil, err
	}
	return refs.Revisions, err
}

// Projects returns a list of projects for a given host.
func (c *ProdClient) Projects(ctx context.Context, host string) ([]string, error) {
	gc, err := c.getHostClient(host)
	if err != nil {
		return nil, err
	}
	req := &gitilespb.ProjectsRequest{}
	resp, err := gc.Projects(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetProjects(), err
}

// ListFiles returns a list of files/directories for a given host/project/ref/path.
func (c *ProdClient) ListFiles(ctx context.Context, host, project, ref, path string) ([]string, error) {
	gc, err := c.getHostClient(host)
	if err != nil {
		return nil, err
	}
	req := &gitilespb.ListFilesRequest{
		Project:    project,
		Committish: ref,
		Path:       path,
	}
	resp, err := gc.ListFiles(ctx, req)
	if err != nil {
		return nil, err
	}
	files := resp.GetFiles()
	names := make([]string, len(files))
	for i, file := range files {
		names[i] = file.GetPath()
	}
	return names, err
}
