// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/appengine/gaetesting"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"

	"infra/appengine/unified-fleet/app/config"
)

const (
	nilEntity         string = "Invalid input - no Entity to add/update"
	emptyID           string = "Invalid input - Entity ID/Name is empty"
	invalidCharacters string = "Invalid input - Entity ID/Name must contain only 4-63 characters, ASCII letters, numbers, dash and underscore."
	alreadyExists     string = "already exists"
	notFound          string = "Not found"
)

type testFixture struct {
	T *testing.T
	C context.Context

	Fleet *FleetServerImpl
}

func newTestFixtureWithContext(ctx context.Context, t *testing.T) (testFixture, func()) {
	tf := testFixture{T: t, C: ctx}
	mc := gomock.NewController(t)

	tf.Fleet = &FleetServerImpl{}

	validate := func() {
		mc.Finish()
	}
	return tf, validate
}

func testingContext() context.Context {
	c := gaetesting.TestingContextWithAppID("dev~infra-unified-fleet-system")
	c = gologger.StdConfig.Use(c)
	c = logging.SetLevel(c, logging.Debug)
	c = config.Use(c, &config.Config{})
	datastore.GetTestable(c).Consistent(true)
	return c
}
