// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: nolus/mint/v1beta1/query.proto

package mintv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Params_FullMethodName          = "/nolus.mint.v1beta1.Query/Params"
	Query_MintState_FullMethodName       = "/nolus.mint.v1beta1.Query/MintState"
	Query_AnnualInflation_FullMethodName = "/nolus.mint.v1beta1.Query/AnnualInflation"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the total set of minting parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// MintState returns the current minting state value.
	MintState(ctx context.Context, in *QueryMintStateRequest, opts ...grpc.CallOption) (*QueryMintStateResponse, error)
	// AnnualInflation returns the current minting inflation rate for the next 12
	// months.
	AnnualInflation(ctx context.Context, in *QueryAnnualInflationRequest, opts ...grpc.CallOption) (*QueryAnnualInflationResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MintState(ctx context.Context, in *QueryMintStateRequest, opts ...grpc.CallOption) (*QueryMintStateResponse, error) {
	out := new(QueryMintStateResponse)
	err := c.cc.Invoke(ctx, Query_MintState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AnnualInflation(ctx context.Context, in *QueryAnnualInflationRequest, opts ...grpc.CallOption) (*QueryAnnualInflationResponse, error) {
	out := new(QueryAnnualInflationResponse)
	err := c.cc.Invoke(ctx, Query_AnnualInflation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params returns the total set of minting parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// MintState returns the current minting state value.
	MintState(context.Context, *QueryMintStateRequest) (*QueryMintStateResponse, error)
	// AnnualInflation returns the current minting inflation rate for the next 12
	// months.
	AnnualInflation(context.Context, *QueryAnnualInflationRequest) (*QueryAnnualInflationResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) MintState(context.Context, *QueryMintStateRequest) (*QueryMintStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintState not implemented")
}
func (UnimplementedQueryServer) AnnualInflation(context.Context, *QueryAnnualInflationRequest) (*QueryAnnualInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnualInflation not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MintState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMintStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MintState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MintState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MintState(ctx, req.(*QueryMintStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AnnualInflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAnnualInflationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AnnualInflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AnnualInflation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AnnualInflation(ctx, req.(*QueryAnnualInflationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nolus.mint.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "MintState",
			Handler:    _Query_MintState_Handler,
		},
		{
			MethodName: "AnnualInflation",
			Handler:    _Query_AnnualInflation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nolus/mint/v1beta1/query.proto",
}
