package captcha

import (
	"container/list"
	"sync"
	"time"
)

// expValue stores timestamp and id of captchas. It is used in the list inside
// MemoryStore for indexing generated captchas by timestamp to enable garbage
// collection of expired captchas.
type idByTimeValue struct {
	timestamp time.Time
	id        string
}

// 内存存储对象
type MemoryStore struct {
	sync.RWMutex
	digitsById map[string]string
	idByTime   *list.List
	numStored  int           // 存储数据的个数
	collectNum int           // 收集数据的个数
	Expire     time.Duration // 过期时间
}

// 生成基于内存的存储结构体对象
func NewMemoryStore(collectNum int, Expire time.Duration) *MemoryStore {
	s := new(MemoryStore)
	s.digitsById = make(map[string]string)
	s.idByTime = list.New()
	s.collectNum = collectNum
	s.Expire = Expire
	return s
}

// 设置ID对应的验证码
func (s *MemoryStore) Set(id string, value string) error {
	s.Lock()
	s.digitsById[id] = value
	s.idByTime.PushBack(idByTimeValue{time.Now(), id})
	s.numStored++
	s.Unlock()
	if s.numStored > s.collectNum {
		go s.collect()
	}
	return nil
}

// 校验验证码
func (s *MemoryStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}

// 获取验证码
func (s *MemoryStore) Get(id string, clear bool) (value string) {
	if !clear {
		s.RLock()
		defer s.RUnlock()
	} else {
		s.Lock()
		defer s.Unlock()
	}
	value, ok := s.digitsById[id]
	if !ok {
		return
	}
	if clear {
		delete(s.digitsById, id)
	}
	return
}

// 收集
func (s *MemoryStore) collect() {
	now := time.Now()
	s.Lock()
	defer s.Unlock()
	for e := s.idByTime.Front(); e != nil; {
		e = s.collectOne(e, now)
	}
}

// 收集一次
func (s *MemoryStore) collectOne(e *list.Element, specifyTime time.Time) *list.Element {

	ev, ok := e.Value.(idByTimeValue)
	if !ok {
		return nil
	}

	if ev.timestamp.Add(s.Expire).Before(specifyTime) {
		delete(s.digitsById, ev.id)
		next := e.Next()
		s.idByTime.Remove(e)
		s.numStored--
		return next
	}
	return nil
}
