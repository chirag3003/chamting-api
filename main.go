package main

import (
	"github.com/XxThunderBlastxX/chamting-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	//Initiated the fiber instance
	app := fiber.New()

	app.Use(cors.New())

	//Setup Router
	router.Route(app)

	//Application listening to the port
	log.Fatal(app.Listen(":8080"))
}
