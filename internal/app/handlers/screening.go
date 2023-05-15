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
		sexSubject, err := strconv.Atoi(c.FormValue(""))
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
