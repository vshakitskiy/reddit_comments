package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var secret = []byte(getSecret())

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "abmop3456cdegqrstu789fvwxyz12hijkl0"
	}
	return secret
}

func Generate(id string) (string, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&JwtCustomClaims{
			ID: id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		},
	)

	token, err := t.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Validate(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		token,
		&JwtCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return secret, nil
		},
	)
}
