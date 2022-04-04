package zgo_jwt

import "github.com/zhangdapeng520/zdpgo_jwt/core/config"

type Generate interface {
	// CreateToken 创建TOKEN
	CreateToken(claims config.Claims) (string, error)
}

type Parse interface {
	// ParseToken 解析TOKEN
	ParseToken(tokenString string) (*config.Claims, error)
}

type Refresh interface {
	// RefreshToken 刷新TOKEN
	RefreshToken(tokenString string) (string, error)
}

// Token TOKEN接口
type Token interface {
	Generate
	Parse
	Refresh
}
