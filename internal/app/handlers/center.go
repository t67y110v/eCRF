package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Handlers) AddNewCenter() fiber.Handler {

	return func(c *fiber.Ctx) error {
		_, _, _, _, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}

		if err := h.pgStore.Repository().AddNewCenter(c.FormValue("name")); err != nil {
			return utils.ErrorPage(c, err)
		}

		return c.Redirect("/admin/panel")

	}

}

func (h *Handlers) UpdateCenter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, _, _, _, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}

		centerID, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		if err := h.pgStore.Repository().UpdateCenter(centerID, c.FormValue("name")); err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Redirect("/admin/panel")
	}
}
