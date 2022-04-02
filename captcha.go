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

// New 创建新的验证码对象
func New(config config.CaptchaConfig) *Captcha {
	c := Captcha{}

	// 初始化配置
	if config.DriverType == "" {
		config.DriverType = "digit" // 数字类型
	}
	if config.StoreType == "" {
		config.StoreType = "memory" // 内存存储
	}
	c.config = &config

	// 初始化存储器
	switch config.StoreType {
	case "memory":
		c.store = base64captcha.DefaultMemStore
	default:
		c.store = base64captcha.DefaultMemStore
	}

	// 初始验证码对象
	switch config.DriverType {
	case "audio":
		driver := base64captcha.NewDriverAudio(6, "zh")
		c.captcha = base64captcha.NewCaptcha(driver, c.store)
	default:
		driver := base64captcha.DefaultDriverDigit
		c.captcha = base64captcha.NewCaptcha(driver, c.store)
	}

	// 初始化方法
	c.Generate = c.captcha.Generate
	c.Verify = c.captcha.Store.Verify

	return &c
}
