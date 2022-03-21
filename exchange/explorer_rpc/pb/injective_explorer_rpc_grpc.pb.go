// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: injective_explorer_rpc.proto

package injective_explorer_rpcpb

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

// InjectiveExplorerRPCClient is the client API for InjectiveExplorerRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InjectiveExplorerRPCClient interface {
	// GetAccountTxs returns tranctions involving in an account based upon params.
	GetAccountTxs(ctx context.Context, in *GetAccountTxsRequest, opts ...grpc.CallOption) (*GetAccountTxsResponse, error)
	// GetBlocks returns blocks based upon the request params
	GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error)
	// GetBlock returns block based upon the height or hash
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	// GetValidator returns validator information on the active chain
	GetValidator(ctx context.Context, in *GetValidatorRequest, opts ...grpc.CallOption) (*GetValidatorResponse, error)
	// GetValidatorUptime returns validator uptime information on the active chain
	GetValidatorUptime(ctx context.Context, in *GetValidatorUptimeRequest, opts ...grpc.CallOption) (*GetValidatorUptimeResponse, error)
	// GetCoinPriceData returns price data from CoinGecko API
	GetCoinPriceData(ctx context.Context, in *GetCoinPriceDataRequest, opts ...grpc.CallOption) (*GetCoinPriceDataResponse, error)
	// GetTxs returns transactions based upon the request params
	GetTxs(ctx context.Context, in *GetTxsRequest, opts ...grpc.CallOption) (*GetTxsResponse, error)
	// GetTxByTxHash returns certain transaction information by its tx hash.
	GetTxByTxHash(ctx context.Context, in *GetTxByTxHashRequest, opts ...grpc.CallOption) (*GetTxByTxHashResponse, error)
	// StreamTxs returns transactions based upon the request params
	StreamTxs(ctx context.Context, in *StreamTxsRequest, opts ...grpc.CallOption) (InjectiveExplorerRPC_StreamTxsClient, error)
	// StreamBlocks returns the latest blocks
	StreamBlocks(ctx context.Context, in *StreamBlocksRequest, opts ...grpc.CallOption) (InjectiveExplorerRPC_StreamBlocksClient, error)
	// GetPeggyDepositTxs returns the peggy deposit transactions based upon the
	// request params
	GetPeggyDepositTxs(ctx context.Context, in *GetPeggyDepositTxsRequest, opts ...grpc.CallOption) (*GetPeggyDepositTxsResponse, error)
	// GetPeggyWithdrawalTxs returns the peggy withdrawal transactions based upon
	// the request params
	GetPeggyWithdrawalTxs(ctx context.Context, in *GetPeggyWithdrawalTxsRequest, opts ...grpc.CallOption) (*GetPeggyWithdrawalTxsResponse, error)
	// GetIBCTransferTxs returns the ibc transfer transactions based upon the
	// request params
	GetIBCTransferTxs(ctx context.Context, in *GetIBCTransferTxsRequest, opts ...grpc.CallOption) (*GetIBCTransferTxsResponse, error)
}

type injectiveExplorerRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewInjectiveExplorerRPCClient(cc grpc.ClientConnInterface) InjectiveExplorerRPCClient {
	return &injectiveExplorerRPCClient{cc}
}

func (c *injectiveExplorerRPCClient) GetAccountTxs(ctx context.Context, in *GetAccountTxsRequest, opts ...grpc.CallOption) (*GetAccountTxsResponse, error) {
	out := new(GetAccountTxsResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetAccountTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error) {
	out := new(GetBlocksResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetValidator(ctx context.Context, in *GetValidatorRequest, opts ...grpc.CallOption) (*GetValidatorResponse, error) {
	out := new(GetValidatorResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetValidatorUptime(ctx context.Context, in *GetValidatorUptimeRequest, opts ...grpc.CallOption) (*GetValidatorUptimeResponse, error) {
	out := new(GetValidatorUptimeResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetValidatorUptime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetCoinPriceData(ctx context.Context, in *GetCoinPriceDataRequest, opts ...grpc.CallOption) (*GetCoinPriceDataResponse, error) {
	out := new(GetCoinPriceDataResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetCoinPriceData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetTxs(ctx context.Context, in *GetTxsRequest, opts ...grpc.CallOption) (*GetTxsResponse, error) {
	out := new(GetTxsResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetTxByTxHash(ctx context.Context, in *GetTxByTxHashRequest, opts ...grpc.CallOption) (*GetTxByTxHashResponse, error) {
	out := new(GetTxByTxHashResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetTxByTxHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) StreamTxs(ctx context.Context, in *StreamTxsRequest, opts ...grpc.CallOption) (InjectiveExplorerRPC_StreamTxsClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveExplorerRPC_ServiceDesc.Streams[0], "/injective_explorer_rpc.InjectiveExplorerRPC/StreamTxs", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveExplorerRPCStreamTxsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveExplorerRPC_StreamTxsClient interface {
	Recv() (*StreamTxsResponse, error)
	grpc.ClientStream
}

type injectiveExplorerRPCStreamTxsClient struct {
	grpc.ClientStream
}

func (x *injectiveExplorerRPCStreamTxsClient) Recv() (*StreamTxsResponse, error) {
	m := new(StreamTxsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveExplorerRPCClient) StreamBlocks(ctx context.Context, in *StreamBlocksRequest, opts ...grpc.CallOption) (InjectiveExplorerRPC_StreamBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveExplorerRPC_ServiceDesc.Streams[1], "/injective_explorer_rpc.InjectiveExplorerRPC/StreamBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveExplorerRPCStreamBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveExplorerRPC_StreamBlocksClient interface {
	Recv() (*StreamBlocksResponse, error)
	grpc.ClientStream
}

type injectiveExplorerRPCStreamBlocksClient struct {
	grpc.ClientStream
}

func (x *injectiveExplorerRPCStreamBlocksClient) Recv() (*StreamBlocksResponse, error) {
	m := new(StreamBlocksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveExplorerRPCClient) GetPeggyDepositTxs(ctx context.Context, in *GetPeggyDepositTxsRequest, opts ...grpc.CallOption) (*GetPeggyDepositTxsResponse, error) {
	out := new(GetPeggyDepositTxsResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetPeggyDepositTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetPeggyWithdrawalTxs(ctx context.Context, in *GetPeggyWithdrawalTxsRequest, opts ...grpc.CallOption) (*GetPeggyWithdrawalTxsResponse, error) {
	out := new(GetPeggyWithdrawalTxsResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetPeggyWithdrawalTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveExplorerRPCClient) GetIBCTransferTxs(ctx context.Context, in *GetIBCTransferTxsRequest, opts ...grpc.CallOption) (*GetIBCTransferTxsResponse, error) {
	out := new(GetIBCTransferTxsResponse)
	err := c.cc.Invoke(ctx, "/injective_explorer_rpc.InjectiveExplorerRPC/GetIBCTransferTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjectiveExplorerRPCServer is the server API for InjectiveExplorerRPC service.
// All implementations must embed UnimplementedInjectiveExplorerRPCServer
// for forward compatibility
type InjectiveExplorerRPCServer interface {
	// GetAccountTxs returns tranctions involving in an account based upon params.
	GetAccountTxs(context.Context, *GetAccountTxsRequest) (*GetAccountTxsResponse, error)
	// GetBlocks returns blocks based upon the request params
	GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error)
	// GetBlock returns block based upon the height or hash
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	// GetValidator returns validator information on the active chain
	GetValidator(context.Context, *GetValidatorRequest) (*GetValidatorResponse, error)
	// GetValidatorUptime returns validator uptime information on the active chain
	GetValidatorUptime(context.Context, *GetValidatorUptimeRequest) (*GetValidatorUptimeResponse, error)
	// GetCoinPriceData returns price data from CoinGecko API
	GetCoinPriceData(context.Context, *GetCoinPriceDataRequest) (*GetCoinPriceDataResponse, error)
	// GetTxs returns transactions based upon the request params
	GetTxs(context.Context, *GetTxsRequest) (*GetTxsResponse, error)
	// GetTxByTxHash returns certain transaction information by its tx hash.
	GetTxByTxHash(context.Context, *GetTxByTxHashRequest) (*GetTxByTxHashResponse, error)
	// StreamTxs returns transactions based upon the request params
	StreamTxs(*StreamTxsRequest, InjectiveExplorerRPC_StreamTxsServer) error
	// StreamBlocks returns the latest blocks
	StreamBlocks(*StreamBlocksRequest, InjectiveExplorerRPC_StreamBlocksServer) error
	// GetPeggyDepositTxs returns the peggy deposit transactions based upon the
	// request params
	GetPeggyDepositTxs(context.Context, *GetPeggyDepositTxsRequest) (*GetPeggyDepositTxsResponse, error)
	// GetPeggyWithdrawalTxs returns the peggy withdrawal transactions based upon
	// the request params
	GetPeggyWithdrawalTxs(context.Context, *GetPeggyWithdrawalTxsRequest) (*GetPeggyWithdrawalTxsResponse, error)
	// GetIBCTransferTxs returns the ibc transfer transactions based upon the
	// request params
	GetIBCTransferTxs(context.Context, *GetIBCTransferTxsRequest) (*GetIBCTransferTxsResponse, error)
	mustEmbedUnimplementedInjectiveExplorerRPCServer()
}

// UnimplementedInjectiveExplorerRPCServer must be embedded to have forward compatible implementations.
type UnimplementedInjectiveExplorerRPCServer struct {
}

func (UnimplementedInjectiveExplorerRPCServer) GetAccountTxs(context.Context, *GetAccountTxsRequest) (*GetAccountTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountTxs not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetValidator(context.Context, *GetValidatorRequest) (*GetValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidator not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetValidatorUptime(context.Context, *GetValidatorUptimeRequest) (*GetValidatorUptimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorUptime not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetCoinPriceData(context.Context, *GetCoinPriceDataRequest) (*GetCoinPriceDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinPriceData not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetTxs(context.Context, *GetTxsRequest) (*GetTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxs not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetTxByTxHash(context.Context, *GetTxByTxHashRequest) (*GetTxByTxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByTxHash not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) StreamTxs(*StreamTxsRequest, InjectiveExplorerRPC_StreamTxsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTxs not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) StreamBlocks(*StreamBlocksRequest, InjectiveExplorerRPC_StreamBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamBlocks not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetPeggyDepositTxs(context.Context, *GetPeggyDepositTxsRequest) (*GetPeggyDepositTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeggyDepositTxs not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetPeggyWithdrawalTxs(context.Context, *GetPeggyWithdrawalTxsRequest) (*GetPeggyWithdrawalTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeggyWithdrawalTxs not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) GetIBCTransferTxs(context.Context, *GetIBCTransferTxsRequest) (*GetIBCTransferTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIBCTransferTxs not implemented")
}
func (UnimplementedInjectiveExplorerRPCServer) mustEmbedUnimplementedInjectiveExplorerRPCServer() {}

// UnsafeInjectiveExplorerRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InjectiveExplorerRPCServer will
// result in compilation errors.
type UnsafeInjectiveExplorerRPCServer interface {
	mustEmbedUnimplementedInjectiveExplorerRPCServer()
}

func RegisterInjectiveExplorerRPCServer(s grpc.ServiceRegistrar, srv InjectiveExplorerRPCServer) {
	s.RegisterService(&InjectiveExplorerRPC_ServiceDesc, srv)
}

func _InjectiveExplorerRPC_GetAccountTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetAccountTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetAccountTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetAccountTxs(ctx, req.(*GetAccountTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetBlocks(ctx, req.(*GetBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetValidator(ctx, req.(*GetValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetValidatorUptime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorUptimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetValidatorUptime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetValidatorUptime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetValidatorUptime(ctx, req.(*GetValidatorUptimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetCoinPriceData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinPriceDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetCoinPriceData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetCoinPriceData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetCoinPriceData(ctx, req.(*GetCoinPriceDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetTxs(ctx, req.(*GetTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetTxByTxHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxByTxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetTxByTxHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetTxByTxHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetTxByTxHash(ctx, req.(*GetTxByTxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_StreamTxs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTxsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveExplorerRPCServer).StreamTxs(m, &injectiveExplorerRPCStreamTxsServer{stream})
}

type InjectiveExplorerRPC_StreamTxsServer interface {
	Send(*StreamTxsResponse) error
	grpc.ServerStream
}

type injectiveExplorerRPCStreamTxsServer struct {
	grpc.ServerStream
}

func (x *injectiveExplorerRPCStreamTxsServer) Send(m *StreamTxsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveExplorerRPC_StreamBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamBlocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveExplorerRPCServer).StreamBlocks(m, &injectiveExplorerRPCStreamBlocksServer{stream})
}

type InjectiveExplorerRPC_StreamBlocksServer interface {
	Send(*StreamBlocksResponse) error
	grpc.ServerStream
}

type injectiveExplorerRPCStreamBlocksServer struct {
	grpc.ServerStream
}

func (x *injectiveExplorerRPCStreamBlocksServer) Send(m *StreamBlocksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveExplorerRPC_GetPeggyDepositTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeggyDepositTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetPeggyDepositTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetPeggyDepositTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetPeggyDepositTxs(ctx, req.(*GetPeggyDepositTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetPeggyWithdrawalTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeggyWithdrawalTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetPeggyWithdrawalTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetPeggyWithdrawalTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetPeggyWithdrawalTxs(ctx, req.(*GetPeggyWithdrawalTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveExplorerRPC_GetIBCTransferTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIBCTransferTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveExplorerRPCServer).GetIBCTransferTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_explorer_rpc.InjectiveExplorerRPC/GetIBCTransferTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveExplorerRPCServer).GetIBCTransferTxs(ctx, req.(*GetIBCTransferTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InjectiveExplorerRPC_ServiceDesc is the grpc.ServiceDesc for InjectiveExplorerRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InjectiveExplorerRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "injective_explorer_rpc.InjectiveExplorerRPC",
	HandlerType: (*InjectiveExplorerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountTxs",
			Handler:    _InjectiveExplorerRPC_GetAccountTxs_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _InjectiveExplorerRPC_GetBlocks_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _InjectiveExplorerRPC_GetBlock_Handler,
		},
		{
			MethodName: "GetValidator",
			Handler:    _InjectiveExplorerRPC_GetValidator_Handler,
		},
		{
			MethodName: "GetValidatorUptime",
			Handler:    _InjectiveExplorerRPC_GetValidatorUptime_Handler,
		},
		{
			MethodName: "GetCoinPriceData",
			Handler:    _InjectiveExplorerRPC_GetCoinPriceData_Handler,
		},
		{
			MethodName: "GetTxs",
			Handler:    _InjectiveExplorerRPC_GetTxs_Handler,
		},
		{
			MethodName: "GetTxByTxHash",
			Handler:    _InjectiveExplorerRPC_GetTxByTxHash_Handler,
		},
		{
			MethodName: "GetPeggyDepositTxs",
			Handler:    _InjectiveExplorerRPC_GetPeggyDepositTxs_Handler,
		},
		{
			MethodName: "GetPeggyWithdrawalTxs",
			Handler:    _InjectiveExplorerRPC_GetPeggyWithdrawalTxs_Handler,
		},
		{
			MethodName: "GetIBCTransferTxs",
			Handler:    _InjectiveExplorerRPC_GetIBCTransferTxs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTxs",
			Handler:       _InjectiveExplorerRPC_StreamTxs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamBlocks",
			Handler:       _InjectiveExplorerRPC_StreamBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "injective_explorer_rpc.proto",
}
