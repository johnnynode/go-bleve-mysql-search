package utils

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	SecretCode = "d8346ea2-6017-43ed-ad68-19c0f971738b"
)

// jwt 解码功能
func JwtDecode(token string) (claims jwt.Claims, err error) {
	mySigningKey := []byte(SecretCode) // 设置签名的key
	// 通过字符串token和secret解析出完整的token
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err == nil {
		claims = t.Claims.(jwt.MapClaims)
	} else {
		claims = nil
	}
	return claims, err
}