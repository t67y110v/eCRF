package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
	ID         primitive.ObjectID `bson:"_id"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	Number     string             `bson:"number"`
	CenterID   int                `bson:"center_id" `
	ProtocolId int                `bson:"protocol_id"`
	Initials   string             `bson:"initials"`
	Screening  struct {
		InformaionConsent struct {
			Signed                      bool      `bson:"signed"`
			DateOfSign                  string    `bson:"date_of_sign"`
			TimeOfSign                  time.Time `bson:"time_of_sign"`
			ReceivedAnInsurancePolicy   bool      `bson:"received_an_insurance_policy"`
			ReceivedAnInformaionConsent bool      `bson:"received_an_informaion_consent"`
		}

		Anthropometry struct {
			AnthropometricDataBeenMeasured bool   `bson:"anthropometric_data_been_measured"`
			ReasonIfNot                    string `bson:"reason_if_not"`
			DateOfStartMeasured            string `bson:"date_of_start_measured"`
			WeightOfBody                   int    `bson:"weight_of_body"`
			HightOfBody                    int    `bson:"hight_of_body"`
			IndexWeigthOfBody              int    `bson:"index_weight_of_body"`
		}
		Demography struct {
			Sex       int    `bson:"sex"`
			Race      int    `bson:"race"`
			BirthDate string `bson:"date"`
		}
		InclusionCriteria struct {
			PresenceOfAnInformationPanel                bool `bson:"presence_of_an_information_panel"`
			Aged18To55Years                             bool `bson:"aged_18_to_55_years"`
			NegativeHIVTestResult                       bool `bson:"negative_hiv_test_result"`
			BodyMassIndex                               bool `bson:"body_mass_index"`
			AbsenceOfAcuteInfectiousDiseases            bool `bson:"absence_of_acute_infectious_diseases"`
			ConsentToUseEffectiveMethodsOfContraception bool `bson:"consent_to_use_effective_methods_of_contraception"`
			NegativePregnancyTest                       bool `bson:"negative_pregnancy_test"`
			NegativeAlcoholTest                         bool `bson:"negative_alcohol_test"`
			NoHistoryOfSeverePostVaccinationReactions   bool `bson:"no_history_of_severe_post_vaccination_reactions"`
			IndicatorsBloodTestsAtScreeningWithin       bool `bson:"indicators_blood_tests_at_screening_within"`
			NoMyocardialChanges                         bool `bson:"no_myocardial_changes"`
			NegativeTestResultForCOVID                  bool `bson:"negative_test_result_for_COVID"`
			NoContraindicationsToVaccination            bool `bson:"no_contraindications_to_vaccination"`
		}

		ExclusionСriteria struct {
			LackOfSignedInformedConsent           bool `bson:"lack_of_signed_informed_consent"`
			SteroidTherapy                        bool `bson:"steroid_therapy"`
			TherapyWithImmunosuppressiveDrugs     bool `bson:"therapy_with_immunosuppressive_drugs"`
			FemaleSubjectsDuringPregnancy         bool `bson:"female_subjects_during_pregnancy"`
			StrokeInLessThanOneYear               bool `bson:"stroke_in_less_than_one_year"`
			ChronicSystemicInfections             bool `bson:"chronic_systemic_infections"`
			AggravatedAllergicHistory             bool `bson:"aggravated_allergic_history"`
			PresenceOfAHistoryOfNeoplasms         bool `bson:"presence_of_a_history_of_neoplasms"`
			HistoryOfSplenectomy                  bool `bson:"history_of_splenectomy"`
			Neutropenia                           bool `bson:"neutropenia"`
			SubjectsWithActiveSyphilis            bool `bson:"subjects_with_active_syphilis"`
			Anorexia                              bool `bson:"anorexia"`
			ExtensiveTattoos                      bool `bson:"extensive_tattoos"`
			TakingNarcoticAndPsychostimulantDrugs bool `bson:"taking_narcotic_and_psychostimulant_drugs"`
			SmokingMoretThanTenCigarettesADay     bool `bson:"smoking_more_than_ten_cigarettes_a_day"`
			AlcoholIntake                         bool `bson:"alcohol_intake"`
			PlannedHospitalization                bool `bson:"planned_hospitalization"`
			DonorBloodDonation                    bool `bson:"donor_blood_donation"`
			SubjectParticipationInAnyOtherStudy   bool `bson:"subject_participation_in_any_other_study"`
			AnyVaccinationInTheLastMonth          bool `bson:"any_vaccination_in_the_last_month"`
			InabilityToReadInRussian              bool `bson:"inability_to_read_in_russian"`
			ResearchCenterStaff                   bool `bson:"research_center_staff"`
			AnyOtherStateOfTheSubjectOfTheStudy   bool `bson:"any_other_state_of_the_subject_of_the_study"`
		}
	}
}
