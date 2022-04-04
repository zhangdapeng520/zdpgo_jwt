package config

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId   uint64                 `json:"user_id"`   // 用户ID
	Username string                 `json:"username"`  // 用户名称
	UserType string                 `json:"user_type"` // 用户类型（username,email,phone）
	Role     uint                   `json:"role"`      // 用户角色
	Data     map[string]interface{} `json:"data"`      // 要传递的其他数据
	jwt.StandardClaims
}
