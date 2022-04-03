// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package testverdictingester

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
	"go.chromium.org/luci/server/span"
	"go.chromium.org/luci/server/tq"
	_ "go.chromium.org/luci/server/tq/txn/spanner"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	ctrlpb "infra/appengine/weetbix/internal/ingestion/control/proto"
	"infra/appengine/weetbix/internal/tasks/taskspb"
	"infra/appengine/weetbix/internal/testutil"
)

func init() {
	RegisterTaskClass()
}

func TestSchedule(t *testing.T) {
	Convey(`TestSchedule`, t, func() {
		ctx := testutil.SpannerTestContext(t)
		ctx, skdr := tq.TestingContext(ctx, nil)

		task := &taskspb.IngestTestVerdicts{
			Build:         &ctrlpb.BuildResult{},
			PartitionTime: timestamppb.New(time.Date(2025, time.January, 1, 12, 0, 0, 0, time.UTC)),
		}
		expected := proto.Clone(task).(*taskspb.IngestTestVerdicts)

		_, err := span.ReadWriteTransaction(ctx, func(ctx context.Context) error {
			Schedule(ctx, task)
			return nil
		})
		So(err, ShouldBeNil)
		So(skdr.Tasks().Payloads()[0], ShouldResembleProto, expected)
	})
}
