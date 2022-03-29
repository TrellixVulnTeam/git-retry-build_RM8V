// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tasks

import (
	"context"
	"fmt"

	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/server/tq"
	"google.golang.org/protobuf/proto"

	"infra/appengine/rubber-stamper/internal/reviewer"
	"infra/appengine/rubber-stamper/tasks/taskspb"
)

func init() {
	tq.RegisterTaskClass(tq.TaskClass{
		ID:        "change-review-task",
		Prototype: (*taskspb.ChangeReviewTask)(nil),
		Queue:     "change-review-queue",
		Kind:      tq.NonTransactional,
		Handler: func(ctx context.Context, payload proto.Message) error {
			t := payload.(*taskspb.ChangeReviewTask)

			if err := reviewer.ReviewChange(ctx, t); err != nil {
				info := tq.TaskExecutionInfo(ctx)
				if info != nil && info.ExecutionCount >= 2 {
					logging.WithError(err).Errorf(ctx, "task (host %s, cl %d, revision %s) failed for at least 3 times", t.Host, t.Number, t.Revision)
				}

				return fmt.Errorf("failed to review change for host %s, cl %d, revision %s: %v", t.Host, t.Number, t.Revision, err.Error())
			}
			return nil
		},
	})
}

// EnqueueChangeReviewTask enqueues a change review task.
func EnqueueChangeReviewTask(ctx context.Context, host string, cl *gerritpb.ChangeInfo) error {
	t := &taskspb.ChangeReviewTask{
		Host:               host,
		Number:             cl.Number,
		Revision:           cl.CurrentRevision,
		Repo:               cl.Project,
		AutoSubmit:         (cl.Labels["Auto-Submit"] != nil) && (cl.Labels["Auto-Submit"].Approved != nil),
		RevertOf:           cl.RevertOf,
		CherryPickOfChange: cl.CherryPickOfChange,
		RevisionsCount:     int64(len(cl.Revisions)),
		OwnerEmail:         cl.Owner.Email,
		Hashtags:           cl.Hashtags,
		Created:            cl.Created,
	}
	dedupKey := fmt.Sprintf("change(%s,%d,%s)", t.Host, t.Number, t.Revision)

	return tq.AddTask(ctx, &tq.Task{
		Payload:          t,
		Title:            dedupKey,
		DeduplicationKey: dedupKey,
	})
}
