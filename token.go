package utils

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GetToken 获取token
func GetToken(id string) string {
	// 设置token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	timeUnix := time.Now().Unix()
	str := strconv.FormatInt(timeUnix, 10)
	tokenString, _ := token.SignedString([]byte(str))
	return tokenString
}
