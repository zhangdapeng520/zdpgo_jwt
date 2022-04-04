package jwt

import (
	"errors"
	"github.com/zhangdapeng520/zdpgo_jwt/core/config"
	"github.com/zhangdapeng520/zdpgo_jwt/libs/jwtgo"

	"time"
)

// Jwt Jwt核心对象
type Jwt struct {
	config *config.Config // config配置对象
}

var (
	TokenExpired     = errors.New("Token已过期")
	TokenNotValidYet = errors.New("Token校验失败")
	TokenMalformed   = errors.New("Token格式错误")
	TokenInvalid     = errors.New("Token无效")
)

// New 创建Jwt对象
func New(cfg config.Config) *Jwt {
	j := Jwt{}

	// 获取默认配置
	cfgNew := config.GetDefaultConfig(cfg)
	j.config = cfgNew

	return &j
}

// CreateToken 创建一个token
func (j *Jwt) CreateToken(claims config.Claims) (string, error) {
	// 过期时间
	if claims.ExpiresAt == 0 {
		claims.ExpiresAt = time.Now().Add(time.Duration(j.config.Expired) * time.Second).Unix()
	}

	// 创建token
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	// 生成token字符串
	key := []byte(j.config.Key)
	tokenStr, err := token.SignedString(key)

	// 校验是否成功
	if err != nil {
		return "", err
	}

	// 返回token字符串
	return tokenStr, nil
}

// ParseToken 解析 token
func (j *Jwt) ParseToken(tokenString string) (*config.Claims, error) {
	// 解析参数
	token, err := jwtgo.ParseWithClaims(tokenString, &config.Claims{}, func(token *jwtgo.Token) (i interface{}, e error) {
		return []byte(j.config.Key), nil
	})

	// 处理失败
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// 成功获取到token
	if token != nil {
		if claims, ok := token.Claims.(*config.Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *Jwt) RefreshToken(tokenString string) (string, error) {
	jwtgo.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwtgo.ParseWithClaims(tokenString, &config.Claims{}, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(j.config.Key), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*config.Claims); ok && token.Valid {
		jwtgo.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(j.config.Expired) * time.Second).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
