package routes

import (
	controllers "github.com/Bonifaceebuka/Blockchain-explorer-Backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(router *fiber.App) {

	router.Get("/", controllers.Home)
	router.Get("/getblockinfo/:block_id", controllers.GetBlockInfo)
	router.Get("/getblocks", controllers.GetBlocks)
	// router.Post("/signin", controllers.Signin)
	// router.Use(middlewares.IsAuthenticate)
}
