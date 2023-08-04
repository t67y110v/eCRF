package requests

type AdverseEvents struct {
	ProtocolID                   int    `json:"protocol_id"`
	SubjectNumber                string `json:"subject_num"`
	AdverseEventsRegistered      int    `json:"adverseeventsregistered"`
	DescriptionOfTheAdverseEvent string `json:"descriptionoftheadverseevent"`
	DateOfStartAE                string `json:"dateofstartae"`
	DateOfEndAE                  string `json:"dateofendae"`
	IsContinuesEnd               int    `json:"iscontinius"`
	Severity                     int    `json:"severity"`
	RecurringPhenomenon          int    `json:"recurringphenomenon"`
	AssociationWithTheDrugUsed   int    `json:"associationwiththedrugused"`
	Foresight                    int    `json:"foresight"`
	ConnectionBetweenAEAndDU     int    `json:"connectionbetweenaeanddu"`
	RenewalAfterUse              int    `json:"renewalafteruse"`
	LocalReaction                int    `json:"localreaction"`
	SubjectDropout               int    `json:"subjectdropout"`
	MeasuresTaken                int    `json:"measurestaken"`
	MeasuresTakenOnUD            int    `json:"measurestakenonud"`
	Exodus                       int    `json:"exodus"`
	IsItSerious                  int    `json:"isitserious"`
	SeverityCriterion            int    `json:"severitycriterion"`
	TestImpact                   int    `json:"testimpact"`
	DoseEffect                   int    `json:"doseeffect"`
	ImpactOnHospitalStay         int    `json:"impactonhospitalstay"`
	RelationshipWithMedication   int    `json:"relationshipwithmedication"`
	Expectancy                   int    `json:"expectancy"`
}
