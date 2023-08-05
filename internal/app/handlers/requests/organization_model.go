package requests

type DeleteOrganization struct {
	ID int `json:"organization_id"`
}

type AddNewOrganization struct {
	Name string `json:"name"`
}

type UpdateOrganization struct {
	OrganizationId int    `json:"organization_id"`
	Name           string `json:"name"`
}

type GetOrganizationName struct {
	ID int `json:"organization_id"`
}
