// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: area.proto

package area

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

// AreaClient is the client API for Area service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AreaClient interface {
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error)
}

type areaClient struct {
	cc grpc.ClientConnInterface
}

func NewAreaClient(cc grpc.ClientConnInterface) AreaClient {
	return &areaClient{cc}
}

func (c *areaClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveReply, error) {
	out := new(SaveReply)
	err := c.cc.Invoke(ctx, "/area.Area/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/area.Area/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := c.cc.Invoke(ctx, "/area.Area/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := c.cc.Invoke(ctx, "/area.Area/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AreaServer is the server API for Area service.
// All implementations must embed UnimplementedAreaServer
// for forward compatibility
type AreaServer interface {
	Save(context.Context, *SaveRequest) (*SaveReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	Delete(context.Context, *DeleteRequest) (*DeleteReply, error)
	mustEmbedUnimplementedAreaServer()
}

// UnimplementedAreaServer must be embedded to have forward compatible implementations.
type UnimplementedAreaServer struct {
}

func (UnimplementedAreaServer) Save(context.Context, *SaveRequest) (*SaveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedAreaServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAreaServer) Update(context.Context, *UpdateRequest) (*UpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAreaServer) Delete(context.Context, *DeleteRequest) (*DeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAreaServer) mustEmbedUnimplementedAreaServer() {}

// UnsafeAreaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AreaServer will
// result in compilation errors.
type UnsafeAreaServer interface {
	mustEmbedUnimplementedAreaServer()
}

func RegisterAreaServer(s grpc.ServiceRegistrar, srv AreaServer) {
	s.RegisterService(&Area_ServiceDesc, srv)
}

func _Area_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Area_ServiceDesc is the grpc.ServiceDesc for Area service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Area_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "area.Area",
	HandlerType: (*AreaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Area_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Area_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Area_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Area_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "area.proto",
}