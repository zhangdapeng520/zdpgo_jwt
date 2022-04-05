package config

import "github.com/zhangdapeng520/zdpgo_jwt/libs/jwtgo"

type Claims struct {
	UserId   uint64 `json:"user_id"`   // 用户ID
	Username string `json:"username"`  // 用户名称
	UserType string `json:"user_type"` // 用户类型（username,email,phone）
	Role     uint   `json:"role"`      // 用户角色
	JsonData string `json:"json_data"` // 要传递的其他数据，JSON序列化的字符串
	jwtgo.StandardClaims
}
