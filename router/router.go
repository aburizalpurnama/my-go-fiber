package router

import (
	"github.com/aburizalpurnama/my-go-fiber/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, handler handler.Handler, userHandler handler.UserHandler) {
	api := app.Group("/api", logger.New())
	v1 := api.Group("/v1")
	v1.Get("/greet/:name", handler.Greet)

	// user routes
	user := v1.Group("/user")
	user.Post("/", userHandler.CreateUser)
}
