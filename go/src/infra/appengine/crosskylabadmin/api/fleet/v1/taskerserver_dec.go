// Code generated by svcdec; DO NOT EDIT

package fleet

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedTasker struct {
	// Service is the service to decorate.
	Service TaskerServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(c context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(c context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedTasker) TriggerRepairOnIdle(c context.Context, req *TriggerRepairOnIdleRequest) (rsp *TaskerTasksResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "TriggerRepairOnIdle", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.TriggerRepairOnIdle(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "TriggerRepairOnIdle", rsp, err)
	}
	return
}

func (s *DecoratedTasker) TriggerRepairOnRepairFailed(c context.Context, req *TriggerRepairOnRepairFailedRequest) (rsp *TaskerTasksResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "TriggerRepairOnRepairFailed", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.TriggerRepairOnRepairFailed(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "TriggerRepairOnRepairFailed", rsp, err)
	}
	return
}

func (s *DecoratedTasker) EnsureBackgroundTasks(c context.Context, req *EnsureBackgroundTasksRequest) (rsp *TaskerTasksResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "EnsureBackgroundTasks", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.EnsureBackgroundTasks(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "EnsureBackgroundTasks", rsp, err)
	}
	return
}
