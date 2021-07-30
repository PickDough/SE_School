package services

import (
	"github.com/Pick-Down/BTC_API/models"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type AuthService struct {
}

//GenerateToken Returns signed JWT.
func (service *AuthService) GenerateToken(userEmail string) (models.Token, error) {
	var err error

	//Specifying claims
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = userEmail
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return models.Token{}, err
	}
	return models.Token{Token: token}, nil
}
