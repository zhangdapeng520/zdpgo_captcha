package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_captcha"
	"github.com/zhangdapeng520/zdpgo_captcha/core/config"
	"github.com/zhangdapeng520/zdpgo_captcha/libs/base64captcha"
	"log"
	"net/http"
)

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64captcha.DriverAudio
	DriverString  *base64captcha.DriverString
	DriverChinese *base64captcha.DriverChinese
	DriverMath    *base64captcha.DriverMath
	DriverDigit   *base64captcha.DriverDigit
}

var (
	store  = base64captcha.DefaultMemStore
	driver = base64captcha.DefaultDriverDigit
)

// base64captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	// 创建验证码对象
	//c := base64captcha.NewCaptcha(driver, store)
	c := zdpgo_captcha.New(config.CaptchaConfig{})

	// 生成验证码
	id, b64s, err := c.Generate()

	// 返回消息
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}

	// 设置为json响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// 响应json数据
	json.NewEncoder(w).Encode(body)
}

// base64captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

	//parse request json body
	decoder := json.NewDecoder(r.Body)
	var param configJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if store.Verify(param.Id, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}
	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(body)
}

func main() {
	// 获取验证码
	http.HandleFunc("/", generateCaptchaHandler)

	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	fmt.Println("Server is at :8777")
	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
}
