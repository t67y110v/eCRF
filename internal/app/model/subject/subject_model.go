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
			AnthropometricDataBeenMeasuredCondition struct {
				AnthropometricDataBeenMeasured int    `bson:"anthropometric_data_been_measured"`
				Color                          int    `bson:"color"`
				Reason                         string `bson:"reason"`
				Comment                        string `bson:"comment"`
			}
			ReasonIfNotCondition struct {
				ReasonIfNot string `bson:"reason_if_not"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
			}
			DateOfStartMeasuredCondition struct {
				DateOfStartMeasured string `bson:"date_of_start_measured"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			WeightOfBodyCondition struct {
				WeightOfBody int    `bson:"weight_of_body"`
				Color        int    `bson:"color"`
				Reason       string `bson:"reason"`
				Comment      string `bson:"comment"`
			}

			HeightOfBodyCondition struct {
				HeightOfBody int    `bson:"height_of_body"`
				Color        int    `bson:"color"`
				Reason       string `bson:"reason"`
				Comment      string `bson:"comment"`
			}

			IndexWeigthOfBodyCondition struct {
				IndexWeigthOfBody int    `bson:"index_weight_of_body"`
				Color             int    `bson:"color"`
				Reason            string `bson:"reason"`
				Comment           string `bson:"comment"`
			}
		}
		Demography struct {
			SexCondition struct {
				Sex     int    `bson:"sex"`
				Color   int    `bson:"color"`
				Reason  string `bson:"reason"`
				Comment string `bson:"comment"`
			}
			RaceCondition struct {
				Race    int    `bson:"race"`
				Color   int    `bson:"color"`
				Reason  string `bson:"reason"`
				Comment string `bson:"comment"`
			}
			BirthDateCondition struct {
				BirthDate string `bson:"date"`
				Color     int    `bson:"color"`
				Reason    string `bson:"reason"`
				Comment   string `bson:"comment"`
			}
		}
		InclusionCriteria struct {
			PresenceOfAnInformationPanelCondition struct {
				PresenceOfAnInformationPanel int    `bson:"presence_of_an_information_panel"`
				Color                        int    `bson:"color"`
				Reason                       string `bson:"reason"`
				Comment                      string `bson:"comment"`
			}
			Aged18To55YearsCondition struct {
				Aged18To55Years int    `bson:"aged_18_to_55_years"`
				Color           int    `bson:"color"`
				Reason          string `bson:"reason"`
				Comment         string `bson:"comment"`
			}
			NegativeHIVTestResultCondition struct {
				NegativeHIVTestResult int    `bson:"negative_hiv_test_result"`
				Color                 int    `bson:"color"`
				Reason                string `bson:"reason"`
				Comment               string `bson:"comment"`
			}
			BodyMassIndexCondition struct {
				BodyMassIndex int    `bson:"body_mass_index"`
				Color         int    `bson:"color"`
				Reason        string `bson:"reason"`
				Comment       string `bson:"comment"`
			}
			AbsenceOfAcuteInfectiousDiseasesCondition struct {
				AbsenceOfAcuteInfectiousDiseases int    `bson:"absence_of_acute_infectious_diseases"`
				Color                            int    `bson:"color"`
				Reason                           string `bson:"reason"`
				Comment                          string `bson:"comment"`
			}
			ConsentToUseEffectiveMethodsOfContraceptionCondition struct {
				ConsentToUseEffectiveMethodsOfContraception int    `bson:"consent_to_use_effective_methods_of_contraception"`
				Color                                       int    `bson:"color"`
				Reason                                      string `bson:"reason"`
				Comment                                     string `bson:"comment"`
			}
			NegativePregnancyTestCondition struct {
				NegativePregnancyTest int    `bson:"negative_pregnancy_test"`
				Color                 int    `bson:"color"`
				Reason                string `bson:"reason"`
				Comment               string `bson:"comment"`
			}
			NegativeDrugTestCondition struct {
				NegativeDrugTest int    `bson:"negative_drug_test"`
				Color            int    `bson:"color"`
				Reason           string `bson:"reason"`
				Comment          string `bson:"comment"`
			}
			NegativeAlcoholTestCondition struct {
				NegativeAlcoholTest int    `bson:"negative_alcohol_test"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			NoHistoryOfSeverePostVaccinationReactionsCondition struct {
				NoHistoryOfSeverePostVaccinationReactions int    `bson:"no_history_of_severe_post_vaccination_reactions"`
				Color                                     int    `bson:"color"`
				Reason                                    string `bson:"reason"`
				Comment                                   string `bson:"comment"`
			}
			IndicatorsBloodTestsAtScreeningWithinCondition struct {
				IndicatorsBloodTestsAtScreeningWithin int    `bson:"indicators_blood_tests_at_screening_within"`
				Color                                 int    `bson:"color"`
				Reason                                string `bson:"reason"`
				Comment                               string `bson:"comment"`
			}
			NoMyocardialChangesCondition struct {
				NoMyocardialChanges int    `bson:"no_myocardial_changes"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			NegativeTestResultForCOVIDCondition struct {
				NegativeTestResultForCOVID int    `bson:"negative_test_result_for_COVID"`
				Color                      int    `bson:"color"`
				Reason                     string `bson:"reason"`
				Comment                    string `bson:"comment"`
			}
			NoContraindicationsToVaccinationCondition struct {
				NoContraindicationsToVaccination int    `bson:"no_contraindications_to_vaccination"`
				Color                            int    `bson:"color"`
				Reason                           string `bson:"reason"`
				Comment                          string `bson:"comment"`
			}
		}

		ExclusionСriteria struct {
			LackOfSignedInformedConsentCondition struct {
				LackOfSignedInformedConsent int    `bson:"lack_of_signed_informed_consent"`
				Color                       int    `bson:"color"`
				Reason                      string `bson:"reason"`
				Comment                     string `bson:"comment"`
			}
			SteroidTherapyCondition struct {
				SteroidTherapy int    `bson:"steroid_therapy"`
				Color          int    `bson:"color"`
				Reason         string `bson:"reason"`
				Comment        string `bson:"comment"`
			}
			TherapyWithImmunosuppressiveDrugsCondition struct {
				TherapyWithImmunosuppressiveDrugs int    `bson:"therapy_with_immunosuppressive_drugs"`
				Color                             int    `bson:"color"`
				Reason                            string `bson:"reason"`
				Comment                           string `bson:"comment"`
			}
			FemaleSubjectsDuringPregnancyCondition struct {
				FemaleSubjectsDuringPregnancy int    `bson:"female_subjects_during_pregnancy"`
				Color                         int    `bson:"color"`
				Reason                        string `bson:"reason"`
				Comment                       string `bson:"comment"`
			}
			StrokeInLessThanOneYearCondition struct {
				StrokeInLessThanOneYear int    `bson:"stroke_in_less_than_one_year"`
				Color                   int    `bson:"color"`
				Reason                  string `bson:"reason"`
				Comment                 string `bson:"comment"`
			}
			ChronicSystemicInfectionsCondition struct {
				ChronicSystemicInfections int    `bson:"chronic_systemic_infections"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
			}
			AggravatedAllergicHistoryCondition struct {
				AggravatedAllergicHistory int    `bson:"aggravated_allergic_history"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
			}
			PresenceOfAHistoryOfNeoplasmsCondition struct {
				PresenceOfAHistoryOfNeoplasms int    `bson:"presence_of_a_history_of_neoplasms"`
				Color                         int    `bson:"color"`
				Reason                        string `bson:"reason"`
				Comment                       string `bson:"comment"`
			}
			HistoryOfSplenectomyCondition struct {
				HistoryOfSplenectomy int    `bson:"history_of_splenectomy"`
				Color                int    `bson:"color"`
				Reason               string `bson:"reason"`
				Comment              string `bson:"comment"`
			}
			NeutropeniaCondition struct {
				Neutropenia int    `bson:"neutropenia"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
			}
			SubjectsWithActiveSyphilisCondition struct {
				SubjectsWithActiveSyphilis int    `bson:"subjects_with_active_syphilis"`
				Color                      int    `bson:"color"`
				Reason                     string `bson:"reason"`
				Comment                    string `bson:"comment"`
			}
			AnorexiaCondition struct {
				Anorexia int    `bson:"anorexia"`
				Color    int    `bson:"color"`
				Reason   string `bson:"reason"`
				Comment  string `bson:"comment"`
			}
			ExtensiveTattoosCondition struct {
				ExtensiveTattoos int    `bson:"extensive_tattoos"`
				Color            int    `bson:"color"`
				Reason           string `bson:"reason"`
				Comment          string `bson:"comment"`
			}
			TakingNarcoticAndPsychostimulantDrugsCondition struct {
				TakingNarcoticAndPsychostimulantDrugs int    `bson:"taking_narcotic_and_psychostimulant_drugs"`
				Color                                 int    `bson:"color"`
				Reason                                string `bson:"reason"`
				Comment                               string `bson:"comment"`
			}
			SmokingMoretThanTenCigarettesADayCondition struct {
				SmokingMoretThanTenCigarettesADay int    `bson:"smoking_more_than_ten_cigarettes_a_day"`
				Color                             int    `bson:"color"`
				Reason                            string `bson:"reason"`
				Comment                           string `bson:"comment"`
			}
			AlcoholIntakeCondition struct {
				AlcoholIntake int    `bson:"alcohol_intake"`
				Color         int    `bson:"color"`
				Reason        string `bson:"reason"`
				Comment       string `bson:"comment"`
			}
			PlannedHospitalizationConditiont struct {
				PlannedHospitalization int    `bson:"planned_hospitalization"`
				Color                  int    `bson:"color"`
				Reason                 string `bson:"reason"`
				Comment                string `bson:"comment"`
			}
			DonorBloodDonationCondition struct {
				DonorBloodDonation int    `bson:"donor_blood_donation"`
				Color              int    `bson:"color"`
				Reason             string `bson:"reason"`
				Comment            string `bson:"comment"`
			}
			SubjectParticipationInAnyOtherStudyCondition struct {
				SubjectParticipationInAnyOtherStudy int    `bson:"subject_participation_in_any_other_study"`
				Color                               int    `bson:"color"`
				Reason                              string `bson:"reason"`
				Comment                             string `bson:"comment"`
			}
			AnyVaccinationInTheLastMonthCondition struct {
				AnyVaccinationInTheLastMonth int    `bson:"any_vaccination_in_the_last_month"`
				Color                        int    `bson:"color"`
				Reason                       string `bson:"reason"`
				Comment                      string `bson:"comment"`
			}
			InabilityToReadInRussianCondition struct {
				InabilityToReadInRussian int    `bson:"inability_to_read_in_russian"`
				Color                    int    `bson:"color"`
				Reason                   string `bson:"reason"`
				Comment                  string `bson:"comment"`
			}
			ResearchCenterStaffCondition struct {
				ResearchCenterStaff int    `bson:"research_center_staff"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			AnyOtherStateOfTheSubjectOfTheStudyCondition struct {
				AnyOtherStateOfTheSubjectOfTheStudy int    `bson:"any_other_state_of_the_subject_of_the_study"`
				Color                               int    `bson:"color"`
				Reason                              string `bson:"reason"`
				Comment                             string `bson:"comment"`
			}
		}
		CompletionOfScreening struct {
			VolunteerEligibleCondition struct {
				VolunteerEligible int    `bson:"volunteer_eligible"`
				Color             int    `bson:"color"`
				Reason            string `bson:"reason"`
				Comment           string `bson:"comment"`
			}
			NoExclusionCriteriaCondition struct {
				NoExclusionCriteria int    `bson:"no_exclusion_criteria"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			InformedOfTheRestrictionsCondition struct {
				InformedOfTheRestrictions int    `bson:"informed_of_the_restrictions"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
			}
			VolunteerIncludedCondition struct {
				VolunteerIncluded int    `bson:"volunteer_included"`
				Color             int    `bson:"color"`
				Reason            string `bson:"reason"`
				Comment           string `bson:"comment"`
			}
			ReasonIfNotCondition struct {
				ReasonIfNot string `bson:"reason_if_not"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
			}
			CommentCondition struct {
				CommentValue string `bson:"comment_value"`
				Color        int    `bson:"color"`
				Reason       string `bson:"reason"`
				Comment      string `bson:"comment"`
			}
		}
	}
}

type InformationConsentErrors struct {
	Field    string `bson:"field"`
	Reasons  string `bson:"reasons"`
	Comments string `bson:"comments"`
}
