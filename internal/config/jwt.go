package config

import(
	"time"
	"github.com/golang-jwt/jwt/v5"
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