package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/handlers"
)

func RegisterAuthRoutes(app *fiber.App) {
	app.Post("/api/login", handlers.Login)
}
