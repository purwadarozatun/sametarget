package auth

import (
	"fmt"

	"codeku.id/sametarget/helpers"
	"codeku.id/sametarget/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Profile(*gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		return c.JSON(user)
	}

}

func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		hashed, hashErr := helpers.HashPassword(user.Password)
		if hashErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to hash password",
			})
		}

		user.Password = hashed

		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to register",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Register success",
		})
	}
}

func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		var existingUser model.User
		if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid credentials",
			})
		}

		if !helpers.CheckPasswordHash(user.Password, existingUser.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid credentials",
			})
		}

		claims := map[string]interface{}{
			"username": existingUser.Username,
			"name":     existingUser.Name,
		}

		tokens := helpers.CreateAuthToken(claims)

		return c.JSON(fiber.Map{
			"data":    tokens,
			"message": "Login success",
		})
	}
}

func RefreshToken(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type RefreshTokenRequest struct {
			RefreshToken string `json:"refresh_token"`
		}

		req := new(RefreshTokenRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		fmt.Println(req.RefreshToken)

		claims, err := helpers.Validate(req.RefreshToken)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		tokens := helpers.CreateAuthToken(claims)
		return c.JSON(fiber.Map{
			"data":    tokens,
			"message": "Token refreshed",
		})

	}
}
