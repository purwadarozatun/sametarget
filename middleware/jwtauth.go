package middlewares

import (
	"fmt"

	"codeku.id/sametarget/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	// return jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(secret)},
	// })

	return func(c *fiber.Ctx) error {
		// Get the JWT token from the request header "Authorization
		token := c.Get("Authorization")

		// remove the "Bearer " prefix from the token
		token = token[7:]
		tokenClaims, err := decode(token)
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

func decode(tokenString string) (model.User, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return model.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		user := model.User{
			Username: claims["preferred_username"].(string),
			Email:    claims["email"].(string),
			Name:     claims["name"].(string),
		}

		return user, nil
	} else {
		return model.User{}, fmt.Errorf("invalid token")
	}
}
