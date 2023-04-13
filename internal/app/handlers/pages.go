package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) AuthPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		c.ClearCookie("JWT")
		er := c.Cookies("error")
		if er == "password" {
			return c.Render("auth_page", fiber.Map{
				"WrongPassword": "True",
			})
		} else if er == "login" {
			return c.Render("auth_page", fiber.Map{
				"WrongLogin": "True",
			})
		}

		return c.Render("auth_page", fiber.Map{})

	}

}

func (h *Handlers) MainPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		filter := c.Params("filter")
		c.ClearCookie("error")

		_, userName, userCentrerID, userRole, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}

		p, err := h.pgStore.Repository().GetProtocolsByFilter(filter, userCentrerID)
		if err != nil {
			return err //404
		}
		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return err
		}
		return c.Render("main_page", fiber.Map{
			"Name":         userName,
			"Role":         getUserRole(userRole),
			"CLinicCenter": cName,
			"Protocols":    p,
		})
	}
}
