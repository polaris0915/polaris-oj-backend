package utils

import (
	"errors"
	"polaris-oj-backend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// // 密钥
// var jwtKey = []byte(config.Jwt.Key)

type Claims struct {
	Identity    string `json:"identity"`
	UserAccount string `json:"userAccount"`
	UserRole    string `json:"userRole"`
	jwt.RegisteredClaims
}

// 生成token
func GetToken(ID string, UserAccount string, UserRole string) (string, error) {
	// 有效期，时间有效期定义在package constant中
	expirationTime := time.Now().Add(time.Duration(config.Jwt.ValidTime) * time.Minute)
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
	var jwtKey = []byte(config.Jwt.Key)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 校验token
func ValidateToken(tokenString string) (*Claims, error) {
	userClaims := &Claims{}
	var jwtKey = []byte(config.Jwt.Key)

	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("ValidateToken Failed")
	}
	return userClaims, nil
}
