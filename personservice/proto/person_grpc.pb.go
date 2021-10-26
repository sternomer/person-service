// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_grpc_mongodb_master

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

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	CreatePerson(ctx context.Context, in *CreatePersonReq, opts ...grpc.CallOption) (*CreatePersonRes, error)
	ReadPerson(ctx context.Context, in *ReadPersonReq, opts ...grpc.CallOption) (*ReadPersonRes, error)
	UpdatePerson(ctx context.Context, in *UpdatePersonReq, opts ...grpc.CallOption) (*UpdatePersonRes, error)
	DeletePerson(ctx context.Context, in *DeletePersonReq, opts ...grpc.CallOption) (*DeletePersonRes, error)
	ListPersons(ctx context.Context, in *ListPersonsReq, opts ...grpc.CallOption) (PersonService_ListPersonsClient, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) CreatePerson(ctx context.Context, in *CreatePersonReq, opts ...grpc.CallOption) (*CreatePersonRes, error) {
	out := new(CreatePersonRes)
	err := c.cc.Invoke(ctx, "/blog.PersonService/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) ReadPerson(ctx context.Context, in *ReadPersonReq, opts ...grpc.CallOption) (*ReadPersonRes, error) {
	out := new(ReadPersonRes)
	err := c.cc.Invoke(ctx, "/blog.PersonService/ReadPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) UpdatePerson(ctx context.Context, in *UpdatePersonReq, opts ...grpc.CallOption) (*UpdatePersonRes, error) {
	out := new(UpdatePersonRes)
	err := c.cc.Invoke(ctx, "/blog.PersonService/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) DeletePerson(ctx context.Context, in *DeletePersonReq, opts ...grpc.CallOption) (*DeletePersonRes, error) {
	out := new(DeletePersonRes)
	err := c.cc.Invoke(ctx, "/blog.PersonService/DeletePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) ListPersons(ctx context.Context, in *ListPersonsReq, opts ...grpc.CallOption) (PersonService_ListPersonsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonService_ServiceDesc.Streams[0], "/blog.PersonService/ListPersons", opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceListPersonsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PersonService_ListPersonsClient interface {
	Recv() (*ListPersonsRes, error)
	grpc.ClientStream
}

type personServiceListPersonsClient struct {
	grpc.ClientStream
}

func (x *personServiceListPersonsClient) Recv() (*ListPersonsRes, error) {
	m := new(ListPersonsRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility
type PersonServiceServer interface {
	CreatePerson(context.Context, *CreatePersonReq) (*CreatePersonRes, error)
	ReadPerson(context.Context, *ReadPersonReq) (*ReadPersonRes, error)
	UpdatePerson(context.Context, *UpdatePersonReq) (*UpdatePersonRes, error)
	DeletePerson(context.Context, *DeletePersonReq) (*DeletePersonRes, error)
	ListPersons(*ListPersonsReq, PersonService_ListPersonsServer) error
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (UnimplementedPersonServiceServer) CreatePerson(context.Context, *CreatePersonReq) (*CreatePersonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedPersonServiceServer) ReadPerson(context.Context, *ReadPersonReq) (*ReadPersonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPerson not implemented")
}
func (UnimplementedPersonServiceServer) UpdatePerson(context.Context, *UpdatePersonReq) (*UpdatePersonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (UnimplementedPersonServiceServer) DeletePerson(context.Context, *DeletePersonReq) (*DeletePersonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePerson not implemented")
}
func (UnimplementedPersonServiceServer) ListPersons(*ListPersonsReq, PersonService_ListPersonsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPersons not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.PersonService/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).CreatePerson(ctx, req.(*CreatePersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_ReadPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).ReadPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.PersonService/ReadPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).ReadPerson(ctx, req.(*ReadPersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.PersonService/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).UpdatePerson(ctx, req.(*UpdatePersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_DeletePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).DeletePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.PersonService/DeletePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).DeletePerson(ctx, req.(*DeletePersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_ListPersons_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPersonsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PersonServiceServer).ListPersons(m, &personServiceListPersonsServer{stream})
}

type PersonService_ListPersonsServer interface {
	Send(*ListPersonsRes) error
	grpc.ServerStream
}

type personServiceListPersonsServer struct {
	grpc.ServerStream
}

func (x *personServiceListPersonsServer) Send(m *ListPersonsRes) error {
	return x.ServerStream.SendMsg(m)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _PersonService_CreatePerson_Handler,
		},
		{
			MethodName: "ReadPerson",
			Handler:    _PersonService_ReadPerson_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _PersonService_UpdatePerson_Handler,
		},
		{
			MethodName: "DeletePerson",
			Handler:    _PersonService_DeletePerson_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPersons",
			Handler:       _PersonService_ListPersons_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/person.proto",
}
