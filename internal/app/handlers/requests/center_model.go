package requests

type DeleteCenter struct {
	ID int `json:"center_id"`
}

type AddCenter struct {
	Name string `json:"name"`
}

type UpdateCenter struct {
	CenterID int    `json:"center_id"`
	Name     string `json:"name"`
}

type GetCenterName struct {
	ID int `json:"center_id"`
}
