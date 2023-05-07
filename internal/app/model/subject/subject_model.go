package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Subject struct {
// 	Id         int    `json:"subject_id"  gorm:"primaryKey"`
// 	Number     string `json:"number"`
// 	CenterID   int    `json:"center_id"  gorm:"primaryKey"`
// 	ProtocolId int    `json:"protocol_id"  gorm:"primaryKey"`
// 	Initials   string `json:"initials"`
// }

// type DemographySubject struct {
// 	Id        int    `json:"demography_id" gorm:"primaryKey"`
// 	SubjectId int    `json:"subject_id" gorm:"primaryKey"`
// 	Sex       int    `json:"sex"`
// 	Race      int    `json:"race"`
// 	BirthDate string `json:"date"`
// }

type Subject struct {
	ID                primitive.ObjectID `bson:"_id"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
	Number            string             `bson:"number"`
	CenterID          int                `bson:"center_id" `
	ProtocolId        int                `bson:"protocol_id"`
	Initials          string             `bson:"initials"`
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
	DemographySubject struct {
		Id        int    `bson:"demography_id"`
		SubjectId int    `bson:"subject_id"`
		Sex       int    `bson:"sex"`
		Race      int    `bson:"race"`
		BirthDate string `bson:"date"`
	}
	CriteriaForInclusionInTheStudy bool `bson:"criteria_for_inclusion_in_the_study"`

	ExclusionСriteria bool `bson:"exclusion_criteria"`
}
