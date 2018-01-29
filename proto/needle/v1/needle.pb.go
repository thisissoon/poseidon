// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/needle/v1/needle.proto

/*
Package needle is a generated protocol buffer package.

It is generated from these files:
	proto/needle/v1/needle.proto

It has these top-level messages:
	Service
	Dentry
	PutServiceRequest
	PutServiceResponse
	PutDentryRequest
	PutDentryResponse
	RouteToVersionRequest
	RouteToVersionResponse
*/
package needle

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

// A Service is a service that compass is managing a namerd dentry for
type Service struct {
	// Id - the primary key - but not used for upserts
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// User created date
	CreateDate *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=create_date,json=createDate" json:"create_date,omitempty"`
	// User updated date
	UpdateDate *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=update_date,json=updateDate" json:"update_date,omitempty"`
	// The logical name, e.g HTTP/1.1 Host name or HTTP/2 :authority
	LogicalName string `protobuf:"bytes,4,opt,name=logical_name,json=logicalName" json:"logical_name,omitempty"`
	// The kubernetets namespace the app runs in
	Namespace string `protobuf:"bytes,5,opt,name=namespace" json:"namespace,omitempty"`
	// Optional description of the service
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetCreateDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *Service) GetUpdateDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdateDate
	}
	return nil
}

func (m *Service) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

func (m *Service) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// A Dentry is a single Delegation Table Rule
type Dentry struct {
	// Id - the primary key - but not used for upserts
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Optional description of the service
	ServiceId string `protobuf:"bytes,2,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	// User created date
	CreateDate *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=create_date,json=createDate" json:"create_date,omitempty"`
	// User updated date
	UpdateDate *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=update_date,json=updateDate" json:"update_date,omitempty"`
	// The prefix to match on
	Prefix string `protobuf:"bytes,5,opt,name=prefix" json:"prefix,omitempty"`
	// The destination of the matched prefix
	Destination string `protobuf:"bytes,6,opt,name=destination" json:"destination,omitempty"`
	// The dentry dtab (k8s namespace)
	Dtab string `protobuf:"bytes,7,opt,name=dtab" json:"dtab,omitempty"`
	// Prioity of the dentry - affects the ordering within the delegation table
	Priority int32 `protobuf:"varint,8,opt,name=priority" json:"priority,omitempty"`
}

func (m *Dentry) Reset()                    { *m = Dentry{} }
func (m *Dentry) String() string            { return proto.CompactTextString(m) }
func (*Dentry) ProtoMessage()               {}
func (*Dentry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Dentry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Dentry) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *Dentry) GetCreateDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *Dentry) GetUpdateDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdateDate
	}
	return nil
}

func (m *Dentry) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Dentry) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Dentry) GetDtab() string {
	if m != nil {
		return m.Dtab
	}
	return ""
}

func (m *Dentry) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

// Request for [PutService](compass.needle.v1.NeedleService.PutService)
type PutServiceRequest struct {
	// Service to be created or updated
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *PutServiceRequest) Reset()                    { *m = PutServiceRequest{} }
func (m *PutServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*PutServiceRequest) ProtoMessage()               {}
func (*PutServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PutServiceRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

// Response for [PutService](compass.needle.v1.NeedleService.PutService)
type PutServiceResponse struct {
	// Service crated or updated
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *PutServiceResponse) Reset()                    { *m = PutServiceResponse{} }
func (m *PutServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*PutServiceResponse) ProtoMessage()               {}
func (*PutServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PutServiceResponse) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

// Request for [PutDentry](compass.needle.v1.NeedleService.PutDentry)
type PutDentryRequest struct {
	// Dentry to create or update
	Dentry *Dentry `protobuf:"bytes,1,opt,name=dentry" json:"dentry,omitempty"`
}

func (m *PutDentryRequest) Reset()                    { *m = PutDentryRequest{} }
func (m *PutDentryRequest) String() string            { return proto.CompactTextString(m) }
func (*PutDentryRequest) ProtoMessage()               {}
func (*PutDentryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PutDentryRequest) GetDentry() *Dentry {
	if m != nil {
		return m.Dentry
	}
	return nil
}

// Response for [PutDentry](compass.needle.v1.NeedleService.PutDentry)
type PutDentryResponse struct {
	// Dentry crated or updated
	Dentry *Dentry `protobuf:"bytes,1,opt,name=dentry" json:"dentry,omitempty"`
}

func (m *PutDentryResponse) Reset()                    { *m = PutDentryResponse{} }
func (m *PutDentryResponse) String() string            { return proto.CompactTextString(m) }
func (*PutDentryResponse) ProtoMessage()               {}
func (*PutDentryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PutDentryResponse) GetDentry() *Dentry {
	if m != nil {
		return m.Dentry
	}
	return nil
}

// Request for [RouteToVersion](compass.needle.v1.NeedleService.RouteToVersion)
type RouteToVersionRequest struct {
	// Logical name of the service to route
	LogicalName string `protobuf:"bytes,1,opt,name=logical_name,json=logicalName" json:"logical_name,omitempty"`
	// Version string of the kubernetets service to route too
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *RouteToVersionRequest) Reset()                    { *m = RouteToVersionRequest{} }
func (m *RouteToVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*RouteToVersionRequest) ProtoMessage()               {}
func (*RouteToVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RouteToVersionRequest) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

func (m *RouteToVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Response for [RouteToVersion](compass.needle.v1.NeedleService.RouteToVersion)
type RouteToVersionResponse struct {
}

func (m *RouteToVersionResponse) Reset()                    { *m = RouteToVersionResponse{} }
func (m *RouteToVersionResponse) String() string            { return proto.CompactTextString(m) }
func (*RouteToVersionResponse) ProtoMessage()               {}
func (*RouteToVersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Service)(nil), "compass.needle.v1.Service")
	proto.RegisterType((*Dentry)(nil), "compass.needle.v1.Dentry")
	proto.RegisterType((*PutServiceRequest)(nil), "compass.needle.v1.PutServiceRequest")
	proto.RegisterType((*PutServiceResponse)(nil), "compass.needle.v1.PutServiceResponse")
	proto.RegisterType((*PutDentryRequest)(nil), "compass.needle.v1.PutDentryRequest")
	proto.RegisterType((*PutDentryResponse)(nil), "compass.needle.v1.PutDentryResponse")
	proto.RegisterType((*RouteToVersionRequest)(nil), "compass.needle.v1.RouteToVersionRequest")
	proto.RegisterType((*RouteToVersionResponse)(nil), "compass.needle.v1.RouteToVersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NeedleService service

type NeedleServiceClient interface {
	// PutService upserts services to be managed
	PutService(ctx context.Context, in *PutServiceRequest, opts ...grpc.CallOption) (*PutServiceResponse, error)
	// PutDentry upserts dentries to be managed
	PutDentry(ctx context.Context, in *PutDentryRequest, opts ...grpc.CallOption) (*PutDentryResponse, error)
	// RouteToVersion routes a specific logical name to a specifc lubernetets services
	// based on a services logicalName and version labels
	RouteToVersion(ctx context.Context, in *RouteToVersionRequest, opts ...grpc.CallOption) (*RouteToVersionResponse, error)
}

type needleServiceClient struct {
	cc *grpc.ClientConn
}

func NewNeedleServiceClient(cc *grpc.ClientConn) NeedleServiceClient {
	return &needleServiceClient{cc}
}

func (c *needleServiceClient) PutService(ctx context.Context, in *PutServiceRequest, opts ...grpc.CallOption) (*PutServiceResponse, error) {
	out := new(PutServiceResponse)
	err := grpc.Invoke(ctx, "/compass.needle.v1.NeedleService/PutService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *needleServiceClient) PutDentry(ctx context.Context, in *PutDentryRequest, opts ...grpc.CallOption) (*PutDentryResponse, error) {
	out := new(PutDentryResponse)
	err := grpc.Invoke(ctx, "/compass.needle.v1.NeedleService/PutDentry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *needleServiceClient) RouteToVersion(ctx context.Context, in *RouteToVersionRequest, opts ...grpc.CallOption) (*RouteToVersionResponse, error) {
	out := new(RouteToVersionResponse)
	err := grpc.Invoke(ctx, "/compass.needle.v1.NeedleService/RouteToVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NeedleService service

type NeedleServiceServer interface {
	// PutService upserts services to be managed
	PutService(context.Context, *PutServiceRequest) (*PutServiceResponse, error)
	// PutDentry upserts dentries to be managed
	PutDentry(context.Context, *PutDentryRequest) (*PutDentryResponse, error)
	// RouteToVersion routes a specific logical name to a specifc lubernetets services
	// based on a services logicalName and version labels
	RouteToVersion(context.Context, *RouteToVersionRequest) (*RouteToVersionResponse, error)
}

func RegisterNeedleServiceServer(s *grpc.Server, srv NeedleServiceServer) {
	s.RegisterService(&_NeedleService_serviceDesc, srv)
}

func _NeedleService_PutService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeedleServiceServer).PutService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compass.needle.v1.NeedleService/PutService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeedleServiceServer).PutService(ctx, req.(*PutServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NeedleService_PutDentry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutDentryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeedleServiceServer).PutDentry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compass.needle.v1.NeedleService/PutDentry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeedleServiceServer).PutDentry(ctx, req.(*PutDentryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NeedleService_RouteToVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteToVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeedleServiceServer).RouteToVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compass.needle.v1.NeedleService/RouteToVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeedleServiceServer).RouteToVersion(ctx, req.(*RouteToVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NeedleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "compass.needle.v1.NeedleService",
	HandlerType: (*NeedleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutService",
			Handler:    _NeedleService_PutService_Handler,
		},
		{
			MethodName: "PutDentry",
			Handler:    _NeedleService_PutDentry_Handler,
		},
		{
			MethodName: "RouteToVersion",
			Handler:    _NeedleService_RouteToVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/needle/v1/needle.proto",
}

func init() { proto.RegisterFile("proto/needle/v1/needle.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x25, 0x59, 0x97, 0x2e, 0xb7, 0x30, 0x6d, 0x96, 0x98, 0x42, 0x34, 0x44, 0x09, 0x43, 0x2a,
	0x2f, 0xa9, 0x3a, 0x78, 0xdb, 0x1b, 0x1a, 0x48, 0xe3, 0x61, 0x9a, 0x42, 0x85, 0x10, 0x3c, 0x54,
	0x6e, 0x72, 0x57, 0x59, 0x6a, 0xe2, 0x10, 0x3b, 0x15, 0xfb, 0x2d, 0xfc, 0x08, 0xfe, 0x1d, 0xcf,
	0xa8, 0xfe, 0xc8, 0xb6, 0x26, 0x62, 0x30, 0xde, 0xec, 0xeb, 0x73, 0x4e, 0xce, 0x3d, 0xbe, 0x0e,
	0x1c, 0x96, 0x15, 0x97, 0x7c, 0x5c, 0x20, 0x66, 0x4b, 0x1c, 0xaf, 0x26, 0x66, 0x15, 0xab, 0x32,
	0xd9, 0x4f, 0x79, 0x5e, 0x52, 0x21, 0x62, 0x53, 0x5d, 0x4d, 0xc2, 0x67, 0x0b, 0xce, 0x17, 0x4b,
	0x1c, 0x2b, 0xc0, 0xbc, 0xbe, 0x1c, 0x4b, 0x96, 0xa3, 0x90, 0x34, 0x2f, 0x35, 0x27, 0xfa, 0xe5,
	0x40, 0xff, 0x23, 0x56, 0x2b, 0x96, 0x22, 0xd9, 0x05, 0x97, 0x65, 0x81, 0x33, 0x74, 0x46, 0x7e,
	0xe2, 0xb2, 0x8c, 0x9c, 0xc0, 0x20, 0xad, 0x90, 0x4a, 0x9c, 0x65, 0x54, 0x62, 0xe0, 0x0e, 0x9d,
	0xd1, 0xe0, 0x38, 0x8c, 0xb5, 0x64, 0x6c, 0x25, 0xe3, 0xa9, 0x95, 0x4c, 0x40, 0xc3, 0x4f, 0xa9,
	0xc4, 0x35, 0xb9, 0x2e, 0xb3, 0x86, 0xbc, 0x75, 0x37, 0x59, 0xc3, 0x15, 0xf9, 0x39, 0x3c, 0x5c,
	0xf2, 0x05, 0x4b, 0xe9, 0x72, 0x56, 0xd0, 0x1c, 0x83, 0x9e, 0xf2, 0x34, 0x30, 0xb5, 0x73, 0x9a,
	0x23, 0x39, 0x04, 0x7f, 0x7d, 0x24, 0x4a, 0x9a, 0x62, 0xb0, 0xad, 0xce, 0xaf, 0x0b, 0x64, 0x08,
	0x83, 0x0c, 0x45, 0x5a, 0xb1, 0x52, 0x32, 0x5e, 0x04, 0x9e, 0xe6, 0xdf, 0x28, 0x45, 0x3f, 0x5c,
	0xf0, 0x4e, 0xb1, 0x90, 0xd5, 0x55, 0xab, 0xef, 0xa7, 0x00, 0x42, 0x47, 0x32, 0x63, 0x99, 0x6a,
	0xdb, 0x4f, 0x7c, 0x53, 0x39, 0x6b, 0xc5, 0xb2, 0xf5, 0x3f, 0xb1, 0xf4, 0xfe, 0x29, 0x96, 0x03,
	0xf0, 0xca, 0x0a, 0x2f, 0xd9, 0x77, 0xd3, 0xb0, 0xd9, 0x99, 0x6e, 0x25, 0x2b, 0xe8, 0x46, 0xb7,
	0xb6, 0x44, 0x08, 0xf4, 0x32, 0x49, 0xe7, 0x41, 0x5f, 0x1d, 0xa9, 0x35, 0x09, 0x61, 0xa7, 0xac,
	0x18, 0xaf, 0x98, 0xbc, 0x0a, 0x76, 0x86, 0xce, 0x68, 0x3b, 0x69, 0xf6, 0xd1, 0x19, 0xec, 0x5f,
	0xd4, 0xd2, 0x0c, 0x46, 0x82, 0xdf, 0x6a, 0x14, 0x92, 0xbc, 0x81, 0xbe, 0x49, 0x41, 0x85, 0xb5,
	0xf6, 0xdd, 0x9a, 0xb8, 0xd8, 0x72, 0x2c, 0x34, 0xfa, 0x00, 0xe4, 0xa6, 0x94, 0x28, 0x79, 0x21,
	0xf0, 0x9e, 0x5a, 0xef, 0x60, 0xef, 0xa2, 0x96, 0xfa, 0xda, 0xac, 0xab, 0x09, 0x78, 0x99, 0x2a,
	0x18, 0xa1, 0x27, 0x1d, 0x42, 0x86, 0x61, 0x80, 0xd1, 0x7b, 0xd5, 0x9d, 0x95, 0x31, 0x8e, 0xee,
	0xa1, 0x33, 0x85, 0xc7, 0x09, 0xaf, 0x25, 0x4e, 0xf9, 0x27, 0xac, 0x04, 0xe3, 0x85, 0xf5, 0xb4,
	0x39, 0xbf, 0x4e, 0x7b, 0x7e, 0x03, 0xe8, 0xaf, 0x34, 0x49, 0x4d, 0x90, 0x9f, 0xd8, 0x6d, 0x14,
	0xc0, 0xc1, 0xa6, 0xaa, 0xb6, 0x78, 0xfc, 0xd3, 0x85, 0x47, 0xe7, 0xca, 0x8c, 0x7d, 0xb2, 0x5f,
	0x01, 0xae, 0xc3, 0x25, 0x47, 0x1d, 0x96, 0x5b, 0xd7, 0x18, 0xbe, 0xbc, 0x03, 0xa5, 0x3f, 0x16,
	0x3d, 0x20, 0x9f, 0xc1, 0x6f, 0x62, 0x22, 0x2f, 0xba, 0x59, 0xb7, 0xee, 0x22, 0x3c, 0xfa, 0x33,
	0xa8, 0x51, 0x5e, 0xc0, 0xee, 0xed, 0x16, 0xc9, 0xa8, 0x83, 0xd9, 0x99, 0x6d, 0xf8, 0xea, 0x2f,
	0x90, 0xf6, 0x43, 0x6f, 0xc9, 0x97, 0xbd, 0xe6, 0x67, 0x79, 0xa2, 0x57, 0x73, 0x4f, 0xbd, 0xb2,
	0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x60, 0x6d, 0xe4, 0xeb, 0x4d, 0x05, 0x00, 0x00,
}
