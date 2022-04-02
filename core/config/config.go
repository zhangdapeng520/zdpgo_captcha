package config

// CaptchaConfig 验证码配置
type CaptchaConfig struct {
	StoreType  string `yaml:"store_type" json:"store_type"`   // 存储方式
	DriverType string `yaml:"driver_type" json:"driver_type"` // 驱动类型
}
