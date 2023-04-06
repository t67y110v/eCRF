package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
)

// @Summary User Registration
// @Description registration of user
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Registration  true "registration"
// @Success 200 {object} responses.Registration
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/register [post]
func (h *Handlers) GetProtocols() fiber.Handler {

	return func(c *fiber.Ctx) error {

		req := &requests.Registration{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		p, err := h.pgStore.UserRepository().GetProtocols()
		if err != nil {
			return err
		}

		return c.JSON(p)

	}

}
