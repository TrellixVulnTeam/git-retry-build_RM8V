// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package engine provides struts and functionality of recovery engine.
// For more details please read go/paris-recovery-engine.
package engine

import (
	"context"
	"fmt"
	"time"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/luciexe/build"

	"infra/cros/recovery/config"
	"infra/cros/recovery/internal/execs"
	"infra/cros/recovery/internal/log"
	"infra/cros/recovery/logger/metrics"
)

// recoveryEngine holds info required for running a recovery plan.
type recoveryEngine struct {
	planName string
	plan     *config.Plan
	args     *execs.RunArgs
	// Caches
	actionResultsCache map[string]error
	recoveryUsageCache map[recoveryUsageKey]error
}

// Error tag to track error with request to start critical actions over.
var startOverTag = errors.BoolTag{Key: errors.NewTagKey("start-over")}

// Run runs the recovery plan.
func Run(ctx context.Context, planName string, plan *config.Plan, args *execs.RunArgs) error {
	r := &recoveryEngine{
		planName: planName,
		plan:     plan,
		args:     args,
	}
	r.initCache()
	defer func() { r.close() }()
	log.Debugf(ctx, "Received plan %s for %s \n%s", r.planName, r.args.ResourceName, r.describe())
	return r.runPlan(ctx)
}

// close free up used resources.
func (r *recoveryEngine) close() {
	if r.actionResultsCache != nil {
		r.actionResultsCache = nil
	}
	// TODO(otabek@): Close the caches.
}

// runPlan executes recovery plan with critical-actions.
func (r *recoveryEngine) runPlan(ctx context.Context) (rErr error) {
	log.Infof(ctx, "Plan %q: started", r.planName)

	var restartTally int64
	var forgivenFailureTally int64
	if r.args != nil && r.args.Metrics != nil {
		var mErr error
		var closer execs.CloserFunc
		action := &metrics.Action{}
		closer, mErr = r.args.NewMetric(
			ctx,
			fmt.Sprintf("plan:%s", r.planName),
			action,
		)
		if mErr == nil {
			defer func() {
				if action != nil {
					action.Observations = append(
						action.Observations,
						metrics.NewInt64Observation("restarts", restartTally),
						metrics.NewInt64Observation("forgiven_failures", forgivenFailureTally),
					)
				}
				closer(ctx, rErr)
			}()
		}
	}

	for {
		if err := r.runCriticalActionAttempt(ctx, restartTally); err != nil {
			if startOverTag.In(err) {
				log.Infof(ctx, "Plan %q for %s: received request to start over.", r.planName, r.args.ResourceName)
				r.resetCacheAfterSuccessfulRecoveryAction()
				restartTally++
				continue
			}
			if r.plan.GetAllowFail() {
				log.Infof(ctx, "Plan %q for %s: failed with error: %s.", r.planName, r.args.ResourceName, err)
				log.Infof(ctx, "Plan %q for %s: is allowed to fail, continue.", r.planName, r.args.ResourceName)
				forgivenFailureTally++
				return nil
			}
			return errors.Annotate(err, "run plan %q", r.planName).Err()
		}
		break
	}
	log.Infof(ctx, "Plan %q: finished successfully.", r.planName)
	log.Infof(ctx, "Plan %q: recorded %d restarts during execution.", r.planName, restartTally)
	log.Infof(ctx, "Plan %q: recorded %d forgiven failures during execution.", r.planName, forgivenFailureTally)
	return nil
}

// runCriticalActionAttempt runs critical action of the plan with wrapper step to show plan restart attempts.
func (r *recoveryEngine) runCriticalActionAttempt(ctx context.Context, attempt int64) (err error) {
	if r.args.ShowSteps {
		var step *build.Step
		stepName := fmt.Sprintf("First run of critical actions for %s", r.planName)
		if attempt > 0 {
			stepName = fmt.Sprintf("Attempt %d to run critical actions for %s", attempt, r.planName)
		}
		step, ctx = build.StartStep(ctx, stepName)
		defer func() { step.End(err) }()
	}
	return r.runActions(ctx, r.plan.GetCriticalActions(), r.args.EnableRecovery)
}

// runActions runs actions in order.
func (r *recoveryEngine) runActions(ctx context.Context, actions []string, enableRecovery bool) error {
	for _, actionName := range actions {
		if err := r.runAction(ctx, actionName, enableRecovery); err != nil {
			return errors.Annotate(err, "run actions").Err()
		}
	}
	return nil
}

// recordActionCloser is a function that takes an error (the ultimate error produced by an action) and records
// it inside a defer block.
type recordActionCloser = func(error)

// recordAction takes a context and an action name and records the initial action for a record.
// The parameter action is assumed NOT to be nil. Also, this function indirectly mutates its parameter action.
func (r *recoveryEngine) recordAction(ctx context.Context, actionName string, action *metrics.Action) recordActionCloser {
	if r == nil {
		log.Debugf(ctx, "RecoveryEngine is nil, skipping")
		return nil
	}
	if r.args == nil {
		log.Debugf(ctx, "Metrics is nil, skipping")
		return nil
	}
	if r.args.Metrics != nil {
		log.Debugf(ctx, "Recording metrics for action %q", actionName)
		// Create the metric up front. Allow 30 seconds to talk to Karte.
		createMetricCtx, createMetricCloser := context.WithTimeout(ctx, 30*time.Second)
		defer createMetricCloser()
		u, err := r.args.NewMetric(
			createMetricCtx,
			// TODO(gregorynisbet): Consider adding a new field to Karte to explicitly track the name
			//                      assigned to an action by recoverylib.
			fmt.Sprintf("action:%s", actionName),
			action,
		)
		if err != nil {
			log.Errorf(ctx, "Encountered error when creating action: %s", err)
			return nil
		}
		// Here we intentionally close over the context "early", before the deadline is applied inside
		// runAction.
		return func(rErr error) {
			// Update the metric. This contains information that we will not know until after the action ran.
			updateMetricCtx, updateMetricCloser := context.WithTimeout(ctx, 30*time.Second)
			defer updateMetricCloser()
			u(updateMetricCtx, rErr)
		}
	} else {
		log.Debugf(ctx, "Skipping metrics for action %q", actionName)
		return nil
	}
}

// runAction runs single action.
// Execution steps:
// 1) Check action's result in cache.
// 2) Check if the action is applicable based on conditions. Skip if any fail.
// 3) Run dependencies of the action. Fail if any fails.
// 4) Run action exec function. Fail if any fail.
func (r *recoveryEngine) runAction(ctx context.Context, actionName string, enableRecovery bool) (rErr error) {
	action := &metrics.Action{}
	if r.args != nil {
		if actionCloser := r.recordAction(ctx, actionName, action); actionCloser != nil {
			defer actionCloser(rErr)
		}
		if r.args.ShowSteps {
			var step *build.Step
			step, ctx = build.StartStep(ctx, fmt.Sprintf("Run %s", actionName))
			defer func() { step.End(rErr) }()
		}
		if r.args.Logger != nil {
			r.args.Logger.IndentLogging()
			defer func() { r.args.Logger.DedentLogging() }()
		}
	}
	log.Infof(ctx, "Action %q: started.", actionName)
	defer func() {
		if rErr != nil {
			log.Debugf(ctx, "Action %q: finished with error %s.", actionName, rErr)
		} else {
			log.Debugf(ctx, "Action %q: finished.", actionName)
		}
	}()
	a := r.getAction(actionName)
	if aErr, ok := r.actionResultFromCache(actionName); ok {
		if aErr == nil {
			log.Infof(ctx, "Action %q: pass (cached).", actionName)
			// Return nil error so we can continue execution of next actions...
			return nil
		}
		if a.GetAllowFailAfterRecovery() {
			log.Infof(ctx, "Action %q: fail (cached). Error: %s", actionName, aErr)
			log.Debugf(ctx, "Action %q: error ignored as action is allowed to fail.", actionName)
			// Return nil error so we can continue execution of next actions...
			return nil
		}
		return errors.Annotate(aErr, "run action %q: (cached)", actionName).Err()
	}
	conditionName, err := r.runActionConditions(ctx, actionName)
	if err != nil {
		log.Infof(ctx, "Action %q: one of conditions %q failed, skipping...", actionName, conditionName)
		log.Debugf(ctx, "Action %q: conditions fail with %s", actionName, err)
		// Return nil error so we can continue execution of next actions...
		return nil
	}
	if err := r.runDependencies(ctx, actionName, enableRecovery); err != nil {
		if startOverTag.In(err) {
			return errors.Annotate(err, "run action %q", actionName).Err()
		}
		if a.GetAllowFailAfterRecovery() {
			log.Infof(ctx, "Action %q: one of dependencies fail. Error: %s", actionName, err)
			log.Debugf(ctx, "Action %q: error ignored as action is allowed to fail.", actionName)
			return nil
		} else {
			return errors.Annotate(err, "run action %q", actionName).Err()
		}
	}
	if err := r.runActionExec(ctx, actionName, enableRecovery); err != nil {
		if startOverTag.In(err) {
			return errors.Annotate(err, "run action %q", actionName).Err()
		}
		if a.GetAllowFailAfterRecovery() {
			log.Infof(ctx, "Action %q: fail. Error: %s", actionName, err)
			log.Debugf(ctx, "Action %q: error ignored as action is allowed to fail.", actionName)
		} else {
			return errors.Annotate(err, "run action %q", actionName).Err()
		}
	} else {
		log.Infof(ctx, "Action %q: finished successfully.", actionName)
	}
	// Return nil error so we can continue execution of next actions...
	return nil
}

// runActionExec runs action's exec function and initiates recovery flow if exec fails.
// The recover flow start only recoveries is enabled.
func (r *recoveryEngine) runActionExec(ctx context.Context, actionName string, enableRecovery bool) error {
	a := r.getAction(actionName)
	if err := r.runActionExecWithTimeout(ctx, a); err != nil {
		if enableRecovery && len(a.GetRecoveryActions()) > 0 {
			log.Infof(ctx, "Action %q: starting recovery actions.", actionName)
			log.Debugf(ctx, "Action %q: fail. Error: %s", actionName, err)
			if rErr := r.runRecoveries(ctx, actionName); rErr != nil {
				return errors.Annotate(rErr, "run action %q exec", actionName).Err()
			}
			log.Infof(ctx, "Run action %q exec: no recoveries left to try", actionName)
		}
		// Cache the action error only after running recoveries.
		// If no recoveries were run, we still cache the action.
		r.cacheActionResult(actionName, err)
		return errors.Annotate(err, "run action %q exec", actionName).Err()
	}
	r.cacheActionResult(actionName, nil)
	return nil
}

// Default time limit per action exec function.
const defaultExecTimeout = 60 * time.Second

func actionExecTimeout(a *config.Action) time.Duration {
	if a.ExecTimeout != nil {
		return a.ExecTimeout.AsDuration()
	}
	return defaultExecTimeout
}

// runActionExecWithTimeout runs action's exec function with timeout.
func (r *recoveryEngine) runActionExecWithTimeout(ctx context.Context, a *config.Action) error {
	timeout := actionExecTimeout(a)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer func() { cancel() }()
	cw := make(chan error, 1)
	go func() {
		err := execs.Run(ctx, &execs.ExecInfo{
			RunArgs:       r.args,
			Name:          a.ExecName,
			ActionArgs:    a.GetExecExtraArgs(),
			ActionTimeout: timeout,
		})
		cw <- err
	}()
	select {
	case err := <-cw:
		return errors.Annotate(err, "run exec %q with timeout %s", a.ExecName, timeout).Err()
	case <-ctx.Done():
		log.Infof(ctx, "Run exec %q with timeout %s: excited timeout", a.ExecName, timeout)
		return errors.Reason("run exec %q with timeout %s: excited timeout", a.ExecName, timeout).Err()
	}
}

// runActionConditions checks if action is applicable based on condition actions.
// If return err then not applicable, if nil then applicable.
func (r *recoveryEngine) runActionConditions(ctx context.Context, actionName string) (conditionName string, err error) {
	a := r.getAction(actionName)
	if len(a.GetConditions()) == 0 {
		return "", nil
	}
	if r.args != nil {
		if r.args.ShowSteps {
			var step *build.Step
			step, ctx = build.StartStep(ctx, "Run continions")
			defer func() { step.End(err) }()
		}
		if r.args.Logger != nil {
			r.args.Logger.IndentLogging()
			defer func() { r.args.Logger.DedentLogging() }()
		}
	}
	log.Debugf(ctx, "Action %q: running conditions...", actionName)
	enableRecovery := false
	for _, condition := range a.GetConditions() {
		if err := r.runAction(ctx, condition, enableRecovery); err != nil {
			log.Debugf(ctx, "Action %q: condition %q fails. Error: %s", actionName, condition, err)
			return condition, errors.Annotate(err, "run conditions").Err()
		}
	}
	log.Debugf(ctx, "Action %q: all conditions passed.", actionName)
	return "", nil
}

// runDependencies runs action's dependencies.
func (r *recoveryEngine) runDependencies(ctx context.Context, actionName string, enableRecovery bool) (rErr error) {
	a := r.getAction(actionName)
	if len(a.GetDependencies()) == 0 {
		return nil
	}
	if r.args != nil {
		if r.args.ShowSteps {
			var step *build.Step
			step, ctx = build.StartStep(ctx, "Run dependencies")
			defer func() { step.End(rErr) }()
		}
		if r.args.Logger != nil {
			r.args.Logger.IndentLogging()
			defer func() { r.args.Logger.DedentLogging() }()
		}
	}
	err := r.runActions(ctx, a.GetDependencies(), enableRecovery)
	return errors.Annotate(err, "run dependencies").Err()
}

// runRecoveries runs action's recoveries.
// Recovery actions are expected to fail. If recovery action fails then next will be attempted.
// Finishes with nil if no recovery action provided or nether succeeded.
// Finishes with start-over request if any recovery succeeded.
// Recovery action will skip if used before.
func (r *recoveryEngine) runRecoveries(ctx context.Context, actionName string) (rErr error) {
	a := r.getAction(actionName)
	if len(a.GetRecoveryActions()) == 0 {
		return nil
	}
	if r.args != nil {
		if r.args.ShowSteps {
			var step *build.Step
			step, ctx = build.StartStep(ctx, "Run recoveries")
			defer func() { step.End(rErr) }()
		}
		if r.args.Logger != nil {
			r.args.Logger.IndentLogging()
			defer func() { r.args.Logger.DedentLogging() }()
		}
	}
	for _, recoveryName := range a.GetRecoveryActions() {
		if r.isRecoveryUsed(actionName, recoveryName) {
			// Engine allows to use each recovery action only once in scope of the action.
			continue
		}
		if err := r.runActions(ctx, []string{recoveryName}, false); err != nil {
			log.Debugf(ctx, "Recovery action %q: fail. Error: %s ", recoveryName, err)
			r.registerRecoveryUsage(actionName, recoveryName, err)
			continue
		}
		r.registerRecoveryUsage(actionName, recoveryName, nil)
		log.Infof(ctx, "Recovery action %q: request to start-over.", recoveryName)
		return errors.Reason("recovery action %q: request to start over", recoveryName).Tag(startOverTag).Err()
	}
	return nil
}

// getAction finds and provides action from the plan collection.
func (r *recoveryEngine) getAction(name string) *config.Action {
	if a, ok := r.plan.Actions[name]; ok {
		return a
	}
	// If we reach this place then we have issues with plan validation logic.
	panic(fmt.Sprintf("action %q not found in the plan", name))
}

// describe describes the plan details with critical actions.
func (r *recoveryEngine) describe() string {
	d := fmt.Sprintf("Plan %q, AllowFail: %v ", r.planName, r.plan.AllowFail)
	if len(r.plan.GetCriticalActions()) > 0 {
		prefix := "\n "
		d += fmt.Sprintf("%sCritical-actions:", prefix)
		for i, a := range r.plan.GetCriticalActions() {
			d += fmt.Sprintf("%s %d: %s", prefix, i, r.describeAction(a, prefix+"  "))
		}
	} else {
		d += "\n No critical-actions"
	}
	return d
}

// describeAction describes the action structure recursively.
func (r *recoveryEngine) describeAction(actionName string, prefix string) string {
	a := r.getAction(actionName)
	ap := fmt.Sprintf("Action %q, AllowFailAfterRecovery: %v, RunControl: %v",
		actionName, a.GetAllowFailAfterRecovery(), a.GetRunControl())
	if len(a.GetConditions()) > 0 {
		ap += fmt.Sprintf("%sConditions:", prefix)
		for i, d := range a.GetConditions() {
			ap += fmt.Sprintf("%s%d: %s", prefix, i, r.describeAction(d, prefix+"  "))
		}
	}
	ap += fmt.Sprintf("%sExec: %s", prefix, r.describeActionExec(actionName))
	if len(a.GetDependencies()) > 0 {
		ap += fmt.Sprintf("%sDependencies:", prefix)
		for i, d := range a.GetDependencies() {
			ap += fmt.Sprintf("%s%d: %s", prefix, i, r.describeAction(d, prefix+"  "))
		}
	}
	if len(a.GetRecoveryActions()) > 0 {
		ap += fmt.Sprintf("%sRecoveryActions:", prefix)
		for i, d := range a.GetRecoveryActions() {
			ap += fmt.Sprintf("%s%d: %s", prefix, i, r.describeAction(d, prefix+"  "))
		}
	}
	return ap
}

// describeActionExec describes the action's exec function with details.
func (r *recoveryEngine) describeActionExec(actionName string) string {
	a := r.getAction(actionName)
	er := a.GetExecName()
	if len(a.GetExecExtraArgs()) > 0 {
		er += fmt.Sprintf(", Args: %s", a.GetExecExtraArgs())
	}
	return er
}

// initCache initializes cache on engine.
// The function extracted to supported testing.
func (r *recoveryEngine) initCache() {
	r.actionResultsCache = make(map[string]error, len(r.plan.GetActions()))
	r.recoveryUsageCache = make(map[recoveryUsageKey]error)
}

// actionResultFromCache reads action's result from cache.
func (r *recoveryEngine) actionResultFromCache(actionName string) (err error, ok bool) {
	err, ok = r.actionResultsCache[actionName]
	return err, ok
}

// cacheActionResult sets action's result to the cache.
func (r *recoveryEngine) cacheActionResult(actionName string, err error) {
	switch r.getAction(actionName).GetRunControl() {
	case config.RunControl_RERUN_AFTER_RECOVERY, config.RunControl_RUN_ONCE:
		r.actionResultsCache[actionName] = err
	case config.RunControl_ALWAYS_RUN:
		// Do not cache the value
	}
}

// resetCacheAfterSuccessfulRecoveryAction resets cache for actions
// with run-control=RERUN_AFTER_RECOVERY.
func (r *recoveryEngine) resetCacheAfterSuccessfulRecoveryAction() {
	for name, a := range r.plan.GetActions() {
		if a.GetRunControl() == config.RunControl_RERUN_AFTER_RECOVERY {
			delete(r.actionResultsCache, name)
		}
	}
}

// isRecoveryUsed checks if recovery action is used in plan or action level scope.
func (r *recoveryEngine) isRecoveryUsed(actionName, recoveryName string) bool {
	k := recoveryUsageKey{
		action:   actionName,
		recovery: recoveryName,
	}
	// If the recovery has been used in previous actions then it can be in
	// the action result cache.
	if err, ok := r.actionResultsCache[recoveryName]; ok {
		r.recoveryUsageCache[k] = err
	}
	_, ok := r.recoveryUsageCache[k]
	return ok
}

// registerRecoveryUsage sets recovery action usage to the cache.
func (r *recoveryEngine) registerRecoveryUsage(actionName, recoveryName string, err error) {
	r.recoveryUsageCache[recoveryUsageKey{
		action:   actionName,
		recovery: recoveryName,
	}] = err
}

// recoveryUsageKey holds action and action's recovery name as key for recovery-usage cache.
type recoveryUsageKey struct {
	action   string
	recovery string
}
