package handlers

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-openapi/runtime/middleware"
	"greeenfield-bsc-archiver/models"
	"greeenfield-bsc-archiver/service"
	"greeenfield-bsc-archiver/types"
	"greeenfield-bsc-archiver/util"

	"greeenfield-bsc-archiver/restapi/operations/block"
)

func HandleGetBundleNameByBlockNumber() func(params block.GetBundleNameByBlockNumberParams) middleware.Responder {
	return func(params block.GetBundleNameByBlockNumberParams) middleware.Responder {
		blockID := params.BlockNumber
		blockNum, err := util.HexToUint64(blockID)
		if err != nil {
			return block.NewGetBundleNameByBlockNumberBadRequest().WithPayload(service.BadRequestWithError(fmt.Errorf("invalid block number")))
		}
		name, err := service.BlockSvc.GetBundleNameByBlockID(blockNum)
		if err != nil {
			return block.NewGetBundleNameByBlockNumberInternalServerError().WithPayload(service.InternalErrorWithError(err))
		}

		response := &models.GetBundleNameByBlockNumberRPCResponse{
			Data: name,
		}
		return block.NewGetBundleNameByBlockNumberOK().WithPayload(response)
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
			var (
				blockNum     uint64
				err          error
				isFullTxList bool
			)
			if len(rpcRequest.Params) != 2 {
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
			isFullTxList, err = strconv.ParseBool(rpcRequest.Params[1])
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
			switch rpcRequest.Params[0] {
			case "latest":
				blockNum, err = service.BlockSvc.GetLatestVerifiedBlockNumber()
				if err != nil {
					return block.NewGetBlockByNumberInternalServerError().WithPayload(service.InternalErrorWithError(err))
				}
			default:
				blockNum, err = util.HexToUint64(rpcRequest.Params[0])
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
			}
			blockInfo, err := service.BlockSvc.GetBlockByBlockNumber(blockNum)
			if err != nil {
				return block.NewGetBlockByNumberInternalServerError().WithPayload(service.InternalErrorWithError(err))
			}

			if !isFullTxList {
				response := &models.GetBlockByNumberSimplifiedRPCResponse{
					ID:      rpcRequest.ID,
					Jsonrpc: rpcRequest.Jsonrpc,
					Result:  types.BlockToSimplifiedBlock(blockInfo),
				}
				return block.NewGetBlockByNumberSimplifiedOK().WithPayload(response)
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
