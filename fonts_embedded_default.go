package zdpgo_captcha

import "embed"

// 用于将fonts下的静态文件嵌入到打包编译的二进制文件中
//go:embed fonts/*.ttf
//go:embed fonts/*.ttc
var defaultEmbeddedFontsFS embed.FS
var DefaultEmbeddedFonts = NewEmbeddedFontsStorage(defaultEmbeddedFontsFS)
