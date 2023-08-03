package model

type Center struct {
	CenterId       int    `json:"center_id" gorm:"primaryKey"`
	OrganizationID int    `json:"organization_id"`
	Name           string `json:"center_name"`
}
