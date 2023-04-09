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
		p, err := h.pgStore.Repository().GetProtocolsByFilter(filter)
		if err != nil {
			return err //404
		}

		id, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}
		u, err := h.pgStore.Repository().FindByID(id)
		if err != nil {
			return loginError(c)
		}

		return c.Render("main_page", fiber.Map{
			"Name":         u.Name,
			"Role":         u.Role,
			"CLinicCenter": u.CenterID,
			"Protocols":    p,
		})
	}
}
