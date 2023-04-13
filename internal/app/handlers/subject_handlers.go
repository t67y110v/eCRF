package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) NewSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")
		_, _, _, _, err := checkToken(cookie)
		if err != nil {
			return loginError(c)
		}

		centerId, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return err
		}
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return err
		}
		if err := h.mgStore.Repository().AddSubject(c.FormValue("number"), c.FormValue("initials"), centerId, protocolId); err != nil {
			return err
		}

		if err != nil {
			return err
		}

		return c.Redirect(fmt.Sprintf("/protocol/%s", c.FormValue("protocol_id")))
	}
}
