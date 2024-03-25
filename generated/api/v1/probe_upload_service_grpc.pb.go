// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/probe_upload_service.proto

package v1

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
	ProbeUploadService_GetExistingProbes_FullMethodName = "/v1.ProbeUploadService/GetExistingProbes"
)

// ProbeUploadServiceClient is the client API for ProbeUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProbeUploadServiceClient interface {
	GetExistingProbes(ctx context.Context, in *GetExistingProbesRequest, opts ...grpc.CallOption) (*GetExistingProbesResponse, error)
}

type probeUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProbeUploadServiceClient(cc grpc.ClientConnInterface) ProbeUploadServiceClient {
	return &probeUploadServiceClient{cc}
}

func (c *probeUploadServiceClient) GetExistingProbes(ctx context.Context, in *GetExistingProbesRequest, opts ...grpc.CallOption) (*GetExistingProbesResponse, error) {
	out := new(GetExistingProbesResponse)
	err := c.cc.Invoke(ctx, ProbeUploadService_GetExistingProbes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProbeUploadServiceServer is the server API for ProbeUploadService service.
// All implementations should embed UnimplementedProbeUploadServiceServer
// for forward compatibility
type ProbeUploadServiceServer interface {
	GetExistingProbes(context.Context, *GetExistingProbesRequest) (*GetExistingProbesResponse, error)
}

// UnimplementedProbeUploadServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProbeUploadServiceServer struct {
}

func (UnimplementedProbeUploadServiceServer) GetExistingProbes(context.Context, *GetExistingProbesRequest) (*GetExistingProbesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExistingProbes not implemented")
}

// UnsafeProbeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProbeUploadServiceServer will
// result in compilation errors.
type UnsafeProbeUploadServiceServer interface {
	mustEmbedUnimplementedProbeUploadServiceServer()
}

func RegisterProbeUploadServiceServer(s grpc.ServiceRegistrar, srv ProbeUploadServiceServer) {
	s.RegisterService(&ProbeUploadService_ServiceDesc, srv)
}

func _ProbeUploadService_GetExistingProbes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExistingProbesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProbeUploadServiceServer).GetExistingProbes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProbeUploadService_GetExistingProbes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProbeUploadServiceServer).GetExistingProbes(ctx, req.(*GetExistingProbesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProbeUploadService_ServiceDesc is the grpc.ServiceDesc for ProbeUploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProbeUploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProbeUploadService",
	HandlerType: (*ProbeUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExistingProbes",
			Handler:    _ProbeUploadService_GetExistingProbes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/probe_upload_service.proto",
}
