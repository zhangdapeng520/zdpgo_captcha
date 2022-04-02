package base64captcha

import (
	"image/color"
	"math/rand"
	"strings"

	"github.com/golang/freetype/truetype"
)

//DriverChinese 中文驱动
type DriverChinese struct {
	Height          int          // png图片像素高度
	Width           int          // png图片像素宽度
	NoiseCount      int          // 文本noise数量
	ShowLineOptions int          // 线条数量
	Length          int          // 字符长度
	Source          string       // 字符串源
	BgColor         *color.RGBA  // 背景颜色
	fontsStorage    FontsStorage // 字体
	Fonts           []string     // 字体列表
	fontsArray      []*truetype.Font
}

//NewDriverChinese 创建中文验证码驱动
func NewDriverChinese(height int, width int, noiseCount int, showLineOptions int, length int, source string, bgColor *color.RGBA, fontsStorage FontsStorage, fonts []string) *DriverChinese {
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

	return &DriverChinese{Height: height, Width: width, NoiseCount: noiseCount, ShowLineOptions: showLineOptions, Length: length, Source: source, BgColor: bgColor, fontsStorage: fontsStorage, fontsArray: tfs}
}

//ConvertFonts 加载字体
func (d *DriverChinese) ConvertFonts() *DriverChinese {
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

//GenerateIdQuestionAnswer 生成验证码ID，问题和答案
func (d *DriverChinese) GenerateIdQuestionAnswer() (id, content, answer string) {
	id = RandomId()

	ss := strings.Split(d.Source, ",")
	length := len(ss)
	if length == 1 {
		c := RandText(d.Length, ss[0])
		return id, c, c
	}
	if length <= d.Length {
		c := RandText(d.Length, TxtNumbers+TxtAlphabet)
		return id, c, c
	}

	res := make([]string, d.Length)
	for k := range res {
		res[k] = ss[rand.Intn(length)]
	}

	content = strings.Join(res, "")
	return id, content, content
}

//DrawCaptcha 生成验证码
func (d *DriverChinese) DrawCaptcha(content string) (item Item, err error) {

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
