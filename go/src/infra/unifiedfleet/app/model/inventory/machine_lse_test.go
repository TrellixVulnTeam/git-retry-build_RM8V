// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package inventory

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/appengine/gaetesting"
	. "go.chromium.org/luci/common/testing/assertions"
	proto "infra/unifiedfleet/api/v1/proto"
	. "infra/unifiedfleet/app/model/datastore"
)

func mockMachineLSE(id string) *proto.MachineLSE {
	return &proto.MachineLSE{
		Name: id,
	}
}

func TestCreateMachineLSE(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContextWithAppID("go-test")
	datastore.GetTestable(ctx).Consistent(true)
	machineLSE1 := mockMachineLSE("machineLSE-1")
	machineLSE2 := mockMachineLSE("")
	Convey("CreateMachineLSE", t, func() {
		Convey("Create new machineLSE", func() {
			resp, err := CreateMachineLSE(ctx, machineLSE1)
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSE1)
		})
		Convey("Create existing machineLSE", func() {
			resp, err := CreateMachineLSE(ctx, machineLSE1)
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, AlreadyExists)
		})
		Convey("Create machineLSE - invalid ID", func() {
			resp, err := CreateMachineLSE(ctx, machineLSE2)
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, InternalError)
		})
	})
}

func TestUpdateMachineLSE(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContextWithAppID("go-test")
	machineLSE1 := mockMachineLSE("machineLSE-1")
	machineLSE2 := mockMachineLSE("machineLSE-1")
	machineLSE2.Hostname = "Linux Server"
	machineLSE3 := mockMachineLSE("machineLSE-3")
	machineLSE4 := mockMachineLSE("")
	Convey("UpdateMachineLSE", t, func() {
		Convey("Update existing machineLSE", func() {
			resp, err := CreateMachineLSE(ctx, machineLSE1)
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSE1)

			resp, err = UpdateMachineLSE(ctx, machineLSE2)
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSE2)
		})
		Convey("Update non-existing machineLSE", func() {
			resp, err := UpdateMachineLSE(ctx, machineLSE3)
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, NotFound)
		})
		Convey("Update machineLSE - invalid ID", func() {
			resp, err := UpdateMachineLSE(ctx, machineLSE4)
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, InternalError)
		})
	})
}

func TestGetMachineLSE(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContextWithAppID("go-test")
	machineLSE1 := mockMachineLSE("machineLSE-1")
	Convey("GetMachineLSE", t, func() {
		Convey("Get machineLSE by existing ID", func() {
			resp, err := CreateMachineLSE(ctx, machineLSE1)
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSE1)
			resp, err = GetMachineLSE(ctx, "machineLSE-1")
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSE1)
		})
		Convey("Get machineLSE by non-existing ID", func() {
			resp, err := GetMachineLSE(ctx, "machineLSE-2")
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, NotFound)
		})
		Convey("Get machineLSE - invalid ID", func() {
			resp, err := GetMachineLSE(ctx, "")
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, InternalError)
		})
	})
}

func TestListMachineLSEs(t *testing.T) {
	t.Parallel()
	Convey("ListMachineLSEs", t, func() {
		ctx := gaetesting.TestingContextWithAppID("go-test")
		datastore.GetTestable(ctx).Consistent(true)
		machineLSEs := make([]*proto.MachineLSE, 0, 4)
		for i := 0; i < 4; i++ {
			machineLSE1 := mockMachineLSE(fmt.Sprintf("machineLSE-%d", i))
			resp, err := CreateMachineLSE(ctx, machineLSE1)
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSE1)
			machineLSEs = append(machineLSEs, resp)
		}
		Convey("List machineLSEs - page_token invalid", func() {
			resp, nextPageToken, err := ListMachineLSEs(ctx, 5, "abc")
			So(resp, ShouldBeNil)
			So(nextPageToken, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, InvalidPageToken)
		})

		Convey("List machineLSEs - Full listing with no pagination", func() {
			resp, nextPageToken, err := ListMachineLSEs(ctx, 4, "")
			So(resp, ShouldNotBeNil)
			So(nextPageToken, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSEs)
		})

		Convey("List machineLSEs - listing with pagination", func() {
			resp, nextPageToken, err := ListMachineLSEs(ctx, 3, "")
			So(resp, ShouldNotBeNil)
			So(nextPageToken, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSEs[:3])

			resp, _, err = ListMachineLSEs(ctx, 2, nextPageToken)
			So(resp, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, machineLSEs[3:])
		})
	})
}
