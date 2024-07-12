/**
 * @Author: Nan
 * @Date: 2023/3/14 22:42
 */

package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTTokenGen struct {
	privateKey []byte
	issuer     string
}

func NewJWTTokenGen() *JWTTokenGen {
	return &JWTTokenGen{issuer: "master-nan", privateKey: []byte("123455")}
}

func (t *JWTTokenGen) GenerateToken(id string) (token string, err error) {
	nowSec := time.Now()
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		Issuer:    t.issuer,
		IssuedAt:  jwt.NewNumericDate(nowSec),
		ExpiresAt: jwt.NewNumericDate(nowSec.Add(3 * time.Hour * time.Duration(1))),
		NotBefore: jwt.NewNumericDate(nowSec),
		Subject:   id,
	})
	return tkn.SignedString(t.privateKey)
}

func (t *JWTTokenGen) ParseToken(token string) (id string, err error) {
	res, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return t.privateKey, nil
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", errors.New("token 错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return "", errors.New("token已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", errors.New("token未激活")
			} else {
				return "", errors.New("token 错误")
			}
		}
	}
	if !res.Valid {
		return "", fmt.Errorf("token not valid")
	}
	if claims, ok := res.Claims.(jwt.MapClaims); ok {
		return claims["sub"].(string), nil
	} else {
		return "", errors.New("token 错误")

	}
}
