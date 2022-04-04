package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_jwt"
	"github.com/zhangdapeng520/zdpgo_jwt/core/config"
	"github.com/zhangdapeng520/zdpgo_jwt/libs/jwtgo"
	"time"
)

func main() {
	j := zdpgo_jwt.New(config.Config{})

	// 1、创建token
	data := map[string]interface{}{
		"a": 111,
		"b": 2.222,
		"c": true,
	}

	maxAge := 60 * 60 * 24
	token, err := j.CreateToken(config.Claims{
		UserId:   1,
		Username: "zhangdapeng",
		UserType: "username",
		Role:     1,
		Data:     data,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
		},
	})
	if err != nil {
		fmt.Println("创建token失败：", err)
		return
	}
	fmt.Println("得到token内容：", token)

	// 2、验证token
	claims, err := j.ParseToken(token)
	if err != nil {
		fmt.Println("解析token失败：", err)
		return
	}
	fmt.Println("得到token内容：", claims)

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
	fmt.Println("得到旧token内容：", claims1)

	claims2, err := j.ParseToken(token)
	if err != nil {
		fmt.Println("解析token失败：", err)
		return
	}
	fmt.Println("得到刷新token内容：", claims2)
}
