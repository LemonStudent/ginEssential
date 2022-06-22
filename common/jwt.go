package common

import (
	"github.com/dgrijalva/jwt-go"
	"orangezoom.cn/ginessential/model"
	"time"
)

var jwtKey = []byte("zqqtoAwGYOqrFYI")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func CreateToken(user model.User) (string, error) {
	nowTime := time.Now()
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenValue string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
