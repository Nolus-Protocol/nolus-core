// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nolus/mint/v1beta1/query.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0819bb52a62656e, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0819bb52a62656e, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryMintStateRequest is the request type for the Query/State RPC method.
type QueryMintStateRequest struct {
}

func (m *QueryMintStateRequest) Reset()         { *m = QueryMintStateRequest{} }
func (m *QueryMintStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMintStateRequest) ProtoMessage()    {}
func (*QueryMintStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0819bb52a62656e, []int{2}
}
func (m *QueryMintStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMintStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMintStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMintStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMintStateRequest.Merge(m, src)
}
func (m *QueryMintStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMintStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMintStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMintStateRequest proto.InternalMessageInfo

// QueryMintStateResponse is the response type for the Query/State RPC
// method.
type QueryMintStateResponse struct {
	NormTimePassed cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=norm_time_passed,json=normTimePassed,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"norm_time_passed"`
	TotalMinted    cosmossdk_io_math.Uint      `protobuf:"bytes,2,opt,name=total_minted,json=totalMinted,proto3,customtype=cosmossdk.io/math.Uint" json:"total_minted"`
}

func (m *QueryMintStateResponse) Reset()         { *m = QueryMintStateResponse{} }
func (m *QueryMintStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMintStateResponse) ProtoMessage()    {}
func (*QueryMintStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0819bb52a62656e, []int{3}
}
func (m *QueryMintStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMintStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMintStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMintStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMintStateResponse.Merge(m, src)
}
func (m *QueryMintStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMintStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMintStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMintStateResponse proto.InternalMessageInfo

// QueryAnnualInflationRequest is the request type for the Query/AnnualInflation
// RPC method.
type QueryAnnualInflationRequest struct {
}

func (m *QueryAnnualInflationRequest) Reset()         { *m = QueryAnnualInflationRequest{} }
func (m *QueryAnnualInflationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAnnualInflationRequest) ProtoMessage()    {}
func (*QueryAnnualInflationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0819bb52a62656e, []int{4}
}
func (m *QueryAnnualInflationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAnnualInflationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAnnualInflationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAnnualInflationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAnnualInflationRequest.Merge(m, src)
}
func (m *QueryAnnualInflationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAnnualInflationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAnnualInflationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAnnualInflationRequest proto.InternalMessageInfo

// QueryAnnualInflationResponse is the response type for the
// Query/AnnualInflation RPC method.
type QueryAnnualInflationResponse struct {
	// inflation is the current minting inflation value.
	AnnualInflation cosmossdk_io_math.Uint `protobuf:"bytes,1,opt,name=annual_inflation,json=annualInflation,proto3,customtype=cosmossdk.io/math.Uint" json:"annual_inflation"`
}

func (m *QueryAnnualInflationResponse) Reset()         { *m = QueryAnnualInflationResponse{} }
func (m *QueryAnnualInflationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAnnualInflationResponse) ProtoMessage()    {}
func (*QueryAnnualInflationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0819bb52a62656e, []int{5}
}
func (m *QueryAnnualInflationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAnnualInflationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAnnualInflationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAnnualInflationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAnnualInflationResponse.Merge(m, src)
}
func (m *QueryAnnualInflationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAnnualInflationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAnnualInflationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAnnualInflationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "nolus.mint.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "nolus.mint.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryMintStateRequest)(nil), "nolus.mint.v1beta1.QueryMintStateRequest")
	proto.RegisterType((*QueryMintStateResponse)(nil), "nolus.mint.v1beta1.QueryMintStateResponse")
	proto.RegisterType((*QueryAnnualInflationRequest)(nil), "nolus.mint.v1beta1.QueryAnnualInflationRequest")
	proto.RegisterType((*QueryAnnualInflationResponse)(nil), "nolus.mint.v1beta1.QueryAnnualInflationResponse")
}

func init() { proto.RegisterFile("nolus/mint/v1beta1/query.proto", fileDescriptor_c0819bb52a62656e) }

var fileDescriptor_c0819bb52a62656e = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0x8e, 0x0b, 0x54, 0xe2, 0x5a, 0xd1, 0xea, 0x28, 0xa5, 0x38, 0xa9, 0x0b, 0x06, 0x95, 0x0f,
	0x51, 0x1f, 0x29, 0x0b, 0x6b, 0x23, 0x96, 0x0a, 0x02, 0x21, 0xc0, 0xc2, 0x12, 0x5d, 0x9c, 0xc3,
	0x3d, 0x61, 0xdf, 0xeb, 0xfa, 0xce, 0x88, 0x0c, 0x2c, 0x48, 0xec, 0x48, 0xfc, 0x03, 0x46, 0x56,
	0xfe, 0x44, 0xc7, 0x4a, 0x2c, 0x88, 0xa1, 0x42, 0x09, 0x3f, 0x04, 0xf9, 0xf5, 0x05, 0xa9, 0x89,
	0xab, 0x76, 0xb3, 0xee, 0x79, 0x9f, 0x8f, 0x7b, 0x1f, 0x1f, 0xf1, 0x14, 0xc4, 0xb9, 0x66, 0x89,
	0x54, 0x86, 0xbd, 0x6f, 0xf6, 0x85, 0xe1, 0x4d, 0xb6, 0x9f, 0x8b, 0x6c, 0x18, 0xa4, 0x19, 0x18,
	0xa0, 0x14, 0xf1, 0xa0, 0xc0, 0x03, 0x8b, 0xbb, 0x2b, 0x11, 0x44, 0x80, 0x30, 0x2b, 0xbe, 0xca,
	0x49, 0xb7, 0x11, 0x01, 0x44, 0xb1, 0x60, 0x3c, 0x95, 0x8c, 0x2b, 0x05, 0x86, 0x1b, 0x09, 0x4a,
	0x5b, 0x74, 0xbd, 0xc2, 0x07, 0x45, 0x11, 0xf6, 0x57, 0x08, 0x7d, 0x51, 0xb8, 0x76, 0x78, 0xc6,
	0x13, 0xdd, 0x15, 0xfb, 0xb9, 0xd0, 0xc6, 0x7f, 0x4e, 0x2e, 0x1f, 0x3b, 0xd5, 0x29, 0x28, 0x2d,
	0xe8, 0x23, 0x32, 0x9f, 0xe2, 0xc9, 0x9a, 0x73, 0xdd, 0xb9, 0xb3, 0xb0, 0xed, 0x06, 0xb3, 0x21,
	0x83, 0x92, 0xd3, 0x3a, 0x7f, 0x70, 0xb4, 0x51, 0xeb, 0xda, 0x79, 0xff, 0x2a, 0xb9, 0x82, 0x82,
	0x6d, 0xa9, 0xcc, 0x4b, 0xc3, 0x8d, 0x98, 0x38, 0x7d, 0x77, 0xc8, 0xea, 0x34, 0x62, 0xdd, 0xda,
	0x64, 0x59, 0x41, 0x96, 0xf4, 0x8c, 0x4c, 0x44, 0x2f, 0xe5, 0x5a, 0x8b, 0x01, 0xfa, 0x2e, 0xb6,
	0x6e, 0x16, 0xda, 0xbf, 0x8f, 0x36, 0xea, 0x21, 0xe8, 0x04, 0xb4, 0x1e, 0xbc, 0x0b, 0x24, 0xb0,
	0x84, 0x9b, 0xbd, 0xe0, 0xa9, 0x88, 0x78, 0x38, 0x7c, 0x2c, 0xc2, 0xee, 0xa5, 0x82, 0xfc, 0x4a,
	0x26, 0xa2, 0x83, 0x54, 0xba, 0x43, 0x16, 0x0d, 0x18, 0x1e, 0xf7, 0x8a, 0xb4, 0x62, 0xb0, 0x36,
	0x87, 0x52, 0x9e, 0x95, 0x5a, 0x9d, 0x95, 0x7a, 0x2d, 0x95, 0xe9, 0x2e, 0x20, 0xa7, 0x8d, 0x14,
	0x7f, 0x9d, 0xd4, 0x31, 0xeb, 0x8e, 0x52, 0x39, 0x8f, 0x77, 0xd5, 0xdb, 0x18, 0x57, 0x3d, 0xb9,
	0x8b, 0x24, 0x8d, 0x6a, 0xd8, 0x5e, 0x68, 0x97, 0x2c, 0x73, 0x84, 0x7a, 0x72, 0x82, 0xd9, 0x0b,
	0x9d, 0x96, 0x62, 0x89, 0x1f, 0x97, 0xdc, 0xfe, 0x71, 0x8e, 0x5c, 0x40, 0x2f, 0xfa, 0x91, 0xcc,
	0x97, 0x1b, 0xa7, 0x9b, 0x55, 0x6d, 0xcc, 0x96, 0xeb, 0xde, 0x3e, 0x75, 0xae, 0xcc, 0xeb, 0xfb,
	0x9f, 0x7e, 0xfe, 0xfd, 0x3a, 0xd7, 0xa0, 0x2e, 0xab, 0xf8, 0x87, 0xca, 0x62, 0xe9, 0x67, 0x87,
	0x5c, 0xfc, 0x5f, 0x1d, 0xbd, 0x7b, 0xa2, 0xf4, 0x74, 0xf1, 0xee, 0xbd, 0xb3, 0x8c, 0xda, 0x20,
	0x37, 0x30, 0x48, 0x9d, 0x5e, 0xab, 0x0a, 0xa2, 0xd1, 0xf9, 0x9b, 0x43, 0x96, 0xa6, 0xf6, 0x4e,
	0xd9, 0x89, 0x16, 0xd5, 0x05, 0xba, 0x0f, 0xce, 0x4e, 0xb0, 0xc9, 0xee, 0x63, 0xb2, 0x4d, 0x7a,
	0xab, 0x2a, 0xd9, 0x74, 0xd9, 0xad, 0x27, 0x07, 0x23, 0xcf, 0x39, 0x1c, 0x79, 0xce, 0x9f, 0x91,
	0xe7, 0x7c, 0x19, 0x7b, 0xb5, 0xc3, 0xb1, 0x57, 0xfb, 0x35, 0xf6, 0x6a, 0x6f, 0x9a, 0x91, 0x34,
	0x7b, 0x79, 0x3f, 0x08, 0x21, 0x61, 0xcf, 0x0a, 0xa5, 0xad, 0x4e, 0xf1, 0x3c, 0x43, 0x88, 0x4b,
	0xe1, 0xad, 0x10, 0x32, 0xc1, 0x3e, 0x94, 0xfa, 0x66, 0x98, 0x0a, 0xdd, 0x9f, 0xc7, 0x07, 0xfc,
	0xf0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x73, 0x0e, 0x40, 0x49, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/nolus.mint.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MintState(ctx context.Context, in *QueryMintStateRequest, opts ...grpc.CallOption) (*QueryMintStateResponse, error) {
	out := new(QueryMintStateResponse)
	err := c.cc.Invoke(ctx, "/nolus.mint.v1beta1.Query/MintState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AnnualInflation(ctx context.Context, in *QueryAnnualInflationRequest, opts ...grpc.CallOption) (*QueryAnnualInflationResponse, error) {
	out := new(QueryAnnualInflationResponse)
	err := c.cc.Invoke(ctx, "/nolus.mint.v1beta1.Query/AnnualInflation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the total set of minting parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// MintState returns the current minting state value.
	MintState(context.Context, *QueryMintStateRequest) (*QueryMintStateResponse, error)
	// AnnualInflation returns the current minting inflation rate for the next 12
	// months.
	AnnualInflation(context.Context, *QueryAnnualInflationRequest) (*QueryAnnualInflationResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) MintState(ctx context.Context, req *QueryMintStateRequest) (*QueryMintStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintState not implemented")
}
func (*UnimplementedQueryServer) AnnualInflation(ctx context.Context, req *QueryAnnualInflationRequest) (*QueryAnnualInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnualInflation not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
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
		FullMethod: "/nolus.mint.v1beta1.Query/Params",
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
		FullMethod: "/nolus.mint.v1beta1.Query/MintState",
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
		FullMethod: "/nolus.mint.v1beta1.Query/AnnualInflation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AnnualInflation(ctx, req.(*QueryAnnualInflationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
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

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryMintStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMintStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMintStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryMintStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMintStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMintStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalMinted.Size()
		i -= size
		if _, err := m.TotalMinted.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.NormTimePassed.Size()
		i -= size
		if _, err := m.NormTimePassed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAnnualInflationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAnnualInflationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAnnualInflationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAnnualInflationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAnnualInflationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAnnualInflationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualInflation.Size()
		i -= size
		if _, err := m.AnnualInflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryMintStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryMintStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NormTimePassed.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.TotalMinted.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAnnualInflationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAnnualInflationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AnnualInflation.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryMintStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryMintStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMintStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryMintStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryMintStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMintStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NormTimePassed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NormTimePassed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalMinted", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalMinted.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAnnualInflationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAnnualInflationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAnnualInflationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAnnualInflationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAnnualInflationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAnnualInflationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualInflation", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualInflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
