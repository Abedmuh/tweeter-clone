package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateToken(userid string) (string, error) {
	secretKey := viper.GetString("JWT_SECRET_KEY")
	timeExp := viper.GetDuration("JWT_TIME_EXP")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userid,
		"exp":  time.Now().Add(time.Duration(timeExp) * time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}