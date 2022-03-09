package zgo_jwt

// JwtConfig jwt配置
type JwtConfig struct {
	Key         string `mapstructure:"key" json:"key"`                     // jwt加密的key
	Expired     uint16 `mapstructure:"expired" json:"expired"`             // token过期时间，秒，默认15分钟
	LogFilePath string `mapstructure:"log_file_path" json:"log_file_path"` // 日志存放路径
	Debug       bool   `mapstructure:"debug" json:"debug"`                 // 是否为开发模式
}
