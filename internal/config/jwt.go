package config

import(
	"time"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
)

var SecretKey = []byte("smart-dispatch-secret")

func GenerateJWT(userID uint, email string) (string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
		"email": email,
		"exp": time.Now().Add(24* time.Hour).Unix(),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ValidateJWT(tokenString string) (*jwt.Token, error){
	return jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error){

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method")
		}

		return SecretKey, nil
	})
}