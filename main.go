package main

import (
	"log"

	"codeku.id/sametarget/handler"
	"codeku.id/sametarget/helpers"
	middlewares "codeku.id/sametarget/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	engine := mustache.New("./views", ".mustache")

	db := helpers.ConnectDb()
	helpers.MigrateDb(db)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	jwt := middlewares.NewAuthMiddleware("secret")

	app.Get("/profile", jwt, handler.Profile(db))

	app.Get("/", jwt, func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
