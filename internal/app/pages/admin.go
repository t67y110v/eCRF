package pages

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/t67y110v/web/internal/app/model/user"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Pages) AdminPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)
		cName, err := h.pgStore.Repository().GetCenterName(user.CenterID)
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
			"Name":         user.Name,
			"Role":         utils.GetUserRole(user.Role),
			"CLinicCenter": cName,
			"Centers":      centers,
			"Users":        users,
		})
	}
}

func (h *Pages) UpdatePage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		centers, err := h.pgStore.Repository().GetAllCenters()
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		return c.Render("admin/edit_user_page", fiber.Map{
			"Name":    user.Name,
			"Role":    utils.GetUserRole(user.Role),
			"User":    user,
			"Centers": centers,
		})
	}
}
