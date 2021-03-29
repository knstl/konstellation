package types

import (
	context "context"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	grpc "google.golang.org/grpc"
)

// MsgVerifyInvariantResponse defines the Msg/VerifyInvariant response type.
type MsgVerifyInvariantResponse struct {
}

// MsgVerifyInvariant represents a message to verify a particular invariance.
type MsgVerifyInvariant struct {
	Sender              string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	InvariantModuleName string `protobuf:"bytes,2,opt,name=invariant_module_name,json=invariantModuleName,proto3" json:"invariant_module_name,omitempty" yaml:"invariant_module_name"`
	InvariantRoute      string `protobuf:"bytes,3,opt,name=invariant_route,json=invariantRoute,proto3" json:"invariant_route,omitempty" yaml:"invariant_route"`
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// VerifyInvariant defines a method to verify a particular invariance.
	VerifyInvariant(context.Context, *MsgVerifyInvariant) (*MsgVerifyInvariantResponse, error)
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.crisis.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyInvariant",
			Handler:    _Msg_VerifyInvariant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/crisis/v1beta1/tx.proto",
}

func _Msg_VerifyInvariant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVerifyInvariant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VerifyInvariant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.crisis.v1beta1.Msg/VerifyInvariant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VerifyInvariant(ctx, req.(*MsgVerifyInvariant))
	}
	return interceptor(ctx, in, info, handler)
}
