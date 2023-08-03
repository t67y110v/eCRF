package model

type Protocol struct {
	Id             int    `json:"protocol_id"  gorm:"primaryKey"`
	Number         int    `json:"number" `
	OrganizationID int    `json:"orgranization_id"`
	Name           string `json:"name"`
	Status         int    `json:"status"`
	CenterId       int    `json:"center_id"  gorm:"primaryKey"`
}
