package zdpgo_captcha

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_captcha/core/config"
	"testing"
)

func getCaptcha() *Captcha {
	return New(config.CaptchaConfig{
		//DriverType: "math",
		//DriverType: "chinese",
		DriverType: "digit",
	})
}

// 测试校验和生成
func TestCaptcha_Verify(t *testing.T) {
	c := getCaptcha()

	// 生成验证码
	id, captcha, err := c.Generate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, captcha)

	// 校验验证码
	var flag bool

	// 获取答案
	answer := c.captcha.Store.Get(id, false)
	fmt.Println("answer===", answer)

	// 第一次校验成功
	flag = c.Verify(id, answer, true)
	fmt.Println(flag)

	// 第二次校验失败，因为该验证码已经在内存中被清空了
	flag = c.Verify(id, answer, true)
	fmt.Println(flag)
}
