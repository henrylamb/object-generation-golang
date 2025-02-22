// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: object-generation.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	JSONSchemaService_GenerateObject_FullMethodName         = "/jsonSchema.JSONSchemaService/GenerateObject"
	JSONSchemaService_StreamGeneratedObjects_FullMethodName = "/jsonSchema.JSONSchemaService/StreamGeneratedObjects"
)

// JSONSchemaServiceClient is the client API for JSONSchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The JSONSchemaService defines a service for generating JSON objects based on a schema definition.
type JSONSchemaServiceClient interface {
	// Standard request-response RPC
	GenerateObject(ctx context.Context, in *RequestBody, opts ...grpc.CallOption) (*Response, error)
	// New method: Server-side streaming RPC
	StreamGeneratedObjects(ctx context.Context, in *RequestBody, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamingResponse], error)
}

type jSONSchemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJSONSchemaServiceClient(cc grpc.ClientConnInterface) JSONSchemaServiceClient {
	return &jSONSchemaServiceClient{cc}
}

func (c *jSONSchemaServiceClient) GenerateObject(ctx context.Context, in *RequestBody, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, JSONSchemaService_GenerateObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jSONSchemaServiceClient) StreamGeneratedObjects(ctx context.Context, in *RequestBody, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamingResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &JSONSchemaService_ServiceDesc.Streams[0], JSONSchemaService_StreamGeneratedObjects_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RequestBody, StreamingResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JSONSchemaService_StreamGeneratedObjectsClient = grpc.ServerStreamingClient[StreamingResponse]

// JSONSchemaServiceServer is the server API for JSONSchemaService service.
// All implementations must embed UnimplementedJSONSchemaServiceServer
// for forward compatibility.
//
// The JSONSchemaService defines a service for generating JSON objects based on a schema definition.
type JSONSchemaServiceServer interface {
	// Standard request-response RPC
	GenerateObject(context.Context, *RequestBody) (*Response, error)
	// New method: Server-side streaming RPC
	StreamGeneratedObjects(*RequestBody, grpc.ServerStreamingServer[StreamingResponse]) error
	mustEmbedUnimplementedJSONSchemaServiceServer()
}

// UnimplementedJSONSchemaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJSONSchemaServiceServer struct{}

func (UnimplementedJSONSchemaServiceServer) GenerateObject(context.Context, *RequestBody) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateObject not implemented")
}
func (UnimplementedJSONSchemaServiceServer) StreamGeneratedObjects(*RequestBody, grpc.ServerStreamingServer[StreamingResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamGeneratedObjects not implemented")
}
func (UnimplementedJSONSchemaServiceServer) mustEmbedUnimplementedJSONSchemaServiceServer() {}
func (UnimplementedJSONSchemaServiceServer) testEmbeddedByValue()                           {}

// UnsafeJSONSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JSONSchemaServiceServer will
// result in compilation errors.
type UnsafeJSONSchemaServiceServer interface {
	mustEmbedUnimplementedJSONSchemaServiceServer()
}

func RegisterJSONSchemaServiceServer(s grpc.ServiceRegistrar, srv JSONSchemaServiceServer) {
	// If the following call pancis, it indicates UnimplementedJSONSchemaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JSONSchemaService_ServiceDesc, srv)
}

func _JSONSchemaService_GenerateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JSONSchemaServiceServer).GenerateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JSONSchemaService_GenerateObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JSONSchemaServiceServer).GenerateObject(ctx, req.(*RequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _JSONSchemaService_StreamGeneratedObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestBody)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JSONSchemaServiceServer).StreamGeneratedObjects(m, &grpc.GenericServerStream[RequestBody, StreamingResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JSONSchemaService_StreamGeneratedObjectsServer = grpc.ServerStreamingServer[StreamingResponse]

// JSONSchemaService_ServiceDesc is the grpc.ServiceDesc for JSONSchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JSONSchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jsonSchema.JSONSchemaService",
	HandlerType: (*JSONSchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateObject",
			Handler:    _JSONSchemaService_GenerateObject_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamGeneratedObjects",
			Handler:       _JSONSchemaService_StreamGeneratedObjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "object-generation.proto",
}
