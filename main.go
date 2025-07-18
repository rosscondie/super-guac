package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rosscondie/photo-blog/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONTEND_ORIGIN"),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))

	app.Static("/", "./public")
	app.Static("/photos", "./content/photos")

	// Mount routes
	routes.RegisterPhotoRoutes(app)
	routes.RegisterPostRoutes(app)
	routes.RegisterAlbumRoutes(app)
	routes.RegisterAuthRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
