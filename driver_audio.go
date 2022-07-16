// Copyright 2017 Eric Zhou. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zdpgo_captcha

//DriverAudio 音频引擎
type DriverAudio struct {
	Length   int    // 默认长度
	Language string // 语言："en", "ja", "ru", "zh".
}

//DefaultDriverAudio 默认音频驱动
var DefaultDriverAudio = NewDriverAudio(6, "zh")

//NewDriverAudio 创建音频驱动
func NewDriverAudio(length int, language string) *DriverAudio {
	return &DriverAudio{Length: length, Language: language}
}

//DrawCaptcha 创建音频验证码
func (d *DriverAudio) DrawCaptcha(content string) (item Item, err error) {
	digits := stringToFakeByte(content)
	audio := newAudio("", digits, d.Language)
	return audio, nil
}

//GenerateIdQuestionAnswer 创建ID，音频问题，和答案
func (d *DriverAudio) GenerateIdQuestionAnswer() (id, q, a string) {
	id = RandomId()
	digits := randomDigits(d.Length)
	a = parseDigitsToString(digits)
	return id, a, a
}
