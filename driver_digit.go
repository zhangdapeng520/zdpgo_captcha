package zdpgo_captcha

import (
	"math/rand"
)

//DriverDigit 数字验证码驱动
type DriverDigit struct {
	Height   int     // 图片高度
	Width    int     // 图片宽度
	Length   int     // 图片长度
	MaxSkew  float64 // skew最大
	DotCount int     // 点的数量
}

//NewDriverDigit 创建验证码数字驱动
func NewDriverDigit(cfg CaptchaConfig) *DriverDigit {
	return &DriverDigit{
		Height:   cfg.Height,
		Width:    cfg.Width,
		Length:   cfg.Length,
		MaxSkew:  cfg.MaxSkew,
		DotCount: cfg.DotCount,
	}
}

//DefaultDriverDigit 默认的验证码数字驱动
var DefaultDriverDigit = NewDriverDigit(CaptchaConfig{})

//GenerateIdQuestionAnswer 生成验证码的ID，问题和答案
func (d *DriverDigit) GenerateIdQuestionAnswer() (id, q, a string) {
	id = RandomId()
	digits := randomDigits(d.Length)
	a = parseDigitsToString(digits)
	return id, a, a
}

//DrawCaptcha 创建数字验证码对象
func (d *DriverDigit) DrawCaptcha(content string) (item Item, err error) {
	// Initialize PRNG.
	itemDigit := NewItemDigit(d.Width, d.Height, d.DotCount, d.MaxSkew)
	//parse digits to string
	digits := stringToFakeByte(content)

	itemDigit.calculateSizes(d.Width, d.Height, len(digits))
	// Randomly position captcha inside the image.
	maxx := d.Width - (itemDigit.width+itemDigit.dotSize)*len(digits) - itemDigit.dotSize
	maxy := d.Height - itemDigit.height - itemDigit.dotSize*2
	var border int
	if d.Width > d.Height {
		border = d.Height / 5
	} else {
		border = d.Width / 5
	}
	x := rand.Intn(maxx-border*2) + border
	y := rand.Intn(maxy-border*2) + border
	// Draw digits.
	for _, n := range digits {
		itemDigit.drawDigit(digitFontData[n], x, y)
		x += itemDigit.width + itemDigit.dotSize
	}
	// Draw strike-through line.
	itemDigit.strikeThrough()
	// Apply wave distortion.
	itemDigit.distort(rand.Float64()*(10-5)+5, rand.Float64()*(200-100)+100)
	// Fill image with random circles.
	itemDigit.fillWithCircles(d.DotCount, itemDigit.dotSize)
	return itemDigit, nil
}
