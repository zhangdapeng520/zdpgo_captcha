package captcha

// 存储数据的接口
type Store interface {
	// 设置ID对应的验证码
	Set(id string, value string) error

	// 获取ID的验证码
	Get(id string, clear bool) string

	// 校验验证码
	Verify(id, answer string, clear bool) bool
}
