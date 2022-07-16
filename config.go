package zdpgo_jwt

// Config jwt配置
type Config struct {
	Key     string `yaml:"key" json:"key"`         // jwt加密的key
	Expired uint16 `yaml:"expired" json:"expired"` // token过期时间（秒），默认15分钟
}
