package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/handlers"
)

func RegisterPostRoutes(app *fiber.App) {
	app.Get("/api/posts", handlers.GetAllPosts)
	app.Get("/api/posts/:slug", handlers.GetPostBySlug)
}
