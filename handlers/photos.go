package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/data"
)

func GetAllPhotos(c *fiber.Ctx) error {
	photos, err := data.GetAllPhotos()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not load photos",
		})
	}
	return c.JSON(photos)
}
