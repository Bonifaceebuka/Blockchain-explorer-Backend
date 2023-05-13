package controllers

import (
	"context"
	"fmt"

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
		fmt.Println("one")
	case "address":
		fmt.Println("two")
	case "hash":
		fmt.Println("three")
	}
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
