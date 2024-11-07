package auth

import (
	middlewares "codeku.id/sametarget/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoute(app *fiber.App, db *gorm.DB) {

	jwt := middlewares.NewAuthMiddleware("secret")
	// app.Post("/register", Register(db))
	app.Post("/login", Login(db))
	app.Post("/register", Register(db))
	app.Post("/refresh-token", RefreshToken(db))

	app.Get("/profile", jwt, Profile(db))
}
