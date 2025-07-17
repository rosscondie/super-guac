package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/handlers"
)

func RegisterAlbumRoutes(app *fiber.App) {
	album := app.Group("/api/albums")

	// Public routes
	album.Get("/", handlers.GetAllAlbums)
	album.Get("/:slug", handlers.GetAlbumBySlug)

	// Protected routes
	album.Post("/", handlers.Protected(), handlers.CreateAlbum)
	album.Post("/:slug/photos", handlers.Protected(), handlers.UploadPhotoToAlbum)
	album.Delete("/:slug", handlers.Protected(), handlers.DeleteAlbumHandler)
	album.Delete("/:slug/photos/:filename", handlers.Protected(), handlers.DeletePhotoFromAlbumHandler)
	album.Put("/:slug", handlers.Protected(), handlers.UpdateAlbumMetadataHandler)
}
