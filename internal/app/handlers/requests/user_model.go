package requests

type Registration struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
	CenterID int    `json:"center_id"`
	Password string `json:"password,omitempty"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckToken struct {
	Cookie string `json:"token"`
}

type Update struct {
	ID       int    `json:"id"`
	Role     int    `json:"role"`
	CenterID int    `json:"center_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Paswword string `json:"password"`
}

type Delete struct {
	ID int `json:"id"`
}
