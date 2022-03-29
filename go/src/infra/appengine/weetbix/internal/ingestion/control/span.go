// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package control

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/server/span"
	"google.golang.org/protobuf/proto"

	"infra/appengine/weetbix/internal/config"
	ctlpb "infra/appengine/weetbix/internal/ingestion/control/proto"
	spanutil "infra/appengine/weetbix/internal/span"
)

// JoinStatsHours is the number of previous hours
// ReadPresubmitRunJoinStatistics/ReadBuildJoinStatistics reads statistics for.
const JoinStatsHours = 36

// Entry is an ingestion control record, used to de-duplicate build ingestions
// and synchronise them with presubmit results (if required).
type Entry struct {
	// The identity of the build which is being ingested.
	// The scheme is: {buildbucket host name}/{build id}.
	BuildID string

	// Project is the LUCI Project the build belongs to. Used for
	// metrics monitoring build/presubmit join performance.
	BuildProject string

	// BuildResult is the result of the build bucket build, to be passed
	// to the result ingestion task. This is nil if the result is
	// not yet known.
	BuildResult *ctlpb.BuildResult

	// BuildJoinedTime is the Spanner commit time the build result was
	// populated. If join has not yet occurred, this is the zero time.
	BuildJoinedTime time.Time

	// IsPresubmit records whether the build is part of a presubmit run.
	// If true, ingestion should wait for the presubmit result to be
	// populated (in addition to the build result) before commencing
	// ingestion.
	IsPresubmit bool

	// PresubmitProject is the LUCI Project the presubmit run belongs to.
	// This may differ from the LUCI Project teh build belongs to. Used for
	// metrics monitoring build/presubmit join performance.
	PresubmitProject string

	// PresubmitResult is result of the presubmit run, to be passed to the
	// result ingestion task. This is nil if the result is
	// not yet known.
	PresubmitResult *ctlpb.PresubmitResult

	// PresubmitJoinedTime is the Spanner commit time the presubmit result was
	// populated. If join has not yet occurred, this is the zero time.
	PresubmitJoinedTime time.Time

	// LastUpdated is the Spanner commit time the row was last updated.
	LastUpdated time.Time
}

// Read reads ingestion control records for the specified build IDs.
// Exactly one *Entry is returned for each build ID. The result entry
// at index i corresponds to the buildIDs[i].
// If a record does not exist for the given build ID, an *Entry of
// nil is returned for that build ID.
func Read(ctx context.Context, buildIDs []string) ([]*Entry, error) {
	uniqueIDs := make(map[string]struct{})
	var keys []spanner.Key
	for _, buildID := range buildIDs {
		keys = append(keys, spanner.Key{buildID})
		if _, ok := uniqueIDs[buildID]; ok {
			return nil, fmt.Errorf("duplicate build ID %s", buildID)
		}
		uniqueIDs[buildID] = struct{}{}
	}
	cols := []string{
		"BuildID",
		"BuildProject",
		"BuildResult",
		"BuildJoinedTime",
		"IsPresubmit",
		"PresubmitProject",
		"PresubmitResult",
		"PresubmitJoinedTime",
		"LastUpdated",
	}
	entryByBuildID := make(map[string]*Entry)
	rows := span.Read(ctx, "Ingestions", spanner.KeySetFromKeys(keys...), cols)
	f := func(r *spanner.Row) error {
		var buildID string
		var buildProject, presubmitProject spanner.NullString
		var buildResultBytes []byte
		var isPresubmit spanner.NullBool
		var presubmitResultBytes []byte
		var buildJoinedTime, presubmitJoinedTime spanner.NullTime
		var lastUpdated time.Time

		err := r.Columns(
			&buildID,
			&buildProject,
			&buildResultBytes,
			&buildJoinedTime,
			&isPresubmit,
			&presubmitProject,
			&presubmitResultBytes,
			&presubmitJoinedTime,
			&lastUpdated)
		if err != nil {
			return errors.Annotate(err, "read Ingestions row").Err()
		}
		var buildResult *ctlpb.BuildResult
		if buildResultBytes != nil {
			buildResult = &ctlpb.BuildResult{}
			if err := proto.Unmarshal(buildResultBytes, buildResult); err != nil {
				return errors.Annotate(err, "unmarshal build result").Err()
			}
		}
		var presubmitResult *ctlpb.PresubmitResult
		if presubmitResultBytes != nil {
			presubmitResult = &ctlpb.PresubmitResult{}
			if err := proto.Unmarshal(presubmitResultBytes, presubmitResult); err != nil {
				return errors.Annotate(err, "unmarshal presubmit result").Err()
			}
		}

		entryByBuildID[buildID] = &Entry{
			BuildID:         buildID,
			BuildProject:    buildProject.StringVal,
			BuildResult:     buildResult,
			BuildJoinedTime: buildJoinedTime.Time,
			// IsPresubmit uses NULL to indicate false.
			IsPresubmit:         isPresubmit.Valid && isPresubmit.Bool,
			PresubmitProject:    presubmitProject.StringVal,
			PresubmitResult:     presubmitResult,
			PresubmitJoinedTime: presubmitJoinedTime.Time,
			LastUpdated:         lastUpdated,
		}
		return nil
	}

	if err := rows.Do(f); err != nil {
		return nil, err
	}

	var result []*Entry
	for _, buildID := range buildIDs {
		// If the entry does not exist, return nil for that build ID.
		entry := entryByBuildID[buildID]
		result = append(result, entry)
	}
	return result, nil
}

// SetBuildResult sets the build result on an ingestion record,
// creating it if necessary.
// This sets the BuildProject, BuildResult, BuildJoinedTime
// fields, as well as the basic identify fields (BuildId, IsPresubmit).
func SetBuildResult(ctx context.Context, e *Entry) error {
	if err := validateEntry(e); err != nil {
		return err
	}
	m := spanutil.InsertOrUpdateMap("Ingestions", map[string]interface{}{
		"BuildId":         e.BuildID,
		"IsPresubmit":     spanner.NullBool{Valid: e.IsPresubmit, Bool: e.IsPresubmit},
		"BuildProject":    spanner.NullString{Valid: e.BuildProject != "", StringVal: e.BuildProject},
		"BuildResult":     e.BuildResult,
		"BuildJoinedTime": spanner.CommitTimestamp,
		"LastUpdated":     spanner.CommitTimestamp,
	})
	span.BufferWrite(ctx, m)
	return nil
}

// SetPresubmitResult sets the build result on an ingestion record,
// creating it if necessary.
// This sets the PresubmitProject, PresubmitResult, PresubmitJoinedTime
// fields, as well as the basic identify fields (BuildId, IsPresubmit).
func SetPresubmitResult(ctx context.Context, e *Entry) error {
	if err := validateEntry(e); err != nil {
		return err
	}
	if !e.IsPresubmit {
		return errors.New("IsPresubmit must be true if calling SetPresumitResult()")
	}
	m := spanutil.InsertOrUpdateMap("Ingestions", map[string]interface{}{
		"BuildId":             e.BuildID,
		"IsPresubmit":         spanner.NullBool{Valid: true, Bool: true},
		"PresubmitProject":    spanner.NullString{Valid: e.PresubmitProject != "", StringVal: e.PresubmitProject},
		"PresubmitResult":     e.PresubmitResult,
		"PresubmitJoinedTime": spanner.CommitTimestamp,
		"LastUpdated":         spanner.CommitTimestamp,
	})
	span.BufferWrite(ctx, m)
	return nil
}

// JoinStatistics captures indicators of how well buildbucket build
// completions are being joined to presubmit run completions.
type JoinStatistics struct {
	// TotalByHour captures the number of presubmit builds in the ingestions
	// table eligible to be joined.
	//
	// Data is broken down by by hours since the presubmit build became
	// eligible for joining. Index 0 indicates the period
	// from ]-1 hour, now], index 1 indicates [-2 hour, -1 hour] and so on.
	TotalByHour []int64

	// JoinedByHour captures the number of presubmit builds in the ingestions
	// table eligible to be joined, which were successfully joined (have
	// both presubmit run and buildbucket build completion present).
	//
	// Data is broken down by by hours since the presubmit build became
	// eligible for joining. Index 0 indicates the period
	// from ]-1 hour, now], index 1 indicates [-2 hour, -1 hour] and so on.
	JoinedByHour []int64
}

// ReadPresubmitJoinStatistics measures the performance joining presubmit runs.
//
// The statistics returned uses presubmit builds with a buildbucket
// build result received as the denominator for measuring join performance.
// The performance joining to presubmit run results is then measured.
// Data is broken down by the project of the buildbucket build.
// The last 36 hours of data for each project is returned. Hours are
// measured since the buildbucket build result was received.
func ReadPresubmitRunJoinStatistics(ctx context.Context) (map[string]JoinStatistics, error) {
	stmt := spanner.NewStatement(`
		SELECT
		  BuildProject as project,
		  TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), BuildJoinedTime, HOUR) as hour,
		  COUNT(*) as total,
		  COUNTIF(HasPresubmitResult) as joined,
		FROM Ingestions@{FORCE_INDEX=IngestionsByIsPresubmit, spanner_emulator.disable_query_null_filtered_index_check=true}
		WHERE IsPresubmit
		  AND BuildJoinedTime >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL @hours HOUR)
		GROUP BY project, hour
	`)
	stmt.Params["hours"] = JoinStatsHours
	return readJoinStatistics(ctx, stmt)
}

// ReadPresubmitJoinStatistics reads indicators of how well buildbucket build
// completions are being joined to presubmit run completions.
//
// The statistics returned uses builds with a presubmit run
// received as the denominator for measuring join performance.
// The performance joining to buildbucket build results is then measured.
// Data is broken down by the project of the presubmit run.
// The last 36 hours of data for each project is returned. Hours are
// measured since the presubmit run result was received.
func ReadBuildJoinStatistics(ctx context.Context) (map[string]JoinStatistics, error) {
	stmt := spanner.NewStatement(`
		SELECT
		  PresubmitProject as project,
		  TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), PresubmitJoinedTime, HOUR) as hour,
		  COUNT(*) as total,
		  COUNTIF(HasBuildResult) as joined,
		FROM Ingestions@{FORCE_INDEX=IngestionsByIsPresubmit, spanner_emulator.disable_query_null_filtered_index_check=true}
		WHERE IsPresubmit
		  AND PresubmitJoinedTime >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL @hours HOUR)
		GROUP BY project, hour
	`)
	stmt.Params["hours"] = JoinStatsHours
	return readJoinStatistics(ctx, stmt)
}

func readJoinStatistics(ctx context.Context, stmt spanner.Statement) (map[string]JoinStatistics, error) {
	result := make(map[string]JoinStatistics)
	it := span.Query(ctx, stmt)
	err := it.Do(func(r *spanner.Row) error {
		var project string
		var hour int64
		var total, joined int64

		err := r.Columns(&project, &hour, &total, &joined)
		if err != nil {
			return errors.Annotate(err, "read row").Err()
		}

		stats, ok := result[project]
		if !ok {
			stats = JoinStatistics{
				// Add zero data for all hours.
				TotalByHour:  make([]int64, JoinStatsHours),
				JoinedByHour: make([]int64, JoinStatsHours),
			}
		}
		stats.TotalByHour[hour] = total
		stats.JoinedByHour[hour] = joined

		result[project] = stats
		return nil
	})
	if err != nil {
		return nil, errors.Annotate(err, "query presubmit join stats by project").Err()
	}
	return result, nil
}

func validateEntry(e *Entry) error {
	if e.BuildID == "" {
		return errors.New("build ID must be specified")
	}
	if e.BuildResult != nil {
		if err := validateBuildResult(e.BuildResult); err != nil {
			return errors.Annotate(err, "build result").Err()
		}
		if !config.ProjectRe.MatchString(e.BuildProject) {
			return errors.New("build project must be valid")
		}
	} else {
		if e.BuildProject != "" {
			return errors.New("build project must only be specified" +
				" if build result is specified")
		}
	}

	if e.PresubmitResult != nil {
		if !e.IsPresubmit {
			return errors.New("presubmit result must not be set unless IsPresubmit is set")
		}
		if err := validatePresubmitResult(e.PresubmitResult); err != nil {
			return errors.Annotate(err, "presubmit result").Err()
		}
		if !config.ProjectRe.MatchString(e.PresubmitProject) {
			return errors.New("presubmit project must be valid")
		}
	} else {
		if e.PresubmitProject != "" {
			return errors.New("presubmit project must only be specified" +
				" if presubmit result is specified")
		}
	}
	return nil
}

func validateBuildResult(r *ctlpb.BuildResult) error {
	switch {
	case r.Host == "":
		return errors.New("host must be specified")
	case r.Id == 0:
		return errors.New("id must be specified")
	case !r.CreationTime.IsValid():
		return errors.New("creation time must be specified")
	}
	return nil
}

func validatePresubmitResult(r *ctlpb.PresubmitResult) error {
	switch {
	case r.PresubmitRunId == nil:
		return errors.New("presubmit run ID must be specified")
	case r.PresubmitRunId.System != "luci-cv":
		// LUCI CV is currently the only supported system.
		return errors.New("presubmit run system must be 'luci-cv'")
	case r.PresubmitRunId.Id == "":
		return errors.New("presubmit run system-specific ID must be specified")
	case !r.CreationTime.IsValid():
		return errors.New("creation time must be specified and valid")
	}
	return nil
}
