// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BeaconDebugClient is the client API for BeaconDebug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeaconDebugClient interface {
	// GetBeaconState returns full BeaconState object for given stateId.
	GetBeaconState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*BeaconStateResponse, error)
	// ListForkChoiceHeads retrieves all possible chain heads (leaves of fork choice tree).
	ListForkChoiceHeads(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ForkChoiceHeadsResponse, error)
}

type beaconDebugClient struct {
	cc grpc.ClientConnInterface
}

func NewBeaconDebugClient(cc grpc.ClientConnInterface) BeaconDebugClient {
	return &beaconDebugClient{cc}
}

func (c *beaconDebugClient) GetBeaconState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*BeaconStateResponse, error) {
	out := new(BeaconStateResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconDebug/GetBeaconState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconDebugClient) ListForkChoiceHeads(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ForkChoiceHeadsResponse, error) {
	out := new(ForkChoiceHeadsResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconDebug/ListForkChoiceHeads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconDebugServer is the server API for BeaconDebug service.
// All implementations must embed UnimplementedBeaconDebugServer
// for forward compatibility
type BeaconDebugServer interface {
	// GetBeaconState returns full BeaconState object for given stateId.
	GetBeaconState(context.Context, *StateRequest) (*BeaconStateResponse, error)
	// ListForkChoiceHeads retrieves all possible chain heads (leaves of fork choice tree).
	ListForkChoiceHeads(context.Context, *empty.Empty) (*ForkChoiceHeadsResponse, error)
	mustEmbedUnimplementedBeaconDebugServer()
}

// UnimplementedBeaconDebugServer must be embedded to have forward compatible implementations.
type UnimplementedBeaconDebugServer struct {
}

func (UnimplementedBeaconDebugServer) GetBeaconState(context.Context, *StateRequest) (*BeaconStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeaconState not implemented")
}
func (UnimplementedBeaconDebugServer) ListForkChoiceHeads(context.Context, *empty.Empty) (*ForkChoiceHeadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListForkChoiceHeads not implemented")
}
func (UnimplementedBeaconDebugServer) mustEmbedUnimplementedBeaconDebugServer() {}

// UnsafeBeaconDebugServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeaconDebugServer will
// result in compilation errors.
type UnsafeBeaconDebugServer interface {
	mustEmbedUnimplementedBeaconDebugServer()
}

func RegisterBeaconDebugServer(s grpc.ServiceRegistrar, srv BeaconDebugServer) {
	s.RegisterService(&_BeaconDebug_serviceDesc, srv)
}

func _BeaconDebug_GetBeaconState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconDebugServer).GetBeaconState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconDebug/GetBeaconState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconDebugServer).GetBeaconState(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconDebug_ListForkChoiceHeads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconDebugServer).ListForkChoiceHeads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconDebug/ListForkChoiceHeads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconDebugServer).ListForkChoiceHeads(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconDebug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1.BeaconDebug",
	HandlerType: (*BeaconDebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBeaconState",
			Handler:    _BeaconDebug_GetBeaconState_Handler,
		},
		{
			MethodName: "ListForkChoiceHeads",
			Handler:    _BeaconDebug_ListForkChoiceHeads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eth/v1/beacon_debug_service.proto",
}