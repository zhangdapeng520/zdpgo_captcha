package zdpgo_captcha

import "github.com/golang/freetype/truetype"

// FontsStorage 字体接口
type FontsStorage interface {
	// LoadFontByName 根据名字加载字体
	LoadFontByName(name string) *truetype.Font

	// LoadFontsByNames 根据名称列表加载字体列表
	LoadFontsByNames(assetFontNames []string) []*truetype.Font
}
