package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handlers) AuthPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		c.ClearCookie("JWT")
		return c.Render("auth_page", fiber.Map{
			"Test": "Test",
		})
	}

}

func (h *Handlers) MainPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")
		if cookie == "" {
			return c.Redirect("/auth")
		}
		tokenString := cookie

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.JSON(fiber.Map{
				"message": "token id is nil",
			})
		}

		id := float64(claims["id"].(float64))

		u, err := h.pgStore.UserRepository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			return err
		}
		var role string = "Менеджер"
		if u.Isadmin {
			role = "Администратор"
		}
		return c.Render("main_page", fiber.Map{
			"Name": u.Name,
			"Role": role,
		})
	}
}
