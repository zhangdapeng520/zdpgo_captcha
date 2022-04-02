package zdpgo_captcha

// Generate 验证码生成接口
type Generate interface {
	// Generate 生成随机的ID，图片验证码的base64字符串
	Generate() (id, b64s string, err error)
}

// Verify 验证码校验接口
type Verify interface {
	// Verify 校验ID和验证码，返回校验结果，可以选择校验后是否清空，防止二次校验
	Verify(id, answer string, clear bool) bool
}
