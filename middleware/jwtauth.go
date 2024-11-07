package middlewares

import (
	"fmt"

	"codeku.id/sametarget/helpers"
	"github.com/gofiber/fiber/v2"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	// return jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(secret)},
	// })

	return func(c *fiber.Ctx) error {
		// Get the JWT token from the request header "Authorization
		token := c.Get("Authorization")

		fmt.Println(token)

		// remove the "Bearer " prefix from the token
		token = token[7:]
		tokenClaims, err := helpers.Decode(token)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		c.Locals("user", tokenClaims)
		return c.Next()
	}
}
