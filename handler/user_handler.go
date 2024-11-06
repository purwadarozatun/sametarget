package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Profile(*gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		return c.JSON(user)
	}
}
