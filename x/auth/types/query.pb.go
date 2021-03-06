// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/auth/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// QueryAccountRequest is request type for the Query/Account RPC method
type QueryAccountRequest struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
}

func (m *QueryAccountRequest) Reset()         { *m = QueryAccountRequest{} }
func (m *QueryAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAccountRequest) ProtoMessage()    {}
func (*QueryAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bc52f4cb65abdb, []int{0}
}
func (m *QueryAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAccountRequest.Merge(m, src)
}
func (m *QueryAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAccountRequest proto.InternalMessageInfo

func (m *QueryAccountRequest) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

// QueryAccountResponse is response type for the Query/Account RPC method
type QueryAccountResponse struct {
	Account *types.Any `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *QueryAccountResponse) Reset()         { *m = QueryAccountResponse{} }
func (m *QueryAccountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAccountResponse) ProtoMessage()    {}
func (*QueryAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bc52f4cb65abdb, []int{1}
}
func (m *QueryAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAccountResponse.Merge(m, src)
}
func (m *QueryAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAccountResponse proto.InternalMessageInfo

func (m *QueryAccountResponse) GetAccount() *types.Any {
	if m != nil {
		return m.Account
	}
	return nil
}

// QueryParametersRequest is request type for the Query/Parameters RPC method
type QueryParametersRequest struct {
}

func (m *QueryParametersRequest) Reset()         { *m = QueryParametersRequest{} }
func (m *QueryParametersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParametersRequest) ProtoMessage()    {}
func (*QueryParametersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bc52f4cb65abdb, []int{2}
}
func (m *QueryParametersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParametersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParametersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParametersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParametersRequest.Merge(m, src)
}
func (m *QueryParametersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParametersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParametersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParametersRequest proto.InternalMessageInfo

// QueryParametersResponse is response type for the Query/Parameters RPC method
type QueryParametersResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParametersResponse) Reset()         { *m = QueryParametersResponse{} }
func (m *QueryParametersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParametersResponse) ProtoMessage()    {}
func (*QueryParametersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1bc52f4cb65abdb, []int{3}
}
func (m *QueryParametersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParametersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParametersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParametersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParametersResponse.Merge(m, src)
}
func (m *QueryParametersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParametersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParametersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParametersResponse proto.InternalMessageInfo

func (m *QueryParametersResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryAccountRequest)(nil), "cosmos.auth.QueryAccountRequest")
	proto.RegisterType((*QueryAccountResponse)(nil), "cosmos.auth.QueryAccountResponse")
	proto.RegisterType((*QueryParametersRequest)(nil), "cosmos.auth.QueryParametersRequest")
	proto.RegisterType((*QueryParametersResponse)(nil), "cosmos.auth.QueryParametersResponse")
}

func init() { proto.RegisterFile("cosmos/auth/query.proto", fileDescriptor_e1bc52f4cb65abdb) }

var fileDescriptor_e1bc52f4cb65abdb = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0xdb, 0xe4, 0xfb, 0xc0, 0x2c, 0x9c, 0x16, 0x02, 0xd8, 0x43, 0xc1, 0xea, 0x41, 0x0f,
	0x6c, 0x03, 0x5e, 0xbd, 0xb4, 0x9e, 0x8c, 0x1e, 0xb0, 0xf1, 0xa2, 0x17, 0xd3, 0x96, 0xb5, 0x18,
	0xa5, 0x5b, 0x3a, 0xdb, 0x44, 0xde, 0xc2, 0x87, 0xe1, 0x21, 0x88, 0x27, 0x8e, 0x9e, 0x88, 0x81,
	0xb7, 0xf0, 0x64, 0xba, 0xbb, 0x8d, 0x10, 0x0c, 0x97, 0xb6, 0x3b, 0x33, 0xbf, 0xff, 0x7f, 0x66,
	0xa7, 0xa8, 0x19, 0x32, 0x18, 0x33, 0xb0, 0xfd, 0x8c, 0x8f, 0xec, 0x49, 0x46, 0xd3, 0x29, 0x49,
	0x52, 0xc6, 0x19, 0xae, 0xc8, 0x04, 0xc9, 0x13, 0x46, 0x3d, 0x62, 0x11, 0x13, 0x71, 0x3b, 0xff,
	0x92, 0x25, 0xc6, 0x61, 0xc4, 0x58, 0xf4, 0x4a, 0x6d, 0x71, 0x0a, 0xb2, 0x27, 0xdb, 0x8f, 0x15,
	0x6d, 0x34, 0x36, 0x65, 0xf3, 0x47, 0x81, 0xc8, 0xf8, 0xa3, 0xd4, 0x52, 0x16, 0xe2, 0x60, 0x05,
	0xa8, 0x76, 0x9b, 0xfb, 0x3b, 0x61, 0xc8, 0xb2, 0x98, 0x7b, 0x74, 0x92, 0x51, 0xe0, 0xf8, 0x1a,
	0x95, 0xfd, 0xe1, 0x30, 0xa5, 0x00, 0x2d, 0xbd, 0xa3, 0x9f, 0x56, 0xdd, 0xde, 0xf7, 0xb2, 0xdd,
	0x8d, 0x9e, 0xf9, 0x28, 0x0b, 0x48, 0xc8, 0xc6, 0x4a, 0x44, 0xbd, 0xba, 0x30, 0x7c, 0xb1, 0xf9,
	0x34, 0xa1, 0x40, 0x9c, 0x30, 0x74, 0x24, 0xe8, 0x15, 0x0a, 0xd6, 0x1d, 0xaa, 0x6f, 0x7b, 0x40,
	0xc2, 0x62, 0xa0, 0xf8, 0x02, 0x95, 0x7d, 0x19, 0x12, 0x26, 0x95, 0x7e, 0x9d, 0xc8, 0xd9, 0x48,
	0x31, 0x1b, 0x71, 0xe2, 0xa9, 0x5b, 0xfd, 0x98, 0x75, 0x0f, 0x14, 0x7b, 0xe5, 0x15, 0x88, 0xd5,
	0x42, 0x0d, 0xa1, 0x3a, 0xf0, 0x53, 0x7f, 0x4c, 0x39, 0x4d, 0x41, 0x35, 0x6f, 0xdd, 0xa0, 0xe6,
	0x4e, 0x46, 0x59, 0xf6, 0x50, 0x29, 0xc9, 0xa3, 0xa0, 0x1c, 0x6b, 0x64, 0xe3, 0xc2, 0x89, 0x00,
	0xc0, 0xfd, 0x37, 0x5f, 0xb6, 0x35, 0x4f, 0x15, 0xf6, 0x67, 0x3a, 0xfa, 0x2f, 0xe4, 0xf0, 0x00,
	0x95, 0x55, 0x1b, 0xb8, 0xb3, 0xc5, 0xfd, 0x71, 0x83, 0xc6, 0xd1, 0x9e, 0x0a, 0xd9, 0x8c, 0xa5,
	0xe1, 0x7b, 0x84, 0x7e, 0x9b, 0xc4, 0xc7, 0xbb, 0xc8, 0xce, 0x70, 0xc6, 0xc9, 0xfe, 0xa2, 0x42,
	0xda, 0xbd, 0x9c, 0xaf, 0x4c, 0x7d, 0xb1, 0x32, 0xf5, 0xaf, 0x95, 0xa9, 0xbf, 0xaf, 0x4d, 0x6d,
	0xb1, 0x36, 0xb5, 0xcf, 0xb5, 0xa9, 0x3d, 0x9c, 0xed, 0x5d, 0xe3, 0x9b, 0xfc, 0x7b, 0xc4, 0x36,
	0x83, 0x92, 0x58, 0xc4, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0xbe, 0xe5, 0xa3, 0xb0,
	0x02, 0x00, 0x00,
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
	// Account returns account details based on address
	Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*QueryAccountResponse, error)
	// Parameters queries all params
	Parameters(ctx context.Context, in *QueryParametersRequest, opts ...grpc.CallOption) (*QueryParametersResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*QueryAccountResponse, error) {
	out := new(QueryAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.auth.Query/Account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Parameters(ctx context.Context, in *QueryParametersRequest, opts ...grpc.CallOption) (*QueryParametersResponse, error) {
	out := new(QueryParametersResponse)
	err := c.cc.Invoke(ctx, "/cosmos.auth.Query/Parameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Account returns account details based on address
	Account(context.Context, *QueryAccountRequest) (*QueryAccountResponse, error)
	// Parameters queries all params
	Parameters(context.Context, *QueryParametersRequest) (*QueryParametersResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Account(ctx context.Context, req *QueryAccountRequest) (*QueryAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (*UnimplementedQueryServer) Parameters(ctx context.Context, req *QueryParametersRequest) (*QueryParametersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parameters not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.auth.Query/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Account(ctx, req.(*QueryAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Parameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Parameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.auth.Query/Parameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Parameters(ctx, req.(*QueryParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.auth.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Account",
			Handler:    _Query_Account_Handler,
		},
		{
			MethodName: "Parameters",
			Handler:    _Query_Parameters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/auth/query.proto",
}

func (m *QueryAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParametersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParametersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParametersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParametersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParametersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParametersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *QueryAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParametersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParametersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAccountRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryAccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			if m.Account == nil {
				m.Account = &types.Any{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryParametersRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParametersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParametersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryParametersResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParametersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParametersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
