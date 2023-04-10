package requests

type Registration struct {
	Id                int    `json:"user_id" gorm:"primaryKey"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Role              int    `json:"role"`
	CenterID          int    `json:"center_id" gorm:"primaryKey"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckToken struct {
	Cookie string `json:"token"`
}
