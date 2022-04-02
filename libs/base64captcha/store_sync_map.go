package base64captcha

import (
	"sync"
	"time"
)

//StoreSyncMap 使用线程安全的map存储
type StoreSyncMap struct {
	liveTime time.Duration
	m        *sync.Map
}

//NewStoreSyncMap 创建新的实例
func NewStoreSyncMap(liveTime time.Duration) *StoreSyncMap {
	return &StoreSyncMap{liveTime: liveTime, m: new(sync.Map)}
}

//smv 值类型
type smv struct {
	t     time.Time
	Value string
}

//newSmv 创建是类型
func newSmv(v string) *smv {
	return &smv{t: time.Now(), Value: v}
}

//rmExpire 移除过期的数据
func (s StoreSyncMap) rmExpire() {
	expireTime := time.Now().Add(-s.liveTime)
	s.m.Range(func(key, value interface{}) bool {
		if sv, ok := value.(*smv); ok && sv.t.Before(expireTime) {
			s.m.Delete(key)
		}
		return true
	})
}

//Set 设置验证码答案
func (s StoreSyncMap) Set(id string, value string) {
	s.rmExpire()
	s.m.Store(id, newSmv(value))
}

//Get 获取验证码答案
func (s StoreSyncMap) Get(id string, clear bool) string {
	v, ok := s.m.Load(id)
	if !ok {
		return ""
	}
	s.m.Delete(id)
	if sv, ok := v.(*smv); ok {
		return sv.Value
	}
	return ""
}

//Verify 校验验证码答案
func (s StoreSyncMap) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}
