package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/handlers"
)

func RegisterPhotoRoutes(app *fiber.App) {
	app.Get("/api/photos", handlers.GetAllPhotos)
}
