package routes

import (
	controllers "github.com/Bonifaceebuka/Blockchain-explorer-Backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(router *fiber.App) {

	router.Get("/", controllers.Home)
	router.Get("/getBlockInfo/:block_id", controllers.GetBlockInfo)
	router.Get("/getBlocks", controllers.GetBlocks)
	router.Get("/getTransactions", controllers.GetTransactions)
	router.Get("/getSearchResults", controllers.GetSearchResults)
	// router.Post("/signin", controllers.Signin)
	// router.Use(middlewares.IsAuthenticate)
}
