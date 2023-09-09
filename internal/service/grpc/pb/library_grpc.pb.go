// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/service/grpc/pb/library.proto

package pb

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

// LibraryClient is the client API for Library service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryClient interface {
	FindBook(ctx context.Context, in *AuthorRequest, opts ...grpc.CallOption) (*BookResponse, error)
	FindAuthor(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*AuthorResponse, error)
}

type libraryClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryClient(cc grpc.ClientConnInterface) LibraryClient {
	return &libraryClient{cc}
}

func (c *libraryClient) FindBook(ctx context.Context, in *AuthorRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/library.Library/FindBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) FindAuthor(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*AuthorResponse, error) {
	out := new(AuthorResponse)
	err := c.cc.Invoke(ctx, "/library.Library/FindAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServer is the server API for Library service.
// All implementations must embed UnimplementedLibraryServer
// for forward compatibility
type LibraryServer interface {
	FindBook(context.Context, *AuthorRequest) (*BookResponse, error)
	FindAuthor(context.Context, *BookRequest) (*AuthorResponse, error)
	mustEmbedUnimplementedLibraryServer()
}

// UnimplementedLibraryServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServer struct {
}

func (UnimplementedLibraryServer) FindBook(context.Context, *AuthorRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBook not implemented")
}
func (UnimplementedLibraryServer) FindAuthor(context.Context, *BookRequest) (*AuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAuthor not implemented")
}
func (UnimplementedLibraryServer) mustEmbedUnimplementedLibraryServer() {}

// UnsafeLibraryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServer will
// result in compilation errors.
type UnsafeLibraryServer interface {
	mustEmbedUnimplementedLibraryServer()
}

func RegisterLibraryServer(s grpc.ServiceRegistrar, srv LibraryServer) {
	s.RegisterService(&Library_ServiceDesc, srv)
}

func _Library_FindBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).FindBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.Library/FindBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).FindBook(ctx, req.(*AuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_FindAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).FindAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.Library/FindAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).FindAuthor(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Library_ServiceDesc is the grpc.ServiceDesc for Library service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Library_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "library.Library",
	HandlerType: (*LibraryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBook",
			Handler:    _Library_FindBook_Handler,
		},
		{
			MethodName: "FindAuthor",
			Handler:    _Library_FindAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/service/grpc/pb/library.proto",
}
