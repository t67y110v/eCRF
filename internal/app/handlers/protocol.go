package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Handlers) SaveProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		status := c.FormValue("status")
		center_id := c.FormValue("center")
		protocol_id := c.FormValue("id")

		p_id, err := strconv.Atoi(protocol_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		s, err := strconv.Atoi(status)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		c_id, err := strconv.Atoi(center_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		err = h.pgStore.Repository().UpdateProtocolById(p_id, s, name, c_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}

		return c.Redirect("/main/filter=0")
	}
}

func (h *Handlers) AddProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		status := c.FormValue("status")
		center_id := c.FormValue("center")

		s, err := strconv.Atoi(status)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		c_id, err := strconv.Atoi(center_id)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		err = h.pgStore.Repository().AddProtocol(name, s, c_id)
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		return c.Redirect("/main/filter=0")
	}
}
