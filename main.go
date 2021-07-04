package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rilladev/rilla/routes"
)
func main(){
	app := fiber.New()
	
	routes.Setup(app)

	app.Listen(":4000")
}