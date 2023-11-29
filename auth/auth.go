package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"
	"u-sushi/models"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/guregu/null.v4"
)

var (
	Key *rsa.PrivateKey
)

func LoadKey() error {
	wrapErr := func(err error) error {
		return fmt.Errorf("load key: %v", err)
	}
	keyFile, err := os.ReadFile("secrets/key.pem")
	if err != nil {
		return wrapErr(err)
	}
	Key, err = jwt.ParseRSAPrivateKeyFromPEM(keyFile)
	if err != nil {
		return wrapErr(err)
	}
	return nil
}

func CreateJWT(userType models.UserType, table null.Int) (string, error) {
	wrapErr := func(err error) error {
		return fmt.Errorf("create jwt: %v", err)
	}
	now := time.Now()
	claims := jwt.MapClaims{
		"iss":      "u-sushi",
		"iat":      now.Unix(),
		"exp":      now.Add(time.Hour).Unix(),
		"userType": userType,
	}
	if userType == models.UserClient {
		claims["table"] = table
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := t.SignedString(Key)
	if err != nil {
		return "", wrapErr(err)
	}
	return token, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	wrapErr := func(err error) error {
		return fmt.Errorf("parsing token: %v", err)
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return &Key.PublicKey, nil
	})
	if err != nil {
		return nil, wrapErr(err)
	}
	return token, nil
}
