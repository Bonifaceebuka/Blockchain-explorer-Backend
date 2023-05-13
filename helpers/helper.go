package helpers

import (
	"context"
	"math/big"

	"github.com/Bonifaceebuka/Blockchain-explorer-Backend/config"
)

const cantFetchLatestBlock = 1
const cantFetchABlock = 2

// const cantFetchLatestBlock = 1
// const cantFetchLatestBlock = 1

func GetBlockDetail(block_id string) any {
	newBigInt := new(big.Int)
	newBigInt, ok := newBigInt.SetString(block_id, 10)

	if !ok {
		return cantFetchABlock
	}

	infuraConnection := config.GetInfuraConnection()
	blockData, err := infuraConnection.BlockByNumber(context.Background(), newBigInt)

	if err != nil {
		return cantFetchLatestBlock
	}

	return blockData.Transactions()
}
