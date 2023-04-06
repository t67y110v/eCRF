package model

type Center struct {
	CenterId int    `json:"center_id" gorm:"primaryKey"`
	Name     string `json:"center_name"`
}
