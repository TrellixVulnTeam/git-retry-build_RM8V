// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"context"
	"infra/tools/kitchen/cookflags"
	"strings"

	swarming "go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.chromium.org/luci/common/errors"
	logdog_types "go.chromium.org/luci/logdog/common/types"
)

func (s *Systemland) genSwarmingTask(ctx context.Context, uid string) (st *swarming.SwarmingRpcsNewTaskRequest, args *cookflags.CookFlags, err error) {
	st = &(*s.SwarmingTask)
	st.TaskSlices[0] = &(*st.TaskSlices[0])
	st.TaskSlices[0].Properties.Env = exfiltrateMap(s.Env)

	if s.KitchenArgs != nil {
		args = &(*s.KitchenArgs)

		// generate AnnotationURL, if needed, and add it to tags
		if strings.Contains(string(args.LogDogFlags.AnnotationURL.Path), generateLogdogToken) {
			var prefix logdog_types.StreamName
			prefix, err = generateLogdogStream(ctx, uid)
			if err != nil {
				err = errors.Annotate(err, "generating logdog prefix").Err()
				return
			}
			args.LogDogFlags.AnnotationURL.Path = logdog_types.StreamPath(strings.Replace(
				string(args.LogDogFlags.AnnotationURL.Path), generateLogdogToken,
				string(prefix), -1))
		}
		if !args.LogDogFlags.AnnotationURL.IsZero() {
			st.Tags = append(st.Tags,
				"allow_milo:1",
				"log_location:"+args.LogDogFlags.AnnotationURL.String(),
			)
		}
	}

	if len(s.CipdPkgs) > 0 {
		if st.TaskSlices[0].Properties.CipdInput == nil {
			st.TaskSlices[0].Properties.CipdInput = &swarming.SwarmingRpcsCipdInput{}
		}

		for subdir, pkgsVers := range s.CipdPkgs {
			for pkg, ver := range pkgsVers {
				st.TaskSlices[0].Properties.CipdInput.Packages = append(
					st.TaskSlices[0].Properties.CipdInput.Packages,
					&swarming.SwarmingRpcsCipdPackage{
						Path:        subdir,
						PackageName: pkg,
						Version:     ver,
					})
			}
		}
	}

	return
}
