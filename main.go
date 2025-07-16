package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rosscondie/photo-blog/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_ORIGIN"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/", "./public")
	app.Static("/photos", "./content/photos")

	// Mount routes
	routes.RegisterPhotoRoutes(app)
	routes.RegisterPostRoutes(app)
	routes.RegisterAlbumRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
