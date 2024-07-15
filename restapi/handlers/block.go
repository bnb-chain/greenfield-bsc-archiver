package handlers

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-openapi/runtime/middleware"
	"greeenfield-bsc-archiver/models"
	"greeenfield-bsc-archiver/service"
	"greeenfield-bsc-archiver/util"

	"greeenfield-bsc-archiver/restapi/operations/block"
)

func HandleGetBundledBlockByNumber() func(params block.GetBundledBlockByNumberParams) middleware.Responder {
	return func(params block.GetBundledBlockByNumberParams) middleware.Responder {
		rpcRequest := params.Body
		if rpcRequest.Params == nil {
			return block.NewGetBundledBlockByNumberOK().WithPayload(
				&models.GetBundledBlockByNumberRPCResponse{
					ID:      rpcRequest.ID,
					Jsonrpc: rpcRequest.Jsonrpc,
					Error: &models.RPCError{
						Code:    -32600,
						Message: "Invalid request",
					},
				},
			)
		}

		switch rpcRequest.Method {

		default:
			return block.NewGetBundledBlockByNumberOK().WithPayload(
				&models.GetBundledBlockByNumberRPCResponse{
					ID:      rpcRequest.ID,
					Jsonrpc: rpcRequest.Jsonrpc,
					Error: &models.RPCError{
						Code:    -32601,
						Message: "method not supported",
					},
				},
			)
		}
	}
}

func HandleGetBundleStartBlockID() func(params block.GetBundleStartBlockIDParams) middleware.Responder {
	return func(params block.GetBundleStartBlockIDParams) middleware.Responder {
		blockID := params.BlockID
		blockNum, err := util.HexToUint64(blockID)
		if err != nil {
			return block.NewGetBundleStartBlockIDBadRequest().WithPayload(service.BadRequestWithError(fmt.Errorf("invalid block id")))
		}
		name, err := service.BlockSvc.GetBundleNameByBlockID(blockNum)
		if err != nil {
			return block.NewGetBundleStartBlockIDInternalServerError().WithPayload(service.InternalErrorWithError(err))
		}

		response := &models.GetBundleStartBlockIDRPCResponse{
			Data: name,
		}
		return block.NewGetBundleStartBlockIDOK().WithPayload(response)
	}
}

func HandleGetBlockByNumber() func(params block.GetBlockByNumberParams) middleware.Responder {
	return func(params block.GetBlockByNumberParams) middleware.Responder {
		rpcRequest := params.Body
		if rpcRequest.Params == nil {
			return block.NewGetBlockByNumberOK().WithPayload(
				&models.GetBlockByNumberRPCResponse{
					ID:      rpcRequest.ID,
					Jsonrpc: rpcRequest.Jsonrpc,
					Error: &models.RPCError{
						Code:    -32600,
						Message: "Invalid request",
					},
				},
			)
		}

		switch rpcRequest.Method {
		case "eth_getBlockByNumber":
			blockNum, err := util.HexToUint64(rpcRequest.Params[0])
			if err != nil {
				return block.NewGetBlockByNumberOK().WithPayload(
					&models.GetBlockByNumberRPCResponse{
						ID:      rpcRequest.ID,
						Jsonrpc: rpcRequest.Jsonrpc,
						Error: &models.RPCError{
							Code:    -32602,
							Message: "invalid argument",
						},
					},
				)
			}
			blockInfo, err := service.BlockSvc.GetBlockByBlockNumber(blockNum)
			if err != nil {
				return block.NewGetBlockByNumberInternalServerError().WithPayload(service.InternalErrorWithError(err))
			}

			response := &models.GetBlockByNumberRPCResponse{
				ID:      rpcRequest.ID,
				Jsonrpc: rpcRequest.Jsonrpc,
				Result:  blockInfo,
			}
			return block.NewGetBlockByNumberOK().WithPayload(response)
		case "eth_getBundledBlockByNumber":
			blockNum, err := util.HexToUint64(rpcRequest.Params[0])
			if err != nil {
				return block.NewGetBundledBlockByNumberOK().WithPayload(
					&models.GetBundledBlockByNumberRPCResponse{
						ID:      rpcRequest.ID,
						Jsonrpc: rpcRequest.Jsonrpc,
						Error: &models.RPCError{
							Code:    -32602,
							Message: "invalid argument",
						},
					},
				)
			}
			blocksInfo, err := service.BlockSvc.GetBundledBlockByBlockNumber(blockNum)
			if err != nil {
				return block.NewGetBundledBlockByNumberInternalServerError().WithPayload(service.InternalErrorWithError(err))
			}

			response := &models.GetBundledBlockByNumberRPCResponse{
				ID:      rpcRequest.ID,
				Jsonrpc: rpcRequest.Jsonrpc,
				Result:  blocksInfo,
			}
			return block.NewGetBundledBlockByNumberOK().WithPayload(response)
		case "eth_getBlockByHash":
			valid := util.IsHexHash(rpcRequest.Params[0])
			if !valid {
				return block.NewGetBlockByHashOK().WithPayload(
					&models.GetBlockByHashRPCResponse{
						ID:      rpcRequest.ID,
						Jsonrpc: rpcRequest.Jsonrpc,
						Error: &models.RPCError{
							Code:    -32602,
							Message: "invalid argument",
						},
					},
				)
			}
			blockInfo, err := service.BlockSvc.GetBlockByBlockHash(common.HexToHash(rpcRequest.Params[0]))
			if err != nil {
				return block.NewGetBlockByHashInternalServerError().WithPayload(service.InternalErrorWithError(err))
			}

			response := &models.GetBlockByHashRPCResponse{
				ID:      rpcRequest.ID,
				Jsonrpc: rpcRequest.Jsonrpc,
				Result:  blockInfo,
			}
			return block.NewGetBlockByHashOK().WithPayload(response)
		case "eth_blockNumber":
			number, err := service.BlockSvc.GetLatestBlockNumber()
			if err != nil {
				return block.NewGetBlockNumberInternalServerError().WithPayload(service.InternalErrorWithError(err))
			}

			response := &models.GetBlockNumberRPCResponse{
				ID:      rpcRequest.ID,
				Jsonrpc: rpcRequest.Jsonrpc,
				Result:  util.Uint64ToHex(number),
			}
			return block.NewGetBlockNumberOK().WithPayload(response)
		default:
			return block.NewGetBlockNumberOK().WithPayload(
				&models.GetBlockNumberRPCResponse{
					ID:      rpcRequest.ID,
					Jsonrpc: rpcRequest.Jsonrpc,
					Error: &models.RPCError{
						Code:    -32601,
						Message: "method not supported",
					},
				},
			)
		}
	}
}
