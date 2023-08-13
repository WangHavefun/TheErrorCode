package utils

import (
	"TheErrorCode/constant"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

// 创建JWT
func createJWT(id int64, username string) (string, error) {
	claims := &Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(constant.JWTCONFIG.JwtTtl)).Unix(), // 有效期 1 小时
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(constant.JWTCONFIG.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 验证 JWT
func validateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return constant.JWTCONFIG.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
