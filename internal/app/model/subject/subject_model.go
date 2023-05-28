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
				DateOfSign string `bson:"dateofsign"`
				Color      int    `bson:"color"`
				Reason     string `bson:"reason"`
				Comment    string `bson:"comment"`
			}
			TimeOfSignCondition struct {
				TimeOfSign string `bson:"timeofsign"`
				Color      int    `bson:"color"`
				Reason     string `bson:"reason"`
				Comment    string `bson:"comment"`
			}
			ReceivedAnInsurancePolicyCondition struct {
				ReceivedAnInsurancePolicy int    `bson:"receivedaninsurancepolicy"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
			}
			ReceivedAnInformaionConsentCondition struct {
				ReceivedAnInformaionConsent int    `bson:"receivedaninformaionconsent"`
				Color                       int    `bson:"color"`
				Reason                      string `bson:"reason"`
				Comment                     string `bson:"comment"`
			}
		}
		Anthropometry struct {
			AnthropometricDataBeenMeasuredCondition struct {
				AnthropometricDataBeenMeasured int    `bson:"anthropometricdatabeenmeasured"`
				Color                          int    `bson:"color"`
				Reason                         string `bson:"reason"`
				Comment                        string `bson:"comment"`
			}
			ReasonIfNotCondition struct {
				ReasonIfNot string `bson:"reasonifnot"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
			}
			DateOfStartMeasuredCondition struct {
				DateOfStartMeasured string `bson:"dateofstartmeasured"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			WeightOfBodyCondition struct {
				WeightOfBody int    `bson:"weightofbody"`
				Color        int    `bson:"color"`
				Reason       string `bson:"reason"`
				Comment      string `bson:"comment"`
			}

			HeightOfBodyCondition struct {
				HeightOfBody int    `bson:"heightofbody"`
				Color        int    `bson:"color"`
				Reason       string `bson:"reason"`
				Comment      string `bson:"comment"`
			}

			IndexWeigthOfBodyCondition struct {
				IndexWeigthOfBody int    `bson:"indexweightofbody"`
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
				PresenceOfAnInformationPanel int    `bson:"presenceofaninformationpanel"`
				Color                        int    `bson:"color"`
				Reason                       string `bson:"reason"`
				Comment                      string `bson:"comment"`
			}
			Aged18To55YearsCondition struct {
				Aged18To55Years int    `bson:"aged18to55years"`
				Color           int    `bson:"color"`
				Reason          string `bson:"reason"`
				Comment         string `bson:"comment"`
			}
			NegativeHIVTestResultCondition struct {
				NegativeHIVTestResult int    `bson:"negativehivtestresult"`
				Color                 int    `bson:"color"`
				Reason                string `bson:"reason"`
				Comment               string `bson:"comment"`
			}
			BodyMassIndexCondition struct {
				BodyMassIndex int    `bson:"bodymassindex"`
				Color         int    `bson:"color"`
				Reason        string `bson:"reason"`
				Comment       string `bson:"comment"`
			}
			AbsenceOfAcuteInfectiousDiseasesCondition struct {
				AbsenceOfAcuteInfectiousDiseases int    `bson:"absenceofacuteinfectiousdiseases"`
				Color                            int    `bson:"color"`
				Reason                           string `bson:"reason"`
				Comment                          string `bson:"comment"`
			}
			ConsentToUseEffectiveMethodsOfContraceptionCondition struct {
				ConsentToUseEffectiveMethodsOfContraception int    `bson:"consenttouseeffectivemethodsofcontraception"`
				Color                                       int    `bson:"color"`
				Reason                                      string `bson:"reason"`
				Comment                                     string `bson:"comment"`
			}
			NegativePregnancyTestCondition struct {
				NegativePregnancyTest int    `bson:"negativepregnancytest"`
				Color                 int    `bson:"color"`
				Reason                string `bson:"reason"`
				Comment               string `bson:"comment"`
			}
			NegativeDrugTestCondition struct {
				NegativeDrugTest int    `bson:"negativedrugtest"`
				Color            int    `bson:"color"`
				Reason           string `bson:"reason"`
				Comment          string `bson:"comment"`
			}
			NegativeAlcoholTestCondition struct {
				NegativeAlcoholTest int    `bson:"negativealcoholtest"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			NoHistoryOfSeverePostVaccinationReactionsCondition struct {
				NoHistoryOfSeverePostVaccinationReactions int    `bson:"nohistoryofseverepostvaccinationreactions"`
				Color                                     int    `bson:"color"`
				Reason                                    string `bson:"reason"`
				Comment                                   string `bson:"comment"`
			}
			IndicatorsBloodTestsAtScreeningWithinCondition struct {
				IndicatorsBloodTestsAtScreeningWithin int    `bson:"indicatorsbloodtestsatscreeningwithin"`
				Color                                 int    `bson:"color"`
				Reason                                string `bson:"reason"`
				Comment                               string `bson:"comment"`
			}
			NoMyocardialChangesCondition struct {
				NoMyocardialChanges int    `bson:"nomyocardialchanges"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			NegativeTestResultForCOVIDCondition struct {
				NegativeTestResultForCOVID int    `bson:"negativetestresultforCOVID"`
				Color                      int    `bson:"color"`
				Reason                     string `bson:"reason"`
				Comment                    string `bson:"comment"`
			}
			NoContraindicationsToVaccinationCondition struct {
				NoContraindicationsToVaccination int    `bson:"nocontraindicationstovaccination"`
				Color                            int    `bson:"color"`
				Reason                           string `bson:"reason"`
				Comment                          string `bson:"comment"`
			}
		}

		ExclusionСriteria struct {
			LackOfSignedInformedConsentCondition struct {
				LackOfSignedInformedConsent int    `bson:"lackofsignedinformedconsent"`
				Color                       int    `bson:"color"`
				Reason                      string `bson:"reason"`
				Comment                     string `bson:"comment"`
			}
			SteroidTherapyCondition struct {
				SteroidTherapy int    `bson:"steroidtherapy"`
				Color          int    `bson:"color"`
				Reason         string `bson:"reason"`
				Comment        string `bson:"comment"`
			}
			TherapyWithImmunosuppressiveDrugsCondition struct {
				TherapyWithImmunosuppressiveDrugs int    `bson:"therapywithimmunosuppressivedrugs"`
				Color                             int    `bson:"color"`
				Reason                            string `bson:"reason"`
				Comment                           string `bson:"comment"`
			}
			FemaleSubjectsDuringPregnancyCondition struct {
				FemaleSubjectsDuringPregnancy int    `bson:"femalesubjectsduringpregnancy"`
				Color                         int    `bson:"color"`
				Reason                        string `bson:"reason"`
				Comment                       string `bson:"comment"`
			}
			StrokeInLessThanOneYearCondition struct {
				StrokeInLessThanOneYear int    `bson:"strokeinlessthanoneyear"`
				Color                   int    `bson:"color"`
				Reason                  string `bson:"reason"`
				Comment                 string `bson:"comment"`
			}
			ChronicSystemicInfectionsCondition struct {
				ChronicSystemicInfections int    `bson:"chronicsystemicinfections"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
			}
			AggravatedAllergicHistoryCondition struct {
				AggravatedAllergicHistory int    `bson:"aggravatedallergichistory"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
			}
			PresenceOfAHistoryOfNeoplasmsCondition struct {
				PresenceOfAHistoryOfNeoplasms int    `bson:"presenceofahistoryofneoplasms"`
				Color                         int    `bson:"color"`
				Reason                        string `bson:"reason"`
				Comment                       string `bson:"comment"`
			}
			HistoryOfSplenectomyCondition struct {
				HistoryOfSplenectomy int    `bson:"historyofsplenectomy"`
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
				SubjectsWithActiveSyphilis int    `bson:"subjectswithactivesyphilis"`
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
				ExtensiveTattoos int    `bson:"extensivetattoos"`
				Color            int    `bson:"color"`
				Reason           string `bson:"reason"`
				Comment          string `bson:"comment"`
			}
			TakingNarcoticAndPsychostimulantDrugsCondition struct {
				TakingNarcoticAndPsychostimulantDrugs int    `bson:"takingnarcoticandpsychostimulantdrugs"`
				Color                                 int    `bson:"color"`
				Reason                                string `bson:"reason"`
				Comment                               string `bson:"comment"`
			}
			SmokingMoretThanTenCigarettesADayCondition struct {
				SmokingMoretThanTenCigarettesADay int    `bson:"smokingmorethantencigarettesaday"`
				Color                             int    `bson:"color"`
				Reason                            string `bson:"reason"`
				Comment                           string `bson:"comment"`
			}
			AlcoholIntakeCondition struct {
				AlcoholIntake int    `bson:"alcoholintake"`
				Color         int    `bson:"color"`
				Reason        string `bson:"reason"`
				Comment       string `bson:"comment"`
			}
			PlannedHospitalizationConditiont struct {
				PlannedHospitalization int    `bson:"plannedhospitalization"`
				Color                  int    `bson:"color"`
				Reason                 string `bson:"reason"`
				Comment                string `bson:"comment"`
			}
			DonorBloodDonationCondition struct {
				DonorBloodDonation int    `bson:"donorblooddonation"`
				Color              int    `bson:"color"`
				Reason             string `bson:"reason"`
				Comment            string `bson:"comment"`
			}
			SubjectParticipationInAnyOtherStudyCondition struct {
				SubjectParticipationInAnyOtherStudy int    `bson:"subjectparticipationinanyotherstudy"`
				Color                               int    `bson:"color"`
				Reason                              string `bson:"reason"`
				Comment                             string `bson:"comment"`
			}
			AnyVaccinationInTheLastMonthCondition struct {
				AnyVaccinationInTheLastMonth int    `bson:"anyvaccinationinthelastmonth"`
				Color                        int    `bson:"color"`
				Reason                       string `bson:"reason"`
				Comment                      string `bson:"comment"`
			}
			InabilityToReadInRussianCondition struct {
				InabilityToReadInRussian int    `bson:"inabilitytoreadinrussian"`
				Color                    int    `bson:"color"`
				Reason                   string `bson:"reason"`
				Comment                  string `bson:"comment"`
			}
			ResearchCenterStaffCondition struct {
				ResearchCenterStaff int    `bson:"researchcenterstaff"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			AnyOtherStateOfTheSubjectOfTheStudyCondition struct {
				AnyOtherStateOfTheSubjectOfTheStudy int    `bson:"anyotherstateofthesubjectofthestudy"`
				Color                               int    `bson:"color"`
				Reason                              string `bson:"reason"`
				Comment                             string `bson:"comment"`
			}
		}
		CompletionOfScreening struct {
			VolunteerEligibleCondition struct {
				VolunteerEligible int    `bson:"volunteereligible"`
				Color             int    `bson:"color"`
				Reason            string `bson:"reason"`
				Comment           string `bson:"comment"`
			}
			NoExclusionCriteriaCondition struct {
				NoExclusionCriteria int    `bson:"noexclusioncriteria"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
			}
			InformedOfTheRestrictionsCondition struct {
				InformedOfTheRestrictions int    `bson:"informedoftherestrictions"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
			}
			VolunteerIncludedCondition struct {
				VolunteerIncluded int    `bson:"volunteerincluded"`
				Color             int    `bson:"color"`
				Reason            string `bson:"reason"`
				Comment           string `bson:"comment"`
			}
			ReasonIfNotCondition struct {
				ReasonIfNot string `bson:"reasonifnot"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
			}
			CommentCondition struct {
				CommentValue string `bson:"commentvalue"`
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
