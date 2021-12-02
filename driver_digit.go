package captcha

import "math/rand"

// 数字驱动器结构体
type DriverDigit struct {
	// 图片的高度
	Height int

	// 图片的宽度
	Width int

	// 验证码的个数
	Length int

	// 倾斜银子
	MaxSkew float64

	// 点的数量
	DotCount int
}

// 生成数字驱动的对象
func NewDriverDigit(height int, width int, length int, maxSkew float64, dotCount int) *DriverDigit {
	return &DriverDigit{Height: height, Width: width, Length: length, MaxSkew: maxSkew, DotCount: dotCount}
}

// 生成ID和问题的答案
func (d *DriverDigit) GenerateIdQuestionAnswer() (id, q, a string) {
	id = RandomId()
	digits := randomDigits(d.Length)
	a = parseDigitsToString(digits)
	return id, a, a
}

// 绘制图片验证码
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
