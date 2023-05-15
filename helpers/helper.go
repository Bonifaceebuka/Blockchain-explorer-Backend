package helpers

import (
	"context"
	"math/big"

	"github.com/Bonifaceebuka/Blockchain-explorer-Backend/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const cantConvertValue = 1 // unable to fetch latest block
const cantFetchABlock = 2  // unable to fetch a particular block
const tnxIsPending = 3     // Transaction is till pending
const invalidHash = 4      // Invalid TNX hash

// const cantFetchLatestBlock = 1
// const cantFetchLatestBlock = 1

func GetBlockDetail(block_id string) any {
	newBigInt := new(big.Int)
	newBigInt, ok := newBigInt.SetString(block_id, 10)

	if !ok {
		return cantConvertValue
	}

	infuraConnection := config.GetInfuraConnection()
	blockData, err := infuraConnection.BlockByNumber(context.Background(), newBigInt)

	if err != nil {
		return cantFetchABlock
	}

	return blockData.Transactions()
}

func GetTnxByHash(hash string) any {
	_, err := hexutil.Decode(hash)
	if err != nil {
		return invalidHash
	}
	infuraConnection := config.GetInfuraConnection()
	tnx, isPending, _ := infuraConnection.TransactionByHash(context.Background(), common.HexToHash(hash))

	if isPending {
		return tnxIsPending
	}

	return tnx
}

// func GetTnxByAddress(address string) any {
// 	covalentBaseUrl := "https://api.covalenthq.com/v1/{chainName}/transaction_v2/{txHash}/"
// 	_, err := hexutil.Decode(address)
// 	if err != nil {
// 		return invalidAddress
// 	}
// 	infuraConnection := config.GetInfuraConnection()
// 	tnx, isPending, _ := infuraConnection.TransactionByAddress(context.Background(), common.HexToAddress(address))

// 	if isPending {
// 		return tnxIsPending
// 	}

// 	return tnx
// }
