// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/tsmon/field"
	"go.chromium.org/luci/common/tsmon/metric"
	cvv0 "go.chromium.org/luci/cv/api/v0"
	cvv1 "go.chromium.org/luci/cv/api/v1"
	"go.chromium.org/luci/server/router"
	"google.golang.org/protobuf/encoding/protojson"

	"infra/appengine/weetbix/internal/cv"
	ctlpb "infra/appengine/weetbix/internal/ingestion/control/proto"
	pb "infra/appengine/weetbix/proto/v1"
)

const (
	// TODO(chanli@@) Removing the hosts after CVPubSub and GetRun RPC added them.
	// Host name of buildbucket.
	bbHost = "cr-buildbucket.appspot.com"

	// maximumCLs is the maximum number of CLs to capture from any completed
	// CV run, after which the CL list is truncated. This avoids CV Runs with
	// an excessive number of included CLs from storing an excessive amount
	// of data per failure.
	maximumCLs = 10
)

var (
	cvRunCounter = metric.NewCounter(
		"weetbix/ingestion/pubsub/cv_runs",
		"The number of CV runs received by Weetbix from PubSub.",
		nil,
		// The LUCI Project.
		field.String("project"),
		// "success", "transient-failure", "permanent-failure" or "ignored".
		field.String("status"))

	runIDRe = regexp.MustCompile(`^projects/(.*)/runs/(.*)$`)

	// Automation service accounts.
	automationAccountRE = regexp.MustCompile(`^.*@.*\.gserviceaccount\.com$`)
)

// CVRunPubSubHandler accepts and processes CV Pub/Sub messages.
func CVRunPubSubHandler(ctx *router.Context) {
	status := "unknown"
	project := "unknown"
	defer func() {
		// Closure for late binding.
		cvRunCounter.Add(ctx.Context, 1, project, status)
	}()
	project, processed, err := cvPubSubHandlerImpl(ctx.Context, ctx.Request)

	switch {
	case err != nil:
		errors.Log(ctx.Context, errors.Annotate(err, "handling cv pubsub event").Err())
		status = processErr(ctx, err)
		return
	case !processed:
		status = "ignored"
		// Use subtly different "success" response codes to surface in
		// standard GAE logs whether an ingestion was ignored or not,
		// while still acknowledging the pub/sub.
		// See https://cloud.google.com/pubsub/docs/push#receiving_messages.
		ctx.Writer.WriteHeader(http.StatusNoContent) // 204
	default:
		status = "success"
		ctx.Writer.WriteHeader(http.StatusOK)
	}
}

func cvPubSubHandlerImpl(ctx context.Context, request *http.Request) (project string, processed bool, err error) {
	psRun, err := extractPubSubRun(request)
	if err != nil {
		return "unknown", false, errors.Annotate(err, "failed to extract run").Err()
	}

	project, runID, err := parseRunID(psRun.Id)
	if err != nil {
		return "unknown", false, errors.Annotate(err, "failed to parse run ID").Err()
	}

	// We do not check if the project is configured in Weetbix,
	// as CV runs can include projects from other projects.
	// It is the projects of these builds that the data
	// gets ingested into.

	run, err := getRun(ctx, psRun)
	switch {
	case err != nil:
		return project, false, errors.Annotate(err, "failed to get run").Err()
	case run.GetCreateTime() == nil:
		return project, false, errors.New("could not get create time for the run")
	}

	owner := "user"
	if automationAccountRE.MatchString(run.Owner) {
		owner = "automation"
	}

	presubmitResultByBuildID := make(map[string]*ctlpb.PresubmitResult)
	for _, tj := range run.Tryjobs {
		b := tj.GetResult().GetBuildbucket()
		if b == nil {
			// Non build-bucket result.
			continue
		}

		if tj.Reuse {
			// Do not ingest re-used tryjobs.
			// Builds should be ingested with the CV run that initiated
			// them, not a CV run that re-used them.
			// Tryjobs can also be marked re-used if they were user
			// initiated through gerrit. In this case, the build would
			// not have been tagged as being part of a CV run (e.g.
			// through user_agent: cq), so it will not expect to be
			// joined to a CV run.
			continue
		}

		buildID := buildID(bbHost, b.Id)
		if _, ok := presubmitResultByBuildID[buildID]; ok {
			logging.Warningf(ctx, "CV Run %s has build %s as tryjob multiple times, ignoring the second occurances", psRun.Id, buildID)
			continue
		}

		presubmitResultByBuildID[buildID] = &ctlpb.PresubmitResult{
			PresubmitRunId: &pb.PresubmitRunId{
				System: "luci-cv",
				Id:     fmt.Sprintf("%s/%s", project, runID),
			},
			PresubmitRunSucceeded: run.Status == cvv0.Run_SUCCEEDED,
			Mode:                  run.GetMode(),
			Owner:                 owner,
			Cls:                   extractRunChangelists(run.Cls),
			CreationTime:          run.CreateTime,
			Critical:              tj.Critical,
		}
	}

	if err := JoinPresubmitResult(ctx, presubmitResultByBuildID, project); err != nil {
		return project, true, errors.Annotate(err, "joining presubmit results").Err()
	}

	return project, true, nil
}

func extractRunChangelists(cls []*cvv0.GerritChange) []*pb.Changelist {
	result := make([]*pb.Changelist, 0, len(cls))
	for _, cl := range cls {
		result = append(result, &pb.Changelist{
			Host:     cl.Host,
			Change:   cl.Change,
			Patchset: cl.Patchset,
		})
	}
	// Sort changelists in ascending order by host, then change,
	// then patchset. This ensures CLs appear in a stable order for
	// multi-CL CV runs.
	sortChangelists(result)

	if len(result) > maximumCLs {
		result = result[:maximumCLs]
	}
	return result
}

func sortChangelists(cls []*pb.Changelist) {
	less := func(i, j int) bool {
		if cls[i].Host < cls[j].Host {
			return true
		}
		if cls[i].Host == cls[j].Host &&
			cls[i].Change < cls[j].Change {
			return true
		}
		if cls[i].Host == cls[j].Host &&
			cls[i].Change == cls[j].Change &&
			cls[i].Patchset < cls[j].Patchset {
			return true
		}
		return false
	}
	sort.Slice(cls, less)
}

func extractPubSubRun(r *http.Request) (*cvv1.PubSubRun, error) {
	var msg pubsubMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		return nil, errors.Annotate(err, "could not decode cv pubsub message").Err()
	}

	var run cvv1.PubSubRun
	err := protojson.Unmarshal(msg.Message.Data, &run)
	if err != nil {
		return nil, errors.Annotate(err, "could not parse cv pubsub message data").Err()
	}
	return &run, nil
}

func parseRunID(runID string) (project string, run string, err error) {
	m := runIDRe.FindStringSubmatch(runID)
	if m == nil {
		return "", "", errors.Reason("run ID does not match %s", runIDRe).Err()
	}
	return m[1], m[2], nil
}

// getRun gets the full Run message by make a GetRun RPC to CV.
//
// Currently we're calling cv.v0.Runs.GetRun, and should switch to v1 when it's
// ready to use.
func getRun(ctx context.Context, psRun *cvv1.PubSubRun) (*cvv0.Run, error) {
	c, err := cv.NewClient(ctx, psRun.Hostname)
	if err != nil {
		return nil, errors.Annotate(err, "failed to create cv client").Err()
	}
	req := &cvv0.GetRunRequest{
		Id: psRun.Id,
	}
	return c.GetRun(ctx, req)
}
