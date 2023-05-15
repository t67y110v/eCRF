package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
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

		return c.Redirect(fmt.Sprintf("/protocol/%d/%s", protocolId, subjectNumber))
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
		return c.Redirect(fmt.Sprintf("/protocol/%d/%s", protocolId, subjectNumber))
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

		return c.Redirect(fmt.Sprintf("/protocol/%d/%s", protocolId, subjectNumber))
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

		if err := h.mgStore.Screening().InclusionCriteria(c.Context(), subject.ID, inclusion1, inclusion2, inclusion3, inclusion4, inclusion5, inclusion6, inclusion7, inclusion8, inclusion9, inclusion10, inclusion11, inclusion12, inclusion13, inclusion14); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		return c.Redirect(fmt.Sprintf("/protocol/%d/%s", protocolId, subjectNumber))
	}
}

func (h *Handlers) ExclusionСriteriaSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, 1)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		if err := h.mgStore.Screening().ExclusionСriteria(c.Context(), subject.ID, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		return c.Redirect("/protocol/1/1")

	}
}
