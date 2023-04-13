package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) AdminPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//	//id, name, centerId, role, nil
		_, userName, userCentrerID, userRole, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}

		if err != nil {
			return c.Redirect("/auth")
		}
		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return err
		}
		centers, err := h.pgStore.Repository().GetAllCenters()
		if err != nil {
			return err
		}
		users, err := h.pgStore.Repository().GetUsers()
		if err != nil {
			return err
		}
		return c.Render("admin/admin_page", fiber.Map{
			"Name":         userName,
			"Role":         getUserRole(userRole),
			"CLinicCenter": cName,
			"Centers":      centers,
			"Users":        users,
		})
	}
}

func (h *Handlers) UpdatePage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, userName, _, userRole, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}

		user, err := h.pgStore.Repository().FindByEmail(c.Params("email"))
		if err != nil {
			return err
		}

		centers, err := h.pgStore.Repository().GetAllCenters()
		if err != nil {
			return err
		}

		return c.Render("admin/edit_user_page", fiber.Map{
			"Name":    userName,
			"Role":    getUserRole(userRole),
			"User":    user,
			"Centers": centers,
		})
	}
}
