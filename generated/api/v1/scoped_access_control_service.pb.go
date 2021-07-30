// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/scoped_access_control_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	storage "github.com/stackrox/rox/generated/storage"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetAuthzPluginConfigsResponse struct {
	Configs              []*storage.AuthzPluginConfig `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetAuthzPluginConfigsResponse) Reset()         { *m = GetAuthzPluginConfigsResponse{} }
func (m *GetAuthzPluginConfigsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthzPluginConfigsResponse) ProtoMessage()    {}
func (*GetAuthzPluginConfigsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecec806f567593cb, []int{0}
}
func (m *GetAuthzPluginConfigsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthzPluginConfigsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthzPluginConfigsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthzPluginConfigsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthzPluginConfigsResponse.Merge(m, src)
}
func (m *GetAuthzPluginConfigsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthzPluginConfigsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthzPluginConfigsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthzPluginConfigsResponse proto.InternalMessageInfo

func (m *GetAuthzPluginConfigsResponse) GetConfigs() []*storage.AuthzPluginConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *GetAuthzPluginConfigsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetAuthzPluginConfigsResponse) Clone() *GetAuthzPluginConfigsResponse {
	if m == nil {
		return nil
	}
	cloned := new(GetAuthzPluginConfigsResponse)
	*cloned = *m

	if m.Configs != nil {
		cloned.Configs = make([]*storage.AuthzPluginConfig, len(m.Configs))
		for idx, v := range m.Configs {
			cloned.Configs[idx] = v.Clone()
		}
	}
	return cloned
}

type UpsertAuthzPluginConfigRequest struct {
	Config *storage.AuthzPluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// When false, use the stored credentials of an existing scoped access control configuration given its ID.
	UpdatePassword       bool     `protobuf:"varint,2,opt,name=update_password,json=updatePassword,proto3" json:"update_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertAuthzPluginConfigRequest) Reset()         { *m = UpsertAuthzPluginConfigRequest{} }
func (m *UpsertAuthzPluginConfigRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertAuthzPluginConfigRequest) ProtoMessage()    {}
func (*UpsertAuthzPluginConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecec806f567593cb, []int{1}
}
func (m *UpsertAuthzPluginConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpsertAuthzPluginConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpsertAuthzPluginConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpsertAuthzPluginConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertAuthzPluginConfigRequest.Merge(m, src)
}
func (m *UpsertAuthzPluginConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpsertAuthzPluginConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertAuthzPluginConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertAuthzPluginConfigRequest proto.InternalMessageInfo

func (m *UpsertAuthzPluginConfigRequest) GetConfig() *storage.AuthzPluginConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *UpsertAuthzPluginConfigRequest) GetUpdatePassword() bool {
	if m != nil {
		return m.UpdatePassword
	}
	return false
}

func (m *UpsertAuthzPluginConfigRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *UpsertAuthzPluginConfigRequest) Clone() *UpsertAuthzPluginConfigRequest {
	if m == nil {
		return nil
	}
	cloned := new(UpsertAuthzPluginConfigRequest)
	*cloned = *m

	cloned.Config = m.Config.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*GetAuthzPluginConfigsResponse)(nil), "v1.GetAuthzPluginConfigsResponse")
	proto.RegisterType((*UpsertAuthzPluginConfigRequest)(nil), "v1.UpsertAuthzPluginConfigRequest")
}

func init() {
	proto.RegisterFile("api/v1/scoped_access_control_service.proto", fileDescriptor_ecec806f567593cb)
}

var fileDescriptor_ecec806f567593cb = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x9d, 0x8b, 0x18, 0x60, 0x24, 0x40, 0x06, 0xb4, 0x12, 0x46, 0xd4, 0x45, 0xc0, 0xaa, 0x82,
	0x1c, 0xb5, 0x70, 0xda, 0xad, 0x5b, 0x11, 0x42, 0x5c, 0x46, 0xa6, 0x4a, 0x68, 0x97, 0x60, 0x9c,
	0x8f, 0x12, 0xad, 0x8d, 0x8d, 0xed, 0x04, 0x0a, 0xea, 0x05, 0x7e, 0x02, 0x17, 0x24, 0xfe, 0x0c,
	0x47, 0x8e, 0x08, 0xfe, 0x00, 0x2a, 0xfc, 0x10, 0x94, 0xd8, 0x63, 0x48, 0x2d, 0x01, 0x21, 0x4e,
	0x89, 0x9e, 0x9f, 0xbf, 0xf7, 0xbe, 0xf7, 0x12, 0xdc, 0x61, 0x32, 0x0d, 0x8b, 0x6e, 0xa8, 0xb9,
	0x90, 0x90, 0xc4, 0x8c, 0x73, 0xd0, 0x3a, 0xe6, 0x22, 0x33, 0x4a, 0x8c, 0x63, 0x0d, 0xaa, 0x48,
	0x39, 0x50, 0xa9, 0x84, 0x11, 0xa4, 0x51, 0x74, 0xbd, 0xf5, 0x91, 0x10, 0xa3, 0x31, 0x84, 0xe5,
	0x35, 0x96, 0x65, 0xc2, 0x30, 0x93, 0x8a, 0x4c, 0x5b, 0x86, 0x47, 0xdc, 0x34, 0x98, 0x48, 0x33,
	0x75, 0xd8, 0x79, 0x87, 0x71, 0x31, 0x99, 0x88, 0xcc, 0x81, 0x9e, 0x36, 0x42, 0xb1, 0x11, 0x84,
	0x2c, 0x37, 0x4f, 0x5f, 0xc6, 0x72, 0x9c, 0x8f, 0x52, 0x77, 0x16, 0x0c, 0xf1, 0x95, 0xbb, 0x60,
	0xfa, 0xe5, 0xc1, 0x6e, 0x85, 0xef, 0x88, 0xec, 0x49, 0x3a, 0xd2, 0x11, 0x68, 0x29, 0x32, 0x0d,
	0xe4, 0x36, 0x3e, 0xc1, 0x2d, 0xd4, 0x44, 0xad, 0x63, 0xed, 0xd3, 0x3d, 0x8f, 0xba, 0x71, 0x74,
	0xe1, 0x56, 0x74, 0x48, 0x0d, 0x66, 0xd8, 0x1f, 0x4a, 0x0d, 0x6a, 0x71, 0x72, 0x04, 0xcf, 0x72,
	0xd0, 0x86, 0xf4, 0xf0, 0xaa, 0x25, 0x37, 0x51, 0x0b, 0xfd, 0x61, 0xac, 0x63, 0x92, 0x4d, 0x7c,
	0x36, 0x97, 0x09, 0x33, 0x10, 0x4b, 0xa6, 0xf5, 0x73, 0xa1, 0x92, 0x66, 0xa3, 0x85, 0xda, 0x27,
	0xa3, 0x33, 0x16, 0xde, 0x75, 0x68, 0xef, 0xf3, 0x71, 0xec, 0xed, 0x55, 0x21, 0xf7, 0xab, 0x8c,
	0x77, 0x6c, 0xc4, 0x7b, 0x36, 0x61, 0xf2, 0x1e, 0xe1, 0xb5, 0x81, 0x9a, 0x46, 0x79, 0xb6, 0xa0,
	0x45, 0x02, 0x5a, 0x74, 0x69, 0xbd, 0x77, 0xef, 0x54, 0xc9, 0xb9, 0x53, 0xc6, 0x1e, 0x3c, 0x78,
	0xfd, 0xe5, 0xfb, 0xdb, 0xc6, 0xfd, 0xe0, 0xd2, 0x51, 0xaf, 0xb6, 0x56, 0x6e, 0xd4, 0x38, 0x34,
	0xa0, 0xcd, 0x96, 0x73, 0xbf, 0x7f, 0x3d, 0xd8, 0xf8, 0x2d, 0x29, 0xb4, 0x0b, 0x24, 0x5b, 0xa8,
	0x43, 0x0e, 0xf0, 0xc5, 0xa5, 0x95, 0x90, 0x23, 0x59, 0x6f, 0xa3, 0x7c, 0xad, 0x2d, 0x2e, 0xb8,
	0x5a, 0x39, 0xf3, 0xc9, 0xfa, 0x52, 0x51, 0x57, 0x14, 0x79, 0x83, 0xf0, 0x85, 0x7e, 0x92, 0xfc,
	0x5b, 0x0e, 0x35, 0x9d, 0x05, 0x37, 0x2a, 0xf9, 0x6b, 0xc1, 0xe5, 0x1a, 0xf9, 0xc3, 0x68, 0xc8,
	0x07, 0x84, 0xd7, 0x86, 0x55, 0x02, 0xff, 0xdf, 0x08, 0x54, 0x46, 0x62, 0x6f, 0xb3, 0xc6, 0x48,
	0xf8, 0xca, 0x3e, 0x69, 0x9a, 0xcc, 0x7e, 0xf6, 0x75, 0xb3, 0xf7, 0xd7, 0x57, 0x50, 0x87, 0x3c,
	0xc2, 0x6b, 0x03, 0x18, 0xc3, 0xb2, 0x0d, 0xce, 0x95, 0x1b, 0x44, 0xa0, 0x45, 0xae, 0x38, 0x6c,
	0x4f, 0xef, 0x0d, 0x7e, 0xfd, 0x80, 0xda, 0x95, 0xbd, 0xa0, 0xd3, 0xaa, 0xd5, 0x4a, 0x93, 0xd9,
	0x36, 0xfd, 0x38, 0xf7, 0xd1, 0xa7, 0xb9, 0x8f, 0xbe, 0xce, 0x7d, 0xf4, 0xee, 0x9b, 0xbf, 0x82,
	0x9b, 0xa9, 0xa0, 0xda, 0x30, 0x7e, 0xa0, 0xc4, 0x0b, 0xfb, 0x3f, 0x53, 0x26, 0x53, 0x5a, 0x74,
	0xf7, 0x1b, 0x45, 0xf7, 0xe1, 0xca, 0xe3, 0xd5, 0x0a, 0xbb, 0xf5, 0x23, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0x73, 0xd9, 0xcf, 0x78, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ScopedAccessControlServiceClient is the client API for ScopedAccessControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type ScopedAccessControlServiceClient interface {
	// DryRunAuthzPluginConfig checks if the given scoped access control plugin is correctly configured.
	DryRunAuthzPluginConfig(ctx context.Context, in *UpsertAuthzPluginConfigRequest, opts ...grpc.CallOption) (*Empty, error)
	// GetAuthzPluginConfigs returns all scoped access control plugins.
	GetAuthzPluginConfigs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAuthzPluginConfigsResponse, error)
	// AddAuthzPluginConfig creates a scoped access control plugin.
	AddAuthzPluginConfig(ctx context.Context, in *UpsertAuthzPluginConfigRequest, opts ...grpc.CallOption) (*storage.AuthzPluginConfig, error)
	// UpdateAuthzPluginConfig modifies a scoped access control plugin.
	UpdateAuthzPluginConfig(ctx context.Context, in *UpsertAuthzPluginConfigRequest, opts ...grpc.CallOption) (*storage.AuthzPluginConfig, error)
	// DeleteAuthzPluginConfig removes a scoped access control plugin.
	DeleteAuthzPluginConfig(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*Empty, error)
}

type scopedAccessControlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScopedAccessControlServiceClient(cc grpc.ClientConnInterface) ScopedAccessControlServiceClient {
	return &scopedAccessControlServiceClient{cc}
}

func (c *scopedAccessControlServiceClient) DryRunAuthzPluginConfig(ctx context.Context, in *UpsertAuthzPluginConfigRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/v1.ScopedAccessControlService/DryRunAuthzPluginConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopedAccessControlServiceClient) GetAuthzPluginConfigs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAuthzPluginConfigsResponse, error) {
	out := new(GetAuthzPluginConfigsResponse)
	err := c.cc.Invoke(ctx, "/v1.ScopedAccessControlService/GetAuthzPluginConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopedAccessControlServiceClient) AddAuthzPluginConfig(ctx context.Context, in *UpsertAuthzPluginConfigRequest, opts ...grpc.CallOption) (*storage.AuthzPluginConfig, error) {
	out := new(storage.AuthzPluginConfig)
	err := c.cc.Invoke(ctx, "/v1.ScopedAccessControlService/AddAuthzPluginConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopedAccessControlServiceClient) UpdateAuthzPluginConfig(ctx context.Context, in *UpsertAuthzPluginConfigRequest, opts ...grpc.CallOption) (*storage.AuthzPluginConfig, error) {
	out := new(storage.AuthzPluginConfig)
	err := c.cc.Invoke(ctx, "/v1.ScopedAccessControlService/UpdateAuthzPluginConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopedAccessControlServiceClient) DeleteAuthzPluginConfig(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/v1.ScopedAccessControlService/DeleteAuthzPluginConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScopedAccessControlServiceServer is the server API for ScopedAccessControlService service.
type ScopedAccessControlServiceServer interface {
	// DryRunAuthzPluginConfig checks if the given scoped access control plugin is correctly configured.
	DryRunAuthzPluginConfig(context.Context, *UpsertAuthzPluginConfigRequest) (*Empty, error)
	// GetAuthzPluginConfigs returns all scoped access control plugins.
	GetAuthzPluginConfigs(context.Context, *Empty) (*GetAuthzPluginConfigsResponse, error)
	// AddAuthzPluginConfig creates a scoped access control plugin.
	AddAuthzPluginConfig(context.Context, *UpsertAuthzPluginConfigRequest) (*storage.AuthzPluginConfig, error)
	// UpdateAuthzPluginConfig modifies a scoped access control plugin.
	UpdateAuthzPluginConfig(context.Context, *UpsertAuthzPluginConfigRequest) (*storage.AuthzPluginConfig, error)
	// DeleteAuthzPluginConfig removes a scoped access control plugin.
	DeleteAuthzPluginConfig(context.Context, *ResourceByID) (*Empty, error)
}

// UnimplementedScopedAccessControlServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScopedAccessControlServiceServer struct {
}

func (*UnimplementedScopedAccessControlServiceServer) DryRunAuthzPluginConfig(ctx context.Context, req *UpsertAuthzPluginConfigRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DryRunAuthzPluginConfig not implemented")
}
func (*UnimplementedScopedAccessControlServiceServer) GetAuthzPluginConfigs(ctx context.Context, req *Empty) (*GetAuthzPluginConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthzPluginConfigs not implemented")
}
func (*UnimplementedScopedAccessControlServiceServer) AddAuthzPluginConfig(ctx context.Context, req *UpsertAuthzPluginConfigRequest) (*storage.AuthzPluginConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAuthzPluginConfig not implemented")
}
func (*UnimplementedScopedAccessControlServiceServer) UpdateAuthzPluginConfig(ctx context.Context, req *UpsertAuthzPluginConfigRequest) (*storage.AuthzPluginConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthzPluginConfig not implemented")
}
func (*UnimplementedScopedAccessControlServiceServer) DeleteAuthzPluginConfig(ctx context.Context, req *ResourceByID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthzPluginConfig not implemented")
}

func RegisterScopedAccessControlServiceServer(s *grpc.Server, srv ScopedAccessControlServiceServer) {
	s.RegisterService(&_ScopedAccessControlService_serviceDesc, srv)
}

func _ScopedAccessControlService_DryRunAuthzPluginConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAuthzPluginConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedAccessControlServiceServer).DryRunAuthzPluginConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ScopedAccessControlService/DryRunAuthzPluginConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedAccessControlServiceServer).DryRunAuthzPluginConfig(ctx, req.(*UpsertAuthzPluginConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopedAccessControlService_GetAuthzPluginConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedAccessControlServiceServer).GetAuthzPluginConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ScopedAccessControlService/GetAuthzPluginConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedAccessControlServiceServer).GetAuthzPluginConfigs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopedAccessControlService_AddAuthzPluginConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAuthzPluginConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedAccessControlServiceServer).AddAuthzPluginConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ScopedAccessControlService/AddAuthzPluginConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedAccessControlServiceServer).AddAuthzPluginConfig(ctx, req.(*UpsertAuthzPluginConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopedAccessControlService_UpdateAuthzPluginConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAuthzPluginConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedAccessControlServiceServer).UpdateAuthzPluginConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ScopedAccessControlService/UpdateAuthzPluginConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedAccessControlServiceServer).UpdateAuthzPluginConfig(ctx, req.(*UpsertAuthzPluginConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopedAccessControlService_DeleteAuthzPluginConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedAccessControlServiceServer).DeleteAuthzPluginConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ScopedAccessControlService/DeleteAuthzPluginConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedAccessControlServiceServer).DeleteAuthzPluginConfig(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScopedAccessControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ScopedAccessControlService",
	HandlerType: (*ScopedAccessControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DryRunAuthzPluginConfig",
			Handler:    _ScopedAccessControlService_DryRunAuthzPluginConfig_Handler,
		},
		{
			MethodName: "GetAuthzPluginConfigs",
			Handler:    _ScopedAccessControlService_GetAuthzPluginConfigs_Handler,
		},
		{
			MethodName: "AddAuthzPluginConfig",
			Handler:    _ScopedAccessControlService_AddAuthzPluginConfig_Handler,
		},
		{
			MethodName: "UpdateAuthzPluginConfig",
			Handler:    _ScopedAccessControlService_UpdateAuthzPluginConfig_Handler,
		},
		{
			MethodName: "DeleteAuthzPluginConfig",
			Handler:    _ScopedAccessControlService_DeleteAuthzPluginConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/scoped_access_control_service.proto",
}

func (m *GetAuthzPluginConfigsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthzPluginConfigsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthzPluginConfigsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Configs) > 0 {
		for iNdEx := len(m.Configs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Configs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintScopedAccessControlService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UpsertAuthzPluginConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpsertAuthzPluginConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpsertAuthzPluginConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdatePassword {
		i--
		if m.UpdatePassword {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScopedAccessControlService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScopedAccessControlService(dAtA []byte, offset int, v uint64) int {
	offset -= sovScopedAccessControlService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetAuthzPluginConfigsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for _, e := range m.Configs {
			l = e.Size()
			n += 1 + l + sovScopedAccessControlService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpsertAuthzPluginConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovScopedAccessControlService(uint64(l))
	}
	if m.UpdatePassword {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovScopedAccessControlService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScopedAccessControlService(x uint64) (n int) {
	return sovScopedAccessControlService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetAuthzPluginConfigsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScopedAccessControlService
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
			return fmt.Errorf("proto: GetAuthzPluginConfigsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthzPluginConfigsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScopedAccessControlService
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
				return ErrInvalidLengthScopedAccessControlService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScopedAccessControlService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Configs = append(m.Configs, &storage.AuthzPluginConfig{})
			if err := m.Configs[len(m.Configs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScopedAccessControlService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScopedAccessControlService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpsertAuthzPluginConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScopedAccessControlService
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
			return fmt.Errorf("proto: UpsertAuthzPluginConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpsertAuthzPluginConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScopedAccessControlService
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
				return ErrInvalidLengthScopedAccessControlService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScopedAccessControlService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &storage.AuthzPluginConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatePassword", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScopedAccessControlService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UpdatePassword = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipScopedAccessControlService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScopedAccessControlService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipScopedAccessControlService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScopedAccessControlService
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
					return 0, ErrIntOverflowScopedAccessControlService
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
					return 0, ErrIntOverflowScopedAccessControlService
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
				return 0, ErrInvalidLengthScopedAccessControlService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScopedAccessControlService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScopedAccessControlService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScopedAccessControlService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScopedAccessControlService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScopedAccessControlService = fmt.Errorf("proto: unexpected end of group")
)
