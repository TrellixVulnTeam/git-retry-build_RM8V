// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package verdicts

import (
	"testing"

	"infra/appengine/weetbix/internal/testutil"
)

func TestMain(m *testing.M) {
	testutil.SpannerTestMain(m)
}
