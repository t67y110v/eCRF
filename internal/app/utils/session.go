package utils

// import (
// 	"strconv"

// 	"github.com/golang-jwt/jwt"
// 	model "github.com/t67y110v/web/internal/app/model/user"
// 	"github.com/t67y110v/web/internal/app/store"
// )

// func CheckToken(token string) (*model.User, error ){

// 	tokenString := token

// 		claims := jwt.MapClaims{}

// 		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("secret"), nil
// 		})

// 		if claims["id"] == nil {
// 			return nil, err
// 		}

// 		id := float64(claims["id"].(float64))
//         u, err :=   store.PostgresStore.Repository.FindByID(strconv.Itoa(int(id)))

// 		if err != nil {
// 			return err
// 		}

// 		return c.JSON(u)
// }
