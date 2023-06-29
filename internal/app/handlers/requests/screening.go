package requests

type StartOfScreening struct {
	ProtocolID           int    `json:"protocol_id"`
	SubjectNumber        string `json:"subject_num"`
	DateStartOfScreening string `json:"dateofstart"`
	TimeStartOfScreening string `json:"timeofstart"`
}

type InformationConsent struct {
	ProtocolID                  int    `json:"protocol_id"`
	SubjectNumber               string `json:"subject_num"`
	DateOfSign                  string `json:"date_of_sign"`
	TimeOfSign                  string `json:"time_of_sign"`
	IsSigned                    int    `json:"is_signed"`
	ReceivedAnInsurancePolicy   int    `json:"insurance_policy"`
	ReceivedAnInformaionConsent int    `json:"information_consent"`
}

type Demography struct {
	ProtocolID    int    `json:"protocol_id"`
	SubjectNumber string `json:"subject_num"`
	Sex           int    `json:"sex"`
	Race          int    `json:"race"`
	Date          string `json:"date_of_birth"`
}

type Anthropometry struct {
	ProtocolID                     int    `json:"protocol_id"`
	SubjectNumber                  string `json:"subject_num"`
	AnthropometricDataBeenMeasured int    `json:"anthropometricdatabeenmeasured"`
	ReasonIfNot                    string `json:"reasonifnot"`
	DateOfStartMeasured            string `json:"dateofstartmeasured"`
	WeightOfBody                   int    `json:"weightofbody"`
	HightOfBody                    int    `json:"heightofbody"`
	IndexWeigthOfBody              int    `json:"indexweightofbody"`
}

type InclusionCriteria struct {
	ProtocolID                                  int    `json:"protocol_id"`
	SubjectNumber                               string `json:"subject_num"`
	PresenceOfAnInformationPanel                int    `json:"presenceofaninformationpanel"`
	Aged18To55Years                             int    `json:"aged18to55years"`
	NegativeHIVTestResult                       int    `json:"negativehivtestresult"`
	BodyMassIndex                               int    `json:"bodymassindex"`
	AbsenceOfAcuteInfectiousDiseases            int    `json:"absenceofacuteinfectiousdiseases"`
	ConsentToUseEffectiveMethodsOfContraception int    `json:"consenttouseeffectivemethodsofcontraception"`
	NegativePregnancyTest                       int    `json:"negativepregnancytest"`
	NegativeDrugTest                            int    `json:"negativedrugtest"`
	NegativeAlcoholTest                         int    `json:"negativealcoholtest"`
	NoHistoryOfSeverePostVaccinationReactions   int    `json:"nohistoryofseverepostvaccinationreactions"`
	IndicatorsBloodTestsAtScreeningWithin       int    `json:"indicatorsbloodtestsatscreeningwithin"`
	NoMyocardialChanges                         int    `json:"nomyocardialchanges"`
	NegativeTestResultForCOVID                  int    `json:"negativetestresultforCOVID"`
	NoContraindicationsToVaccination            int    `json:"nocontraindicationstovaccination"`
}

type ExclusionСriteria struct {
	ProtocolID                            int    `json:"protocol_id"`
	SubjectNumber                         string `json:"subject_num"`
	LackOfSignedInformedConsent           int    `json:"lackofsignedinformedconsent"`
	SteroidTherapy                        int    `json:"steroidtherapy"`
	TherapyWithImmunosuppressiveDrugs     int    `json:"therapywithimmunosuppressivedrugs"`
	FemaleSubjectsDuringPregnancy         int    `json:"femalesubjectsduringpregnancy"`
	StrokeInLessThanOneYear               int    `json:"strokeinlessthanoneyear"`
	ChronicSystemicInfections             int    `json:"chronicsystemicinfections"`
	AggravatedAllergicHistory             int    `json:"aggravatedallergichistory"`
	PresenceOfAHistoryOfNeoplasms         int    `json:"presenceofahistoryofneoplasms"`
	HistoryOfSplenectomy                  int    `json:"historyofsplenectomy"`
	Neutropenia                           int    `json:"neutropenia"`
	SubjectsWithActiveSyphilis            int    `json:"subjectswithactivesyphilis"`
	Anorexia                              int    `json:"anorexia"`
	ExtensiveTattoos                      int    `json:"extensivetattoos"`
	TakingNarcoticAndPsychostimulantDrugs int    `json:"takingnarcoticandpsychostimulantdrugs"`
	SmokingMoretThanTenCigarettesADay     int    `json:"smokingmorethantencigarettesaday"`
	AlcoholIntake                         int    `json:"alcoholintake"`
	PlannedHospitalization                int    `json:"plannedhospitalization"`
	DonorBloodDonation                    int    `json:"donorblooddonation"`
	SubjectParticipationInAnyOtherStudy   int    `json:"subjectparticipationinanyotherstudy"`
	AnyVaccinationInTheLastMonth          int    `json:"anyvaccinationinthelastmonth"`
	InabilityToReadInRussian              int    `json:"inabilitytoreadinrussian"`
	ResearchCenterStaff                   int    `json:"researchcenterstaff"`
	AnyOtherStateOfTheSubjectOfTheStudy   int    `json:"anyotherstateofthesubjectofthestudy"`
}

type CompletionOfScreening struct {
	ProtocolID                int    `json:"protocol_id"`
	SubjectNumber             string `json:"subject_num"`
	VolunteerEligible         int    `json:"volunteereligible"`
	NoExclusionCriteria       int    `json:"noexclusioncriteria"`
	InformedOfTheRestrictions int    `json:"informedoftherestrictions"`
	VolunteerIncluded         int    `json:"volunteerincluded"`
	ReasonIfNot               string `json:"reasonifnot"`
	CommentValue              string `json:"commentvalue"`
}
