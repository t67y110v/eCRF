package model

type Subject struct {
	Id         int    `json:"subject_id"  gorm:"primaryKey"`
	Number     string `json:"number"`
	CenterID   int    `json:"center_id"  gorm:"primaryKey"`
	ProtocolId int    `json:"protocol_id"  gorm:"primaryKey"`
	Initials   string `json:"initials"`
}

type DemographySubject struct {
	Id        int    `json:"demography_id" gorm:"primaryKey"`
	SubjectId int    `json:"subject_id" gorm:"primaryKey"`
	Sex       int    `json:"sex"`
	Race      int    `json:"race"`
	BirthDate string `json:"date"`
}
