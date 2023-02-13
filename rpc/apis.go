// Package rpc contains RPC handler methods and utilities to start
// Stratos Web3-compatibly JSON-RPC server.
package rpc

import (
	"fmt"

	"github.com/tendermint/tendermint/node"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/stratosnet/stratos-chain/rpc/backend"
	"github.com/stratosnet/stratos-chain/rpc/namespaces/ethereum/debug"
	"github.com/stratosnet/stratos-chain/rpc/namespaces/ethereum/eth"
	"github.com/stratosnet/stratos-chain/rpc/namespaces/ethereum/eth/filters"
	"github.com/stratosnet/stratos-chain/rpc/namespaces/ethereum/miner"
	"github.com/stratosnet/stratos-chain/rpc/namespaces/ethereum/net"
	"github.com/stratosnet/stratos-chain/rpc/namespaces/ethereum/personal"
	"github.com/stratosnet/stratos-chain/rpc/namespaces/ethereum/txpool"
	"github.com/stratosnet/stratos-chain/rpc/namespaces/ethereum/web3"
	"github.com/stratosnet/stratos-chain/rpc/types"

	evmkeeper "github.com/stratosnet/stratos-chain/x/evm/keeper"
)

// RPC namespaces and API version
const (
	// Cosmos namespaces

	CosmosNamespace = "cosmos"

	// Ethereum namespaces

	Web3Namespace     = "web3"
	EthNamespace      = "eth"
	PersonalNamespace = "personal"
	NetNamespace      = "net"
	TxPoolNamespace   = "txpool"
	DebugNamespace    = "debug"
	MinerNamespace    = "miner"

	apiVersion = "1.0"
)

// APICreator creates the JSON-RPC API implementations.
type APICreator = func(*server.Context, *node.Node, *evmkeeper.Keeper, storetypes.MultiStore, client.Context) []rpc.API

// apiCreators defines the JSON-RPC API namespaces.
var apiCreators map[string]APICreator

func init() {
	apiCreators = map[string]APICreator{
		EthNamespace: func(ctx *server.Context, tmNode *node.Node, evmKeeper *evmkeeper.Keeper, ms storetypes.MultiStore, clientCtx client.Context) []rpc.API {
			nonceLock := new(types.AddrLocker)
			evmBackend := backend.NewBackend(ctx, tmNode, evmKeeper, ms, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: EthNamespace,
					Version:   apiVersion,
					Service:   eth.NewPublicAPI(ctx.Logger, clientCtx, evmBackend, nonceLock),
					Public:    true,
				},
				{
					Namespace: EthNamespace,
					Version:   apiVersion,
					Service:   filters.NewPublicAPI(ctx.Logger, clientCtx, tmNode.EventBus(), evmBackend),
					Public:    true,
				},
			}
		},
		Web3Namespace: func(*server.Context, *node.Node, *evmkeeper.Keeper, storetypes.MultiStore, client.Context) []rpc.API {
			return []rpc.API{
				{
					Namespace: Web3Namespace,
					Version:   apiVersion,
					Service:   web3.NewPublicAPI(),
					Public:    true,
				},
			}
		},
		NetNamespace: func(ctx *server.Context, tmNode *node.Node, evmKeeper *evmkeeper.Keeper, ms storetypes.MultiStore, clientCtx client.Context) []rpc.API {
			evmBackend := backend.NewBackend(ctx, tmNode, evmKeeper, ms, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: NetNamespace,
					Version:   apiVersion,
					Service:   net.NewPublicAPI(evmBackend),
					Public:    true,
				},
			}
		},
		PersonalNamespace: func(ctx *server.Context, tmNode *node.Node, evmKeeper *evmkeeper.Keeper, ms storetypes.MultiStore, clientCtx client.Context) []rpc.API {
			evmBackend := backend.NewBackend(ctx, tmNode, evmKeeper, ms, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: PersonalNamespace,
					Version:   apiVersion,
					Service:   personal.NewAPI(ctx.Logger, clientCtx, evmBackend),
					Public:    false,
				},
			}
		},
		TxPoolNamespace: func(ctx *server.Context, tmNode *node.Node, evmKeeper *evmkeeper.Keeper, ms storetypes.MultiStore, clientCtx client.Context) []rpc.API {
			evmBackend := backend.NewBackend(ctx, tmNode, evmKeeper, ms, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: TxPoolNamespace,
					Version:   apiVersion,
					Service:   txpool.NewPublicAPI(ctx.Logger, clientCtx, evmBackend),
					Public:    true,
				},
			}
		},
		DebugNamespace: func(ctx *server.Context, tmNode *node.Node, evmKeeper *evmkeeper.Keeper, ms storetypes.MultiStore, clientCtx client.Context) []rpc.API {
			evmBackend := backend.NewBackend(ctx, tmNode, evmKeeper, ms, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: DebugNamespace,
					Version:   apiVersion,
					Service:   debug.NewAPI(ctx, evmBackend, clientCtx),
					Public:    true,
				},
			}
		},
		MinerNamespace: func(ctx *server.Context, tmNode *node.Node, evmKeeper *evmkeeper.Keeper, ms storetypes.MultiStore, clientCtx client.Context) []rpc.API {
			evmBackend := backend.NewBackend(ctx, tmNode, evmKeeper, ms, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: MinerNamespace,
					Version:   apiVersion,
					Service:   miner.NewPrivateAPI(ctx, clientCtx, evmBackend),
					Public:    false,
				},
			}
		},
	}
}

// GetRPCAPIs returns the list of all APIs
func GetRPCAPIs(ctx *server.Context, tmNode *node.Node, evmKeeper *evmkeeper.Keeper, ms storetypes.MultiStore, clientCtx client.Context, selectedAPIs []string) []rpc.API {
	var apis []rpc.API

	for _, ns := range selectedAPIs {
		if creator, ok := apiCreators[ns]; ok {
			apis = append(apis, creator(ctx, tmNode, evmKeeper, ms, clientCtx)...)
		} else {
			ctx.Logger.Error("invalid namespace value", "namespace", ns)
		}
	}

	return apis
}

// RegisterAPINamespace registers a new API namespace with the API creator.
// This function fails if the namespace is already registered.
func RegisterAPINamespace(ns string, creator APICreator) error {
	if _, ok := apiCreators[ns]; ok {
		return fmt.Errorf("duplicated api namespace %s", ns)
	}
	apiCreators[ns] = creator
	return nil
}
