package base64captcha

import "embed"

// 默认的字体
var defaultEmbeddedFontsFS embed.FS
var DefaultEmbeddedFonts = NewEmbeddedFontsStorage(defaultEmbeddedFontsFS)
