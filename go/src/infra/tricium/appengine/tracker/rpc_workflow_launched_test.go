// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tracker

import (
	"testing"

	"golang.org/x/net/context"

	ds "go.chromium.org/gae/service/datastore"

	. "github.com/smartystreets/goconvey/convey"

	admin "infra/tricium/api/admin/v1"
	"infra/tricium/api/v1"
	"infra/tricium/appengine/common/track"
	"infra/tricium/appengine/common/triciumtest"
)

const (
	clangIsolatorUbuntu  = "ClangIsolator_Ubuntu"
	clangIsolatorWindows = "ClangIsolator_Windows"
	fileIsolator         = "GitFileIsolator_Ubuntu"
)

// mockWorkflowProvider mocks common.WorkflowProvider.
type mockWorkflowProvider struct {
}

func (mockWorkflowProvider) GetWorkflow(c context.Context, runID int64) (*admin.Workflow, error) {
	return &admin.Workflow{
		Workers: []*admin.Worker{
			{
				Name:  clangIsolatorUbuntu,
				Needs: tricium.Data_FILES,
			},
			{
				Name:  clangIsolatorWindows,
				Needs: tricium.Data_FILES,
			},
			{
				Name:  fileIsolator,
				Needs: tricium.Data_GIT_FILE_DETAILS,
				Next: []string{
					clangIsolatorUbuntu,
					clangIsolatorWindows,
				},
			},
		},
	}, nil
}

func TestWorkflowLaunchedRequest(t *testing.T) {
	Convey("Test Environment", t, func() {
		ctx := triciumtest.Context()

		Convey("Workflow request", func() {
			// Add pending workflow run entity.
			request := &track.AnalyzeRequest{}
			So(ds.Put(ctx, request), ShouldBeNil)
			requestKey := ds.KeyForObj(ctx, request)
			run := &track.WorkflowRun{ID: 1, Parent: requestKey}
			So(ds.Put(ctx, run), ShouldBeNil)
			runKey := ds.KeyForObj(ctx, run)
			runResult := &track.WorkflowRunResult{
				ID:     1,
				Parent: runKey,
				State:  tricium.State_PENDING,
			}
			So(ds.Put(ctx, runResult), ShouldBeNil)

			// Mark workflow as launched.
			err := workflowLaunched(ctx, &admin.WorkflowLaunchedRequest{
				RunId: request.ID,
			}, mockWorkflowProvider{})
			So(err, ShouldBeNil)

			name, _, err := track.ExtractFunctionPlatform(fileIsolator)

			Convey("Marks workflow run as launched", func() {
				// Run entry is marked as launched.
				So(ds.Get(ctx, runResult), ShouldBeNil)
				So(runResult.State, ShouldEqual, tricium.State_RUNNING)
			})

			Convey("Adds list of functions to WorkflowRun", func() {
				So(ds.Get(ctx, run), ShouldBeNil)
				So(len(run.Functions), ShouldEqual, 3)
			})

			Convey("Worker and function is marked pending", func() {
				So(err, ShouldBeNil)
				functionRunKey := ds.NewKey(ctx, "FunctionRun", name, 0, runKey)
				workerKey := ds.NewKey(ctx, "WorkerRun", fileIsolator, 0, functionRunKey)
				wr := &track.WorkerRunResult{ID: 1, Parent: workerKey}
				So(ds.Get(ctx, wr), ShouldBeNil)
				So(wr.State, ShouldEqual, tricium.State_PENDING)
				fr := &track.FunctionRunResult{ID: 1, Parent: functionRunKey}
				So(ds.Get(ctx, fr), ShouldBeNil)
				So(fr.State, ShouldEqual, tricium.State_PENDING)
			})
		})
	})
}
