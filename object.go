package zgo_jwt

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	LoginName          string // 登录名称
	LoginType          string // 登录类型（username,email,phone）
	Role               uint   // 用户角色
	jwt.StandardClaims        // 权限相关
}
