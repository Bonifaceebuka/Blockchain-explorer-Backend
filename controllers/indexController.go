package controllers

import (
	"github.com/Bonifaceebuka/Blockchain-explorer-Backend/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

var infuraConnection *ethclient.Client

func init() {
	infuraConnection = config.GetInfuraConnection()
	config.LoadEnv()
}

func Home(c *fiber.Ctx) error {
	c.Status(200)

	return c.JSON(fiber.Map{
		"msg": "API service is fully up!",
	})

}
