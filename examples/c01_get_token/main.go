package main

import (
	"errors"
	"fmt"
	jwt "github.com/zhangdapeng520/zdpgo_jwt"
	"time"
)

var secret = "hello"

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetToken 生成Token
func GetToken(username string) (string, error) {
	c := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Second * 1800)},
			Issuer:    "张大鹏",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString([]byte(secret))
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的Token")
}

func main() {
	// 生成token
	token, err := GetToken("zhangdapeng")
	if err != nil {
		fmt.Println("生成Token失败：", err)
		return
	}
	fmt.Println("生成token成功：", token)

	// 解析token
	claims, err := ParseToken(token)
	if err != nil {
		fmt.Println("解析Token失败：", err)
		return
	}
	fmt.Println("解析Token成功：", claims)
}
