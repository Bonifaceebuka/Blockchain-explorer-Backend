package main

import (
	"os"

	"github.com/Bonifaceebuka/Blockchain-explorer-Backend/config"
	"github.com/Bonifaceebuka/Blockchain-explorer-Backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New()

	config.LoadEnv()
	port := os.Getenv("PORT")
	routes.Router(router)
	router.Listen(":" + port)

}
