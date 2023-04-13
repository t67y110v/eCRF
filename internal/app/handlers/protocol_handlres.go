package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) SaveProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")
		name := c.FormValue("name")
		status := c.FormValue("status")
		center_id := c.FormValue("center")
		protocol_id := c.FormValue("id")
		id, _, _, _, err := checkToken(cookie)
		if err != nil {
			return loginError(c)
		}

		_, err = h.pgStore.Repository().FindByID(id)
		if err != nil {
			return err
		}

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
		cookie := c.Cookies("JWT")
		name := c.FormValue("name")
		status := c.FormValue("status")
		center_id := c.FormValue("center")
		id, _, _, _, err := checkToken(cookie)
		if err != nil {
			return loginError(c)
		}
		_, err = h.pgStore.Repository().FindByID(id)
		if err != nil {
			return err
		}

		if err != nil {
			return c.Redirect("/main/filter=0")
		}
		s, err := strconv.Atoi(status)
		if err != nil {
			return c.Redirect("/main/filter=0")
		}
		c_id, err := strconv.Atoi(center_id)
		if err != nil {
			return c.Redirect("/main/filter=0")
		}
		err = h.pgStore.Repository().AddProtocol(name, s, c_id)
		if err != nil {
			return c.Redirect("/main/filter=0")
		}

		return c.Redirect("/main/filter=0")
	}
}
