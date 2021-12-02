package captcha

// 驱动器接口
type Driver interface {
	// 绘制验证码的方法
	DrawCaptcha(content string) (item Item, err error)

	// 生成ID和问题答案的方法
	GenerateIdQuestionAnswer() (id, q, a string)
}
