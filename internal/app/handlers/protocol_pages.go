package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) ProtocolPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		id, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}
		u, err := h.pgStore.Repository().FindByID(id)
		if err != nil {
			return err
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
		return c.Render("protocol/protocol_index", fiber.Map{
			"Name":         u.Name,
			"Role":         u.Role,
			"CLinicCenter": u.CenterID,
			"Protocol":     p,
		})

	}

}

func (h *Handlers) ProtocolEdit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}
		u, err := h.pgStore.Repository().FindByID(id)
		if err != nil {
			return err
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

		return c.Render("protocol/protocol_edit", fiber.Map{
			"Name":         u.Name,
			"Role":         u.Role,
			"CLinicCenter": u.CenterID,
			"Protocol":     p,
		})
	}
}

func (h *Handlers) ProtocolNew() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := checkToken(c.Cookies("JWT"))
		if err != nil {
			return loginError(c)
		}
		u, err := h.pgStore.Repository().FindByID(id)
		if err != nil {
			return err // 404
		}

		if err != nil {
			return c.Redirect("/auth")
		}

		return c.Render("protocol/protocol_new", fiber.Map{
			"Name":         u.Name,
			"Role":         u.Role,
			"CLinicCenter": u.CenterID,
		})
	}
}
