package captcha

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 测试设置和获取验证码
func TestSetGet(t *testing.T) {
	// 创建内存存储
	s := NewMemoryStore(LimitNumber, Expire)

	// 创建ID和内容
	id := "captch1a id"
	d := "random-string"

	// 设置
	_ = s.Set(id, d)
	fmt.Println(id, d)

	// 获取
	d2 := s.Get(id, false)
	fmt.Println(id, d2)

	if d2 != d {
		t.Errorf("保存的值： %v, 获取的值： %v", d, d2)
	}
}

// 测试设置和获取验证码，使用默认内存
func TestSetGetDefault(t *testing.T) {
	// 创建ID和内容
	id := "name"
	d := "张大鹏"

	// 设置
	_ = DefaultMemoryStore.Set(id, d)
	fmt.Println(id, d)

	// 获取
	d2 := DefaultMemoryStore.Get(id, false)
	fmt.Println(id, d2)

	if d2 != d {
		t.Errorf("保存的值： %v, 获取的值： %v", d, d2)
	}
}

// 测试获取并清空验证码
func TestGetClear(t *testing.T) {
	// 创建内存存储
	s := NewMemoryStore(LimitNumber, Expire)
	id := "captcha id"
	d := "932839jfffjkdss"

	// 设置值
	_ = s.Set(id, d)
	d2 := s.Get(id, true)
	fmt.Println("第一次设置和获取：", d, d2)
	if d != d2 {
		t.Errorf("保存的值： %v, 获取的值： %v", d, d2)
	}

	// 再次获取
	d2 = s.Get(id, false)
	fmt.Println("第二次设置和获取：", d, d2)
	if d2 != "" {
		t.Errorf("没有清空： (%q=%v)", id, d2)
	}
}

func BenchmarkSetCollect(b *testing.B) {
	b.StopTimer()
	d := "fdskfew9832232r"
	s := NewMemoryStore(9999, -1)
	ids := make([]string, 1000)
	for i := range ids {
		ids[i] = fmt.Sprintf("%d", rand.Int63())
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			_ = s.Set(ids[j], d)
		}
	}
}

func TestMemoryStore_SetGoCollect(t *testing.T) {
	s := NewMemoryStore(10, -1)
	for i := 0; i <= 100; i++ {
		_ = s.Set(fmt.Sprint(i), fmt.Sprint(i))
	}
}

func TestMemoryStore_CollectNotExpire(t *testing.T) {
	s := NewMemoryStore(10, time.Hour)
	for i := 0; i < 50; i++ {
		_ = s.Set(fmt.Sprint(i), fmt.Sprint(i))
	}

	// let background goroutine to go
	time.Sleep(time.Second)

	if v := s.Get("0", false); v != "0" {
		t.Error("mem store get failed")
	}
}

func TestNewMemoryStore(t *testing.T) {
	type args struct {
		collectNum int
		Expire     time.Duration
	}
	tests := []struct {
		name string
		args args
		want Store
	}{
		{"", args{20, time.Hour}, nil},
		{"", args{20, time.Hour * 5}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryStore(tt.args.collectNum, tt.args.Expire); got == nil {
				t.Errorf("NewMemoryStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MemoryStore_Set(t *testing.T) {
	thisStore := NewMemoryStore(10, time.Hour)
	type args struct {
		id    string
		value string
	}
	tests := []struct {
		name string
		s    Store
		args args
	}{
		{"", thisStore, args{RandomId(), RandomId()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.s.Set(tt.args.id, tt.args.value)
		})
	}
}

func Test_MemoryStore_Verify(t *testing.T) {
	thisStore := NewMemoryStore(10, time.Hour)
	_ = thisStore.Set("xx", "xx")
	got := thisStore.Verify("xx", "xx", false)
	if !got {
		t.Error("failed1")
	}
	got = thisStore.Verify("xx", "xx", true)

	if !got {
		t.Error("failed2")
	}
	got = thisStore.Verify("xx", "xx", true)

	if got {
		t.Error("failed3")
	}
}

func Test_MemoryStore_Get(t *testing.T) {
	thisStore := NewMemoryStore(10, time.Hour)
	_ = thisStore.Set("xx", "xx")
	got := thisStore.Get("xx", false)
	if got != "xx" {
		t.Error("failed1")
	}
	got = thisStore.Get("xx", true)
	if got != "xx" {
		t.Error("failed2")
	}
	got = thisStore.Get("xx", false)
	if got == "xx" {
		t.Error("failed3")
	}

}
