// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
package manifestutil

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"infra/cros/internal/assert"
	"infra/cros/internal/gerrit"
	"infra/cros/internal/repo"
	"infra/cros/internal/util"

	"github.com/golang/mock/gomock"
	gitilespb "go.chromium.org/luci/common/proto/gitiles"
	"go.chromium.org/luci/common/proto/gitiles/mock_gitiles"
)

var (
	fooManifest = &repo.Manifest{
		Default: repo.Default{
			RemoteName: "chromeos",
			Revision:   "123",
		},
		Remotes: []repo.Remote{
			{
				Fetch: "https://chromium.org/remote",
				Name:  "chromium",
				Alias: "chromeos",
			},
			{
				Fetch:    "https://google.com/remote",
				Name:     "google",
				Revision: "125",
			},
		},
		Projects: []repo.Project{
			{Path: "baz/", Name: "baz", RemoteName: "chromium"},
			{Path: "fiz/", Name: "fiz", Revision: "124"},
			{Name: "buz", RemoteName: "google",
				Annotations: []repo.Annotation{
					{Name: "branch-mode", Value: "pin"},
				},
			},
		},
		Includes: []repo.Include{
			{Name: "bar.xml"},
		},
	}
	fooXML = `<?xml version="1.0" encoding="UTF-8"?>
<manifest>
  <include name="bar.xml" />
  <default remote="chromeos" revision="123"/>
  <remote fetch="https://chromium.org/remote" name="chromium" alias="chromeos"/>
  <remote fetch="https://google.com/remote" name="google" revision="125"/>
  <project name="baz" path="baz/" remote="chromium"/>
  <project name="fiz" path="fiz/" revision="124" />
  <project name="buz" remote="google">
    <annotation name="branch-mode" value="pin"/>
  </project>
</manifest>`
)

func ManifestEq(a, b *repo.Manifest) bool {
	if len(a.Projects) != len(b.Projects) {
		return false
	}
	for i := range a.Projects {
		if !reflect.DeepEqual(&a.Projects[i], &b.Projects[i]) {
			return false
		}
	}
	if len(a.Includes) != len(b.Includes) {
		return false
	}
	for i := range a.Includes {
		if a.Includes[i] != b.Includes[i] {
			return false
		}
	}
	return true
}

func ManifestMapEq(expected, actual map[string]*repo.Manifest) error {
	for file := range expected {
		if _, ok := actual[file]; !ok {
			return fmt.Errorf("missing manifest %s", file)
		}
		if !ManifestEq(expected[file], actual[file]) {
			return fmt.Errorf("expected %v, found %v", expected[file], actual[file])
		}
	}
	return nil
}

func TestLoadManifestTreeFromFile_success(t *testing.T) {
	expectedResults := make(map[string]*repo.Manifest)
	expectedResults["foo.xml"] = fooManifest
	expectedResults["bar.xml"] = &repo.Manifest{
		Projects: []repo.Project{
			{Path: "baz/", Name: "baz"},
			{Path: "project/", Name: "project"},
		},
	}

	res, err := LoadManifestTreeFromFile("test_data/foo.xml")
	assert.NilError(t, err)
	if err = ManifestMapEq(expectedResults, res); err != nil {
		t.Errorf(err.Error())
	}
}

func TestLoadManifestTreeFromFile_bad_include(t *testing.T) {
	_, err := LoadManifestTreeFromFile("test_data/bogus.xml")
	assert.ErrorContains(t, err, "bad-include.xml")
}

func TestLoadManifestTreeFromFile_bad_xml(t *testing.T) {
	_, err := LoadManifestTreeFromFile("test_data/invalid.xml")
	assert.ErrorContains(t, err, "unmarshal")
}

func TestLoadManifestFromFile(t *testing.T) {
	manifest, err := LoadManifestFromFile("test_data/foo.xml")
	assert.NilError(t, err)
	assert.Assert(t, ManifestEq(fooManifest, manifest))
}

func TestLoadManifestFromFileRaw(t *testing.T) {
	data, err := LoadManifestFromFileRaw("test_data/foo.xml")
	assert.NilError(t, err)
	assert.StringsEqual(t, fooXML, string(data))
}

func TestLoadManifestFromFileWithIncludes(t *testing.T) {
	// We have decent coverage of the code in other TestLoadManifestFromFile...
	// tests so in this test (to reduce complexity and size) we just look to see
	// if all the projects from child manifests have been included.
	expectedProjectNames := []string{"baz", "fiz", "buz", "project"}

	res, err := LoadManifestFromFileWithIncludes("test_data/foo.xml")
	assert.NilError(t, err)

	projectNames := make([]string, len(res.Projects))
	for i, project := range res.Projects {
		projectNames[i] = project.Name
	}
	assert.Assert(t, util.UnorderedEqual(expectedProjectNames, projectNames))
}

func TestLoadManifestFromGitiles(t *testing.T) {
	// Mock Gitiles controller
	ctl := gomock.NewController(t)
	gitilesMock := mock_gitiles.NewMockGitilesClient(ctl)

	project := "foo"
	branch := "refs/heads/foo"

	manifestXML := `
<?xml version="1.0" encoding="UTF-8"?>
<manifest>
  <default remote="chromeos" revision="123"/>
  <remote name="chromeos" />
  <include name="sub.xml" />

  <project name="foo" path="foo/" />
  <project name="bar" path="bar/" />
</manifest>
	`
	subXML := `
<?xml version="1.0" encoding="UTF-8"?>
<manifest>
  <default remote="chromium" />
  <remote name="chromium" />

  <project name="baz" path="baz/" />
</manifest>
	`

	reqManifest := &gitilespb.DownloadFileRequest{
		Project:    project,
		Path:       "manifest.xml",
		Committish: branch,
		Format:     gitilespb.DownloadFileRequest_TEXT,
	}
	gitilesMock.EXPECT().DownloadFile(gomock.Any(), gerrit.DownloadFileRequestEq(reqManifest)).AnyTimes().Return(
		&gitilespb.DownloadFileResponse{
			Contents: manifestXML,
		},
		nil,
	)
	reqSubManifest := &gitilespb.DownloadFileRequest{
		Project:    project,
		Path:       "sub.xml",
		Committish: branch,
		Format:     gitilespb.DownloadFileRequest_TEXT,
	}
	gitilesMock.EXPECT().DownloadFile(gomock.Any(), gerrit.DownloadFileRequestEq(reqSubManifest)).AnyTimes().Return(
		&gitilespb.DownloadFileResponse{
			Contents: subXML,
		},
		nil,
	)

	expected := map[string]*repo.Manifest{
		"manifest.xml": {
			Default: repo.Default{
				RemoteName: "chromeos",
				Revision:   "123",
			},
			Remotes: []repo.Remote{
				{
					Name: "chromeos",
				},
			},
			Includes: []repo.Include{
				{
					Name: "sub.xml",
				},
			},
			Projects: []repo.Project{
				{
					Name: "foo",
					Path: "foo/",
				},
				{
					Name: "bar",
					Path: "bar/",
				},
			},
		},
		"sub.xml": {
			Default: repo.Default{
				RemoteName: "chromium",
			},
			Remotes: []repo.Remote{
				{
					Name: "chromium",
				},
			},
			Projects: []repo.Project{
				{
					Name: "baz",
					Path: "baz/",
				},
			},
		},
	}
	expectedMerged := &repo.Manifest{
		Default: repo.Default{
			RemoteName: "chromeos",
			Revision:   "123",
		},
		Remotes: []repo.Remote{
			{
				Name: "chromeos",
			},
			{
				Name: "chromium",
			},
		},
		Projects: []repo.Project{
			{
				Name:       "foo",
				Path:       "foo/",
				Revision:   "123",
				RemoteName: "chromeos",
			},
			{
				Name:       "bar",
				Path:       "bar/",
				Revision:   "123",
				RemoteName: "chromeos",
			},
			{
				Name:       "baz",
				Path:       "baz/",
				Revision:   "123",
				RemoteName: "chromium",
			},
		},
	}
	gerrit.MockGitiles = gitilesMock
	ctx := context.Background()

	// Test LoadManifestFromGitiles
	got, err := LoadManifestFromGitiles(ctx, nil, "host", project, branch, "manifest.xml")
	assert.NilError(t, err)
	assert.Assert(t, ManifestEq(got, expected["manifest.xml"]))

	// Test LoadManifestFromGitilesWithIncludes
	got, err = LoadManifestFromGitilesWithIncludes(ctx, nil, "host", project, branch, "manifest.xml")
	assert.NilError(t, err)
	assert.Assert(t, ManifestEq(got, expectedMerged))

	// Test LoadManifestTreeFromGitiles
	gotMap, err := LoadManifestTreeFromGitiles(ctx, nil, "host", project, branch, "manifest.xml")
	assert.NilError(t, err)
	assert.NilError(t, ManifestMapEq(gotMap, expected))
}
