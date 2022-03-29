// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
package gerrit

import (
	"testing"

	"infra/cros/internal/assert"

	gitilespb "go.chromium.org/luci/common/proto/gitiles"
)

func TestDownloadFileRequestMatcher(t *testing.T) {
	req := &gitilespb.DownloadFileRequest{
		Project:    "chromeos/manifest-versions",
		Path:       "foo",
		Committish: "HEAD",
	}
	matcher := DownloadFileRequestEq(req)

	a := &gitilespb.DownloadFileRequest{
		Project:    "chromeos/manifest-versions",
		Path:       "foo",
		Committish: "HEAD",
	}
	b := &gitilespb.DownloadFileRequest{
		Project:    "chromeos/manifest-versions",
		Path:       "bar",
		Committish: "HEAD^1",
	}
	assert.Assert(t, matcher.Matches(a))
	assert.Assert(t, !matcher.Matches(b))
}
func TestRefsRequestMatcher(t *testing.T) {
	req := &gitilespb.RefsRequest{
		Project:  "chromeos/manifest-versions",
		RefsPath: "refs/heads",
	}
	matcher := RefsRequestEq(req)

	a := &gitilespb.RefsRequest{
		Project:  "chromeos/manifest-versions",
		RefsPath: "refs/heads",
	}
	b := &gitilespb.RefsRequest{
		Project: "chromeos/manifest-versions",
	}
	assert.Assert(t, matcher.Matches(a))
	assert.Assert(t, !matcher.Matches(b))
}

func TestListFilesRequestMatcher(t *testing.T) {
	req := &gitilespb.ListFilesRequest{
		Project:    "chromeos/manifest-versions",
		Committish: "refs/heads",
		Path:       "foo/",
	}
	matcher := ListFilesRequestEq(req)

	a := &gitilespb.ListFilesRequest{
		Project:    "chromeos/manifest-versions",
		Committish: "refs/heads",
		Path:       "foo/",
	}
	b := &gitilespb.ListFilesRequest{
		Project:    "chromeos/manifest-versions",
		Committish: "refs/heads",
	}
	assert.Assert(t, matcher.Matches(a))
	assert.Assert(t, !matcher.Matches(b))
}
