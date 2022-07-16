package zdpgo_captcha

import (
	"embed"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

// 用于将fonts下的静态文件嵌入到打包编译的二进制文件中
//go:embed fonts/*.ttf
//go:embed fonts/*.ttc
var defaultEmbeddedFontsFS embed.FS
var DefaultEmbeddedFonts = NewEmbeddedFontsStorage(defaultEmbeddedFontsFS)

type EmbeddedFontsStorage struct {
	fs embed.FS
}

// LoadFontByName 根据名字加载字体
func (s *EmbeddedFontsStorage) LoadFontByName(name string) *truetype.Font {
	fontBytes, err := s.fs.ReadFile(name)
	if err != nil {
		panic(err)
	}

	//font file bytes to trueTypeFont
	trueTypeFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	return trueTypeFont
}

// LoadFontsByNames 根据名称列表加载字体列表
func (s *EmbeddedFontsStorage) LoadFontsByNames(assetFontNames []string) []*truetype.Font {
	fonts := make([]*truetype.Font, 0)
	for _, assetName := range assetFontNames {
		f := s.LoadFontByName(assetName)
		fonts = append(fonts, f)
	}
	return fonts
}

// NewEmbeddedFontsStorage 创建字体存储器
func NewEmbeddedFontsStorage(fs embed.FS) *EmbeddedFontsStorage {
	return &EmbeddedFontsStorage{
		fs: fs,
	}
}
