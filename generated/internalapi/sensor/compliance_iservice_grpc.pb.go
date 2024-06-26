// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: internalapi/sensor/compliance_iservice.proto

package sensor

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
	ComplianceService_Communicate_FullMethodName = "/sensor.ComplianceService/Communicate"
)

// ComplianceServiceClient is the client API for ComplianceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComplianceServiceClient interface {
	Communicate(ctx context.Context, opts ...grpc.CallOption) (ComplianceService_CommunicateClient, error)
}

type complianceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceServiceClient(cc grpc.ClientConnInterface) ComplianceServiceClient {
	return &complianceServiceClient{cc}
}

func (c *complianceServiceClient) Communicate(ctx context.Context, opts ...grpc.CallOption) (ComplianceService_CommunicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &ComplianceService_ServiceDesc.Streams[0], ComplianceService_Communicate_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &complianceServiceCommunicateClient{stream}
	return x, nil
}

type ComplianceService_CommunicateClient interface {
	Send(*MsgFromCompliance) error
	Recv() (*MsgToCompliance, error)
	grpc.ClientStream
}

type complianceServiceCommunicateClient struct {
	grpc.ClientStream
}

func (x *complianceServiceCommunicateClient) Send(m *MsgFromCompliance) error {
	return x.ClientStream.SendMsg(m)
}

func (x *complianceServiceCommunicateClient) Recv() (*MsgToCompliance, error) {
	m := new(MsgToCompliance)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComplianceServiceServer is the server API for ComplianceService service.
// All implementations should embed UnimplementedComplianceServiceServer
// for forward compatibility
type ComplianceServiceServer interface {
	Communicate(ComplianceService_CommunicateServer) error
}

// UnimplementedComplianceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedComplianceServiceServer struct {
}

func (UnimplementedComplianceServiceServer) Communicate(ComplianceService_CommunicateServer) error {
	return status.Errorf(codes.Unimplemented, "method Communicate not implemented")
}

// UnsafeComplianceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplianceServiceServer will
// result in compilation errors.
type UnsafeComplianceServiceServer interface {
	mustEmbedUnimplementedComplianceServiceServer()
}

func RegisterComplianceServiceServer(s grpc.ServiceRegistrar, srv ComplianceServiceServer) {
	s.RegisterService(&ComplianceService_ServiceDesc, srv)
}

func _ComplianceService_Communicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ComplianceServiceServer).Communicate(&complianceServiceCommunicateServer{stream})
}

type ComplianceService_CommunicateServer interface {
	Send(*MsgToCompliance) error
	Recv() (*MsgFromCompliance, error)
	grpc.ServerStream
}

type complianceServiceCommunicateServer struct {
	grpc.ServerStream
}

func (x *complianceServiceCommunicateServer) Send(m *MsgToCompliance) error {
	return x.ServerStream.SendMsg(m)
}

func (x *complianceServiceCommunicateServer) Recv() (*MsgFromCompliance, error) {
	m := new(MsgFromCompliance)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComplianceService_ServiceDesc is the grpc.ServiceDesc for ComplianceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComplianceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.ComplianceService",
	HandlerType: (*ComplianceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Communicate",
			Handler:       _ComplianceService_Communicate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internalapi/sensor/compliance_iservice.proto",
}
