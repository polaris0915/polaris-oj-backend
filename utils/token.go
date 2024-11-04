package utils

import (
	"fmt"
	"polaris-oj-backend/constant"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 密钥
var jwtKey = []byte(constant.Key)

type Claims struct {
	Identity    string `json:"identity"`
	UserAccount string `json:"userAccount"`
	UserRole    string `json:"userRole"`
	jwt.RegisteredClaims
}

// 生成token
func GetToken(ID string, UserAccount string, UserRole string) (string, error) {
	// 有效期，时间有效期定义在package constant中
	expirationTime := time.Now().Add(time.Duration(constant.ValidTime) * time.Minute)
	// fmt.Println("expiration: ", expirationTime)

	claims := &Claims{
		Identity:    ID,
		UserAccount: UserAccount,
		UserRole:    UserRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 校验token
func ValidateToken(tokenString string) (*Claims, error) {
	userClaims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("ValidateToken Failed")
	}
	return userClaims, nil
}
