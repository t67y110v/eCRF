package pages

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/t67y110v/web/internal/app/model/user"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Pages) AuthPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		c.ClearCookie("JWT")
		er := c.Cookies("error")
		c.ClearCookie("error")
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

func (h *Pages) MainPage() fiber.Handler {

	return func(c *fiber.Ctx) error {

		filter := c.Params("filter")

		user := c.Locals("user").(*model.User)

		p, err := h.pgStore.Repository().GetProtocolsByFilter(filter, user.CenterID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		cName, err := h.pgStore.Repository().GetCenterName(user.CenterID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		return c.Render("main_page", fiber.Map{
			"Name":         user.Name,
			"Role":         utils.GetUserRole(user.Role),
			"CLinicCenter": cName,
			"Protocols":    p,
		})
	}
}
