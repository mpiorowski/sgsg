// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: main.proto

package proto

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Auth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthResponse, error)
	GetNotes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Service_GetNotesClient, error)
	GetNoteById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Note, error)
	CreateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error)
	DeleteNote(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Auth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Service/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNotes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Service_GetNotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], "/proto.Service/GetNotes", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceGetNotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_GetNotesClient interface {
	Recv() (*Note, error)
	grpc.ClientStream
}

type serviceGetNotesClient struct {
	grpc.ClientStream
}

func (x *serviceGetNotesClient) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) GetNoteById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/proto.Service/GetNoteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/proto.Service/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteNote(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Service/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Auth(context.Context, *Empty) (*AuthResponse, error)
	GetNotes(*Empty, Service_GetNotesServer) error
	GetNoteById(context.Context, *Id) (*Note, error)
	CreateNote(context.Context, *Note) (*Note, error)
	DeleteNote(context.Context, *Id) (*Empty, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Auth(context.Context, *Empty) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedServiceServer) GetNotes(*Empty, Service_GetNotesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNotes not implemented")
}
func (UnimplementedServiceServer) GetNoteById(context.Context, *Id) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteById not implemented")
}
func (UnimplementedServiceServer) CreateNote(context.Context, *Note) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedServiceServer) DeleteNote(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Auth(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetNotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).GetNotes(m, &serviceGetNotesServer{stream})
}

type Service_GetNotesServer interface {
	Send(*Note) error
	grpc.ServerStream
}

type serviceGetNotesServer struct {
	grpc.ServerStream
}

func (x *serviceGetNotesServer) Send(m *Note) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_GetNoteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetNoteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/GetNoteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetNoteById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteNote(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Service_Auth_Handler,
		},
		{
			MethodName: "GetNoteById",
			Handler:    _Service_GetNoteById_Handler,
		},
		{
			MethodName: "CreateNote",
			Handler:    _Service_CreateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _Service_DeleteNote_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNotes",
			Handler:       _Service_GetNotes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "main.proto",
}