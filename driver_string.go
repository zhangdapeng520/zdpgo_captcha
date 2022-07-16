package zdpgo_captcha

import (
	"image/color"
	"strings"

	"github.com/golang/freetype/truetype"
)

//DriverString 字符串驱动
type DriverString struct {
	Height          int         // 图片高度
	Width           int         // 图片宽度
	NoiseCount      int         // noise数量
	ShowLineOptions int         // 线条数量
	Length          int         // 字符串长度
	Source          string      // 源
	BgColor         *color.RGBA // 背景颜色
	fontsStorage    FontsStorage
	Fonts           []string
	fontsArray      []*truetype.Font
}

//NewDriverString 创建字符串驱动
func NewDriverString(height int, width int, noiseCount int, showLineOptions int, length int, source string, bgColor *color.RGBA, fontsStorage FontsStorage, fonts []string) *DriverString {
	if fontsStorage == nil {
		fontsStorage = DefaultEmbeddedFonts
	}

	tfs := []*truetype.Font{}
	for _, fff := range fonts {
		tf := fontsStorage.LoadFontByName("fonts/" + fff)
		tfs = append(tfs, tf)
	}

	if len(tfs) == 0 {
		tfs = fontsAll
	}

	return &DriverString{Height: height, Width: width, NoiseCount: noiseCount, ShowLineOptions: showLineOptions, Length: length, Source: source, BgColor: bgColor, fontsStorage: fontsStorage, fontsArray: tfs}
}

//ConvertFonts 加载字体
func (d *DriverString) ConvertFonts() *DriverString {
	if d.fontsStorage == nil {
		d.fontsStorage = DefaultEmbeddedFonts
	}

	tfs := []*truetype.Font{}
	for _, fff := range d.Fonts {
		tf := d.fontsStorage.LoadFontByName("fonts/" + fff)
		tfs = append(tfs, tf)
	}
	if len(tfs) == 0 {
		tfs = fontsAll
	}

	d.fontsArray = tfs

	return d
}

//GenerateIdQuestionAnswer 创建ID，问题和答案
func (d *DriverString) GenerateIdQuestionAnswer() (id, content, answer string) {
	id = RandomId()
	content = RandText(d.Length, d.Source)
	return id, content, content
}

//DrawCaptcha 创建验证码
func (d *DriverString) DrawCaptcha(content string) (item Item, err error) {

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
		source := TxtNumbers + TxtAlphabet + ",.[]<>"
		noise := RandText(d.NoiseCount, strings.Repeat(source, d.NoiseCount))
		err = itemChar.drawNoise(noise, d.fontsArray)
		if err != nil {
			return
		}
	}

	//draw content
	err = itemChar.drawText(content, d.fontsArray)
	if err != nil {
		return
	}

	return itemChar, nil
}
