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
	DemographySubject struct {
		Id        int    `bson:"demography_id"`
		SubjectId int    `bson:"subject_id"`
		Sex       int    `bson:"sex"`
		Race      int    `bson:"race"`
		BirthDate string `bson:"date"`
	}
}
