package base64captcha

import "io"

//Item 二进制验证码接口
type Item interface {
	//WriteTo 写入到一个输出对象
	WriteTo(w io.Writer) (n int64, err error)

	// EncodeB64string 转换为base64字符串
	EncodeB64string() string
}
