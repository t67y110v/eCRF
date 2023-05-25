package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Handlers) InformaionConsentSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		signed, err := strconv.Atoi(c.FormValue("signed"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		dateOfSign := c.FormValue("date_of_sign")
		timeOfSign := c.FormValue("time_of_sign")

		original, err := strconv.Atoi(c.FormValue("original"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		consent, err := strconv.Atoi(c.FormValue("consent"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, protocolId)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		if err := h.mgStore.Screening().InformaionConsent(c.Context(), subject.ID, dateOfSign, timeOfSign, signed, original, consent); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		return c.Redirect(fmt.Sprintf("/protocol/%d/%s#item-1-1", protocolId, subjectNumber))
	}
}

func (h *Handlers) DemographySubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, protocolId)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		sexSubject, err := strconv.Atoi(c.FormValue("sex"))
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		raceSubject, err := strconv.Atoi(c.FormValue("race"))
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		birthDateSubject := c.FormValue("birth_date")

		if err := h.mgStore.Screening().Demography(c.Context(), subject.ID, sexSubject, raceSubject, birthDateSubject); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		return c.Redirect(fmt.Sprintf("/protocol/%d/%s#item-1-2", protocolId, subjectNumber))
	}
}

func (h *Handlers) AnthropometrySubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, protocolId)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		measured, err := strconv.Atoi(c.FormValue("data_been_measured"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		ifNot := c.FormValue("if_not")
		dateOfStart := c.FormValue("date_of_start")
		weight, err := strconv.Atoi(c.FormValue("weight"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		height, err := strconv.Atoi(c.FormValue("height"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		index, err := strconv.Atoi(c.FormValue("index"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		if err := h.mgStore.Screening().Anthropometry(c.Context(), subject.ID, measured, ifNot, dateOfStart, weight, height, index); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		return c.Redirect(fmt.Sprintf("/protocol/%d/%s#item-1-3", protocolId, subjectNumber))
	}
}

func (h *Handlers) InclusionCriteriaSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, protocolId)
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		inclusion1, _ := strconv.Atoi(c.FormValue("inclusion1"))
		inclusion2, _ := strconv.Atoi(c.FormValue("inclusion2"))
		inclusion3, _ := strconv.Atoi(c.FormValue("inclusion3"))
		inclusion4, _ := strconv.Atoi(c.FormValue("inclusion4"))
		inclusion5, _ := strconv.Atoi(c.FormValue("inclusion5"))
		inclusion6, _ := strconv.Atoi(c.FormValue("inclusion6"))
		inclusion7, _ := strconv.Atoi(c.FormValue("inclusion7"))
		inclusion8, _ := strconv.Atoi(c.FormValue("inclusion8"))
		inclusion9, _ := strconv.Atoi(c.FormValue("inclusion9"))
		inclusion10, _ := strconv.Atoi(c.FormValue("inclusion10"))
		inclusion11, _ := strconv.Atoi(c.FormValue("inclusion11"))
		inclusion12, _ := strconv.Atoi(c.FormValue("inclusion12"))
		inclusion13, _ := strconv.Atoi(c.FormValue("inclusion13"))
		inclusion14, _ := strconv.Atoi(c.FormValue("inclusion14"))

		if err := h.mgStore.Screening().InclusionCriteria(
			c.Context(),
			subject.ID,
			inclusion1, inclusion2,
			inclusion3, inclusion4,
			inclusion5, inclusion6,
			inclusion7, inclusion8,
			inclusion9, inclusion10,
			inclusion11, inclusion12,
			inclusion13, inclusion14,
		); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		return c.Redirect(fmt.Sprintf("/protocol/%d/%s#item-1-4", protocolId, subjectNumber))
	}
}

func (h *Handlers) ExclusionСriteriaSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, protocolId)
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		exclusion1, _ := strconv.Atoi(c.FormValue("exclusion1"))
		exclusion2, _ := strconv.Atoi(c.FormValue("exclusion2"))
		exclusion3, _ := strconv.Atoi(c.FormValue("exclusion3"))
		exclusion4, _ := strconv.Atoi(c.FormValue("exclusion4"))
		exclusion5, _ := strconv.Atoi(c.FormValue("exclusion5"))
		exclusion6, _ := strconv.Atoi(c.FormValue("exclusion6"))
		exclusion7, _ := strconv.Atoi(c.FormValue("exclusion7"))
		exclusion8, _ := strconv.Atoi(c.FormValue("exclusion8"))
		exclusion9, _ := strconv.Atoi(c.FormValue("exclusion9"))
		exclusion10, _ := strconv.Atoi(c.FormValue("exclusion10"))
		exclusion11, _ := strconv.Atoi(c.FormValue("exclusion11"))
		exclusion12, _ := strconv.Atoi(c.FormValue("exclusion12"))
		exclusion13, _ := strconv.Atoi(c.FormValue("exclusion13"))
		exclusion14, _ := strconv.Atoi(c.FormValue("exclusion14"))
		exclusion15, _ := strconv.Atoi(c.FormValue("exclusion15"))
		exclusion16, _ := strconv.Atoi(c.FormValue("exclusion16"))
		exclusion17, _ := strconv.Atoi(c.FormValue("exclusion17"))
		exclusion18, _ := strconv.Atoi(c.FormValue("exclusion18"))
		exclusion19, _ := strconv.Atoi(c.FormValue("exclusion19"))
		exclusion20, _ := strconv.Atoi(c.FormValue("exclusion20"))
		exclusion21, _ := strconv.Atoi(c.FormValue("exclusion21"))
		exclusion22, _ := strconv.Atoi(c.FormValue("exclusion22"))
		exclusion23, _ := strconv.Atoi(c.FormValue("exclusion23"))

		if err := h.mgStore.Screening().ExclusionСriteria(
			c.Context(),
			subject.ID,
			exclusion1, exclusion2,
			exclusion3, exclusion4,
			exclusion5, exclusion6,
			exclusion7, exclusion8,
			exclusion9, exclusion10,
			exclusion11, exclusion12,
			exclusion13, exclusion14,
			exclusion15, exclusion16,
			exclusion17, exclusion18,
			exclusion19, exclusion20,
			exclusion21, exclusion22,
			exclusion23,
		); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		return c.Redirect(fmt.Sprintf("/protocol/%d/%s#item-1-5", protocolId, subjectNumber))

	}
}

func (h *Handlers) CompletionOfScreening() fiber.Handler {
	return func(c *fiber.Ctx) error {

		subjectNumber := c.FormValue("subject_number")
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, protocolId)
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		completion1, _ := strconv.Atoi(c.FormValue("completion1"))
		completion2, _ := strconv.Atoi(c.FormValue("completion2"))
		completion3, _ := strconv.Atoi(c.FormValue("completion3"))
		completion4, _ := strconv.Atoi(c.FormValue("completion4"))
		completion5 := c.FormValue("completion5")
		completion6 := c.FormValue("completion6")

		if err := h.mgStore.Screening().CompletionOfScreening(c.Context(), subject.ID, completion1, completion2, completion3, completion4, completion5, completion6); err != nil {
			h.logger.Warning(err)
			return utils.ErrorPage(c, err)
		}
		return c.Redirect(fmt.Sprintf("/protocol/%d/%s#item-1-6", protocolId, subjectNumber))
	}
}

func (h *Handlers) UpdateColor() fiber.Handler {
	return func(c *fiber.Ctx) error {

		req := &requests.UpdateColorRequest{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle filter by category, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		subjectNumber := req.SubjectNumber
		protocolId, err := strconv.Atoi(req.ProtocolID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, protocolId)
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		field := req.FieldName
		value := req.Value

		fieldName := utils.GetFieldName(field)
		if err := h.mgStore.Screening().UpdateColor(c.Context(), subject.ID, fieldName, value); err != nil {
			return utils.ErrorPage(c, err)
		}
		return nil

	}
}

func (h *Handlers) UpdateColorWithComment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &requests.UpdateColorWithCommentRequest{}
		reader := bytes.NewReader(c.Body())
		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle filter by category, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		subjectNumber := req.SubjectNumber
		protocolId, err := strconv.Atoi(req.ProtocolID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, protocolId)
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		field := req.FieldName

		value := req.Value
		fieldName := utils.GetFieldName(field)

		reason := req.Reason
		comment := req.Comment

		//fmt.Println("s - ", subject.ID, "p - ", protocolId, "field - ", field, "fieldNAme - ", fieldName, "v - ", value, "r - ", reason, "c - ", comment)
		if err := h.mgStore.Screening().UpdateColorWithComment(c.Context(), subject.ID, fieldName, reason, comment, value); err != nil {
			return utils.ErrorPage(c, err)
		}

		return nil

	}
}
