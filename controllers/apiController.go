package controllers

import (
	"context"
	"log"
	"math/big"

	"github.com/Bonifaceebuka/Blockchain-explorer-Backend/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

var infuraConnection *ethclient.Client

func init() {
	infuraConnection = config.GetInfuraConnection()
	config.LoadEnv()
}

type Block struct {
	BlockId       int    `json:"block_id"`
	BlockNumber   uint64 `json:"block_number"`
	BlockTime     uint64 `json:"block_time"`
	BlockTotalTnx int    `json:"block_total_tnx"`
}

func Home(c *fiber.Ctx) error {
	c.Status(200)

	return c.JSON(fiber.Map{
		"msg": "API service is fully up!",
	})

}

func GetBlockInfo(c *fiber.Ctx) error {
	block_id := c.Params("block_id")
	newBigInt := new(big.Int)
	newBigInt, ok := newBigInt.SetString(block_id, 10)

	if !ok {

	}

	blockData, err := infuraConnection.BlockByNumber(context.Background(), newBigInt)

	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"msg": "Unable to fetch the latest block information!",
		})

	}

	return c.JSON(fiber.Map{
		"blockData": blockData.Transactions(),
	})
}

func GetBlocks(c *fiber.Ctx) error {
	blockHeader, err := infuraConnection.HeaderByNumber(context.Background(), nil)

	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"msg": "Unable to fetch the latest block information!",
		})

	}

	latestBlockNumber := blockHeader.Number.Uint64()
	latestBlockNumber = latestBlockNumber - 1
	pagingBlock := uint64(latestBlockNumber - 12)

	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"msg": "Unable to fetch the data of the latest block!",
		})

	}

	var blocks Block

	if err != nil {

		log.Fatal(err)
	}

	j := 0
	var arrBlock []Block
	for i := latestBlockNumber; i <= latestBlockNumber && i >= pagingBlock; i-- {
		currentBlock := new(big.Int).SetUint64(i)
		latestBlockData, err := infuraConnection.BlockByNumber(context.Background(), currentBlock)

		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{
				"msg": "Unable to fetch the data of this block: " + currentBlock.String(),
			})
		}

		j = j + 1
		blocks.BlockId = j
		blocks.BlockNumber = i
		blocks.BlockTime = latestBlockData.Time()
		blocks.BlockTotalTnx = len(latestBlockData.Transactions())

		arrBlock = append(arrBlock, blocks)
	}

	return c.JSON(fiber.Map{
		"latestBlocks": arrBlock,
	})
}
