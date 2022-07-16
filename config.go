package zdpgo_captcha

// BackgroundColor 背景颜色
type BackgroundColor struct {
	R uint8 `yaml:"r" json:"r"`
	G uint8 `yaml:"g" json:"g"`
	B uint8 `yaml:"b" json:"b"`
	A uint8 `yaml:"a" json:"a"`
}

// CaptchaConfig 验证码配置
type CaptchaConfig struct {
	StoreType       string          `yaml:"store_type" json:"store_type"`   // 存储方式
	DriverType      string          `yaml:"driver_type" json:"driver_type"` // 驱动类型
	Width           int             `yaml:"width" json:"width"`             // 验证码宽度
	Height          int             `yaml:"height" json:"height"`           // 验证码高度
	Length          int             `yaml:"length" json:"length"`           // 验证码长度
	MaxSkew         float64         `yaml:"max_skew" json:"max_skew"`
	DotCount        int             `yaml:"dot_count" json:"dot_count"` // 点的数量
	NoiseCount      int             `yaml:"noise_count" json:"noise_count"`
	ShowLineOptions int             `yaml:"show_line_options" json:"show_line_options"`
	BgColor         BackgroundColor `yaml:"background_color" json:"background_color"` // 图片背景颜色
}
