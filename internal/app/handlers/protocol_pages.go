package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) ProtocolPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		_, userName, userCentrerID, userRole, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}

		protocol_id := c.Params("id")
		p_id, err := strconv.Atoi(protocol_id)

		if err != nil {
			return c.Redirect("/auth")
		}

		p, err := h.pgStore.Repository().GetProtocolById(p_id)
		if err != nil {
			return c.Redirect("/auth")
		}
		s, err := h.mgStore.Repository().GetSubjectsByProtocolId(p_id)
		if err != nil {
			return err // 502
		}

		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return err
		}
		return c.Render("protocol/protocol_index", fiber.Map{
			"Name":         userName,
			"Role":         getUserRole(userRole),
			"CLinicCenter": cName,
			"ClinicId":     userCentrerID,
			"Protocol":     p,
			"Subjects":     s, //s,
		})

	}

}

func (h *Handlers) ProtocolEdit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, userName, userCentrerID, userRole, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}

		if err != nil {
			return c.Redirect("/auth")
		}

		protocol_id := c.Params("id")
		p_id, err := strconv.Atoi(protocol_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		p, err := h.pgStore.Repository().GetProtocolById(p_id)
		if err != nil {
			return c.Redirect("/auth") // 404
		}
		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return err
		}
		return c.Render("protocol/protocol_edit", fiber.Map{
			"Name":         userName,
			"Role":         getUserRole(userRole),
			"CLinicCenter": cName,
			"Protocol":     p,
		})
	}
}

func (h *Handlers) ProtocolNew() fiber.Handler {
	return func(c *fiber.Ctx) error {
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
		return c.Render("protocol/protocol_new", fiber.Map{
			"Name":         userName,
			"Role":         getUserRole(userRole),
			"CLinicCenter": cName,
		})
	}
}
