package requests

type DeleteCenter struct {
	ID int `json:"center_id"`
}

type AddCenter struct {
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
}

type UpdateCenter struct {
	CenterID       int    `json:"center_id"`
	OrganizationID int    `json:"organization_id"`
	Name           string `json:"name"`
}

type GetCenterName struct {
	ID int `json:"center_id"`
}

// TODO :  при создании указывать организацию, после этого написать метод, который будет добавлять текущий центр к организации
