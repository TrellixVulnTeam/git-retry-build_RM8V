// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cmd

import (
	"fmt"

	"go.chromium.org/chromiumos/infra/proto/go/test_platform"

	"infra/cmd/cros_test_platform/internal/site"

	"github.com/golang/protobuf/proto"
	"github.com/maruel/subcommands"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/steps"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/errors"
)

// SchedulerTrafficSplit implements the `scheduler-traffic-split` subcommand.
var SchedulerTrafficSplit = &subcommands.Command{
	UsageLine: "scheduler-traffic-split -input_json /path/to/input.json -output_json /path/to/output.json",
	ShortDesc: "Determine traffic split between backend schedulers.",
	LongDesc: `Determine traffic split between backend schedulers, i.e. Autotest vs Skylab.

Step input and output is JSON encoded protobuf defined at
https://chromium.googlesource.com/chromiumos/infra/proto/+/master/src/test_platform/steps/scheduler_traffic_split.proto`,
	CommandRun: func() subcommands.CommandRun {
		c := &schedulerTrafficSplitRun{}
		c.authFlags.Register(&c.Flags, site.DefaultAuthOptions)
		c.Flags.StringVar(&c.inputPath, "input_json", "", "Path that contains JSON encoded test_platform.steps.SchedulerTrafficSplitRequests")
		c.Flags.StringVar(&c.outputPath, "output_json", "", "Path where JSON encoded test_platform.steps.SchedulerTrafficSplitResponses should be written.")
		c.Flags.BoolVar(&c.directAllToSkylab, "rip-cautotest", true, "Cautotest is now at peace. Use a simple forwarding rule to send all traffic to Skylab.")
		return c
	},
}

type schedulerTrafficSplitRun struct {
	subcommands.CommandRunBase
	authFlags authcli.Flags

	inputPath  string
	outputPath string

	// A fast-path flag that replaces the traffic splitter logic with a mostly
	// trivial redirection to Skylab.
	directAllToSkylab bool

	// TODO(crbug.com/1002941) Internally transition to tagged requests only.
	orderedTags []string
}

func (c *schedulerTrafficSplitRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	err := c.innerRun(a, args, env)
	if err != nil {
		fmt.Fprintf(a.GetErr(), "%s\n", err)
	}
	return exitCode(err)
}

func (c *schedulerTrafficSplitRun) innerRun(a subcommands.Application, args []string, env subcommands.Env) error {
	if err := c.processCLIArgs(args); err != nil {
		return err
	}
	ctx := cli.GetContext(a, c, env)
	ctx = setupLogging(ctx)

	requests, err := c.readRequests()
	if err != nil {
		return err
	}
	if len(requests) == 0 {
		return errors.Reason("zero requests").Err()
	}

	if !c.directAllToSkylab {
		return errors.Reason("traffic split via config is deprecated").Err()
	}
	return c.sendAllToSkylab(requests)
}

func (c *schedulerTrafficSplitRun) sendAllToSkylab(requests []*steps.SchedulerTrafficSplitRequest) error {
	resps := make([]*steps.SchedulerTrafficSplitResponse, len(requests))
	for i, r := range requests {
		resps[i] = c.sendToSkylab(r.Request)
	}
	return c.writeResponses(resps)
}

func (c *schedulerTrafficSplitRun) sendToSkylab(req *test_platform.Request) *steps.SchedulerTrafficSplitResponse {
	var dst test_platform.Request
	proto.Merge(&dst, req)
	setQuotaAccountForLegacyPools(&dst)
	return &steps.SchedulerTrafficSplitResponse{
		SkylabRequest: &dst,
	}
}

// TODO(crbug.com/1026367) Once CTP stops receiving requests targeted at these
// legacy pools, drop the traffic splitter step entirely.
// The main source of these legacy requests are release builders on old
// branches.
var quotaAccountsForLegacyPools = map[test_platform.Request_Params_Scheduling_ManagedPool]string{
	test_platform.Request_Params_Scheduling_MANAGED_POOL_BVT:        "legacypool-bvt",
	test_platform.Request_Params_Scheduling_MANAGED_POOL_CONTINUOUS: "pfq",
	test_platform.Request_Params_Scheduling_MANAGED_POOL_CQ:         "cq",
	test_platform.Request_Params_Scheduling_MANAGED_POOL_SUITES:     "legacypool-suites",
}

func setQuotaAccountForLegacyPools(req *test_platform.Request) {
	if qa, ok := quotaAccountsForLegacyPools[req.GetParams().GetScheduling().GetManagedPool()]; ok {
		req.Params.Scheduling.Pool = &test_platform.Request_Params_Scheduling_QuotaAccount{
			QuotaAccount: qa,
		}
	}
}

func (c *schedulerTrafficSplitRun) processCLIArgs(args []string) error {
	if len(args) > 0 {
		return errors.Reason("have %d positional args, want 0", len(args)).Err()
	}
	if c.inputPath == "" {
		return errors.Reason("-input_json not specified").Err()
	}
	if c.outputPath == "" {
		return errors.Reason("-output_json not specified").Err()
	}
	return nil
}

func (c *schedulerTrafficSplitRun) readRequests() ([]*steps.SchedulerTrafficSplitRequest, error) {
	var rs steps.SchedulerTrafficSplitRequests
	if err := readRequest(c.inputPath, &rs); err != nil {
		return nil, err
	}
	ts, reqs := c.unzipTaggedRequests(rs.TaggedRequests)
	c.orderedTags = ts
	return reqs, nil
}

func (c *schedulerTrafficSplitRun) unzipTaggedRequests(trs map[string]*steps.SchedulerTrafficSplitRequest) ([]string, []*steps.SchedulerTrafficSplitRequest) {
	var ts []string
	var rs []*steps.SchedulerTrafficSplitRequest
	for t, r := range trs {
		ts = append(ts, t)
		rs = append(rs, r)
	}
	return ts, rs
}

func (c *schedulerTrafficSplitRun) writeResponses(resps []*steps.SchedulerTrafficSplitResponse) error {
	return writeResponse(c.outputPath, &steps.SchedulerTrafficSplitResponses{
		TaggedResponses: c.zipTaggedResponses(c.orderedTags, resps),
	})
}

func (c *schedulerTrafficSplitRun) zipTaggedResponses(ts []string, rs []*steps.SchedulerTrafficSplitResponse) map[string]*steps.SchedulerTrafficSplitResponse {
	if len(ts) != len(rs) {
		panic(fmt.Sprintf("got %d responses for %d tags (%s)", len(rs), len(ts), ts))
	}
	m := make(map[string]*steps.SchedulerTrafficSplitResponse)
	for i := range ts {
		m[ts[i]] = rs[i]
	}
	return m
}
