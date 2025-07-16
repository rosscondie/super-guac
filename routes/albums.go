package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/handlers"
)

func RegisterAlbumRoutes(app *fiber.App) {
	album := app.Group("/api/albums")

	album.Get("/", handlers.GetAllAlbums)
	album.Get("/:slug", handlers.GetAlbumBySlug)

	album.Post("/", handlers.CreateAlbum)
	album.Post("/:slug/photos", handlers.UploadPhotoToAlbum)

	album.Delete("/:slug", handlers.DeleteAlbumHandler)
	album.Delete("/:slug/photos/:filename", handlers.DeletePhotoFromAlbumHandler)
}
