package config

// GetDefaultCaptchaConfig 获取默认的配置
func GetDefaultCaptchaConfig(config CaptchaConfig) CaptchaConfig {
	if config.Width == 0 {
		config.Width = 240
	}
	if config.Height == 0 {
		config.Height = 80
	}
	if config.Length == 0 {
		config.Length = 6
	}
	if config.MaxSkew == 0 {
		config.MaxSkew = 0.7
	}
	if config.DotCount == 0 {
		config.DotCount = 80
	}
	return config
}
