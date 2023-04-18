package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Pages) AdminPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, userName, userCentrerID, userRole, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}
		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		centers, err := h.pgStore.Repository().GetAllCenters()
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		users, err := h.pgStore.Repository().GetUsers()
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Render("admin/admin_page", fiber.Map{
			"Name":         userName,
			"Role":         utils.GetUserRole(userRole),
			"CLinicCenter": cName,
			"Centers":      centers,
			"Users":        users,
		})
	}
}

func (h *Pages) UpdatePage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, userName, _, userRole, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}

		user, err := h.pgStore.Repository().FindByEmail(c.Params("email"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		centers, err := h.pgStore.Repository().GetAllCenters()
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		return c.Render("admin/edit_user_page", fiber.Map{
			"Name":    userName,
			"Role":    utils.GetUserRole(userRole),
			"User":    user,
			"Centers": centers,
		})
	}
}
