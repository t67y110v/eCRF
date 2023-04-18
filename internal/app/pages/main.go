package pages

import (
	"github.com/gofiber/fiber/v2"
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

		_, userName, userCentrerID, userRole, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}

		p, err := h.pgStore.Repository().GetProtocolsByFilter(filter, userCentrerID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Render("main_page", fiber.Map{
			"Name":         userName,
			"Role":         utils.GetUserRole(userRole),
			"CLinicCenter": cName,
			"Protocols":    p,
		})
	}
}
