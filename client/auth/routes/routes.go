package routes

import (
	"github.com/gofiber/fiber/v2"
	"golangtest.com/client/auth/handler"
)

func Routes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/login", handler.Login)
}
