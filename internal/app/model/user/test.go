package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Pas() {
	b, _ := bcrypt.GenerateFromPassword([]byte("string"), bcrypt.MinCost)

	fmt.Println(string(b))

}
