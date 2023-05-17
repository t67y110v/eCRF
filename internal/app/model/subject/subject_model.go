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
			SignedCondition struct {
				Signed  int    `bson:"signed"`
				Color   int    `bson:"color"`
				Reason  string `bson:"reason"`
				Comment string `bson:"comment"`
			}
			DateOfSignCondition struct {
				DateOfSign string `bson:"date_of_sign"`
				Color      int    `bson:"color"`
				Reason     string `bson:"reason"`
				Comment    string `bson:"comment"`
			}
			TimeOfSignCondition struct {
				TimeOfSign string `bson:"time_of_sign"`
				Color      int    `bson:"color"`
				Reason     string `bson:"reason"`
				Comment    string `bson:"comment"`
			}
			ReceivedAnInsurancePolicyCondition struct {
				ReceivedAnInsurancePolicy int    `bson:"received_an_insurance_policy"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
			}
			ReceivedAnInformaionConsentCondition struct {
				ReceivedAnInformaionConsent int    `bson:"received_an_informaion_consent"`
				Color                       int    `bson:"color"`
				Reason                      string `bson:"reason"`
				Comment                     string `bson:"comment"`
			}
		}

		Anthropometry struct {
			AnthropometricDataBeenMeasured int    `bson:"anthropometric_data_been_measured"`
			ReasonIfNot                    string `bson:"reason_if_not"`
			DateOfStartMeasured            string `bson:"date_of_start_measured"`
			WeightOfBody                   int    `bson:"weight_of_body"`
			HeightOfBody                   int    `bson:"height_of_body"`
			IndexWeigthOfBody              int    `bson:"index_weight_of_body"`
		}
		Demography struct {
			Sex       int    `bson:"sex"`
			Race      int    `bson:"race"`
			BirthDate string `bson:"date"`
		}
		InclusionCriteria struct {
			PresenceOfAnInformationPanel                int `bson:"presence_of_an_information_panel"`
			Aged18To55Years                             int `bson:"aged_18_to_55_years"`
			NegativeHIVTestResult                       int `bson:"negative_hiv_test_result"`
			BodyMassIndex                               int `bson:"body_mass_index"`
			AbsenceOfAcuteInfectiousDiseases            int `bson:"absence_of_acute_infectious_diseases"`
			ConsentToUseEffectiveMethodsOfContraception int `bson:"consent_to_use_effective_methods_of_contraception"`
			NegativePregnancyTest                       int `bson:"negative_pregnancy_test"`
			NegativeDrugTest                            int `bson:"negative_drug_test"`
			NegativeAlcoholTest                         int `bson:"negative_alcohol_test"`
			NoHistoryOfSeverePostVaccinationReactions   int `bson:"no_history_of_severe_post_vaccination_reactions"`
			IndicatorsBloodTestsAtScreeningWithin       int `bson:"indicators_blood_tests_at_screening_within"`
			NoMyocardialChanges                         int `bson:"no_myocardial_changes"`
			NegativeTestResultForCOVID                  int `bson:"negative_test_result_for_COVID"`
			NoContraindicationsToVaccination            int `bson:"no_contraindications_to_vaccination"`
		}

		ExclusionСriteria struct {
			LackOfSignedInformedConsent           int `bson:"lack_of_signed_informed_consent"`
			SteroidTherapy                        int `bson:"steroid_therapy"`
			TherapyWithImmunosuppressiveDrugs     int `bson:"therapy_with_immunosuppressive_drugs"`
			FemaleSubjectsDuringPregnancy         int `bson:"female_subjects_during_pregnancy"`
			StrokeInLessThanOneYear               int `bson:"stroke_in_less_than_one_year"`
			ChronicSystemicInfections             int `bson:"chronic_systemic_infections"`
			AggravatedAllergicHistory             int `bson:"aggravated_allergic_history"`
			PresenceOfAHistoryOfNeoplasms         int `bson:"presence_of_a_history_of_neoplasms"`
			HistoryOfSplenectomy                  int `bson:"history_of_splenectomy"`
			Neutropenia                           int `bson:"neutropenia"`
			SubjectsWithActiveSyphilis            int `bson:"subjects_with_active_syphilis"`
			Anorexia                              int `bson:"anorexia"`
			ExtensiveTattoos                      int `bson:"extensive_tattoos"`
			TakingNarcoticAndPsychostimulantDrugs int `bson:"taking_narcotic_and_psychostimulant_drugs"`
			SmokingMoretThanTenCigarettesADay     int `bson:"smoking_more_than_ten_cigarettes_a_day"`
			AlcoholIntake                         int `bson:"alcohol_intake"`
			PlannedHospitalization                int `bson:"planned_hospitalization"`
			DonorBloodDonation                    int `bson:"donor_blood_donation"`
			SubjectParticipationInAnyOtherStudy   int `bson:"subject_participation_in_any_other_study"`
			AnyVaccinationInTheLastMonth          int `bson:"any_vaccination_in_the_last_month"`
			InabilityToReadInRussian              int `bson:"inability_to_read_in_russian"`
			ResearchCenterStaff                   int `bson:"research_center_staff"`
			AnyOtherStateOfTheSubjectOfTheStudy   int `bson:"any_other_state_of_the_subject_of_the_study"`
		}
	}
}

type InformationConsentErrors struct {
	Field    string `bson:"field"`
	Reasons  string `bson:"reasons"`
	Comments string `bson:"comments"`
}
