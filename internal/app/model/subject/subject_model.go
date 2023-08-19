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
		StartOfScreening struct {
			DateStartOfScreeningCondition struct {
				DateOfStart string `bson:"dateofstart"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			TimeStartOfScreeningCondition struct {
				TimeOfStart string `bson:"timeofstart"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
		}
		InformaionConsent struct {
			SignedCondition struct {
				Signed      int    `bson:"signed"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			DateOfSignCondition struct {
				DateOfSign  string `bson:"dateofsign"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			TimeOfSignCondition struct {
				TimeOfSign  string `bson:"timeofsign"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			ReceivedAnInsurancePolicyCondition struct {
				ReceivedAnInsurancePolicy int    `bson:"receivedaninsurancepolicy"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
				Sender                    string `bson:"sender"`
				SendersRole               string `bson:"sendersrole"`
			}
			ReceivedAnInformaionConsentCondition struct {
				ReceivedAnInformaionConsent int    `bson:"receivedaninformaionconsent"`
				Color                       int    `bson:"color"`
				Reason                      string `bson:"reason"`
				Comment                     string `bson:"comment"`
				Sender                      string `bson:"sender"`
				SendersRole                 string `bson:"sendersrole"`
			}
		}
		Anthropometry struct {
			AnthropometricDataBeenMeasuredCondition struct {
				AnthropometricDataBeenMeasured int    `bson:"anthropometricdatabeenmeasured"`
				Color                          int    `bson:"color"`
				Reason                         string `bson:"reason"`
				Comment                        string `bson:"comment"`
				Sender                         string `bson:"sender"`
				SendersRole                    string `bson:"sendersrole"`
			}
			ReasonIfNotCondition struct {
				ReasonIfNot string `bson:"reasonifnot"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			DateOfStartMeasuredCondition struct {
				DateOfStartMeasured string `bson:"dateofstartmeasured"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
				Sender              string `bson:"sender"`
				SendersRole         string `bson:"sendersrole"`
			}
			WeightOfBodyCondition struct {
				WeightOfBody int    `bson:"weightofbody"`
				Color        int    `bson:"color"`
				Reason       string `bson:"reason"`
				Comment      string `bson:"comment"`
				Sender       string `bson:"sender"`
				SendersRole  string `bson:"sendersrole"`
			}

			HeightOfBodyCondition struct {
				HeightOfBody int    `bson:"heightofbody"`
				Color        int    `bson:"color"`
				Reason       string `bson:"reason"`
				Comment      string `bson:"comment"`
				Sender       string `bson:"sender"`
				SendersRole  string `bson:"sendersrole"`
			}

			IndexWeigthOfBodyCondition struct {
				IndexWeigthOfBody int    `bson:"indexweightofbody"`
				Color             int    `bson:"color"`
				Reason            string `bson:"reason"`
				Comment           string `bson:"comment"`
				Sender            string `bson:"sender"`
				SendersRole       string `bson:"sendersrole"`
			}
		}
		Demography struct {
			SexCondition struct {
				Sex         int    `bson:"sex"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			RaceCondition struct {
				Race        int    `bson:"race"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			BirthDateCondition struct {
				BirthDate   string `bson:"date"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
		}
		InclusionCriteria struct {
			PresenceOfAnInformationPanelCondition struct {
				PresenceOfAnInformationPanel int    `bson:"presenceofaninformationpanel"`
				Color                        int    `bson:"color"`
				Reason                       string `bson:"reason"`
				Comment                      string `bson:"comment"`
				Sender                       string `bson:"sender"`
				SendersRole                  string `bson:"sendersrole"`
			}
			Aged18To55YearsCondition struct {
				Aged18To55Years int    `bson:"aged18to55years"`
				Color           int    `bson:"color"`
				Reason          string `bson:"reason"`
				Comment         string `bson:"comment"`
				Sender          string `bson:"sender"`
				SendersRole     string `bson:"sendersrole"`
			}
			NegativeHIVTestResultCondition struct {
				NegativeHIVTestResult int    `bson:"negativehivtestresult"`
				Color                 int    `bson:"color"`
				Reason                string `bson:"reason"`
				Comment               string `bson:"comment"`
				Sender                string `bson:"sender"`
				SendersRole           string `bson:"sendersrole"`
			}
			BodyMassIndexCondition struct {
				BodyMassIndex int    `bson:"bodymassindex"`
				Color         int    `bson:"color"`
				Reason        string `bson:"reason"`
				Comment       string `bson:"comment"`
				Sender        string `bson:"sender"`
				SendersRole   string `bson:"sendersrole"`
			}
			AbsenceOfAcuteInfectiousDiseasesCondition struct {
				AbsenceOfAcuteInfectiousDiseases int    `bson:"absenceofacuteinfectiousdiseases"`
				Color                            int    `bson:"color"`
				Reason                           string `bson:"reason"`
				Comment                          string `bson:"comment"`
				Sender                           string `bson:"sender"`
				SendersRole                      string `bson:"sendersrole"`
			}
			ConsentToUseEffectiveMethodsOfContraceptionCondition struct {
				ConsentToUseEffectiveMethodsOfContraception int    `bson:"consenttouseeffectivemethodsofcontraception"`
				Color                                       int    `bson:"color"`
				Reason                                      string `bson:"reason"`
				Comment                                     string `bson:"comment"`
				Sender                                      string `bson:"sender"`
				SendersRole                                 string `bson:"sendersrole"`
			}
			NegativePregnancyTestCondition struct {
				NegativePregnancyTest int    `bson:"negativepregnancytest"`
				Color                 int    `bson:"color"`
				Reason                string `bson:"reason"`
				Comment               string `bson:"comment"`
				Sender                string `bson:"sender"`
				SendersRole           string `bson:"sendersrole"`
			}
			NegativeDrugTestCondition struct {
				NegativeDrugTest int    `bson:"negativedrugtest"`
				Color            int    `bson:"color"`
				Reason           string `bson:"reason"`
				Comment          string `bson:"comment"`
				Sender           string `bson:"sender"`
				SendersRole      string `bson:"sendersrole"`
			}
			NegativeAlcoholTestCondition struct {
				NegativeAlcoholTest int    `bson:"negativealcoholtest"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
				Sender              string `bson:"sender"`
				SendersRole         string `bson:"sendersrole"`
			}
			NoHistoryOfSeverePostVaccinationReactionsCondition struct {
				NoHistoryOfSeverePostVaccinationReactions int    `bson:"nohistoryofseverepostvaccinationreactions"`
				Color                                     int    `bson:"color"`
				Reason                                    string `bson:"reason"`
				Comment                                   string `bson:"comment"`
				Sender                                    string `bson:"sender"`
				SendersRole                               string `bson:"sendersrole"`
			}
			IndicatorsBloodTestsAtScreeningWithinCondition struct {
				IndicatorsBloodTestsAtScreeningWithin int    `bson:"indicatorsbloodtestsatscreeningwithin"`
				Color                                 int    `bson:"color"`
				Reason                                string `bson:"reason"`
				Comment                               string `bson:"comment"`
				Sender                                string `bson:"sender"`
				SendersRole                           string `bson:"sendersrole"`
			}
			NoMyocardialChangesCondition struct {
				NoMyocardialChanges int    `bson:"nomyocardialchanges"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
				Sender              string `bson:"sender"`
				SendersRole         string `bson:"sendersrole"`
			}
			NegativeTestResultForCOVIDCondition struct {
				NegativeTestResultForCOVID int    `bson:"negativetestresultforCOVID"`
				Color                      int    `bson:"color"`
				Reason                     string `bson:"reason"`
				Comment                    string `bson:"comment"`
				Sender                     string `bson:"sender"`
				SendersRole                string `bson:"sendersrole"`
			}
			NoContraindicationsToVaccinationCondition struct {
				NoContraindicationsToVaccination int    `bson:"nocontraindicationstovaccination"`
				Color                            int    `bson:"color"`
				Reason                           string `bson:"reason"`
				Comment                          string `bson:"comment"`
				Sender                           string `bson:"sender"`
				SendersRole                      string `bson:"sendersrole"`
			}
		}

		ExclusionСriteria struct {
			LackOfSignedInformedConsentCondition struct {
				LackOfSignedInformedConsent int    `bson:"lackofsignedinformedconsent"`
				Color                       int    `bson:"color"`
				Reason                      string `bson:"reason"`
				Comment                     string `bson:"comment"`
				Sender                      string `bson:"sender"`
				SendersRole                 string `bson:"sendersrole"`
			}
			SteroidTherapyCondition struct {
				SteroidTherapy int    `bson:"steroidtherapy"`
				Color          int    `bson:"color"`
				Reason         string `bson:"reason"`
				Comment        string `bson:"comment"`
				Sender         string `bson:"sender"`
				SendersRole    string `bson:"sendersrole"`
			}
			TherapyWithImmunosuppressiveDrugsCondition struct {
				TherapyWithImmunosuppressiveDrugs int    `bson:"therapywithimmunosuppressivedrugs"`
				Color                             int    `bson:"color"`
				Reason                            string `bson:"reason"`
				Comment                           string `bson:"comment"`
				Sender                            string `bson:"sender"`
				SendersRole                       string `bson:"sendersrole"`
			}
			FemaleSubjectsDuringPregnancyCondition struct {
				FemaleSubjectsDuringPregnancy int    `bson:"femalesubjectsduringpregnancy"`
				Color                         int    `bson:"color"`
				Reason                        string `bson:"reason"`
				Comment                       string `bson:"comment"`
				Sender                        string `bson:"sender"`
				SendersRole                   string `bson:"sendersrole"`
			}
			StrokeInLessThanOneYearCondition struct {
				StrokeInLessThanOneYear int    `bson:"strokeinlessthanoneyear"`
				Color                   int    `bson:"color"`
				Reason                  string `bson:"reason"`
				Comment                 string `bson:"comment"`
				Sender                  string `bson:"sender"`
				SendersRole             string `bson:"sendersrole"`
			}
			ChronicSystemicInfectionsCondition struct {
				ChronicSystemicInfections int    `bson:"chronicsystemicinfections"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
				Sender                    string `bson:"sender"`
				SendersRole               string `bson:"sendersrole"`
			}
			AggravatedAllergicHistoryCondition struct {
				AggravatedAllergicHistory int    `bson:"aggravatedallergichistory"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
				Sender                    string `bson:"sender"`
				SendersRole               string `bson:"sendersrole"`
			}
			PresenceOfAHistoryOfNeoplasmsCondition struct {
				PresenceOfAHistoryOfNeoplasms int    `bson:"presenceofahistoryofneoplasms"`
				Color                         int    `bson:"color"`
				Reason                        string `bson:"reason"`
				Comment                       string `bson:"comment"`
				Sender                        string `bson:"sender"`
				SendersRole                   string `bson:"sendersrole"`
			}
			HistoryOfSplenectomyCondition struct {
				HistoryOfSplenectomy int    `bson:"historyofsplenectomy"`
				Color                int    `bson:"color"`
				Reason               string `bson:"reason"`
				Comment              string `bson:"comment"`
				Sender               string `bson:"sender"`
				SendersRole          string `bson:"sendersrole"`
			}
			NeutropeniaCondition struct {
				Neutropenia int    `bson:"neutropenia"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			SubjectsWithActiveSyphilisCondition struct {
				SubjectsWithActiveSyphilis int    `bson:"subjectswithactivesyphilis"`
				Color                      int    `bson:"color"`
				Reason                     string `bson:"reason"`
				Comment                    string `bson:"comment"`
				Sender                     string `bson:"sender"`
				SendersRole                string `bson:"sendersrole"`
			}
			AnorexiaCondition struct {
				Anorexia    int    `bson:"anorexia"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			ExtensiveTattoosCondition struct {
				ExtensiveTattoos int    `bson:"extensivetattoos"`
				Color            int    `bson:"color"`
				Reason           string `bson:"reason"`
				Comment          string `bson:"comment"`
				Sender           string `bson:"sender"`
				SendersRole      string `bson:"sendersrole"`
			}
			TakingNarcoticAndPsychostimulantDrugsCondition struct {
				TakingNarcoticAndPsychostimulantDrugs int    `bson:"takingnarcoticandpsychostimulantdrugs"`
				Color                                 int    `bson:"color"`
				Reason                                string `bson:"reason"`
				Comment                               string `bson:"comment"`
				Sender                                string `bson:"sender"`
				SendersRole                           string `bson:"sendersrole"`
			}
			SmokingMoretThanTenCigarettesADayCondition struct {
				SmokingMoretThanTenCigarettesADay int    `bson:"smokingmorethantencigarettesaday"`
				Color                             int    `bson:"color"`
				Reason                            string `bson:"reason"`
				Comment                           string `bson:"comment"`
				Sender                            string `bson:"sender"`
				SendersRole                       string `bson:"sendersrole"`
			}
			AlcoholIntakeCondition struct {
				AlcoholIntake int    `bson:"alcoholintake"`
				Color         int    `bson:"color"`
				Reason        string `bson:"reason"`
				Comment       string `bson:"comment"`
				Sender        string `bson:"sender"`
				SendersRole   string `bson:"sendersrole"`
			}
			PlannedHospitalizationConditiont struct {
				PlannedHospitalization int    `bson:"plannedhospitalization"`
				Color                  int    `bson:"color"`
				Reason                 string `bson:"reason"`
				Comment                string `bson:"comment"`
				Sender                 string `bson:"sender"`
				SendersRole            string `bson:"sendersrole"`
			}
			DonorBloodDonationCondition struct {
				DonorBloodDonation int    `bson:"donorblooddonation"`
				Color              int    `bson:"color"`
				Reason             string `bson:"reason"`
				Comment            string `bson:"comment"`
				Sender             string `bson:"sender"`
				SendersRole        string `bson:"sendersrole"`
			}
			SubjectParticipationInAnyOtherStudyCondition struct {
				SubjectParticipationInAnyOtherStudy int    `bson:"subjectparticipationinanyotherstudy"`
				Color                               int    `bson:"color"`
				Reason                              string `bson:"reason"`
				Comment                             string `bson:"comment"`
				Sender                              string `bson:"sender"`
				SendersRole                         string `bson:"sendersrole"`
			}
			AnyVaccinationInTheLastMonthCondition struct {
				AnyVaccinationInTheLastMonth int    `bson:"anyvaccinationinthelastmonth"`
				Color                        int    `bson:"color"`
				Reason                       string `bson:"reason"`
				Comment                      string `bson:"comment"`
				Sender                       string `bson:"sender"`
				SendersRole                  string `bson:"sendersrole"`
			}
			InabilityToReadInRussianCondition struct {
				InabilityToReadInRussian int    `bson:"inabilitytoreadinrussian"`
				Color                    int    `bson:"color"`
				Reason                   string `bson:"reason"`
				Comment                  string `bson:"comment"`
				Sender                   string `bson:"sender"`
				SendersRole              string `bson:"sendersrole"`
			}
			ResearchCenterStaffCondition struct {
				ResearchCenterStaff int    `bson:"researchcenterstaff"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
				Sender              string `bson:"sender"`
				SendersRole         string `bson:"sendersrole"`
			}
			AnyOtherStateOfTheSubjectOfTheStudyCondition struct {
				AnyOtherStateOfTheSubjectOfTheStudy int    `bson:"anyotherstateofthesubjectofthestudy"`
				Color                               int    `bson:"color"`
				Reason                              string `bson:"reason"`
				Comment                             string `bson:"comment"`
				Sender                              string `bson:"sender"`
				SendersRole                         string `bson:"sendersrole"`
			}
		}
		CompletionOfScreening struct {
			VolunteerEligibleCondition struct {
				VolunteerEligible int    `bson:"volunteereligible"`
				Color             int    `bson:"color"`
				Reason            string `bson:"reason"`
				Comment           string `bson:"comment"`
				Sender            string `bson:"sender"`
				SendersRole       string `bson:"sendersrole"`
			}
			NoExclusionCriteriaCondition struct {
				NoExclusionCriteria int    `bson:"noexclusioncriteria"`
				Color               int    `bson:"color"`
				Reason              string `bson:"reason"`
				Comment             string `bson:"comment"`
				Sender              string `bson:"sender"`
				SendersRole         string `bson:"sendersrole"`
			}
			InformedOfTheRestrictionsCondition struct {
				InformedOfTheRestrictions int    `bson:"informedoftherestrictions"`
				Color                     int    `bson:"color"`
				Reason                    string `bson:"reason"`
				Comment                   string `bson:"comment"`
				Sender                    string `bson:"sender"`
				SendersRole               string `bson:"sendersrole"`
			}
			VolunteerIncludedCondition struct {
				VolunteerIncluded int    `bson:"volunteerincluded"`
				Color             int    `bson:"color"`
				Reason            string `bson:"reason"`
				Comment           string `bson:"comment"`
				Sender            string `bson:"sender"`
				SendersRole       string `bson:"sendersrole"`
			}
			ReasonIfNotCondition struct {
				ReasonIfNot string `bson:"reasonifnot"`
				Color       int    `bson:"color"`
				Reason      string `bson:"reason"`
				Comment     string `bson:"comment"`
				Sender      string `bson:"sender"`
				SendersRole string `bson:"sendersrole"`
			}
			CommentCondition struct {
				CommentValue string `bson:"commentvalue"`
				Color        int    `bson:"color"`
				Reason       string `bson:"reason"`
				Comment      string `bson:"comment"`
				Sender       string `bson:"sender"`
				SendersRole  string `bson:"sendersrole"`
			}
		}
	}
}

type AdverseEventsAr struct {
	AdverseEventsRegisteredCondition struct {
		AdverseEventsRegistered int    `bson:"adverseeventsregistered"`
		Color                   int    `bson:"color"`
		Reason                  string `bson:"reason"`
		Comment                 string `bson:"comment"`
		Sender                  string `bson:"sender"`
		SendersRole             string `bson:"sendersrole"`
	}
	DescriptionOfTheAdverseEventCondition struct {
		DescriptionOfTheAdverseEvent string `bson:"descriptionoftheadverseevent"`
		Color                        int    `bson:"color"`
		Reason                       string `bson:"reason"`
		Comment                      string `bson:"comment"`
		Sender                       string `bson:"sender"`
		SendersRole                  string `bson:"sendersrole"`
	}
	DateOfStartAECondition struct {
		DateOfStartAE string `bson:"dateofstartae"`
		Color         int    `bson:"color"`
		Reason        string `bson:"reason"`
		Comment       string `bson:"comment"`
		Sender        string `bson:"sender"`
		SendersRole   string `bson:"sendersrole"`
	}
	DateOfEndAECondition struct {
		DateOfEndAE string `bson:"dateofendae"`
		IsContinues int    `bson:"iscontinius"`
		Color       int    `bson:"color"`
		Reason      string `bson:"reason"`
		Comment     string `bson:"comment"`
		Sender      string `bson:"sender"`
		SendersRole string `bson:"sendersrole"`
	}
	SeverityCondition struct {
		Severity    int    `bson:"severity"`
		Color       int    `bson:"color"`
		Reason      string `bson:"reason"`
		Comment     string `bson:"comment"`
		Sender      string `bson:"sender"`
		SendersRole string `bson:"sendersrole"`
	}
	RecurringPhenomenonCondition struct {
		RecurringPhenomenon int    `bson:"recurringphenomenon"`
		Color               int    `bson:"color"`
		Reason              string `bson:"reason"`
		Comment             string `bson:"comment"`
		Sender              string `bson:"sender"`
		SendersRole         string `bson:"sendersrole"`
	}

	AssociationWithTheDrugUsedCondition struct {
		AssociationWithTheDrugUsed int    `bson:"associationwiththedrugused"`
		Color                      int    `bson:"color"`
		Reason                     string `bson:"reason"`
		Comment                    string `bson:"comment"`
		Sender                     string `bson:"sender"`
		SendersRole                string `bson:"sendersrole"`
	}
	ForesightCondition struct {
		Foresight   int    `bson:"foresight"`
		Color       int    `bson:"color"`
		Reason      string `bson:"reason"`
		Comment     string `bson:"comment"`
		Sender      string `bson:"sender"`
		SendersRole string `bson:"sendersrole"`
	}

	ConnectionBetweenAEAndDUCondition struct {
		ConnectionBetweenAEAndDU int    `bson:"ConnectionBetweenAEAndDU"`
		Color                    int    `bson:"color"`
		Reason                   string `bson:"reason"`
		Comment                  string `bson:"comment"`
		Sender                   string `bson:"sender"`
		SendersRole              string `bson:"sendersrole"`
	}
	RenewalAfterUseCondition struct {
		RenewalAfterUse int    `bson:"renewalafteruse"`
		Color           int    `bson:"color"`
		Reason          string `bson:"reason"`
		Comment         string `bson:"comment"`
		Sender          string `bson:"sender"`
		SendersRole     string `bson:"sendersrole"`
	}
	LocalReactionCondition struct {
		LocalReaction int    `bson:"LocalReaction"`
		Color         int    `bson:"color"`
		Reason        string `bson:"reason"`
		Comment       string `bson:"comment"`
		Sender        string `bson:"sender"`
		SendersRole   string `bson:"sendersrole"`
	}

	SubjectDropoutCondition struct {
		SubjectDropout int    `bson:"SubjectDropout"`
		Color          int    `bson:"color"`
		Reason         string `bson:"reason"`
		Comment        string `bson:"comment"`
		Sender         string `bson:"sender"`
		SendersRole    string `bson:"sendersrole"`
	}
	MeasuresTakenCondition struct {
		MeasuresTaken int    `bson:"MeasuresTaken"`
		Color         int    `bson:"color"`
		Reason        string `bson:"reason"`
		Comment       string `bson:"comment"`
		Sender        string `bson:"sender"`
		SendersRole   string `bson:"sendersrole"`
	}
	MeasuresTakenOnUDCondition struct {
		MeasuresTakenOnUD int    `bson:"MeasuresTakenOnUD"`
		Color             int    `bson:"color"`
		Reason            string `bson:"reason"`
		Comment           string `bson:"comment"`
		Sender            string `bson:"sender"`
		SendersRole       string `bson:"sendersrole"`
	}
	ExodusCondition struct {
		Exodus      int    `bson:"Exodus"`
		Color       int    `bson:"color"`
		Reason      string `bson:"reason"`
		Comment     string `bson:"comment"`
		Sender      string `bson:"sender"`
		SendersRole string `bson:"sendersrole"`
	}
	IsItSeriousCondition struct {
		IsItSerious int    `bson:"IsItSerious"`
		Color       int    `bson:"color"`
		Reason      string `bson:"reason"`
		Comment     string `bson:"comment"`
		Sender      string `bson:"sender"`
		SendersRole string `bson:"sendersrole"`
	}
	SeverityCriterionCondition struct {
		SeverityCriterion int    `bson:"severitycriterion"`
		Color             int    `bson:"color"`
		Reason            string `bson:"reason"`
		Comment           string `bson:"comment"`
		Sender            string `bson:"sender"`
		SendersRole       string `bson:"sendersrole"`
	}
	TestImpactCondition struct {
		TestImpact  int    `bson:"testimpact"`
		Color       int    `bson:"color"`
		Reason      string `bson:"reason"`
		Comment     string `bson:"comment"`
		Sender      string `bson:"sender"`
		SendersRole string `bson:"sendersrole"`
	}

	DoseEffectCondition struct {
		DoseEffect  int    `bson:"DoseEffect"`
		Color       int    `bson:"color"`
		Reason      string `bson:"reason"`
		Comment     string `bson:"comment"`
		Sender      string `bson:"sender"`
		SendersRole string `bson:"sendersrole"`
	}
	ImpactOnHospitalStayCondition struct {
		ImpactOnHospitalStay int    `bson:"impactonhospitalstay"`
		Color                int    `bson:"color"`
		Reason               string `bson:"reason"`
		Comment              string `bson:"comment"`
		Sender               string `bson:"sender"`
		SendersRole          string `bson:"sendersrole"`
	}
	RelationshipWithMedicationCondition struct {
		RelationshipWithMedication int    `bson:"relationshipwithmedication"`
		Color                      int    `bson:"color"`
		Reason                     string `bson:"reason"`
		Comment                    string `bson:"comment"`
		Sender                     string `bson:"sender"`
		SendersRole                string `bson:"sendersrole"`
	}
	ExpectancyCondition struct {
		Expectancy  int    `bson:"expectancy"`
		Color       int    `bson:"color"`
		Reason      string `bson:"reason"`
		Comment     string `bson:"comment"`
		Sender      string `bson:"sender"`
		SendersRole string `bson:"sendersrole"`
	}
}

type InformationConsentErrors struct {
	Field    string `bson:"field"`
	Reasons  string `bson:"reasons"`
	Comments string `bson:"comments"`
}
