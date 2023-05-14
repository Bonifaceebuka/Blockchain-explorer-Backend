package controllers

import (
	"context"
	"fmt"

	"github.com/Bonifaceebuka/Blockchain-explorer-Backend/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetTransactions(c *fiber.Ctx) error {
	blockData, err := infuraConnection.BlockByNumber(context.Background(), nil)

	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"msg": "Unable to fetch the latest block information!",
		})

	}

	return c.JSON(fiber.Map{
		"latest_transactions": blockData.Transactions(),
	})
}

func GetSearchResults(c *fiber.Ctx) error {
	searchType := c.Query("search_type")
	searchValue := c.Query("search_value")
	var response any
	statusCode := 0

	if searchType == "" {
		c.Status(302)
		return c.JSON(fiber.Map{
			"msg": "Search type is not defined. Possible values are: block, address, hash!",
		})
	}

	if searchValue == "" {
		c.Status(302)
		return c.JSON(fiber.Map{
			"msg": "Search value is not defined!",
		})
	}

	switch searchType {
	case "block":
		response = helpers.GetBlockDetail(searchValue)
		switch response {
		case 1:
			statusCode = 500
			response = helpers.GetMessage("cant_convert_value")
		case 2:
			statusCode = 500
			response = helpers.GetMessage("fetch_a_block")
		}

	case "address":
		fmt.Println("two")
	case "hash":
		response = helpers.GetTnxByHash(searchValue)
		switch response {
		case 3:
			statusCode = 200
			response = helpers.GetMessage("tnx_is_pending")
		case 4:
			statusCode = 302
			response = helpers.GetMessage("invalid_tnx_hash")
		}
	}

	c.Status(statusCode)
	return c.JSON(fiber.Map{
		"response": response,
	})
}
