package responses

type Organization struct {
	ID   int    `json:"organization_id"  gorm:"primaryKey"`
	Name string `json:"name"`
}
