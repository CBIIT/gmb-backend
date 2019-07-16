// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/customer_client_link_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for
// [CustomerClientLinkService.GetCustomerClientLink][google.ads.googleads.v0.services.CustomerClientLinkService.GetCustomerClientLink].
type GetCustomerClientLinkRequest struct {
	// The resource name of the customer client link to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerClientLinkRequest) Reset()         { *m = GetCustomerClientLinkRequest{} }
func (m *GetCustomerClientLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerClientLinkRequest) ProtoMessage()    {}
func (*GetCustomerClientLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_client_link_service_706cd4e851eebd0b, []int{0}
}
func (m *GetCustomerClientLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Unmarshal(m, b)
}
func (m *GetCustomerClientLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Marshal(b, m, deterministic)
}
func (dst *GetCustomerClientLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerClientLinkRequest.Merge(dst, src)
}
func (m *GetCustomerClientLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Size(m)
}
func (m *GetCustomerClientLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerClientLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerClientLinkRequest proto.InternalMessageInfo

func (m *GetCustomerClientLinkRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CustomerClientLinkService.MutateCustomerClientLink][google.ads.googleads.v0.services.CustomerClientLinkService.MutateCustomerClientLink].
type MutateCustomerClientLinkRequest struct {
	// The ID of the customer whose customer link are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform on the individual CustomerClientLink.
	Operation            *CustomerClientLinkOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateCustomerClientLinkRequest) Reset()         { *m = MutateCustomerClientLinkRequest{} }
func (m *MutateCustomerClientLinkRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerClientLinkRequest) ProtoMessage()    {}
func (*MutateCustomerClientLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_client_link_service_706cd4e851eebd0b, []int{1}
}
func (m *MutateCustomerClientLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerClientLinkRequest.Unmarshal(m, b)
}
func (m *MutateCustomerClientLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerClientLinkRequest.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerClientLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerClientLinkRequest.Merge(dst, src)
}
func (m *MutateCustomerClientLinkRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerClientLinkRequest.Size(m)
}
func (m *MutateCustomerClientLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerClientLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerClientLinkRequest proto.InternalMessageInfo

func (m *MutateCustomerClientLinkRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerClientLinkRequest) GetOperation() *CustomerClientLinkOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single operation (create, update) on a CustomerClientLink.
type CustomerClientLinkOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerClientLinkOperation_Create
	//	*CustomerClientLinkOperation_Update
	Operation            isCustomerClientLinkOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CustomerClientLinkOperation) Reset()         { *m = CustomerClientLinkOperation{} }
func (m *CustomerClientLinkOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerClientLinkOperation) ProtoMessage()    {}
func (*CustomerClientLinkOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_client_link_service_706cd4e851eebd0b, []int{2}
}
func (m *CustomerClientLinkOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClientLinkOperation.Unmarshal(m, b)
}
func (m *CustomerClientLinkOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClientLinkOperation.Marshal(b, m, deterministic)
}
func (dst *CustomerClientLinkOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClientLinkOperation.Merge(dst, src)
}
func (m *CustomerClientLinkOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerClientLinkOperation.Size(m)
}
func (m *CustomerClientLinkOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClientLinkOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClientLinkOperation proto.InternalMessageInfo

func (m *CustomerClientLinkOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCustomerClientLinkOperation_Operation interface {
	isCustomerClientLinkOperation_Operation()
}

type CustomerClientLinkOperation_Create struct {
	Create *resources.CustomerClientLink `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerClientLinkOperation_Update struct {
	Update *resources.CustomerClientLink `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

func (*CustomerClientLinkOperation_Create) isCustomerClientLinkOperation_Operation() {}

func (*CustomerClientLinkOperation_Update) isCustomerClientLinkOperation_Operation() {}

func (m *CustomerClientLinkOperation) GetOperation() isCustomerClientLinkOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerClientLinkOperation) GetCreate() *resources.CustomerClientLink {
	if x, ok := m.GetOperation().(*CustomerClientLinkOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CustomerClientLinkOperation) GetUpdate() *resources.CustomerClientLink {
	if x, ok := m.GetOperation().(*CustomerClientLinkOperation_Update); ok {
		return x.Update
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CustomerClientLinkOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CustomerClientLinkOperation_OneofMarshaler, _CustomerClientLinkOperation_OneofUnmarshaler, _CustomerClientLinkOperation_OneofSizer, []interface{}{
		(*CustomerClientLinkOperation_Create)(nil),
		(*CustomerClientLinkOperation_Update)(nil),
	}
}

func _CustomerClientLinkOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CustomerClientLinkOperation)
	// operation
	switch x := m.Operation.(type) {
	case *CustomerClientLinkOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *CustomerClientLinkOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CustomerClientLinkOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _CustomerClientLinkOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CustomerClientLinkOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.CustomerClientLink)
		err := b.DecodeMessage(msg)
		m.Operation = &CustomerClientLinkOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.CustomerClientLink)
		err := b.DecodeMessage(msg)
		m.Operation = &CustomerClientLinkOperation_Update{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CustomerClientLinkOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CustomerClientLinkOperation)
	// operation
	switch x := m.Operation.(type) {
	case *CustomerClientLinkOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CustomerClientLinkOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for a CustomerClientLink mutate.
type MutateCustomerClientLinkResponse struct {
	// A result that identifies the resource affected by the mutate request.
	Result               *MutateCustomerClientLinkResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateCustomerClientLinkResponse) Reset()         { *m = MutateCustomerClientLinkResponse{} }
func (m *MutateCustomerClientLinkResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerClientLinkResponse) ProtoMessage()    {}
func (*MutateCustomerClientLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_client_link_service_706cd4e851eebd0b, []int{3}
}
func (m *MutateCustomerClientLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerClientLinkResponse.Unmarshal(m, b)
}
func (m *MutateCustomerClientLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerClientLinkResponse.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerClientLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerClientLinkResponse.Merge(dst, src)
}
func (m *MutateCustomerClientLinkResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerClientLinkResponse.Size(m)
}
func (m *MutateCustomerClientLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerClientLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerClientLinkResponse proto.InternalMessageInfo

func (m *MutateCustomerClientLinkResponse) GetResult() *MutateCustomerClientLinkResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for a single customer client link mutate.
type MutateCustomerClientLinkResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerClientLinkResult) Reset()         { *m = MutateCustomerClientLinkResult{} }
func (m *MutateCustomerClientLinkResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerClientLinkResult) ProtoMessage()    {}
func (*MutateCustomerClientLinkResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_client_link_service_706cd4e851eebd0b, []int{4}
}
func (m *MutateCustomerClientLinkResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerClientLinkResult.Unmarshal(m, b)
}
func (m *MutateCustomerClientLinkResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerClientLinkResult.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerClientLinkResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerClientLinkResult.Merge(dst, src)
}
func (m *MutateCustomerClientLinkResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerClientLinkResult.Size(m)
}
func (m *MutateCustomerClientLinkResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerClientLinkResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerClientLinkResult proto.InternalMessageInfo

func (m *MutateCustomerClientLinkResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerClientLinkRequest)(nil), "google.ads.googleads.v0.services.GetCustomerClientLinkRequest")
	proto.RegisterType((*MutateCustomerClientLinkRequest)(nil), "google.ads.googleads.v0.services.MutateCustomerClientLinkRequest")
	proto.RegisterType((*CustomerClientLinkOperation)(nil), "google.ads.googleads.v0.services.CustomerClientLinkOperation")
	proto.RegisterType((*MutateCustomerClientLinkResponse)(nil), "google.ads.googleads.v0.services.MutateCustomerClientLinkResponse")
	proto.RegisterType((*MutateCustomerClientLinkResult)(nil), "google.ads.googleads.v0.services.MutateCustomerClientLinkResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClientLinkServiceClient is the client API for CustomerClientLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientLinkServiceClient interface {
	// Returns the requested CustomerClientLink in full detail.
	GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error)
	// Creates or updates a customer client link. Operation statuses are returned.
	MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error)
}

type customerClientLinkServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClientLinkServiceClient(cc *grpc.ClientConn) CustomerClientLinkServiceClient {
	return &customerClientLinkServiceClient{cc}
}

func (c *customerClientLinkServiceClient) GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error) {
	out := new(resources.CustomerClientLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerClientLinkService/GetCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClientLinkServiceClient) MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error) {
	out := new(MutateCustomerClientLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerClientLinkService/MutateCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientLinkServiceServer is the server API for CustomerClientLinkService service.
type CustomerClientLinkServiceServer interface {
	// Returns the requested CustomerClientLink in full detail.
	GetCustomerClientLink(context.Context, *GetCustomerClientLinkRequest) (*resources.CustomerClientLink, error)
	// Creates or updates a customer client link. Operation statuses are returned.
	MutateCustomerClientLink(context.Context, *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error)
}

func RegisterCustomerClientLinkServiceServer(s *grpc.Server, srv CustomerClientLinkServiceServer) {
	s.RegisterService(&_CustomerClientLinkService_serviceDesc, srv)
}

func _CustomerClientLinkService_GetCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerClientLinkService/GetCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, req.(*GetCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerClientLinkService_MutateCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerClientLinkService/MutateCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, req.(*MutateCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.CustomerClientLinkService",
	HandlerType: (*CustomerClientLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClientLink",
			Handler:    _CustomerClientLinkService_GetCustomerClientLink_Handler,
		},
		{
			MethodName: "MutateCustomerClientLink",
			Handler:    _CustomerClientLinkService_MutateCustomerClientLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/customer_client_link_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/customer_client_link_service.proto", fileDescriptor_customer_client_link_service_706cd4e851eebd0b)
}

var fileDescriptor_customer_client_link_service_706cd4e851eebd0b = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x4d, 0x56, 0x0a, 0x3b, 0xd5, 0x4b, 0x40, 0xa8, 0x75, 0xe9, 0x96, 0xb8, 0x87, 0xa5,
	0x87, 0x49, 0xa9, 0x2c, 0x4a, 0xd6, 0x8a, 0x69, 0xd1, 0xae, 0xe0, 0xba, 0x4b, 0x84, 0x22, 0x5a,
	0x28, 0xb3, 0xcd, 0x6c, 0x08, 0x4d, 0x32, 0x71, 0x66, 0x52, 0x0f, 0xeb, 0x5e, 0x04, 0x0f, 0x9e,
	0xfd, 0x00, 0x82, 0x47, 0x3f, 0x8a, 0xe0, 0x49, 0x3f, 0x82, 0x78, 0xf0, 0x53, 0x48, 0x32, 0x33,
	0xe9, 0xea, 0x36, 0xad, 0xec, 0xde, 0x5e, 0x26, 0xef, 0xfd, 0xde, 0xfb, 0xbf, 0x37, 0x6f, 0x40,
	0xdf, 0x27, 0xc4, 0x0f, 0xb1, 0x85, 0x3c, 0x66, 0x09, 0x33, 0xb3, 0x66, 0x6d, 0x8b, 0x61, 0x3a,
	0x0b, 0x26, 0x98, 0x59, 0x93, 0x94, 0x71, 0x12, 0x61, 0x3a, 0x9e, 0x84, 0x01, 0x8e, 0xf9, 0x38,
	0x0c, 0xe2, 0xe9, 0x58, 0xfe, 0x85, 0x09, 0x25, 0x9c, 0x18, 0x4d, 0x11, 0x09, 0x91, 0xc7, 0x60,
	0x01, 0x81, 0xb3, 0x36, 0x54, 0x90, 0xfa, 0xfd, 0xb2, 0x34, 0x14, 0x33, 0x92, 0xd2, 0xb2, 0x3c,
	0x82, 0x5f, 0xdf, 0x50, 0xd1, 0x49, 0x60, 0xa1, 0x38, 0x26, 0x1c, 0xf1, 0x80, 0xc4, 0x4c, 0xfe,
	0x95, 0xd9, 0xad, 0xfc, 0xeb, 0x28, 0x3d, 0xb6, 0x8e, 0x03, 0x1c, 0x7a, 0xe3, 0x08, 0x31, 0x15,
	0xdf, 0xf8, 0xd7, 0xe3, 0x0d, 0x45, 0x49, 0x82, 0xa9, 0x24, 0x98, 0x7d, 0xb0, 0x31, 0xc0, 0xbc,
	0x2f, 0x0b, 0xe8, 0xe7, 0xf9, 0x9f, 0x06, 0xf1, 0xd4, 0xc5, 0xaf, 0x53, 0xcc, 0xb8, 0x71, 0x1b,
	0x5c, 0x57, 0x75, 0x8e, 0x63, 0x14, 0xe1, 0x9a, 0xd6, 0xd4, 0xb6, 0xd7, 0xdd, 0x6b, 0xea, 0xf0,
	0x19, 0x8a, 0xb0, 0xf9, 0x49, 0x03, 0x9b, 0xfb, 0x29, 0x47, 0x1c, 0x97, 0x83, 0x36, 0x41, 0xb5,
	0x90, 0x19, 0x78, 0x12, 0x03, 0xd4, 0xd1, 0x13, 0xcf, 0x78, 0x05, 0xd6, 0x49, 0x82, 0x69, 0xae,
	0xaf, 0xa6, 0x37, 0xb5, 0xed, 0x6a, 0xa7, 0x0b, 0x57, 0x75, 0x17, 0x9e, 0x4f, 0x78, 0xa0, 0x20,
	0xee, 0x9c, 0x67, 0x7e, 0xd0, 0xc1, 0xad, 0x25, 0xae, 0xc6, 0x2e, 0xa8, 0xa6, 0x89, 0x87, 0x38,
	0xce, 0x7b, 0x57, 0xbb, 0x9a, 0xa7, 0xaf, 0xab, 0xf4, 0xaa, 0x79, 0xf0, 0x71, 0xd6, 0xde, 0x7d,
	0xc4, 0xa6, 0x2e, 0x10, 0xee, 0x99, 0x6d, 0x1c, 0x80, 0xca, 0x84, 0x62, 0xc4, 0x45, 0x73, 0xaa,
	0x9d, 0x9d, 0xd2, 0xb2, 0x8b, 0x91, 0x2f, 0xa8, 0x7b, 0xef, 0x8a, 0x2b, 0x31, 0x19, 0x50, 0xe0,
	0x65, 0x1f, 0x2e, 0x0e, 0x14, 0x98, 0x5e, 0xf5, 0x4c, 0x6f, 0xcd, 0xb7, 0xa0, 0x59, 0x3e, 0x2c,
	0x96, 0x90, 0x98, 0x61, 0xe3, 0x05, 0xa8, 0x50, 0xcc, 0xd2, 0x90, 0x4b, 0x49, 0x0f, 0x57, 0x4f,
	0x62, 0x09, 0x33, 0x0d, 0xb9, 0x2b, 0x79, 0xe6, 0x23, 0xd0, 0x58, 0xee, 0xf9, 0x5f, 0x57, 0xae,
	0xf3, 0x63, 0x0d, 0xdc, 0x3c, 0x4f, 0x78, 0x2e, 0x8a, 0x31, 0xbe, 0x69, 0xe0, 0xc6, 0xc2, 0x6b,
	0x6d, 0x3c, 0x58, 0x2d, 0x64, 0xd9, 0x3e, 0xd4, 0x2f, 0x36, 0x0a, 0xb3, 0xfb, 0xee, 0xfb, 0xcf,
	0x8f, 0xfa, 0x5d, 0x63, 0x27, 0x5b, 0xfc, 0x93, 0xbf, 0xe4, 0x75, 0xd5, 0x0e, 0x30, 0xab, 0x55,
	0xbc, 0x04, 0xf3, 0x50, 0x66, 0xb5, 0x4e, 0x8d, 0x5f, 0x1a, 0xa8, 0x95, 0x75, 0xcd, 0x70, 0x2e,
	0x33, 0x1b, 0xa1, 0xaa, 0x77, 0xa9, 0xf1, 0xe6, 0x57, 0xc6, 0xec, 0xe7, 0x12, 0xbb, 0xe6, 0xbd,
	0x4c, 0xe2, 0x5c, 0xd3, 0xc9, 0x99, 0xad, 0xef, 0xb6, 0x4e, 0x17, 0x29, 0xb4, 0xa3, 0x9c, 0x6d,
	0x6b, 0xad, 0xde, 0x7b, 0x1d, 0x6c, 0x4d, 0x48, 0xb4, 0xb2, 0x9c, 0x5e, 0xa3, 0x74, 0xf8, 0x87,
	0xd9, 0xb2, 0x1e, 0x6a, 0x2f, 0xf7, 0x24, 0xc3, 0x27, 0x21, 0x8a, 0x7d, 0x48, 0xa8, 0x6f, 0xf9,
	0x38, 0xce, 0x57, 0x59, 0xbd, 0xc3, 0x49, 0xc0, 0xca, 0x5f, 0xff, 0x5d, 0x65, 0x7c, 0xd6, 0xd7,
	0x06, 0x8e, 0xf3, 0x45, 0x6f, 0x0e, 0x04, 0xd0, 0xf1, 0x18, 0x14, 0x66, 0x66, 0x0d, 0xdb, 0x50,
	0x26, 0x66, 0x5f, 0x95, 0xcb, 0xc8, 0xf1, 0xd8, 0xa8, 0x70, 0x19, 0x0d, 0xdb, 0x23, 0xe5, 0xf2,
	0x5b, 0xdf, 0x12, 0xe7, 0xb6, 0xed, 0x78, 0xcc, 0xb6, 0x0b, 0x27, 0xdb, 0x1e, 0xb6, 0x6d, 0x5b,
	0xb9, 0x1d, 0x55, 0xf2, 0x3a, 0xef, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x00, 0x59, 0x09,
	0xa4, 0x06, 0x00, 0x00,
}
