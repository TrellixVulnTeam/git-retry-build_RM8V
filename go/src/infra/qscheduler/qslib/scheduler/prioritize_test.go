// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// TODO(akeshet): The tests in this file make use of a lot of unexported
// methods and fields. It would be better if they were rewritten to use
// only the exported API of the scheduler. That would also entail building
// an exported API for getting job prioritization.

package scheduler

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"

	. "github.com/smartystreets/goconvey/convey"
)

// TestPrioritizeOne tests that PrioritizeRequests behaves correctly
// for a single request.
func TestPrioritizeOne(t *testing.T) {
	aid := AccountID("a1")
	rid := RequestID("r1")

	Convey("Given a scheduler, with a request for account", t, func() {
		ctx := context.Background()
		tm := time.Unix(0, 0)
		s := New(tm)
		accountCases := []struct {
			accountBalance   balance
			expectedPriority int
		}{
			{
				balance{1, 0, 0},
				0,
			},
			{
				balance{0, 1, 0},
				1,
			},
			{
				balance{},
				FreeBucket,
			},
		}

		s.AddRequest(ctx, rid, NewRequest(aid, nil, tm), tm)
		for _, c := range accountCases {
			Convey(fmt.Sprintf("given account balance is %v", c.accountBalance), func() {
				s.AddAccount(ctx, aid, &AccountConfig{}, c.accountBalance[:])
				Convey("when prioritizing", func() {
					got := s.prioritizeRequests()
					Convey(fmt.Sprintf("then the request gets priority %d.", c.expectedPriority), func() {
						So(got[0].Priority, ShouldEqual, c.expectedPriority)
						So(got[0].RequestID, ShouldEqual, rid)
					})
				})
			})
		}

		Convey("when no such account exists", func() {
			Convey("when prioritizing", func() {
				got := s.prioritizeRequests()
				Convey("then the request is put in free bucket priority.", func() {
					So(got[0].Priority, ShouldEqual, FreeBucket)
					So(got[0].RequestID, ShouldEqual, rid)
				})
			})
		})
	})
}

// TestPrioritizeMany tests that PrioritizeRequests behaves correctly
// for a number of requests.
func TestPrioritizeMany(t *testing.T) {
	nReqs := 10
	aid := AccountID("a1")
	Convey("Given requests with different enqueue times, but inserted in random order", t, func() {
		ctx := context.Background()
		s := New(time.Unix(0, 0))
		// Use a fixed seed, so the test is reproducible and the request order is
		// pseudo-random.
		rand.Seed(10)
		perm := rand.Perm(nReqs)
		for _, i := range perm {
			tm := time.Unix(int64(i), 0)
			s.AddRequest(ctx, RequestID(strconv.Itoa(i)), NewRequest(aid, nil, tm), tm)
		}

		Convey("given no matching account", func() {
			Convey("when prioritizing", func() {
				got := s.prioritizeRequests()
				Convey("then requests are prioritized by enqueue time.", func() {
					times := make([]time.Time, nReqs)
					for i, g := range got {
						times[i] = g.Request.enqueueTime
					}
					So(times, ShouldBeChronological)
					So(got, ShouldHaveLength, nReqs)
				})
			})
		})

		Convey("given an account with no maximum fanout", func() {
			s.AddAccount(ctx, aid, NewAccountConfig(0, 0, nil), nil)
			Convey("when prioritizing", func() {
				got := s.prioritizeRequests()
				Convey("then requests are prioritized by enqueue time.", func() {
					times := make([]time.Time, nReqs)
					for i, g := range got {
						times[i] = g.Request.enqueueTime
					}
					So(times, ShouldBeChronological)
					So(got, ShouldHaveLength, nReqs)
				})
			})
		})

		Convey("given the account specifices a maximum fanout and some requests for that account are already running", func() {
			maxFanout := 5
			s.AddAccount(ctx, aid, &AccountConfig{MaxFanout: int32(maxFanout)}, []float64{0, 1})

			// Two requests are already running.
			addRunningRequest(ctx, s, "11", "11", aid, 1, time.Unix(0, 0))
			addRunningRequest(ctx, s, "12", "12", aid, 1, time.Unix(0, 0))

			Convey("when prioritizing", func() {
				got := s.prioritizeRequests()
				Convey("then requests beyond the maximum fanout are put in the free bucket priority.", func() {
					// First three get the account's available bucket (P1), the remaining
					// ones get the free bucket.
					fb := FreeBucket
					expectedPriorities := []int{1, 1, 1, fb, fb, fb, fb, fb, fb, fb}

					actualPriorities := make([]int, nReqs)
					for i := 0; i < nReqs; i++ {
						actualPriorities[i] = got[i].Priority
					}
					So(actualPriorities, ShouldResemble, expectedPriorities)
					So(got, ShouldHaveLength, nReqs)
				})
			})
		})
	})
}

// addRunningRequest is a test helper to add a new request to a scheduler and
// immediately start it running on a new worker.
func addRunningRequest(ctx context.Context, s *Scheduler, rid RequestID, wid WorkerID, aid AccountID, pri int, tm time.Time) {
	s.AddRequest(ctx, rid, NewRequest(aid, []string{}, tm), tm)
	s.MarkIdle(ctx, wid, []string{}, tm)
	s.state.applyAssignment(&Assignment{Priority: pri, RequestID: rid, WorkerID: wid, Type: AssignmentIdleWorker})
}

// TestForPriority tests that ForPriority method returns the correct
// sub-slices of a prioritized list.
func TestForPriority(t *testing.T) {
	t.Parallel()
	pRequests := orderedRequests([]prioritizedRequest{
		{Priority: 0},
		{Priority: 0},
		{Priority: 1},
		{Priority: 3},
		{Priority: 3},
		{Priority: 4},
	})

	expecteds := []orderedRequests{
		pRequests[0:2],
		pRequests[2:3],
		pRequests[3:3],
		pRequests[3:5],
		pRequests[5:6],
		pRequests[6:6],
	}

	for p := 0; p < 6; p++ {
		actual := pRequests.forPriority(p)
		expected := expecteds[p]
		if diff := pretty.Compare(actual, expected); diff != "" {
			t.Errorf(fmt.Sprintf("P%d slice got unexpected diff (-got +want): %s", p, diff))
		}
	}
}

// getWorkers is a helper function to turn a slice of running tasks
// into a workers map.
func getWorkers(running []*TaskRun) map[string]*Worker {
	workers := make(map[string]*Worker)
	for i, r := range running {
		wid := fmt.Sprintf("w%d", i)
		workers[wid] = &Worker{Labels: []string{}, RunningTask: r}
	}
	return workers
}
