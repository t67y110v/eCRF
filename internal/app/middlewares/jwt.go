package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	model "github.com/t67y110v/web/internal/app/model/user"
	"github.com/t67y110v/web/internal/app/utils"
)

func CheckJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(c.Cookies("JWT"), claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("11we$*9sd*(@!)"), nil
		})
		if claims["id"] == nil || err != nil {
			return utils.ErrorPage(c, errors.New("nil token"))
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
