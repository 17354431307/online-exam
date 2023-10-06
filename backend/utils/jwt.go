package utils

import (
	"backend/global"
	"backend/model/system/request"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenMalformed   = errors.New("That's not even a token")
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.OE_CONFIG.Jwt.SigningKey),
	}
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if errors.As(err, &jwt.ErrTokenMalformed) {
			return nil, TokenMalformed
		} else if errors.As(err, &jwt.ErrTokenExpired) {
			return nil, TokenExpired
		} else if errors.As(err, &jwt.ErrTokenNotValidYet) {
			return nil, TokenNotValidYet
		}
		return nil, TokenInvalid
	}

	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, err
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
