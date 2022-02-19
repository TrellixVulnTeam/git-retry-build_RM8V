// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package inventory

import (
	"context"
	"strings"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/grpcutil"

	fleet "infra/appengine/crosskylabadmin/api/fleet/v1"
	"infra/appengine/crosskylabadmin/internal/app/frontend/internal/datastore/stableversion/satlab"
	"infra/cros/stableversion"
)

// SetSatlabStableVersion replaces a satlab stable version with a new entry.
func (is *ServerImpl) SetSatlabStableVersion(ctx context.Context, req *fleet.SetSatlabStableVersionRequest) (_ *fleet.SetSatlabStableVersionResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	if err := validateSetSatlabStableVersion(req); err != nil {
		return nil, errors.Annotate(err, "set satlab stable version").Err()
	}

	newEntry, err := satlab.MakeSatlabStableVersionEntry(req, true)
	if err != nil {
		return nil, errors.Annotate(err, "set satlab stable version").Err()
	}

	if pErr := satlab.PutSatlabStableVersionEntry(ctx, newEntry); pErr != nil {
		return nil, errors.Annotate(pErr, "set satlab stable version").Err()
	}
	return &fleet.SetSatlabStableVersionResponse{}, nil
}

// DeleteSatlabStableVersion deletes a satlab stable version entry.
func (is *ServerImpl) DeleteSatlabStableVersion(ctx context.Context, req *fleet.DeleteSatlabStableVersionRequest) (_ *fleet.DeleteSatlabStableVersionResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	var hostname string
	var board string
	var model string
	switch v := req.GetStrategy().(type) {
	case *fleet.DeleteSatlabStableVersionRequest_SatlabHostnameDeletionCriterion:
		hostname = v.SatlabHostnameDeletionCriterion.GetHostname()
	case *fleet.DeleteSatlabStableVersionRequest_SatlabBoardModelDeletionCriterion:
		model = v.SatlabBoardModelDeletionCriterion.GetModel()
		board = v.SatlabBoardModelDeletionCriterion.GetBoard()
	}

	id := satlab.MakeSatlabStableVersionID(hostname, board, model)
	if id == "" {
		return nil, errors.Reason("delete satlab stable version: failed to produce identifier").Err()
	}

	if err := satlab.DeleteSatlabStableVersionEntryByRawID(ctx, id); err != nil {
		return nil, errors.Annotate(err, "delete satlab stable version").Err()
	}
	return &fleet.DeleteSatlabStableVersionResponse{}, nil
}

// ValidateSetSatlabStableVersion validates a set satlab stable version request.
func validateSetSatlabStableVersion(req *fleet.SetSatlabStableVersionRequest) error {
	if req == nil {
		return errors.Reason("validate set satlab stable version: request cannot be nil").Err()
	}
	var hostname string
	var board string
	var model string
	switch v := req.GetStrategy().(type) {
	case *fleet.SetSatlabStableVersionRequest_SatlabBoardAndModelStrategy:
		s := v.SatlabBoardAndModelStrategy
		board = strings.TrimSpace(strings.ToLower(s.GetBoard()))
		model = strings.TrimSpace(strings.ToLower(s.GetModel()))
	case *fleet.SetSatlabStableVersionRequest_SatlabHostnameStrategy:
		hostname = strings.TrimSpace(strings.ToLower(v.SatlabHostnameStrategy.GetHostname()))
	}

	if err := shallowValidateKeyFields(hostname, board, model); err != nil {
		return errors.Annotate(err, "validate set satlab stable version").Err()
	}

	osVersion := req.GetCrosVersion()
	fwVersion := req.GetFirmwareVersion()
	fwImage := req.GetFirmwareImage()

	if err := shallowValidateValueFields(osVersion, fwVersion, fwImage); err != nil {
		return errors.Annotate(err, "validate set satlab stable version").Err()
	}

	return nil
}

// ShallowValidateKeyFields validates the key fields of a satlab stable version request. These are the fields
// that are used to look up a record.
//
// This is a shallow validation because it does not consult any sources of truth to see if the information is valid.
func shallowValidateKeyFields(hostname string, board string, model string) error {
	if hostname != "" {
		if board == "" && model == "" {
			return nil
		}
		return errors.Reason("shallow validate key fields: cannot use both hostname %q and board/model %q/%q", hostname, board, model).Err()
	}
	if board != "" && model != "" {
		return nil
	}
	return errors.Reason("shallow validate key fields: expected board %q and model %q to both be non-empty", board, model).Err()
}

// ShallowValidateValueFields validates the value fields, which correspond to fragments of gs:// URLs.
//
// This is a shallow validation because it does not consult any sources of truth to see if the information is valid.
func shallowValidateValueFields(os string, fw string, fwImage string) error {
	if err := stableversion.ValidateCrOSVersion(os); err != nil {
		return errors.Annotate(err, "shallow validate value fields").Err()
	}
	if err := stableversion.ValidateFirmwareVersion(fw); err != nil {
		return errors.Annotate(err, "shallow validate value fields").Err()
	}
	if err := stableversion.ValidateFaftVersion(fwImage); err != nil {
		return errors.Annotate(err, "shallow validate value fields").Err()
	}
	return nil
}
