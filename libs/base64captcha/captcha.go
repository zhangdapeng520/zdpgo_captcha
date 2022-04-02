package base64captcha

import (
	"strings"
)

// Captcha 验证码基本信息
type Captcha struct {
	Driver Driver
	Store  Store
}

//NewCaptcha 创建验证码对象
func NewCaptcha(driver Driver, store Store) *Captcha {
	return &Captcha{Driver: driver, Store: store}
}

//Generate 生成随机的ID，验证码图片的base64编码和错误
func (c *Captcha) Generate() (id, b64s string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}

	// 存储id和正确的验证码
	err = c.Store.Set(id, answer)

	if err != nil {
		return "", "", err
	}
	b64s = item.EncodeB64string()
	return
}

//Verify 通过给定的id键，并删除存储中的验证码值，返回一个布尔值
// 如果您有多个验证码实例共享同一个存储，可以调用 `store.Verify` 方法
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
	vv := c.Store.Get(id, clear)

	// 修复了一些redis key-value字符串值的问题
	vv = strings.TrimSpace(vv)
	return vv == strings.TrimSpace(answer)
}
