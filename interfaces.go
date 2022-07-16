package zdpgo_captcha

import (
	"github.com/golang/freetype/truetype"
	"io"
)

// Driver 驱动器接口
type Driver interface {
	// DrawCaptcha 绘制二进制图片
	DrawCaptcha(content string) (item Item, err error)

	// GenerateIdQuestionAnswer 生成ID，问题和答案
	GenerateIdQuestionAnswer() (id, q, a string)
}

// FontsStorage 字体接口
type FontsStorage interface {
	// LoadFontByName 根据名字加载字体
	LoadFontByName(name string) *truetype.Font

	// LoadFontsByNames 根据名称列表加载字体列表
	LoadFontsByNames(assetFontNames []string) []*truetype.Font
}

//Item 二进制验证码接口
type Item interface {
	//WriteTo 写入到一个输出对象
	WriteTo(w io.Writer) (n int64, err error)

	// EncodeB64string 转换为base64字符串
	EncodeB64string() string
}

// Store 存储器接口
type Store interface {
	// Set 设置验证码答案
	Set(id string, value string) error

	// Get 获取验证码答案
	Get(id string, clear bool) string

	//Verify 验证验证码答案
	Verify(id, answer string, clear bool) bool
}
