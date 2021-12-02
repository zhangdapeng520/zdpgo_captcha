package captcha

import "time"

var (
	// 默认限制数量
	LimitNumber = 4

	// 默认过期时间
	Expire = 10 * time.Minute

	// 默认内存存储器
	DefaultMemoryStore = NewMemoryStore(LimitNumber, Expire)

	// 默认的驱动
	DefaultDriverDigit = NewDriverDigit(80, 240, 5, 0.7, 80)
)
