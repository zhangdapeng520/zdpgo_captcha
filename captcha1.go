package zdpgo_captcha

type ZCaptcha struct {
	store   Store          // 验证码存储对象
	Config  *CaptchaConfig // 配置对象
	captcha *Captcha       //验证码核心对象

	// 方法区
	Generate func() (id, b64s string, err error)
	Verify   func(id, answer string, clear bool) bool
}

// Default 使用默认配置生成验证码对象
func Default() *ZCaptcha {
	return New(CaptchaConfig{})
}

// New 创建新的验证码对象
func New(cf CaptchaConfig) *ZCaptcha {
	c := ZCaptcha{}

	// 初始化配置
	cfg := GetDefaultCaptchaConfig(cf)
	c.Config = &cfg

	// 初始化存储器
	switch c.Config.StoreType {
	case "memory":
		c.store = DefaultMemStore
	default:
		c.store = DefaultMemStore
	}

	// 初始验证码对象
	switch c.Config.DriverType {
	case "audio": // 音频验证码
		driver := NewDriverAudio(6, "zh")
		c.captcha = NewCaptcha(driver, c.store)
	case "math": // 数学验证码
		driver := NewDriverMath(cfg)
		c.captcha = NewCaptcha(driver, c.store)
	case "chinese": // 中文验证码
		driver := NewDriverChinese(cfg)
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
