package jwt

import jwtLib "github.com/golang-jwt/jwt/v5"

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
