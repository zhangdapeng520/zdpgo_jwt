package zgo_jwt

import (
	"errors"
	"github.com/zhangdapeng520/zdpgo_zap"

	"time"

	"github.com/dgrijalva/jwt-go"
)

// Jwt Jwt核心对象
type Jwt struct {
	log    *zdpgo_zap.Zap // zap日志对象
	config *JwtConfig     // config配置对象
}

var (
	TokenExpired     = errors.New("Token已过期")
	TokenNotValidYet = errors.New("Token校验失败")
	TokenMalformed   = errors.New("Token格式错误")
	TokenInvalid     = errors.New("Token无效")
)

// New 创建Jwt对象
func New(config JwtConfig) *Jwt {
	j := Jwt{}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_jwt.log"
	}
	j.log = zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:       config.Debug,
		LogFilePath: config.LogFilePath,
	})

	// 配置
	if config.Key == "" {
		config.Key = "zdpgo_jwt"
	}
	if config.Expired == 0 {
		config.Expired = 60 * 15 // 默认15分钟
	}
	j.config = &config

	return &j
}

// CreateToken 创建一个token
func (j *Jwt) CreateToken(claims Claims) (string, error) {
	// 过期时间
	if claims.ExpiresAt == 0 {
		claims.ExpiresAt = time.Now().Add(time.Duration(j.config.Expired) * time.Second).Unix()
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成token字符串
	key := []byte(j.config.Key)
	tokenStr, err := token.SignedString(key)

	// 校验是否成功
	if err != nil {
		j.log.Error("生成token字符串失败", "error", err.Error())
		return "", err
	}

	// 返回token字符串
	return tokenStr, nil
}

// ParseToken 解析 token
func (j *Jwt) ParseToken(tokenString string) (*Claims, error) {
	// 解析参数
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(j.config.Key), nil
	})

	// 处理失败
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// 成功获取到token
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *Jwt) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Key), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(j.config.Expired) * time.Second).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
