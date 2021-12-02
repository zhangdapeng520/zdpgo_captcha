package captcha

import (
	"encoding/binary"
	"image/color"
	"math"
	"math/rand"
	"strings"
	"time"
)

func init() {
	//init rand seed
	rand.Seed(time.Now().UnixNano())
}

//RandText creates random text of given size.
func RandText(size int, sourceChars string) string {
	if sourceChars == "" || size == 0 {
		return ""
	}

	if size >= len(sourceChars) {
		sourceChars = strings.Repeat(sourceChars, size)
	}

	sourceRunes := []rune(sourceChars)
	sourceLength := len(sourceRunes)

	text := make([]rune, size)
	for i := range text {
		text[i] = sourceRunes[rand.Intn(sourceLength)]
	}
	return string(text)
}

//Random get random number between min and max. 生成指定大小的随机数.
func random(min int64, max int64) float64 {
	return float64(min) + rand.Float64()*float64(max-min)
}

//RandDeepColor get random deep color. 随机生成深色系.
func RandDeepColor() color.RGBA {

	randColor := RandColor()

	increase := float64(30 + rand.Intn(255))

	red := math.Abs(math.Min(float64(randColor.R)-increase, 255))

	green := math.Abs(math.Min(float64(randColor.G)-increase, 255))
	blue := math.Abs(math.Min(float64(randColor.B)-increase, 255))

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//RandLightColor get random ligth color. 随机生成浅色.
func RandLightColor() color.RGBA {
	red := rand.Intn(55) + 200
	green := rand.Intn(55) + 200
	blue := rand.Intn(55) + 200
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//RandColor get random color. 生成随机颜色.
func RandColor() color.RGBA {
	red := rand.Intn(255)
	green := rand.Intn(255)
	var blue int
	if (red + green) > 400 {
		blue = 0
	} else {
		blue = 400 - green - red
	}
	if blue > 255 {
		blue = 255
	}
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

func randIntRange(from, to int) int {
	// rand.Intn panics if n <= 0.
	if to-from <= 0 {
		return from
	}
	return rand.Intn(to-from) + from
}
func randFloat64Range(from, to float64) float64 {
	return rand.Float64()*(to-from) + from
}
func randBytes(n int) []byte {
	// Since we don't have a buffer for generated bytes in siprng state,
	// we just generate enough 8-byte blocks and then cut the result to the
	// required length. Doing it this way, we lose generated bytes, and we
	// don't get the strictly sequential deterministic output from PRNG:
	// calling Uint64() and then Bytes(3) produces different output than
	// when calling them in the reverse order, but for our applications
	// this is OK.
	numBlocks := (n + 8 - 1) / 8
	b := make([]byte, numBlocks*8)
	for i := 0; i < len(b); i += 8 {

		binary.LittleEndian.PutUint64(b[i:], rand.Uint64())
	}
	return b[:n]
}

// RandomId returns a new random id key string.
func RandomId() string {
	b := randomBytesMod(idLen, byte(len(idChars)))
	for i, c := range b {
		b[i] = idChars[c]
	}
	return string(b)
}
