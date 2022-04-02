package base64captcha

import "time"

var (
	GCLimitNumber   = 10240                                     // 创建的验证码数量，用于触发默认存储区使用的垃圾收集。
	Expiration      = 10 * time.Minute                          // 默认过期时间
	DefaultMemStore = NewMemoryStore(GCLimitNumber, Expiration) // 默认内存存储器
)
