package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_captcha"
	"log"
	"net/http"
)

//configJsonBody json request body.
type configJsonBody struct {
	Id          string
	VerifyValue string
}

var (
	captcha = zdpgo_captcha.Default()
)

// base64captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	// 生成验证码
	id, b64s, err := captcha.Generate()

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

// 校验验证码
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	value := r.URL.Query().Get("value")

	// 校验验证码
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if captcha.Verify(id, value, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}

	// 设置返回json
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	json.NewEncoder(w).Encode(body)
}

func main() {
	// 获取验证码
	http.HandleFunc("/", generateCaptchaHandler)

	// 校验验证码
	// 校验示例：http://localhost:8777/verify?id=b9XR0Of9Vy8exHRxyuto&value=2833
	http.HandleFunc("/verify", captchaVerifyHandle)

	// 启动服务
	fmt.Println("启动服务 http://localhost:8777")
	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
}
