package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
)

// @Summary StartOfScreeningSubject
// @Description StartOfScreeningSubject value of subject
// @Tags         Subject.Screening
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.StartOfScreening true  "StartOfScreening"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/screening/startofscreening [patch]
func (h *Handlers) StartOfScreeningSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.StartOfScreening{}
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
		if err := h.mgStore.Screening().StartScreening(c.Context(), subject.ID, req.DateStartOfScreening, req.TimeStartOfScreening); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary InformaitonConsentSubject
// @Description InformaitonConsentSubject value of subject
// @Tags         Subject.Screening
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.InformationConsent true  "InformaionConsent"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/screening/informationconsent [patch]
func (h *Handlers) InformationConsentSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.InformationConsent{}
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

		if err := h.mgStore.Screening().InformaionConsent(
			c.Context(),
			subject.ID,
			req.DateOfSign,
			req.TimeOfSign,
			req.IsSigned,
			req.ReceivedAnInsurancePolicy,
			req.ReceivedAnInformaionConsent,
		); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary DemographySubject
// @Description DemographySubject value of subject
// @Tags         Subject.Screening
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Demography true  "DemographySubject"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/screening/demography [patch]
func (h *Handlers) DemographySubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.Demography{}
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
		if err := h.mgStore.Screening().Demography(c.Context(), subject.ID, req.Sex, req.Race, req.Date); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary AnthropometrySubject
// @Description AnthropometrySubject value of subject
// @Tags         Subject.Screening
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Anthropometry true  "AnthropometrySubject"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/screening/anthropometry [patch]
func (h *Handlers) AnthropometrySubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.Anthropometry{}
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

		if err := h.mgStore.Screening().Anthropometry(c.Context(),
			subject.ID, req.AnthropometricDataBeenMeasured,
			req.ReasonIfNot,
			req.DateOfStartMeasured,
			req.WeightOfBody,
			req.HightOfBody,
			req.IndexWeigthOfBody,
		); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary InclusionCriteriaSubject
// @Description InclusionCriteriaSubject value of subject
// @Tags         Subject.Screening
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.InclusionCriteria true  "InclusionCriteriaSubject"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/screening/inclusioncriteria [patch]
func (h *Handlers) InclusionCriteriaSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.InclusionCriteria{}
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

		if err := h.mgStore.Screening().InclusionCriteria(
			c.Context(),
			subject.ID,
			req.PresenceOfAnInformationPanel, req.Aged18To55Years,
			req.NegativeHIVTestResult, req.BodyMassIndex,
			req.AbsenceOfAcuteInfectiousDiseases, req.ConsentToUseEffectiveMethodsOfContraception,
			req.NegativePregnancyTest, req.NegativeDrugTest,
			req.NegativeAlcoholTest, req.NoHistoryOfSeverePostVaccinationReactions,
			req.IndicatorsBloodTestsAtScreeningWithin, req.NoMyocardialChanges,
			req.NegativeTestResultForCOVID, req.NoContraindicationsToVaccination,
		); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary ExclusionСriteriaSubject
// @Description ExclusionСriteriaSubject value of subject
// @Tags         Subject.Screening
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.ExclusionСriteria true  "ExclusionСriteriaSubject"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/screening/exclusioncriteria [patch]
func (h *Handlers) ExclusionСriteriaSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.ExclusionСriteria{}
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
		if err := h.mgStore.Screening().ExclusionСriteria(
			c.Context(),
			subject.ID,
			req.LackOfSignedInformedConsent, req.SteroidTherapy,
			req.TherapyWithImmunosuppressiveDrugs, req.FemaleSubjectsDuringPregnancy,
			req.StrokeInLessThanOneYear, req.ChronicSystemicInfections,
			req.AggravatedAllergicHistory, req.PresenceOfAHistoryOfNeoplasms,
			req.HistoryOfSplenectomy, req.Neutropenia,
			req.SubjectsWithActiveSyphilis, req.Anorexia,
			req.ExtensiveTattoos, req.TakingNarcoticAndPsychostimulantDrugs,
			req.SmokingMoretThanTenCigarettesADay, req.AlcoholIntake,
			req.PlannedHospitalization, req.DonorBloodDonation,
			req.SubjectParticipationInAnyOtherStudy, req.AnyVaccinationInTheLastMonth,
			req.InabilityToReadInRussian, req.ResearchCenterStaff,
			req.AnyOtherStateOfTheSubjectOfTheStudy,
		); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})

	}
}

// @Summary CompletionOfScreening
// @Description CompletionOfScreening value of subject
// @Tags         Subject.Screening
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.CompletionOfScreening true  "CompletionOfScreening"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/screening/completion [patch]
func (h *Handlers) CompletionOfScreening() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.CompletionOfScreening{}
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

		if err := h.mgStore.Screening().CompletionOfScreening(c.Context(),
			subject.ID, req.VolunteerEligible,
			req.NoExclusionCriteria,
			req.InformedOfTheRestrictions,
			req.VolunteerIncluded,
			req.ReasonIfNot,
			req.CommentValue,
		); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}
