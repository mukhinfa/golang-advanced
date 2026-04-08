package jwt

import (
	"errors"
	"fmt"

	jwtLib "github.com/golang-jwt/jwt/v5"
)

type jwt struct {
	secret string
}

func NewJWT(secret string) *jwt {
	return &jwt{
		secret: secret,
	}
}

func (j *jwt) Create(phone string) (string, error) {
	t := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, jwtLib.MapClaims{
		"phone": phone,
	})
	s, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *jwt) Parse(tokenStr string) (string, error) {
	token, err := jwtLib.Parse(tokenStr, func(t *jwtLib.Token) (any, error) {
		if _, ok := t.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwtLib.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}
	phone, ok := claims["phone"].(string)
	if !ok {
		return "", errors.New("phone claim missing")
	}
	return phone, nil
}
