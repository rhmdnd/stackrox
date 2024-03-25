// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/feature_flag_service.proto

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
	FeatureFlagService_GetFeatureFlags_FullMethodName = "/v1.FeatureFlagService/GetFeatureFlags"
)

// FeatureFlagServiceClient is the client API for FeatureFlagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureFlagServiceClient interface {
	GetFeatureFlags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFeatureFlagsResponse, error)
}

type featureFlagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureFlagServiceClient(cc grpc.ClientConnInterface) FeatureFlagServiceClient {
	return &featureFlagServiceClient{cc}
}

func (c *featureFlagServiceClient) GetFeatureFlags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFeatureFlagsResponse, error) {
	out := new(GetFeatureFlagsResponse)
	err := c.cc.Invoke(ctx, FeatureFlagService_GetFeatureFlags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureFlagServiceServer is the server API for FeatureFlagService service.
// All implementations should embed UnimplementedFeatureFlagServiceServer
// for forward compatibility
type FeatureFlagServiceServer interface {
	GetFeatureFlags(context.Context, *Empty) (*GetFeatureFlagsResponse, error)
}

// UnimplementedFeatureFlagServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFeatureFlagServiceServer struct {
}

func (UnimplementedFeatureFlagServiceServer) GetFeatureFlags(context.Context, *Empty) (*GetFeatureFlagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeatureFlags not implemented")
}

// UnsafeFeatureFlagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureFlagServiceServer will
// result in compilation errors.
type UnsafeFeatureFlagServiceServer interface {
	mustEmbedUnimplementedFeatureFlagServiceServer()
}

func RegisterFeatureFlagServiceServer(s grpc.ServiceRegistrar, srv FeatureFlagServiceServer) {
	s.RegisterService(&FeatureFlagService_ServiceDesc, srv)
}

func _FeatureFlagService_GetFeatureFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureFlagServiceServer).GetFeatureFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureFlagService_GetFeatureFlags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureFlagServiceServer).GetFeatureFlags(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FeatureFlagService_ServiceDesc is the grpc.ServiceDesc for FeatureFlagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureFlagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FeatureFlagService",
	HandlerType: (*FeatureFlagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeatureFlags",
			Handler:    _FeatureFlagService_GetFeatureFlags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/feature_flag_service.proto",
}
