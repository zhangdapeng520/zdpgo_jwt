package zgo_jwt

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId             uint64 // 用户ID
	Username           string // 用户名称
	UserType           string // 用户类型（username,email,phone）
	Role               uint   // 用户角色
	jwt.StandardClaims        // 权限相关
}
