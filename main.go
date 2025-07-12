package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rosscondie/photo-blog/routes"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	// Mount routes
	routes.RegisterPhotoRoutes(app)
	routes.RegisterPostRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
