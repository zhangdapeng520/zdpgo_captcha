package captcha

import (
	"strings"
)

var (
	// 默认的数字验证码对象
	DefaultCaptchaDigit *CaptchaDigit
)

func init() {
	// 初始化数字验证码对象
	DefaultCaptchaDigit = &CaptchaDigit{
		Driver: DefaultDriverDigit,
		Store:  DefaultMemoryStore,
	}
}

// 验证码结构体基本内容
type CaptchaDigit struct {
	Driver *DriverDigit
	Store  *MemoryStore
}

// 生成数字验证码对象
func NewCaptchaDigit(driver *DriverDigit, store *MemoryStore) *CaptchaDigit {
	return &CaptchaDigit{Driver: driver, Store: store}
}

// 生成随机的ID和验证码图片的base64字符串
func (c *CaptchaDigit) Generate() (id, b64s string, err error) {
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
func (c *CaptchaDigit) Verify(id, answer string, clear bool) (match bool) {
	// 获取存储的验证码
	vv := c.Store.Get(id, clear)

	// 去除空格
	vv = strings.TrimSpace(vv)

	// 校验
	return vv == strings.TrimSpace(answer)
}
