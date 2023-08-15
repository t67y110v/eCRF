package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
	"github.com/t67y110v/web/internal/app/utils"
)

// @Summary UpdateColor
// @Description UpdateColor
// @Tags         Subject.Action
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.UpdateColorRequest true  "UpdateColorRequest"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/action/updatecolor [patch]
func (h *Handlers) UpdateColor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.UpdateColorRequest{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		subject, err := h.mgStore.Subject().GetSubjectByNumber(req.SubjectNumber, req.ProtocolID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		fieldName := utils.GetFieldName(req.FieldName, req.Count)
		if err := h.mgStore.Color().UpdateColor(c.Context(), subject.ID, fieldName, req.Value); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Протокол:%d Субьект:%s, Обновление статуса - %d в поле %s", req.ProtocolID, req.SubjectNumber, req.Value, req.FieldName), c.Cookies("token_name"), c.Cookies("token_role"), "patch", req)
		return c.JSON(fiber.Map{
			"message": "success",
		})

	}
}

// @Summary UpdateColorWithComment
// @Description UpdateColorWithComment
// @Tags         Subject.Action
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.UpdateColorWithCommentRequest true  "UpdateColorWithCommentRequest"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/action/updatecolorwithcomment [patch]
func (h *Handlers) UpdateColorWithComment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.UpdateColorWithCommentRequest{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		subject, err := h.mgStore.Subject().GetSubjectByNumber(req.SubjectNumber, req.ProtocolID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if err := h.mgStore.Color().UpdateColorWithComment(c.Context(), subject.ID, utils.GetFieldName(req.FieldName, req.Count), req.Reason, req.Comment, req.Sender, req.SendersRole, req.Value); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Протокол:%d Субьект:%s, Обновление статуса - %d в поле %s с комментарием: %s", req.ProtocolID, req.SubjectNumber, req.Value, req.FieldName, req.Comment), c.Cookies("token_name"), c.Cookies("token_role"), "patch", req)
		return c.JSON(fiber.Map{
			"message": "success",
		})

	}
}

// @Summary UpdateFieldValue
// @Description UpdateFieldValue
// @Tags         Subject.Action
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.UpdateValueWithColor true  "UpdateFieldValue"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/action/updatefield [patch]
func (h *Handlers) UpdateFieldValue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.UpdateValueWithColor{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		subject, err := h.mgStore.Subject().GetSubjectByNumber(req.SubjectNumber, req.ProtocolID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		field := utils.GetFieldName(req.FieldName, req.Count)
		fieldValue := utils.GetFieldNameForUpdate(field)
		value, err := strconv.Atoi(req.Value)
		if err != nil {
			if err := h.mgStore.Color().UpdateFieldStringValue(c.Context(), subject.ID, field, fieldValue, req.Value, req.Color); err != nil {
				c.Status(http.StatusBadRequest)
				return c.JSON(fiber.Map{
					"message": err.Error(),
				})

			}
			go h.journal.SaveAction(c.Context(), fmt.Sprintf("Протокол:%d Субьект:%s, Обновление значения - %s в поле %s", req.ProtocolID, req.SubjectNumber, req.Value, req.FieldName), c.Cookies("token_name"), c.Cookies("token_role"), "patch", req)
		} else {
			if err := h.mgStore.Color().UpdateFieldIntValue(c.Context(), subject.ID, field, fieldValue, value, req.Color); err != nil {
				c.Status(http.StatusBadRequest)
				return c.JSON(fiber.Map{
					"message": err.Error(),
				})
			}
			go h.journal.SaveAction(c.Context(), fmt.Sprintf("Протокол:%d Субьект:%s, Обновление значения - %s в поле %s", req.ProtocolID, req.SubjectNumber, req.Value, req.FieldName), c.Cookies("token_name"), c.Cookies("token_role"), "patch", req)
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}
