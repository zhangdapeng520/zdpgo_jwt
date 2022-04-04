package zgo_jwt

import (
	"zgo_jwt/core/config"
	"zgo_jwt/core/jwt"
)

// Jwt Jwt核心对象
type Jwt struct {
	// 方法区
	CreateToken  func(claims config.Claims) (string, error)
	ParseToken   func(tokenString string) (*config.Claims, error)
	RefreshToken func(tokenString string) (string, error)
}

// New 创建Jwt对象
func New(cfg config.Config) *Jwt {
	j := Jwt{}

	// 创建token对象
	token := jwt.New(cfg)

	// 初始化方法
	j.CreateToken = token.CreateToken
	j.ParseToken = token.ParseToken
	j.RefreshToken = token.RefreshToken

	return &j
}
