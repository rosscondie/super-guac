package handlers

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/data"
)

// GET /api/albums
func GetAllAlbums(c *fiber.Ctx) error {
	albums, err := data.GetAllAlbums()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load albums",
		})
	}
	return c.JSON(albums)
}

// GET /api/albums/:slug
func GetAlbumBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	photos, err := data.GetPhotosByAlbum(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Album not found",
		})
	}

	metadata, err := data.GetAlbumMetadata(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Album metadata not found",
		})
	}

	return c.JSON(fiber.Map{
		"photos":   photos,
		"metadata": metadata,
	})
}

func CreateAlbum(c *fiber.Ctx) error {
	var album data.Album
	if err := c.BodyParser(&album); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	if album.Slug == "" || album.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and slug are required"})
	}
	if err := data.CreateAlbum(album); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusCreated)
}

func UploadPhotoToAlbum(c *fiber.Ctx) error {
	albumSlug := c.Params("slug")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form data"})
	}

	files := form.File["photo"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No photo uploaded"})
	}

	// Only support uploading one photo at a time for now
	file := files[0]
	savePath := filepath.Join("content/photos", albumSlug, file.Filename)

	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save photo"})
	}

	return c.JSON(fiber.Map{"message": "Photo uploaded successfully", "filename": file.Filename})
}

func DeleteAlbumHandler(c *fiber.Ctx) error {
	slug := c.Params("slug")

	err := data.DeleteAlbum(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
}

func DeletePhotoFromAlbumHandler(c *fiber.Ctx) error {
	slug := c.Params("slug")
	filename := c.Params("filename")

	err := data.DeletePhotoFromAlbum(slug, filename)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateAlbumMetadataHandler(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var metadata data.AlbumMetadata
	if err := c.BodyParser(&metadata); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	// Set the slug in metadata to match the URL param (in case it's missing or inconsistent)
	metadata.Slug = slug

	if err := data.UpdateAlbumMetadata(slug, metadata); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Album metadata updated",
	})
}
