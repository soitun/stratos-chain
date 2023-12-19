// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stratos/evm/v1/query.proto

package evmv1

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
	Query_Account_FullMethodName          = "/stratos.evm.v1.Query/Account"
	Query_CosmosAccount_FullMethodName    = "/stratos.evm.v1.Query/CosmosAccount"
	Query_ValidatorAccount_FullMethodName = "/stratos.evm.v1.Query/ValidatorAccount"
	Query_Balance_FullMethodName          = "/stratos.evm.v1.Query/Balance"
	Query_Storage_FullMethodName          = "/stratos.evm.v1.Query/Storage"
	Query_Code_FullMethodName             = "/stratos.evm.v1.Query/Code"
	Query_Params_FullMethodName           = "/stratos.evm.v1.Query/Params"
	Query_EthCall_FullMethodName          = "/stratos.evm.v1.Query/EthCall"
	Query_EstimateGas_FullMethodName      = "/stratos.evm.v1.Query/EstimateGas"
	Query_TraceTx_FullMethodName          = "/stratos.evm.v1.Query/TraceTx"
	Query_TraceBlock_FullMethodName       = "/stratos.evm.v1.Query/TraceBlock"
	Query_BaseFee_FullMethodName          = "/stratos.evm.v1.Query/BaseFee"
	Query_BaseFeeParam_FullMethodName     = "/stratos.evm.v1.Query/BaseFeeParam"
	Query_BlockGas_FullMethodName         = "/stratos.evm.v1.Query/BlockGas"
	Query_AddressConvert_FullMethodName   = "/stratos.evm.v1.Query/AddressConvert"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Account queries an Ethereum account.
	Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*QueryAccountResponse, error)
	// CosmosAccount queries an Ethereum account's Cosmos Address.
	CosmosAccount(ctx context.Context, in *QueryCosmosAccountRequest, opts ...grpc.CallOption) (*QueryCosmosAccountResponse, error)
	// ValidatorAccount queries an Ethereum account's from a validator consensus
	// Address.
	ValidatorAccount(ctx context.Context, in *QueryValidatorAccountRequest, opts ...grpc.CallOption) (*QueryValidatorAccountResponse, error)
	// Balance queries the balance of a the EVM denomination for a single
	// EthAccount.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// Storage queries the balance of all coins for a single account.
	Storage(ctx context.Context, in *QueryStorageRequest, opts ...grpc.CallOption) (*QueryStorageResponse, error)
	// Code queries the balance of all coins for a single account.
	Code(ctx context.Context, in *QueryCodeRequest, opts ...grpc.CallOption) (*QueryCodeResponse, error)
	// Params queries the parameters of x/evm module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// EthCall implements the `eth_call` rpc api
	EthCall(ctx context.Context, in *EthCallRequest, opts ...grpc.CallOption) (*MsgEthereumTxResponse, error)
	// EstimateGas implements the `eth_estimateGas` rpc api
	EstimateGas(ctx context.Context, in *EthCallRequest, opts ...grpc.CallOption) (*EstimateGasResponse, error)
	// TraceTx implements the `debug_traceTransaction` rpc api
	TraceTx(ctx context.Context, in *QueryTraceTxRequest, opts ...grpc.CallOption) (*QueryTraceTxResponse, error)
	// TraceBlock implements the `debug_traceBlockByNumber` and `debug_traceBlockByHash` rpc api
	TraceBlock(ctx context.Context, in *QueryTraceBlockRequest, opts ...grpc.CallOption) (*QueryTraceBlockResponse, error)
	// BaseFee queries the base fee of the parent block of the current block,
	// it's similar to feemarket module's method, but also checks london hardfork status.
	BaseFee(ctx context.Context, in *QueryBaseFeeRequest, opts ...grpc.CallOption) (*QueryBaseFeeResponse, error)
	// BaseFeeParam queries the base fee of the parent block of the current block.
	BaseFeeParam(ctx context.Context, in *QueryBaseFeeRequest, opts ...grpc.CallOption) (*QueryBaseFeeResponse, error)
	// BlockGas queries the gas used at a given block height
	BlockGas(ctx context.Context, in *QueryBlockGasRequest, opts ...grpc.CallOption) (*QueryBlockGasResponse, error)
	AddressConvert(ctx context.Context, in *QueryAddressConvertRequest, opts ...grpc.CallOption) (*QueryAddressConvertResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*QueryAccountResponse, error) {
	out := new(QueryAccountResponse)
	err := c.cc.Invoke(ctx, Query_Account_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CosmosAccount(ctx context.Context, in *QueryCosmosAccountRequest, opts ...grpc.CallOption) (*QueryCosmosAccountResponse, error) {
	out := new(QueryCosmosAccountResponse)
	err := c.cc.Invoke(ctx, Query_CosmosAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorAccount(ctx context.Context, in *QueryValidatorAccountRequest, opts ...grpc.CallOption) (*QueryValidatorAccountResponse, error) {
	out := new(QueryValidatorAccountResponse)
	err := c.cc.Invoke(ctx, Query_ValidatorAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, Query_Balance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Storage(ctx context.Context, in *QueryStorageRequest, opts ...grpc.CallOption) (*QueryStorageResponse, error) {
	out := new(QueryStorageResponse)
	err := c.cc.Invoke(ctx, Query_Storage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Code(ctx context.Context, in *QueryCodeRequest, opts ...grpc.CallOption) (*QueryCodeResponse, error) {
	out := new(QueryCodeResponse)
	err := c.cc.Invoke(ctx, Query_Code_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EthCall(ctx context.Context, in *EthCallRequest, opts ...grpc.CallOption) (*MsgEthereumTxResponse, error) {
	out := new(MsgEthereumTxResponse)
	err := c.cc.Invoke(ctx, Query_EthCall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EstimateGas(ctx context.Context, in *EthCallRequest, opts ...grpc.CallOption) (*EstimateGasResponse, error) {
	out := new(EstimateGasResponse)
	err := c.cc.Invoke(ctx, Query_EstimateGas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TraceTx(ctx context.Context, in *QueryTraceTxRequest, opts ...grpc.CallOption) (*QueryTraceTxResponse, error) {
	out := new(QueryTraceTxResponse)
	err := c.cc.Invoke(ctx, Query_TraceTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TraceBlock(ctx context.Context, in *QueryTraceBlockRequest, opts ...grpc.CallOption) (*QueryTraceBlockResponse, error) {
	out := new(QueryTraceBlockResponse)
	err := c.cc.Invoke(ctx, Query_TraceBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BaseFee(ctx context.Context, in *QueryBaseFeeRequest, opts ...grpc.CallOption) (*QueryBaseFeeResponse, error) {
	out := new(QueryBaseFeeResponse)
	err := c.cc.Invoke(ctx, Query_BaseFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BaseFeeParam(ctx context.Context, in *QueryBaseFeeRequest, opts ...grpc.CallOption) (*QueryBaseFeeResponse, error) {
	out := new(QueryBaseFeeResponse)
	err := c.cc.Invoke(ctx, Query_BaseFeeParam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BlockGas(ctx context.Context, in *QueryBlockGasRequest, opts ...grpc.CallOption) (*QueryBlockGasResponse, error) {
	out := new(QueryBlockGasResponse)
	err := c.cc.Invoke(ctx, Query_BlockGas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AddressConvert(ctx context.Context, in *QueryAddressConvertRequest, opts ...grpc.CallOption) (*QueryAddressConvertResponse, error) {
	out := new(QueryAddressConvertResponse)
	err := c.cc.Invoke(ctx, Query_AddressConvert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Account queries an Ethereum account.
	Account(context.Context, *QueryAccountRequest) (*QueryAccountResponse, error)
	// CosmosAccount queries an Ethereum account's Cosmos Address.
	CosmosAccount(context.Context, *QueryCosmosAccountRequest) (*QueryCosmosAccountResponse, error)
	// ValidatorAccount queries an Ethereum account's from a validator consensus
	// Address.
	ValidatorAccount(context.Context, *QueryValidatorAccountRequest) (*QueryValidatorAccountResponse, error)
	// Balance queries the balance of a the EVM denomination for a single
	// EthAccount.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// Storage queries the balance of all coins for a single account.
	Storage(context.Context, *QueryStorageRequest) (*QueryStorageResponse, error)
	// Code queries the balance of all coins for a single account.
	Code(context.Context, *QueryCodeRequest) (*QueryCodeResponse, error)
	// Params queries the parameters of x/evm module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// EthCall implements the `eth_call` rpc api
	EthCall(context.Context, *EthCallRequest) (*MsgEthereumTxResponse, error)
	// EstimateGas implements the `eth_estimateGas` rpc api
	EstimateGas(context.Context, *EthCallRequest) (*EstimateGasResponse, error)
	// TraceTx implements the `debug_traceTransaction` rpc api
	TraceTx(context.Context, *QueryTraceTxRequest) (*QueryTraceTxResponse, error)
	// TraceBlock implements the `debug_traceBlockByNumber` and `debug_traceBlockByHash` rpc api
	TraceBlock(context.Context, *QueryTraceBlockRequest) (*QueryTraceBlockResponse, error)
	// BaseFee queries the base fee of the parent block of the current block,
	// it's similar to feemarket module's method, but also checks london hardfork status.
	BaseFee(context.Context, *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error)
	// BaseFeeParam queries the base fee of the parent block of the current block.
	BaseFeeParam(context.Context, *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error)
	// BlockGas queries the gas used at a given block height
	BlockGas(context.Context, *QueryBlockGasRequest) (*QueryBlockGasResponse, error)
	AddressConvert(context.Context, *QueryAddressConvertRequest) (*QueryAddressConvertResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Account(context.Context, *QueryAccountRequest) (*QueryAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (UnimplementedQueryServer) CosmosAccount(context.Context, *QueryCosmosAccountRequest) (*QueryCosmosAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CosmosAccount not implemented")
}
func (UnimplementedQueryServer) ValidatorAccount(context.Context, *QueryValidatorAccountRequest) (*QueryValidatorAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorAccount not implemented")
}
func (UnimplementedQueryServer) Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedQueryServer) Storage(context.Context, *QueryStorageRequest) (*QueryStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Storage not implemented")
}
func (UnimplementedQueryServer) Code(context.Context, *QueryCodeRequest) (*QueryCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) EthCall(context.Context, *EthCallRequest) (*MsgEthereumTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthCall not implemented")
}
func (UnimplementedQueryServer) EstimateGas(context.Context, *EthCallRequest) (*EstimateGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateGas not implemented")
}
func (UnimplementedQueryServer) TraceTx(context.Context, *QueryTraceTxRequest) (*QueryTraceTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TraceTx not implemented")
}
func (UnimplementedQueryServer) TraceBlock(context.Context, *QueryTraceBlockRequest) (*QueryTraceBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TraceBlock not implemented")
}
func (UnimplementedQueryServer) BaseFee(context.Context, *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaseFee not implemented")
}
func (UnimplementedQueryServer) BaseFeeParam(context.Context, *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaseFeeParam not implemented")
}
func (UnimplementedQueryServer) BlockGas(context.Context, *QueryBlockGasRequest) (*QueryBlockGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockGas not implemented")
}
func (UnimplementedQueryServer) AddressConvert(context.Context, *QueryAddressConvertRequest) (*QueryAddressConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressConvert not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Account_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Account(ctx, req.(*QueryAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CosmosAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCosmosAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CosmosAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CosmosAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CosmosAccount(ctx, req.(*QueryCosmosAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ValidatorAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorAccount(ctx, req.(*QueryValidatorAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Storage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Storage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Storage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Storage(ctx, req.(*QueryStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Code_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Code(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Code_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Code(ctx, req.(*QueryCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EthCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EthCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EthCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EthCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EthCall(ctx, req.(*EthCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EstimateGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EthCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EstimateGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EstimateGas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EstimateGas(ctx, req.(*EthCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TraceTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTraceTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TraceTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TraceTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TraceTx(ctx, req.(*QueryTraceTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TraceBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTraceBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TraceBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TraceBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TraceBlock(ctx, req.(*QueryTraceBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BaseFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBaseFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BaseFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BaseFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BaseFee(ctx, req.(*QueryBaseFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BaseFeeParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBaseFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BaseFeeParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BaseFeeParam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BaseFeeParam(ctx, req.(*QueryBaseFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BlockGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlockGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BlockGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BlockGas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BlockGas(ctx, req.(*QueryBlockGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AddressConvert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAddressConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AddressConvert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AddressConvert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AddressConvert(ctx, req.(*QueryAddressConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stratos.evm.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Account",
			Handler:    _Query_Account_Handler,
		},
		{
			MethodName: "CosmosAccount",
			Handler:    _Query_CosmosAccount_Handler,
		},
		{
			MethodName: "ValidatorAccount",
			Handler:    _Query_ValidatorAccount_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "Storage",
			Handler:    _Query_Storage_Handler,
		},
		{
			MethodName: "Code",
			Handler:    _Query_Code_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "EthCall",
			Handler:    _Query_EthCall_Handler,
		},
		{
			MethodName: "EstimateGas",
			Handler:    _Query_EstimateGas_Handler,
		},
		{
			MethodName: "TraceTx",
			Handler:    _Query_TraceTx_Handler,
		},
		{
			MethodName: "TraceBlock",
			Handler:    _Query_TraceBlock_Handler,
		},
		{
			MethodName: "BaseFee",
			Handler:    _Query_BaseFee_Handler,
		},
		{
			MethodName: "BaseFeeParam",
			Handler:    _Query_BaseFeeParam_Handler,
		},
		{
			MethodName: "BlockGas",
			Handler:    _Query_BlockGas_Handler,
		},
		{
			MethodName: "AddressConvert",
			Handler:    _Query_AddressConvert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stratos/evm/v1/query.proto",
}
