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
	ConnectionBetweenAEAndDU     int    `json:"ConnectionBetweenAEAndDU"`
	RenewalAfterUse              int    `json:"renewalafteruse"`
	LocalReaction                int    `json:"LocalReaction"`
	SubjectDropout               int    `json:"SubjectDropout"`
	MeasuresTaken                int    `json:"MeasuresTaken"`
	MeasuresTakenOnUD            int    `json:"MeasuresTakenOnUD"`
	Exodus                       int    `json:"Exodus"`
	IsItSerious                  int    `json:"IsItSerious"`
	SeverityCriterion            int    `json:"severitycriterion"`
	TestImpact                   int    `json:"testimpact"`
	DoseEffect                   int    `json:"DoseEffect"`
	ImpactOnHospitalStay         int    `json:"impactonhospitalstay"`
	RelationshipWithMedication   int    `json:"relationshipwithmedication"`
	Expectancy                   int    `json:"expectancy"`
}
