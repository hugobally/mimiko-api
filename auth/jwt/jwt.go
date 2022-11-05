package jwt

import (
	"encoding/base64"
	jwtlib "github.com/golang-jwt/jwt/v4"
	"github.com/hugobally/mimiko/backend/prisma"
	"time"
)

type Claims struct {
	UserId string `json:"userId"`
	jwtlib.StandardClaims
}

type Util struct {
	jwtKey []byte
}

func NewJwtHandler(jwtKeyStr string) *Util {
	key, err := base64.StdEncoding.DecodeString(jwtKeyStr)
	if err != nil {
		panic("auhth.NewMiddleware: error while decoding jwtKey")
	}
	return &Util{
		jwtKey: key,
	}
}

func (j *Util) CreateAccessToken(user string, exp time.Time) (*string, error) {
	claims := &Claims{
		UserId: user,
		StandardClaims: jwtlib.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.jwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (j *Util) ValidateAccessToken(tknStr string) (*prisma.User, error) {
	claims := &Claims{}

	tkn, err := jwtlib.ParseWithClaims(tknStr, claims, func(token *jwtlib.Token) (interface{}, error) {
		return j.jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, jwtlib.ErrSignatureInvalid
	}

	return &prisma.User{
		ID: claims.UserId,
	}, nil
}
