package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/data"
)

func GetAllPosts(c *fiber.Ctx) error {
	posts := data.GetAllPosts()
	return c.JSON(posts)
}

func GetPostBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	post, err := data.GetPostBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}
	return c.JSON(post)
}
