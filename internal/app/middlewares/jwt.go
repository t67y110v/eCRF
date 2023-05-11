package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	model "github.com/t67y110v/web/internal/app/model/user"
)

func CheckJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(c.Cookies("JWT"), claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("11we$*9sd*(@!)"), nil
		})
		if claims["id"] == nil || err != nil {
			return c.Redirect("/auth")
		}

		id := int(claims["id"].(float64))
		role := (int(claims["role"].(float64)))
		name := claims["name"].(string)

		centerId := (int(claims["centerId"].(float64)))
		u := &model.User{
			Name:     name,
			Role:     role,
			Id:       id,
			CenterID: centerId,
		}

		c.Locals("user", u)
		c.Locals("name", name)

		return c.Next()

	}
}
