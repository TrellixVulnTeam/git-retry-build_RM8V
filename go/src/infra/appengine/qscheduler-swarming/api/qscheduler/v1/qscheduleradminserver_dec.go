// Code generated by svcdec; DO NOT EDIT.

package qscheduler

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedQSchedulerAdmin struct {
	// Service is the service to decorate.
	Service QSchedulerAdminServer
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

func (s *DecoratedQSchedulerAdmin) CreateSchedulerPool(c context.Context, req *CreateSchedulerPoolRequest) (rsp *CreateSchedulerPoolResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "CreateSchedulerPool", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateSchedulerPool(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "CreateSchedulerPool", rsp, err)
	}
	return
}

func (s *DecoratedQSchedulerAdmin) CreateAccount(c context.Context, req *CreateAccountRequest) (rsp *CreateAccountResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "CreateAccount", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateAccount(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "CreateAccount", rsp, err)
	}
	return
}
