package zdpgo_jwt

import (
	"errors"
	"time"
)

// Jwt Jwt核心对象
type Jwt struct {
	Config *Config // config配置对象
}

var (
	TokenExpired     = errors.New("token已过期")
	TokenNotValidYet = errors.New("token校验失败")
	TokenMalformed   = errors.New("token格式错误")
	TokenInvalid     = errors.New("token无效")
)

// NewJwt 创建Jwt对象
func NewJwt() *Jwt {
	return NewJwtWithConfig(&Config{})
}

// NewJwtWithConfig 根据配置创建JWT对象
func NewJwtWithConfig(config *Config) *Jwt {
	j := &Jwt{}

	// 配置
	if config.Key == "" {
		config.Key = "123!@#abcABC张大鹏!@#△▲☀"
	}
	if config.Expired == 0 {
		config.Expired = 60 * 15 // 15分钟
	}
	j.Config = config

	// 返回
	return j
}

// CreateToken 创建一个token
func (j *Jwt) CreateToken(claims ClaimsData) (string, error) {
	// 过期时间
	if claims.ExpiresAt == 0 {
		claims.ExpiresAt = time.Now().Add(time.Duration(j.Config.Expired) * time.Second).Unix()
	}

	// 创建token
	token := NewWithClaims(SigningMethodHS256, claims)

	// 生成token字符串
	key := []byte(j.Config.Key)
	tokenStr, err := token.SignedString(key)

	// 校验是否成功
	if err != nil {
		return "", err
	}

	// 返回token字符串
	return tokenStr, nil
}

// ParseToken 解析 token
func (j *Jwt) ParseToken(tokenString string) (*ClaimsData, error) {
	// 解析参数
	token, err := ParseWithClaims(tokenString, &ClaimsData{}, func(token *Token) (i interface{}, e error) {
		return []byte(j.Config.Key), nil
	})

	// 处理失败
	if err != nil {
		if ve, ok := err.(*ValidationError); ok {
			if ve.Errors&ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// 成功获取到token
	if token != nil {
		if claims, ok := token.Claims.(*ClaimsData); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *Jwt) RefreshToken(tokenString string) (string, error) {
	// 实例化时间函数
	TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	// 解析token
	token, err := ParseWithClaims(tokenString, &ClaimsData{}, func(token *Token) (interface{}, error) {
		return []byte(j.Config.Key), nil
	})
	if err != nil {
		return "", err
	}

	// 得到token内容
	if claims, ok := token.Claims.(*ClaimsData); ok && token.Valid {
		TimeFunc = time.Now

		// 刷新token的时间，在当前时间上追加默认的过期时间
		claims.ExpiresAt = time.Now().Add(time.Duration(j.Config.Expired) * time.Second).Unix()

		// 返回新创建的token
		return j.CreateToken(*claims)
	}

	// 解析失败，返回错误
	return "", TokenInvalid
}
