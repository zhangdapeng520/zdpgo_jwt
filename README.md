# zdpgo_jwt
使用Golang创建、刷新和校验JWT Token

项目地址：https://github.com/zhangdapeng520/zdpgo_jwt

## 功能清单
- 创建token
- 校验token
- 刷新token

## 版本历史
- 版本0.1.0 2022年3月9日 基本功能
- 版本0.1.1 2022年4月4日 代码优化
- 版本0.1.2 2022年4月5日 支持传递JSON字符串

## 使用示例

### 创建、校验和刷新token
```go
package main

import (
	"encoding/json"
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
	jsonData, err := json.Marshal(data)

	maxAge := 60 * 60 * 24
	token, err := j.CreateToken(config.Claims{
		UserId:   1,
		Username: "zhangdapeng",
		UserType: "username",
		Role:     1,
		JsonData: string(jsonData),
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
```