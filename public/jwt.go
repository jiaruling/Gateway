package public

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/jiaruling/Gateway/global"
)

func JwtDecode(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JwtSignKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("token is not jwt.StandardClaims")
	}
}

func JwtEncode(claims jwt.StandardClaims) (string, error) {
	mySigningKey := []byte(global.JwtSignKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
