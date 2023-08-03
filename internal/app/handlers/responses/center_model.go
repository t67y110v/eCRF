package responses

type Center struct {
	CenterId       int    `json:"center_id" gorm:"primaryKey"`
	Name           string `json:"center_name"`
	OrganizationID int    `json:"organization_id"`
}
