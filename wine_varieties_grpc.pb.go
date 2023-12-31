// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: wine_varieties.proto

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

const (
	WineClassifierService_GetWineVariety_FullMethodName = "/pb.WineClassifierService/GetWineVariety"
)

// WineClassifierServiceClient is the client API for WineClassifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WineClassifierServiceClient interface {
	GetWineVariety(ctx context.Context, in *WineReviewRequest, opts ...grpc.CallOption) (*WineReviewResponse, error)
}

type wineClassifierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWineClassifierServiceClient(cc grpc.ClientConnInterface) WineClassifierServiceClient {
	return &wineClassifierServiceClient{cc}
}

func (c *wineClassifierServiceClient) GetWineVariety(ctx context.Context, in *WineReviewRequest, opts ...grpc.CallOption) (*WineReviewResponse, error) {
	out := new(WineReviewResponse)
	err := c.cc.Invoke(ctx, WineClassifierService_GetWineVariety_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WineClassifierServiceServer is the server API for WineClassifierService service.
// All implementations must embed UnimplementedWineClassifierServiceServer
// for forward compatibility
type WineClassifierServiceServer interface {
	GetWineVariety(context.Context, *WineReviewRequest) (*WineReviewResponse, error)
	mustEmbedUnimplementedWineClassifierServiceServer()
}

// UnimplementedWineClassifierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWineClassifierServiceServer struct {
}

func (UnimplementedWineClassifierServiceServer) GetWineVariety(context.Context, *WineReviewRequest) (*WineReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWineVariety not implemented")
}
func (UnimplementedWineClassifierServiceServer) mustEmbedUnimplementedWineClassifierServiceServer() {}

// UnsafeWineClassifierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WineClassifierServiceServer will
// result in compilation errors.
type UnsafeWineClassifierServiceServer interface {
	mustEmbedUnimplementedWineClassifierServiceServer()
}

func RegisterWineClassifierServiceServer(s grpc.ServiceRegistrar, srv WineClassifierServiceServer) {
	s.RegisterService(&WineClassifierService_ServiceDesc, srv)
}

func _WineClassifierService_GetWineVariety_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WineReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WineClassifierServiceServer).GetWineVariety(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WineClassifierService_GetWineVariety_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WineClassifierServiceServer).GetWineVariety(ctx, req.(*WineReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WineClassifierService_ServiceDesc is the grpc.ServiceDesc for WineClassifierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WineClassifierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WineClassifierService",
	HandlerType: (*WineClassifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWineVariety",
			Handler:    _WineClassifierService_GetWineVariety_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wine_varieties.proto",
}
