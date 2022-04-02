package zdpgo_captcha

import (
	"github.com/zhangdapeng520/zdpgo_captcha/core/config"
	"github.com/zhangdapeng520/zdpgo_captcha/libs/base64captcha"
)

type Captcha struct {
	store   base64captcha.Store    // 验证码存储对象
	config  *config.CaptchaConfig  // 配置对象
	captcha *base64captcha.Captcha //验证码核心对象

	// 方法区
	Generate func() (id, b64s string, err error)
	Verify   func(id, answer string, clear bool) bool
}

// Default 使用默认配置生成验证码对象
func Default() *Captcha {
	return New(config.CaptchaConfig{})
}

// New 创建新的验证码对象
func New(cf config.CaptchaConfig) *Captcha {
	c := Captcha{}

	// 初始化配置
	cfg := config.GetDefaultCaptchaConfig(cf)
	c.config = &cfg

	// 初始化存储器
	switch c.config.StoreType {
	case "memory":
		c.store = base64captcha.DefaultMemStore
	default:
		c.store = base64captcha.DefaultMemStore
	}

	// 初始验证码对象
	switch c.config.DriverType {
	case "audio": // 音频验证码
		driver := base64captcha.NewDriverAudio(6, "zh")
		c.captcha = base64captcha.NewCaptcha(driver, c.store)
	case "math": // 数学验证码
		driver := base64captcha.NewDriverMath(cfg)
		c.captcha = base64captcha.NewCaptcha(driver, c.store)
	case "chinese": // 中文验证码
		driver := base64captcha.NewDriverChinese(cfg)
		c.captcha = base64captcha.NewCaptcha(driver, c.store)
	default: // 数字验证码
		driver := base64captcha.DefaultDriverDigit
		c.captcha = base64captcha.NewCaptcha(driver, c.store)
	}

	// 初始化方法
	c.Generate = c.captcha.Generate
	c.Verify = c.captcha.Store.Verify

	return &c
}
