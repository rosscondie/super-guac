package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/data"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	// GET /api/posts
	app.Get("/api/posts", func(c *fiber.Ctx) error {
		posts := data.GetAllPosts()
		return c.JSON(posts)
	})

	// Get /api/posts/:slug
	app.Get("/api/posts/:slug", func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		post, err := data.GetPostBySlug(slug)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Post not found",
			})
		}
		return c.JSON(post)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
