package zdpgo_captcha

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/golang/freetype/truetype"
)

// https://en.wikipedia.org/wiki/Unicode_block
// 语言字典映射
var langMap = map[string][]int{
	"zh-CN":   {19968, 40869},
	"latin":   {0x0000, 0x007f},
	"zh":      {0x4e00, 0x9fa5},
	"ko":      {12593, 12686},
	"jp":      {12449, 12531}, //[]int{12353, 12435}
	"ru":      {1025, 1169},
	"th":      {0x0e00, 0x0e7f},
	"greek":   {0x0380, 0x03ff},
	"arabic":  {0x0600, 0x06ff},
	"hebrew":  {0x0590, 0x05ff},
	"emotion": {0x1f601, 0x1f64f},
}

func generateRandomRune(size int, code string) string {
	lang, ok := langMap[code]
	if !ok {
		fmt.Sprintf("can not font language of %s", code)
		lang = langMap["latin"]
	}
	start := lang[0]
	end := lang[1]
	randRune := make([]rune, size)
	for i := range randRune {
		idx := rand.Intn(end-start) + start
		randRune[i] = rune(idx)
	}
	return string(randRune)
}

//DriverLanguage 语言驱动
type DriverLanguage struct {
	Height          int         // 图片高度
	Width           int         // 图片宽度
	NoiseCount      int         // noise数量
	ShowLineOptions int         // 线条
	Length          int         // 验证码长度
	BgColor         *color.RGBA // 背景颜色
	fontsStorage    FontsStorage
	Fonts           []*truetype.Font // 字体列表
	LanguageCode    string
}

//NewDriverLanguage 创建一个语言驱动
func NewDriverLanguage(height int, width int, noiseCount int, showLineOptions int, length int, bgColor *color.RGBA, fontsStorage FontsStorage, fonts []*truetype.Font, languageCode string) *DriverLanguage {
	return &DriverLanguage{Height: height, Width: width, NoiseCount: noiseCount, ShowLineOptions: showLineOptions, Length: length, BgColor: bgColor, fontsStorage: fontsStorage, Fonts: fonts, LanguageCode: languageCode}
}

//GenerateIdQuestionAnswer 创建ID，问题和答案
func (d *DriverLanguage) GenerateIdQuestionAnswer() (id, content, answer string) {
	id = RandomId()
	content = generateRandomRune(d.Length, d.LanguageCode)
	return id, content, content
}

//DrawCaptcha 绘制验证码
func (d *DriverLanguage) DrawCaptcha(content string) (item Item, err error) {
	var bgc color.RGBA
	if d.BgColor != nil {
		bgc = *d.BgColor
	} else {
		bgc = RandLightColor()
	}
	itemChar := NewItemChar(d.Width, d.Height, bgc)

	//draw hollow line
	if d.ShowLineOptions&OptionShowHollowLine == OptionShowHollowLine {
		itemChar.drawHollowLine()
	}

	//draw slime line
	if d.ShowLineOptions&OptionShowSlimeLine == OptionShowSlimeLine {
		itemChar.drawSlimLine(3)
	}

	//draw sine line
	if d.ShowLineOptions&OptionShowSineLine == OptionShowSineLine {
		itemChar.drawSineLine()
	}

	//draw noise
	if d.NoiseCount > 0 {
		noise := RandText(d.NoiseCount, TxtNumbers+TxtAlphabet+",.[]<>")
		err = itemChar.drawNoise(noise, fontsAll)
		if err != nil {
			return
		}
	}

	//draw content
	//use font that match your language
	err = itemChar.drawText(content, []*truetype.Font{fontChinese})
	if err != nil {
		return
	}

	return itemChar, nil
}
