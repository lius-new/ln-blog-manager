// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: analyzer.proto

package analyzer

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

// AnalyzerClient is the client API for Analyzer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyzerClient interface {
	// ================ Record  ================
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
	MergeRecord(ctx context.Context, in *MergeRecordRequest, opts ...grpc.CallOption) (*MergeRecordResponse, error)
	DeleteRecordById(ctx context.Context, in *DeleteRecordByIdRequest, opts ...grpc.CallOption) (*DeleteRecordByIdResponse, error)
	SelectRecordById(ctx context.Context, in *SelectRecordByIdRequest, opts ...grpc.CallOption) (*SelectRecordByIdResponse, error)
	SelectRecordByPage(ctx context.Context, in *SelectRecordByPageRequest, opts ...grpc.CallOption) (*SelectRecordByPageResponse, error)
	// ================ Blocked  ================
	CreateBlocked(ctx context.Context, in *CreateBlockedRequest, opts ...grpc.CallOption) (*CreateBlockedResponse, error)
	ModifyBlockedWithBlockEnd(ctx context.Context, in *ModifyBlockedWithBlockEndRequest, opts ...grpc.CallOption) (*ModifyBlockedWithBlockEndResponse, error)
	ModifyBlockedWithBlockCountAdd(ctx context.Context, in *ModifyBlockedWithBlockCountAddRequest, opts ...grpc.CallOption) (*ModifyBlockedWithBlockCountAddResponse, error)
	DeleteBlockedWithBlockIP(ctx context.Context, in *DeleteBlockedWithBlockIPRequest, opts ...grpc.CallOption) (*DeleteBlockedWithBlockIPResponse, error)
	SelectBlockedByBlockIP(ctx context.Context, in *SelectBlockedByBlockIPRequest, opts ...grpc.CallOption) (*SelectBlockedByBlockIPResponse, error)
	SelectBlockedById(ctx context.Context, in *SelectBlockedByIdRequest, opts ...grpc.CallOption) (*SelectBlockedByIdResponse, error)
	SelectBlockedByPage(ctx context.Context, in *SelectBlockedByPageRequest, opts ...grpc.CallOption) (*SelectBlockedByPageResponse, error)
	// 判断是否被封禁
	JudgeBlockedByIP(ctx context.Context, in *JudgeBlockedByIPRequest, opts ...grpc.CallOption) (*JudgeBlockedByIPResponse, error)
}

type analyzerClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyzerClient(cc grpc.ClientConnInterface) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) MergeRecord(ctx context.Context, in *MergeRecordRequest, opts ...grpc.CallOption) (*MergeRecordResponse, error) {
	out := new(MergeRecordResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/MergeRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) DeleteRecordById(ctx context.Context, in *DeleteRecordByIdRequest, opts ...grpc.CallOption) (*DeleteRecordByIdResponse, error) {
	out := new(DeleteRecordByIdResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/DeleteRecordById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) SelectRecordById(ctx context.Context, in *SelectRecordByIdRequest, opts ...grpc.CallOption) (*SelectRecordByIdResponse, error) {
	out := new(SelectRecordByIdResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/SelectRecordById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) SelectRecordByPage(ctx context.Context, in *SelectRecordByPageRequest, opts ...grpc.CallOption) (*SelectRecordByPageResponse, error) {
	out := new(SelectRecordByPageResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/SelectRecordByPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) CreateBlocked(ctx context.Context, in *CreateBlockedRequest, opts ...grpc.CallOption) (*CreateBlockedResponse, error) {
	out := new(CreateBlockedResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/CreateBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) ModifyBlockedWithBlockEnd(ctx context.Context, in *ModifyBlockedWithBlockEndRequest, opts ...grpc.CallOption) (*ModifyBlockedWithBlockEndResponse, error) {
	out := new(ModifyBlockedWithBlockEndResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/ModifyBlockedWithBlockEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) ModifyBlockedWithBlockCountAdd(ctx context.Context, in *ModifyBlockedWithBlockCountAddRequest, opts ...grpc.CallOption) (*ModifyBlockedWithBlockCountAddResponse, error) {
	out := new(ModifyBlockedWithBlockCountAddResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/ModifyBlockedWithBlockCountAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) DeleteBlockedWithBlockIP(ctx context.Context, in *DeleteBlockedWithBlockIPRequest, opts ...grpc.CallOption) (*DeleteBlockedWithBlockIPResponse, error) {
	out := new(DeleteBlockedWithBlockIPResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/DeleteBlockedWithBlockIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) SelectBlockedByBlockIP(ctx context.Context, in *SelectBlockedByBlockIPRequest, opts ...grpc.CallOption) (*SelectBlockedByBlockIPResponse, error) {
	out := new(SelectBlockedByBlockIPResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/SelectBlockedByBlockIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) SelectBlockedById(ctx context.Context, in *SelectBlockedByIdRequest, opts ...grpc.CallOption) (*SelectBlockedByIdResponse, error) {
	out := new(SelectBlockedByIdResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/SelectBlockedById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) SelectBlockedByPage(ctx context.Context, in *SelectBlockedByPageRequest, opts ...grpc.CallOption) (*SelectBlockedByPageResponse, error) {
	out := new(SelectBlockedByPageResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/SelectBlockedByPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) JudgeBlockedByIP(ctx context.Context, in *JudgeBlockedByIPRequest, opts ...grpc.CallOption) (*JudgeBlockedByIPResponse, error) {
	out := new(JudgeBlockedByIPResponse)
	err := c.cc.Invoke(ctx, "/analyzer.Analyzer/JudgeBlockedByIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyzerServer is the server API for Analyzer service.
// All implementations must embed UnimplementedAnalyzerServer
// for forward compatibility
type AnalyzerServer interface {
	// ================ Record  ================
	CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	MergeRecord(context.Context, *MergeRecordRequest) (*MergeRecordResponse, error)
	DeleteRecordById(context.Context, *DeleteRecordByIdRequest) (*DeleteRecordByIdResponse, error)
	SelectRecordById(context.Context, *SelectRecordByIdRequest) (*SelectRecordByIdResponse, error)
	SelectRecordByPage(context.Context, *SelectRecordByPageRequest) (*SelectRecordByPageResponse, error)
	// ================ Blocked  ================
	CreateBlocked(context.Context, *CreateBlockedRequest) (*CreateBlockedResponse, error)
	ModifyBlockedWithBlockEnd(context.Context, *ModifyBlockedWithBlockEndRequest) (*ModifyBlockedWithBlockEndResponse, error)
	ModifyBlockedWithBlockCountAdd(context.Context, *ModifyBlockedWithBlockCountAddRequest) (*ModifyBlockedWithBlockCountAddResponse, error)
	DeleteBlockedWithBlockIP(context.Context, *DeleteBlockedWithBlockIPRequest) (*DeleteBlockedWithBlockIPResponse, error)
	SelectBlockedByBlockIP(context.Context, *SelectBlockedByBlockIPRequest) (*SelectBlockedByBlockIPResponse, error)
	SelectBlockedById(context.Context, *SelectBlockedByIdRequest) (*SelectBlockedByIdResponse, error)
	SelectBlockedByPage(context.Context, *SelectBlockedByPageRequest) (*SelectBlockedByPageResponse, error)
	// 判断是否被封禁
	JudgeBlockedByIP(context.Context, *JudgeBlockedByIPRequest) (*JudgeBlockedByIPResponse, error)
	mustEmbedUnimplementedAnalyzerServer()
}

// UnimplementedAnalyzerServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyzerServer struct {
}

func (UnimplementedAnalyzerServer) CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedAnalyzerServer) MergeRecord(context.Context, *MergeRecordRequest) (*MergeRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeRecord not implemented")
}
func (UnimplementedAnalyzerServer) DeleteRecordById(context.Context, *DeleteRecordByIdRequest) (*DeleteRecordByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecordById not implemented")
}
func (UnimplementedAnalyzerServer) SelectRecordById(context.Context, *SelectRecordByIdRequest) (*SelectRecordByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectRecordById not implemented")
}
func (UnimplementedAnalyzerServer) SelectRecordByPage(context.Context, *SelectRecordByPageRequest) (*SelectRecordByPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectRecordByPage not implemented")
}
func (UnimplementedAnalyzerServer) CreateBlocked(context.Context, *CreateBlockedRequest) (*CreateBlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlocked not implemented")
}
func (UnimplementedAnalyzerServer) ModifyBlockedWithBlockEnd(context.Context, *ModifyBlockedWithBlockEndRequest) (*ModifyBlockedWithBlockEndResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyBlockedWithBlockEnd not implemented")
}
func (UnimplementedAnalyzerServer) ModifyBlockedWithBlockCountAdd(context.Context, *ModifyBlockedWithBlockCountAddRequest) (*ModifyBlockedWithBlockCountAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyBlockedWithBlockCountAdd not implemented")
}
func (UnimplementedAnalyzerServer) DeleteBlockedWithBlockIP(context.Context, *DeleteBlockedWithBlockIPRequest) (*DeleteBlockedWithBlockIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlockedWithBlockIP not implemented")
}
func (UnimplementedAnalyzerServer) SelectBlockedByBlockIP(context.Context, *SelectBlockedByBlockIPRequest) (*SelectBlockedByBlockIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectBlockedByBlockIP not implemented")
}
func (UnimplementedAnalyzerServer) SelectBlockedById(context.Context, *SelectBlockedByIdRequest) (*SelectBlockedByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectBlockedById not implemented")
}
func (UnimplementedAnalyzerServer) SelectBlockedByPage(context.Context, *SelectBlockedByPageRequest) (*SelectBlockedByPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectBlockedByPage not implemented")
}
func (UnimplementedAnalyzerServer) JudgeBlockedByIP(context.Context, *JudgeBlockedByIPRequest) (*JudgeBlockedByIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JudgeBlockedByIP not implemented")
}
func (UnimplementedAnalyzerServer) mustEmbedUnimplementedAnalyzerServer() {}

// UnsafeAnalyzerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyzerServer will
// result in compilation errors.
type UnsafeAnalyzerServer interface {
	mustEmbedUnimplementedAnalyzerServer()
}

func RegisterAnalyzerServer(s grpc.ServiceRegistrar, srv AnalyzerServer) {
	s.RegisterService(&Analyzer_ServiceDesc, srv)
}

func _Analyzer_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_MergeRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).MergeRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/MergeRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).MergeRecord(ctx, req.(*MergeRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_DeleteRecordById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).DeleteRecordById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/DeleteRecordById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).DeleteRecordById(ctx, req.(*DeleteRecordByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_SelectRecordById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRecordByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).SelectRecordById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/SelectRecordById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).SelectRecordById(ctx, req.(*SelectRecordByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_SelectRecordByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRecordByPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).SelectRecordByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/SelectRecordByPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).SelectRecordByPage(ctx, req.(*SelectRecordByPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_CreateBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).CreateBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/CreateBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).CreateBlocked(ctx, req.(*CreateBlockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_ModifyBlockedWithBlockEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyBlockedWithBlockEndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).ModifyBlockedWithBlockEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/ModifyBlockedWithBlockEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).ModifyBlockedWithBlockEnd(ctx, req.(*ModifyBlockedWithBlockEndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_ModifyBlockedWithBlockCountAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyBlockedWithBlockCountAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).ModifyBlockedWithBlockCountAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/ModifyBlockedWithBlockCountAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).ModifyBlockedWithBlockCountAdd(ctx, req.(*ModifyBlockedWithBlockCountAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_DeleteBlockedWithBlockIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlockedWithBlockIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).DeleteBlockedWithBlockIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/DeleteBlockedWithBlockIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).DeleteBlockedWithBlockIP(ctx, req.(*DeleteBlockedWithBlockIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_SelectBlockedByBlockIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectBlockedByBlockIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).SelectBlockedByBlockIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/SelectBlockedByBlockIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).SelectBlockedByBlockIP(ctx, req.(*SelectBlockedByBlockIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_SelectBlockedById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectBlockedByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).SelectBlockedById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/SelectBlockedById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).SelectBlockedById(ctx, req.(*SelectBlockedByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_SelectBlockedByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectBlockedByPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).SelectBlockedByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/SelectBlockedByPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).SelectBlockedByPage(ctx, req.(*SelectBlockedByPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_JudgeBlockedByIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgeBlockedByIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).JudgeBlockedByIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analyzer.Analyzer/JudgeBlockedByIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).JudgeBlockedByIP(ctx, req.(*JudgeBlockedByIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Analyzer_ServiceDesc is the grpc.ServiceDesc for Analyzer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analyzer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analyzer.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecord",
			Handler:    _Analyzer_CreateRecord_Handler,
		},
		{
			MethodName: "MergeRecord",
			Handler:    _Analyzer_MergeRecord_Handler,
		},
		{
			MethodName: "DeleteRecordById",
			Handler:    _Analyzer_DeleteRecordById_Handler,
		},
		{
			MethodName: "SelectRecordById",
			Handler:    _Analyzer_SelectRecordById_Handler,
		},
		{
			MethodName: "SelectRecordByPage",
			Handler:    _Analyzer_SelectRecordByPage_Handler,
		},
		{
			MethodName: "CreateBlocked",
			Handler:    _Analyzer_CreateBlocked_Handler,
		},
		{
			MethodName: "ModifyBlockedWithBlockEnd",
			Handler:    _Analyzer_ModifyBlockedWithBlockEnd_Handler,
		},
		{
			MethodName: "ModifyBlockedWithBlockCountAdd",
			Handler:    _Analyzer_ModifyBlockedWithBlockCountAdd_Handler,
		},
		{
			MethodName: "DeleteBlockedWithBlockIP",
			Handler:    _Analyzer_DeleteBlockedWithBlockIP_Handler,
		},
		{
			MethodName: "SelectBlockedByBlockIP",
			Handler:    _Analyzer_SelectBlockedByBlockIP_Handler,
		},
		{
			MethodName: "SelectBlockedById",
			Handler:    _Analyzer_SelectBlockedById_Handler,
		},
		{
			MethodName: "SelectBlockedByPage",
			Handler:    _Analyzer_SelectBlockedByPage_Handler,
		},
		{
			MethodName: "JudgeBlockedByIP",
			Handler:    _Analyzer_JudgeBlockedByIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyzer.proto",
}
