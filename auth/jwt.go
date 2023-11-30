package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"u-sushi/models"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/guregu/null.v4"
)

var (
	ErrAuthorizationRequired = errors.New("bearer authorization required")
)

var (
	Key *rsa.PrivateKey
)

type UserClaims struct {
	jwt.MapClaims
	UserType  models.UserType
	SessionID null.Int
}

func LoadKey() error {
	wrapErr := func(err error) error {
		return fmt.Errorf("load key: %v", err)
	}
	keyFilePath := os.Getenv("KEY_FILE")
	keyFile, err := os.ReadFile(keyFilePath)
	if err != nil {
		return wrapErr(err)
	}
	Key, err = jwt.ParseRSAPrivateKeyFromPEM(keyFile)
	if err != nil {
		return wrapErr(err)
	}
	return nil
}

func CreateJWT(userType models.UserType, sessionID null.Int) (string, error) {
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
	if sessionID.Valid {
		claims["sessionID"] = sessionID.Int64
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := t.SignedString(Key)
	if err != nil {
		return "", wrapErr(err)
	}
	return token, nil
}

func ParseJWT(tokenString string) (*jwt.Token, error) {
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

func ParseClaims(token *jwt.Token) UserClaims {
	claims := token.Claims.(jwt.MapClaims)
	userClaims := UserClaims{
		MapClaims: claims,
		UserType:  models.UserType(claims["userType"].(string)),
		SessionID: null.Int{},
	}
	if sessionID, ok := claims["sessionID"]; ok {
		userClaims.SessionID = null.IntFrom(int64(sessionID.(float64)))
	}
	return userClaims
}

const TokenPrefix string = "Bearer "

func ExtractJWT(r *http.Request) (*jwt.Token, error) {
	tokenStr := r.Header.Get("Authorization")
	if !strings.HasPrefix(tokenStr, TokenPrefix) {
		return nil, ErrAuthorizationRequired
	}
	tokenStr = tokenStr[len(TokenPrefix):]
	return ParseJWT(tokenStr)
}

func ExtractClaims(r *http.Request) (*UserClaims, error) {
	token, err := ExtractJWT(r)
	if err != nil {
		return nil, err
	}
	claims := ParseClaims(token)
	return &claims, nil
}
