package responses

type Registration struct {
	Id             int    `json:"user_id" `
	Email          string `json:"email"`
	Name           string `json:"name"`
	Role           int    `json:"role"`
	CenterID       int    `json:"center_id" `
	OrganizationID int    `json:"organization_id"`
}

type Login struct {
	JWT string `json:"JWT"`
}
