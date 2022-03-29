// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package vpython

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.chromium.org/luci/hardcoded/chromeinfra"
	"go.chromium.org/luci/vpython/api/vpython"
	"go.chromium.org/luci/vpython/application"
	"go.chromium.org/luci/vpython/cipd"
	"go.chromium.org/luci/vpython/spec"

	"github.com/mitchellh/go-homedir"
	cipdClient "go.chromium.org/luci/cipd/client/cipd"
	"go.chromium.org/luci/cipd/client/cipd/plugin/host"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/common/system/environ"
)

const (
	// BypassENV is an environment variable that is used to detect if we shouldn't
	// do any vpython stuff at all, but should instead directly invoke the next
	// `python` on PATH.
	BypassENV = "VPYTHON_BYPASS"

	// BypassSentinel must be the BypassENV value (verbatim) in order to trigger
	// vpython bypass.
	BypassSentinel = "manually managed python not supported by chrome operations"
)

var cipdPackageLoader = cipd.PackageLoader{
	Options: cipdClient.ClientOptions{
		ServiceURL: chromeinfra.CIPDServiceURL,
		UserAgent:  fmt.Sprintf("vpython, %s", cipdClient.UserAgent),
	},
	Template: func(c context.Context, tags []*vpython.PEP425Tag) (map[string]string, error) {
		tag := pep425TagSelector(tags)
		if tag == nil {
			return nil, nil
		}
		return getPEP425CIPDTemplateForTag(tag)
	},
}

var defaultConfig = application.Config{
	PackageLoader: &cipdPackageLoader,
	SpecLoader: spec.Loader{
		CommonFilesystemBarriers: []string{
			".gclient",
		},
		CommonSpecNames: []string{
			".vpython",
		},
		PartnerSuffix: ".vpython",
	},
	DefaultSpec: vpython.Spec{
		PythonVersion: "2",
	},
	VENVPackage: vpython.Spec_Package{
		Name:    "infra/3pp/tools/virtualenv",
		Version: "version:2@16.7.10.chromium.7",
	},
	PruneThreshold:          7 * 24 * time.Hour, // One week.
	MaxPrunesPerSweep:       3,
	DefaultVerificationTags: verificationScenarios,
}

func mainImpl(c context.Context, argv []string, env environ.Env, python3only bool) int {
	// Enable plugins support. Important to call it before LoadFromEnv to make
	// sure plugin related env vars are recognized.
	cipdPackageLoader.Options.PluginHost = &host.Host{PluginsContext: c}

	// Initialize our CIPD package loader from the environment.
	//
	// If we don't have an environment-specific CIPD cache directory, use one
	// relative to the user's home directory.
	if err := cipdPackageLoader.Options.LoadFromEnv(env.GetEmpty); err != nil {
		logging.Errorf(c, "Could not inialize CIPD package loader: %s", err)
		return 1
	}
	if cipdPackageLoader.Options.CacheDir == "" {
		hd, err := homedir.Dir()
		if err == nil {
			cipdPackageLoader.Options.CacheDir = filepath.Join(hd, ".vpython_cipd_cache")
		} else {
			logging.WithError(err).Warningf(c,
				"Failed to resolve user home directory. No CIPD cache will be enabled.")
		}
	}

	// Determine if we're bypassing "vpython".
	defaultConfig.Bypass = env.GetEmpty(BypassENV) == BypassSentinel
	// Determine if we're operating in "vpython3" mode (invoked as ./vpython3, ./vpython3.exe,
	// ./python3, or ./python3.exe). Alternately, the `vpython3` binary version
	// will explicitly indicate python3-only behavior.
	if python3only || strings.HasSuffix(argv[0], "python3") || strings.HasSuffix(argv[0], "python3.exe") {
		defaultConfig.SpecLoader.CommonSpecNames = []string{".vpython3"}
		defaultConfig.SpecLoader.PartnerSuffix = ".vpython3"
		defaultConfig.DefaultSpec.PythonVersion = "3.8"
		defaultConfig.DefaultVerificationTags = verificationScenarios38
		defaultConfig.VpythonOptIn = true
	}
	return defaultConfig.Main(c, argv, env)
}

// Main implements the vpython binary.
//
// If `python3only` is false, this will inspect argv[0] to determine if the
// current process "looks like" a python3 invocation and set defaults
// accordingly. Otherwise this will always have python3 behavior.
//
// The argv[0] inspection is for backwards compatibility with original `vpython`
// deployments which symlinked `vpython3` to the actual `vpython` binary.
func Main(python3only bool) {
	c := context.Background()
	c = gologger.StdConfig.Use(logging.SetLevel(c, logging.Warning))
	ret := mainImpl(c, os.Args, environ.System(), python3only)
	// os.Exit seems not to flush logging targets on Windows. The logger stores
	// the logging target as io.Writer which has no mechanism to flush. Knowing
	// gologger.StdConfig is configured to use os.Stderr, flush it directly.
	// https://crbug.com/1017136.
	os.Stderr.Sync()
	os.Exit(ret)
}
