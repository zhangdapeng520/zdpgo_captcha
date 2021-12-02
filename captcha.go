package captcha

import (
	"strings"
)

// 验证码结构体基本内容
type Captcha struct {
	Driver Driver
	Store  Store
}

// 创建验证码对象
func NewCaptcha(driver Driver, store Store) *Captcha {
	return &Captcha{Driver: driver, Store: store}
}

// 生成随机的ID和验证码图片的base64字符串
func (c *Captcha) Generate() (id, b64s string, err error) {
	// 创建ID，内容和答案
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()

	// 绘制图片
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}

	// 存储内容
	err = c.Store.Set(id, answer)
	if err != nil {
		return "", "", err
	}

	// 转换为base64编码
	b64s = item.EncodeB64string()
	return
}

// 校验ID和验证码
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
	// 获取存储的验证码
	vv := c.Store.Get(id, clear)

	// 去除空格
	vv = strings.TrimSpace(vv)

	// 校验
	return vv == strings.TrimSpace(answer)
}
