package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) AddNewCenter() fiber.Handler {

	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")
		_, err := checkToken(cookie)
		if err != nil {
			return loginError(c)
		}

		if err := h.pgStore.Repository().AddNewCenter(c.FormValue("name")); err != nil {
			return err
		}

		return c.Redirect("/admin/panel")

	}

}

func (h *Handlers) UpdateCenter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")
		_, err := checkToken(cookie)
		if err != nil {
			return loginError(c)
		}

		centerID, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return err
		}
		if err := h.pgStore.Repository().UpdateCenter(centerID, c.FormValue("name")); err != nil {
			return err
		}
		return c.Redirect("/admin/panel")
	}
}
