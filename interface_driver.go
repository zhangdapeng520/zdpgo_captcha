package zdpgo_captcha

// Driver 驱动器接口
type Driver interface {
	//DrawCaptcha 绘制二进制图片
	DrawCaptcha(content string) (item Item, err error)

	//GenerateIdQuestionAnswer 生成ID，问题和答案
	GenerateIdQuestionAnswer() (id, q, a string)
}
