package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetSubject struct {
	ID         primitive.ObjectID `json:"_id"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Number     string             `json:"number"`
	CenterID   int                `json:"center_id" `
	ProtocolId int                `json:"protocol_id"`
	Initials   string             `json:"initials"`
}

type GetAdverseEventsAr struct {
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
