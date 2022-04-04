package config

func GetDefaultConfig(cfg Config) *Config {
	if cfg.Key == "" {
		cfg.Key = "123!@#abcABC张大鹏!@#△▲☀"
	}
	if cfg.Expired == 0 {
		cfg.Expired = 60 * 10 // 10分钟
	}
	if cfg.LogFilePath == "" {
		cfg.LogFilePath = "logs/zdpgo/zdpgo_jwt.log"
	}
	if !cfg.Debug {
		cfg.Debug = true
	}
	return &cfg
}
