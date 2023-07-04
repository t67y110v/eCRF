package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id                int    `json:"user_id" gorm:"primaryKey"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Role              int    `json:"role"`
	CenterID          int    `json:"center_id" gorm:"primaryKey"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}
func (u *User) ClearPswd() {
	u.Password = ""
	u.EncryptedPassword = ""
}

func (u *User) ComparePassword(password string) bool {

	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {

	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
