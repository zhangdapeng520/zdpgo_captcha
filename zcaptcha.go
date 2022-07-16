package zdpgo_captcha

import "image/color"

type ZCaptcha struct {
	store   Store          // 验证码存储对象
	Config  *CaptchaConfig // 配置对象
	captcha *Captcha       //验证码核心对象

	// 方法区
	Generate func() (id, b64s string, err error)
	Verify   func(id, answer string, clear bool) bool
}

// New 创建新的验证码对象
func New() *ZCaptcha {
	return NewWitchConfig(CaptchaConfig{})
}

func NewChinese() *ZCaptcha {
	return NewWitchConfig(CaptchaConfig{
		DriverType: "chinese",
	})
}

func NewMath() *ZCaptcha {
	return NewWitchConfig(CaptchaConfig{
		DriverType: "math",
	})
}

func NewStr() *ZCaptcha {
	return NewWitchConfig(CaptchaConfig{
		DriverType: "str",
	})
}

// NewWitchConfig 使用配置对象创建captcha
func NewWitchConfig(config CaptchaConfig) *ZCaptcha {
	c := ZCaptcha{}

	// 初始化配置
	if config.Width == 0 {
		config.Width = 240
	}
	if config.Height == 0 {
		config.Height = 80
	}
	if config.Length == 0 {
		config.Length = 4
	}
	if config.MaxSkew == 0 {
		config.MaxSkew = 0.7
	}
	if config.DotCount == 0 {
		config.DotCount = 80
	}
	c.Config = &config

	// 初始化存储器
	switch c.Config.StoreType {
	case "memory":
		c.store = DefaultMemStore
	default:
		c.store = DefaultMemStore
	}

	// 初始验证码对象
	switch c.Config.DriverType {
	case "math": // 数学验证码
		driver := NewDriverMath(config)
		c.captcha = NewCaptcha(driver, c.store)
	case "chinese": // 中文验证码
		driver := NewDriverChinese(config)
		c.captcha = NewCaptcha(driver, c.store)
	case "str": // 中文验证码
		driver := NewDriverString(
			config.Height,
			config.Width,
			config.NoiseCount,
			config.ShowLineOptions,
			config.Length,
			"23456789abcdefghijkmnpqrstuvwxyzABCDEFGHIJKMNPQRSTUVWXYZ",
			&color.RGBA{0, 0, 0, 0},
			DefaultEmbeddedFonts,
			[]string{
				"3Dumb.ttf",
				"ApothecaryFont.ttf",
				"Comismsh.ttf",
				"DENNEthree-dee.ttf",
				"DeborahFancyDress.ttf",
				"Flim-Flam.ttf",
				"RitaSmith.ttf",
				"actionj.ttf",
				"chromohv.ttf",
			})
		c.captcha = NewCaptcha(driver, c.store)
	default: // 数字验证码
		driver := DefaultDriverDigit
		c.captcha = NewCaptcha(driver, c.store)
	}

	// 初始化方法
	c.Generate = c.captcha.Generate
	c.Verify = c.captcha.Store.Verify

	return &c
}
