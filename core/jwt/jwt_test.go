package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
	"zgo_jwt/core/config"
)

func prepareJwt() *Jwt {
	j := New(config.Config{})
	return j
}

// 测试新建jwt
func TestJWT_New(t *testing.T) {
	j := prepareJwt()
	if j == nil {
		t.Error(j)
	}
	t.Log(j)
}

// 测试新建token
func TestJWT_CreateToken(t *testing.T) {
	j := prepareJwt()
	maxAge := 60 * 60 * 24
	token, err := j.CreateToken(config.Claims{
		UserId:   1,
		Username: "zhangdapeng",
		UserType: "username",
		Role:     1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)

}

// 测试解析token
func TestJWT_ParseToken(t *testing.T) {
	j := prepareJwt()
	maxAge := 60 * 60 * 24
	token, err := j.CreateToken(config.Claims{
		UserId:   1,
		Username: "zhangdapeng",
		UserType: "username",
		Role:     1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	claims, err := j.ParseToken(token)
	t.Log(claims.UserId, claims.Username, claims.UserType, claims.Role)
	if err != nil {
		t.Error(err)
	}
	t.Log(claims)
}

// 测试刷新token
func TestJWT_RefreshToken(t *testing.T) {
	j := prepareJwt()
	maxAge := 60 * 60 * 24
	token, err := j.CreateToken(config.Claims{
		UserId:   1,
		Username: "zhangdapeng",
		UserType: "username",
		Role:     1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	newToken, err := j.RefreshToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(newToken)
}
