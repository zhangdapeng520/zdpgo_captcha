package zdpgo_captcha

// Store 存储器接口
type Store interface {
	// Set 设置验证码答案
	Set(id string, value string) error

	// Get 获取验证码答案
	Get(id string, clear bool) string

	//Verify 验证验证码答案
	Verify(id, answer string, clear bool) bool
}
