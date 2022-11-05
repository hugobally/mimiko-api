package jwt

import (
	"encoding/base64"
	jwtlib "github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	UserID uint `json:"userId"`
	jwtlib.StandardClaims
}

type Util struct {
	JwtKey []byte
}

func NewJwtHandler(jwtKeyStr string) *Util {
	key, err := base64.StdEncoding.DecodeString(jwtKeyStr)
	if err != nil {
		panic("auth.NewMiddleware: error while decoding JwtKey")
	}
	return &Util{
		JwtKey: key,
	}
}

func (j *Util) CreateAccessToken(user uint, exp time.Time) (*string, error) {
	claims := &Claims{
		UserID: user,
		StandardClaims: jwtlib.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.JwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (j *Util) ValidateAccessToken(tknStr string) (uint, error) {
	claims := &Claims{}

	tkn, err := jwtlib.ParseWithClaims(tknStr, claims, func(token *jwtlib.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		return 0, err
	}

	if !tkn.Valid {
		return 0, jwtlib.ErrSignatureInvalid
	}

	return claims.UserID, nil
}
