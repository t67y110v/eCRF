package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CheckJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(c.Cookies("token"), claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("11we$*9sd*(@!)"), nil
		})
		if claims["id"] == nil || err != nil {
			c.Status(http.StatusUnauthorized)
			return nil
		}
		// c.Locals("user", u)
		// c.Locals("name", name)

		return c.Next()

	}
}
