// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/client"
	"github.com/maruel/subcommands"
	"github.com/pkg/errors"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/phosphorus"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/logging"

	"infra/cros/cmd/phosphorus/internal/autotest/atutil"
)

// RunTest subcommand: Run a test against one or multiple DUTs.
var RunTest = &subcommands.Command{
	UsageLine: "run-test -input_json /path/to/input.json",
	ShortDesc: "Run a test against one or multiple DUTs.",
	LongDesc: `Run a test against one or multiple DUTs.

A wrapper around 'autoserv'.`,
	CommandRun: func() subcommands.CommandRun {
		c := &runTestRun{}
		c.Flags.StringVar(&c.InputPath, "input_json", "", "Path that contains JSON encoded test_platform.phosphorus.RunTestRequest")
		c.Flags.StringVar(&c.OutputPath, "output_json", "", "Path to write JSON encoded test_platform.phosphorus.RunTestResponse to")
		return c
	},
}

type runTestRun struct {
	CommonRun
}

type autoservResult struct {
	CmdResult  *atutil.Result
	ResultsDir string
}

func (c *runTestRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.ValidateArgs(); err != nil {
		fmt.Fprintln(a.GetErr(), err.Error())
		c.Flags.Usage()
		return 1
	}

	ctx := cli.GetContext(a, c, env)
	if err := c.innerRun(ctx, args, env); err != nil {
		logApplicationError(ctx, a, err)
		return 1
	}
	return 0
}

func (c *runTestRun) innerRun(ctx context.Context, args []string, env subcommands.Env) error {
	r := &phosphorus.RunTestRequest{}
	if err := ReadJSONPB(c.InputPath, r); err != nil {
		return err
	}
	if err := validateRunTestRequest(r); err != nil {
		return err
	}

	if d := r.Deadline.AsTime(); !d.IsZero() {
		var c context.CancelFunc
		log.Printf("Running with deadline %s (current time: %s)", d, time.Now().UTC())
		ctx, c = context.WithDeadline(ctx, d)
		defer c()
	}
	ar, err := runTestStep(ctx, r)
	if err != nil {
		return err
	}
	return WriteJSONPB(c.OutputPath, runTestResponse(ar))
}

func runTestResponse(r *autoservResult) *phosphorus.RunTestResponse {
	return &phosphorus.RunTestResponse{
		ResultsDir: r.ResultsDir,
		State:      runTestState(r.CmdResult),
	}
}

func runTestState(r *atutil.Result) phosphorus.RunTestResponse_State {
	if r.Success() {
		return phosphorus.RunTestResponse_SUCCEEDED
	}
	if r.RunResult.Aborted {
		return phosphorus.RunTestResponse_ABORTED
	}
	return phosphorus.RunTestResponse_FAILED
}

func validateRunTestRequest(r *phosphorus.RunTestRequest) error {
	missingArgs := getCommonMissingArgs(r.Config)

	if len(r.DutHostnames) == 0 {
		missingArgs = append(missingArgs, "DUT hostname(s)")
	}

	if r.GetAutotest().GetName() == "" {
		missingArgs = append(missingArgs, "test name")
	}

	if len(missingArgs) > 0 {
		return fmt.Errorf("no %s provided", strings.Join(missingArgs, ", "))
	}

	return nil
}

// runTestStep runs an individual test. It is a wrapper around autoserv.
func runTestStep(ctx context.Context, r *phosphorus.RunTestRequest) (*autoservResult, error) {
	j := getMainJob(r.Config)

	dir := filepath.Join(r.Config.Task.ResultsDir, "autoserv_test")

	t := &atutil.Test{
		Args:               r.GetAutotest().GetTestArgs(),
		ClientTest:         r.GetAutotest().GetIsClientTest(),
		ControlName:        r.GetAutotest().GetName(),
		ImageStorageServer: r.GetAutotest().GetImageStorageServer(),
		Hosts:              r.DutHostnames,
		Keyvals:            r.GetAutotest().GetKeyvals(),
		Name:               r.GetAutotest().GetDisplayName(),
		PeerDuts:           r.GetAutotest().GetPeerDuts(),
		RequireSSP:         !r.GetAutotest().GetIsClientTest(),
		ResultsDir:         dir,
		SSPBaseImageName:   r.Config.GetTask().GetSspBaseImageName(),
	}

	var dockerClient *client.Client

	if r.ContainerImageInfo != nil {
		var err error
		registry := r.GetContainerImageInfo().GetRepository().GetHostname()
		if registry == "" {
			return nil, fmt.Errorf("ContainerImageInfo must set repository hostname")
		}

		dockerClient, err = client.NewClientWithOpts(
			client.FromEnv,
			client.WithAPIVersionNegotiation(),
		)
		if err != nil {
			return nil, errors.Wrap(err, "creating Docker client")
		}

		logging.Infof(ctx, "using Docker client version %s", dockerClient.ClientVersion())
	}

	ar, err := atutil.RunAutoserv(ctx, j, t, os.Stdout, dockerClient, r.ContainerImageInfo)
	if err != nil {
		return nil, errors.Wrap(err, "run test")
	}
	return &autoservResult{
		CmdResult:  ar,
		ResultsDir: dir,
	}, nil
}
