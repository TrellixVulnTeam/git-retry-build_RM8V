// Code generated by svcdec; DO NOT EDIT.

package ufspb

import (
	"context"

	proto "github.com/golang/protobuf/proto"

	empty "github.com/golang/protobuf/ptypes/empty"
	status "google.golang.org/genproto/googleapis/rpc/status"
	proto1 "infra/unifiedfleet/api/v1/proto"
)

type DecoratedFleet struct {
	// Service is the service to decorate.
	Service FleetServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(ctx context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(ctx context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedFleet) MachineRegistration(ctx context.Context, req *MachineRegistrationRequest) (rsp *MachineRegistrationResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "MachineRegistration", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.MachineRegistration(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "MachineRegistration", rsp, err)
	}
	return
}

func (s *DecoratedFleet) RackRegistration(ctx context.Context, req *RackRegistrationRequest) (rsp *RackRegistrationResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "RackRegistration", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.RackRegistration(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "RackRegistration", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateChromePlatform(ctx context.Context, req *CreateChromePlatformRequest) (rsp *proto1.ChromePlatform, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateChromePlatform", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateChromePlatform(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateChromePlatform", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateChromePlatform(ctx context.Context, req *UpdateChromePlatformRequest) (rsp *proto1.ChromePlatform, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateChromePlatform", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateChromePlatform(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateChromePlatform", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetChromePlatform(ctx context.Context, req *GetChromePlatformRequest) (rsp *proto1.ChromePlatform, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetChromePlatform", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetChromePlatform(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetChromePlatform", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListChromePlatforms(ctx context.Context, req *ListChromePlatformsRequest) (rsp *ListChromePlatformsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListChromePlatforms", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListChromePlatforms(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListChromePlatforms", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteChromePlatform(ctx context.Context, req *DeleteChromePlatformRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteChromePlatform", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteChromePlatform(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteChromePlatform", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportChromePlatforms(ctx context.Context, req *ImportChromePlatformsRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportChromePlatforms", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportChromePlatforms(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportChromePlatforms", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListOSVersions(ctx context.Context, req *ListOSVersionsRequest) (rsp *ListOSVersionsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListOSVersions", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListOSVersions(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListOSVersions", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportOSVersions(ctx context.Context, req *ImportOSVersionsRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportOSVersions", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportOSVersions(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportOSVersions", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateMachineLSEPrototype(ctx context.Context, req *CreateMachineLSEPrototypeRequest) (rsp *proto1.MachineLSEPrototype, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateMachineLSEPrototype", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateMachineLSEPrototype(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateMachineLSEPrototype", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateMachineLSEPrototype(ctx context.Context, req *UpdateMachineLSEPrototypeRequest) (rsp *proto1.MachineLSEPrototype, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateMachineLSEPrototype", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateMachineLSEPrototype(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateMachineLSEPrototype", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetMachineLSEPrototype(ctx context.Context, req *GetMachineLSEPrototypeRequest) (rsp *proto1.MachineLSEPrototype, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetMachineLSEPrototype", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetMachineLSEPrototype(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetMachineLSEPrototype", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListMachineLSEPrototypes(ctx context.Context, req *ListMachineLSEPrototypesRequest) (rsp *ListMachineLSEPrototypesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListMachineLSEPrototypes", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListMachineLSEPrototypes(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListMachineLSEPrototypes", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteMachineLSEPrototype(ctx context.Context, req *DeleteMachineLSEPrototypeRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteMachineLSEPrototype", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteMachineLSEPrototype(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteMachineLSEPrototype", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateRackLSEPrototype(ctx context.Context, req *CreateRackLSEPrototypeRequest) (rsp *proto1.RackLSEPrototype, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateRackLSEPrototype", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateRackLSEPrototype(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateRackLSEPrototype", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateRackLSEPrototype(ctx context.Context, req *UpdateRackLSEPrototypeRequest) (rsp *proto1.RackLSEPrototype, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateRackLSEPrototype", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateRackLSEPrototype(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateRackLSEPrototype", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetRackLSEPrototype(ctx context.Context, req *GetRackLSEPrototypeRequest) (rsp *proto1.RackLSEPrototype, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetRackLSEPrototype", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetRackLSEPrototype(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetRackLSEPrototype", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListRackLSEPrototypes(ctx context.Context, req *ListRackLSEPrototypesRequest) (rsp *ListRackLSEPrototypesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListRackLSEPrototypes", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListRackLSEPrototypes(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListRackLSEPrototypes", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteRackLSEPrototype(ctx context.Context, req *DeleteRackLSEPrototypeRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteRackLSEPrototype", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteRackLSEPrototype(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteRackLSEPrototype", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateMachine(ctx context.Context, req *CreateMachineRequest) (rsp *proto1.Machine, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateMachine", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateMachine(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateMachine", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateMachine(ctx context.Context, req *UpdateMachineRequest) (rsp *proto1.Machine, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateMachine", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateMachine(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateMachine", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetMachine(ctx context.Context, req *GetMachineRequest) (rsp *proto1.Machine, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetMachine", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetMachine(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetMachine", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListMachines(ctx context.Context, req *ListMachinesRequest) (rsp *ListMachinesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListMachines", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListMachines(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListMachines", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteMachine(ctx context.Context, req *DeleteMachineRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteMachine", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteMachine(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteMachine", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportMachines(ctx context.Context, req *ImportMachinesRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportMachines", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportMachines(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportMachines", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateRack(ctx context.Context, req *CreateRackRequest) (rsp *proto1.Rack, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateRack", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateRack(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateRack", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateRack(ctx context.Context, req *UpdateRackRequest) (rsp *proto1.Rack, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateRack", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateRack(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateRack", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetRack(ctx context.Context, req *GetRackRequest) (rsp *proto1.Rack, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetRack", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetRack(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetRack", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListRacks(ctx context.Context, req *ListRacksRequest) (rsp *ListRacksResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListRacks", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListRacks(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListRacks", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteRack(ctx context.Context, req *DeleteRackRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteRack", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteRack(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteRack", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateMachineLSE(ctx context.Context, req *CreateMachineLSERequest) (rsp *proto1.MachineLSE, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateMachineLSE", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateMachineLSE(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateMachineLSE", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateMachineLSE(ctx context.Context, req *UpdateMachineLSERequest) (rsp *proto1.MachineLSE, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateMachineLSE", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateMachineLSE(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateMachineLSE", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetMachineLSE(ctx context.Context, req *GetMachineLSERequest) (rsp *proto1.MachineLSE, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetMachineLSE", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetMachineLSE(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetMachineLSE", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListMachineLSEs(ctx context.Context, req *ListMachineLSEsRequest) (rsp *ListMachineLSEsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListMachineLSEs", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListMachineLSEs(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListMachineLSEs", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteMachineLSE(ctx context.Context, req *DeleteMachineLSERequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteMachineLSE", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteMachineLSE(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteMachineLSE", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportMachineLSEs(ctx context.Context, req *ImportMachineLSEsRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportMachineLSEs", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportMachineLSEs(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportMachineLSEs", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportOSMachineLSEs(ctx context.Context, req *ImportOSMachineLSEsRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportOSMachineLSEs", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportOSMachineLSEs(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportOSMachineLSEs", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateRackLSE(ctx context.Context, req *CreateRackLSERequest) (rsp *proto1.RackLSE, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateRackLSE", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateRackLSE(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateRackLSE", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateRackLSE(ctx context.Context, req *UpdateRackLSERequest) (rsp *proto1.RackLSE, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateRackLSE", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateRackLSE(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateRackLSE", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetRackLSE(ctx context.Context, req *GetRackLSERequest) (rsp *proto1.RackLSE, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetRackLSE", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetRackLSE(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetRackLSE", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListRackLSEs(ctx context.Context, req *ListRackLSEsRequest) (rsp *ListRackLSEsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListRackLSEs", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListRackLSEs(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListRackLSEs", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteRackLSE(ctx context.Context, req *DeleteRackLSERequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteRackLSE", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteRackLSE(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteRackLSE", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateNic(ctx context.Context, req *CreateNicRequest) (rsp *proto1.Nic, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateNic", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateNic(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateNic", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateNic(ctx context.Context, req *UpdateNicRequest) (rsp *proto1.Nic, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateNic", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateNic(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateNic", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetNic(ctx context.Context, req *GetNicRequest) (rsp *proto1.Nic, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetNic", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetNic(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetNic", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListNics(ctx context.Context, req *ListNicsRequest) (rsp *ListNicsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListNics", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListNics(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListNics", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteNic(ctx context.Context, req *DeleteNicRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteNic", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteNic(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteNic", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportNics(ctx context.Context, req *ImportNicsRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportNics", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportNics(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportNics", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportDatacenters(ctx context.Context, req *ImportDatacentersRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportDatacenters", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportDatacenters(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportDatacenters", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateKVM(ctx context.Context, req *CreateKVMRequest) (rsp *proto1.KVM, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateKVM", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateKVM(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateKVM", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateKVM(ctx context.Context, req *UpdateKVMRequest) (rsp *proto1.KVM, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateKVM", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateKVM(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateKVM", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetKVM(ctx context.Context, req *GetKVMRequest) (rsp *proto1.KVM, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetKVM", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetKVM(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetKVM", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListKVMs(ctx context.Context, req *ListKVMsRequest) (rsp *ListKVMsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListKVMs", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListKVMs(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListKVMs", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteKVM(ctx context.Context, req *DeleteKVMRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteKVM", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteKVM(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteKVM", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateRPM(ctx context.Context, req *CreateRPMRequest) (rsp *proto1.RPM, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateRPM", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateRPM(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateRPM", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateRPM(ctx context.Context, req *UpdateRPMRequest) (rsp *proto1.RPM, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateRPM", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateRPM(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateRPM", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetRPM(ctx context.Context, req *GetRPMRequest) (rsp *proto1.RPM, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetRPM", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetRPM(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetRPM", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListRPMs(ctx context.Context, req *ListRPMsRequest) (rsp *ListRPMsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListRPMs", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListRPMs(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListRPMs", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteRPM(ctx context.Context, req *DeleteRPMRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteRPM", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteRPM(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteRPM", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateDrac(ctx context.Context, req *CreateDracRequest) (rsp *proto1.Drac, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateDrac", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateDrac(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateDrac", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateDrac(ctx context.Context, req *UpdateDracRequest) (rsp *proto1.Drac, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateDrac", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateDrac(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateDrac", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetDrac(ctx context.Context, req *GetDracRequest) (rsp *proto1.Drac, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetDrac", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetDrac(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetDrac", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListDracs(ctx context.Context, req *ListDracsRequest) (rsp *ListDracsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListDracs", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListDracs(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListDracs", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteDrac(ctx context.Context, req *DeleteDracRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteDrac", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteDrac(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteDrac", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateSwitch(ctx context.Context, req *CreateSwitchRequest) (rsp *proto1.Switch, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateSwitch", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateSwitch(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateSwitch", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateSwitch(ctx context.Context, req *UpdateSwitchRequest) (rsp *proto1.Switch, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateSwitch", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateSwitch(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateSwitch", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetSwitch(ctx context.Context, req *GetSwitchRequest) (rsp *proto1.Switch, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetSwitch", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetSwitch(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetSwitch", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListSwitches(ctx context.Context, req *ListSwitchesRequest) (rsp *ListSwitchesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListSwitches", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListSwitches(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListSwitches", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteSwitch(ctx context.Context, req *DeleteSwitchRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteSwitch", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteSwitch(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteSwitch", rsp, err)
	}
	return
}

func (s *DecoratedFleet) CreateVlan(ctx context.Context, req *CreateVlanRequest) (rsp *proto1.Vlan, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateVlan", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateVlan(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateVlan", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateVlan(ctx context.Context, req *UpdateVlanRequest) (rsp *proto1.Vlan, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateVlan", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateVlan(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateVlan", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetVlan(ctx context.Context, req *GetVlanRequest) (rsp *proto1.Vlan, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetVlan", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetVlan(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetVlan", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListVlans(ctx context.Context, req *ListVlansRequest) (rsp *ListVlansResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListVlans", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListVlans(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListVlans", rsp, err)
	}
	return
}

func (s *DecoratedFleet) DeleteVlan(ctx context.Context, req *DeleteVlanRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteVlan", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteVlan(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteVlan", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportVlans(ctx context.Context, req *ImportVlansRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportVlans", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportVlans(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportVlans", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportOSVlans(ctx context.Context, req *ImportOSVlansRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportOSVlans", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportOSVlans(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportOSVlans", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ImportStates(ctx context.Context, req *ImportStatesRequest) (rsp *status.Status, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ImportStates", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ImportStates(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ImportStates", rsp, err)
	}
	return
}

func (s *DecoratedFleet) UpdateState(ctx context.Context, req *UpdateStateRequest) (rsp *proto1.StateRecord, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateState", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateState(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateState", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetState(ctx context.Context, req *GetStateRequest) (rsp *proto1.StateRecord, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetState", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetState(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetState", rsp, err)
	}
	return
}

func (s *DecoratedFleet) GetDHCPConfig(ctx context.Context, req *GetDHCPConfigRequest) (rsp *proto1.DHCPConfig, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetDHCPConfig", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetDHCPConfig(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetDHCPConfig", rsp, err)
	}
	return
}

func (s *DecoratedFleet) ListVMs(ctx context.Context, req *ListVMsRequest) (rsp *ListVMsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListVMs", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListVMs(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListVMs", rsp, err)
	}
	return
}
