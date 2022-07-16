package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_jwt"
	"time"
)

func main() {
	j := zdpgo_jwt.NewJwt()

	// 1、创建token
	token, err := j.CreateToken(zdpgo_jwt.ClaimsData{
		UserId:   "1",
		Username: "zhangdapeng",
		UserType: "username",
		Role:     1,
		Data: map[string]interface{}{
			"a": 111,
			"b": 2.222,
			"c": true,
		},
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 过期时间，必须设置
	})
	if err != nil {
		fmt.Println("创建token失败：", err)
		return
	}
	fmt.Println(token)
	fmt.Println("=====================")

	// 2、验证token
	claims, err := j.ParseToken(token)
	if err != nil {
		fmt.Println("解析token失败：", err)
		return
	}
	fmt.Println("user_id：", claims.UserId)
	fmt.Println("username：", claims.Username)
	fmt.Println("user_type：", claims.UserType)
	fmt.Println("role：", claims.Role)
	fmt.Println("data：", claims.Data)
	fmt.Println("expires_at：", claims.ExpiresAt)
	fmt.Println("=====================")

	// 3、刷新token
	tokenNew, err := j.RefreshToken(token)
	if err != nil {
		fmt.Println("更新token失败：", tokenNew)
		return
	}

	// 4、校验旧的token和新的token
	claims1, err := j.ParseToken(token)
	if err != nil {
		fmt.Println("解析token失败：", err)
		return
	}
	fmt.Println("user_id：", claims1.UserId)
	fmt.Println("username：", claims1.Username)
	fmt.Println("user_type：", claims1.UserType)
	fmt.Println("role：", claims1.Role)
	fmt.Println("data：", claims1.Data)
	fmt.Println("expires_at：", claims1.ExpiresAt)
	fmt.Println("=====================")

	claims2, err := j.ParseToken(token)
	if err != nil {
		fmt.Println("解析token失败：", err)
		return
	}
	fmt.Println("user_id：", claims2.UserId)
	fmt.Println("username：", claims2.Username)
	fmt.Println("user_type：", claims2.UserType)
	fmt.Println("role：", claims2.Role)
	fmt.Println("data：", claims2.Data)
	fmt.Println("expires_at：", claims2.ExpiresAt)
	fmt.Println("=====================")
}
