// Copyright 2017 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package base64captcha supports digits, numbers,alphabet, arithmetic, audio and digit-alphabet captcha.
// base64captcha is used for fast development of RESTful APIs, web apps and backend services in Go. give a string identifier to the package and it returns with a base64-encoding-png-string
package base64captcha

import (
	"fmt"
	"strings"
)

// Captcha captcha basic information.
type Captcha struct {
	Driver Driver
	Store  Store
}

//NewCaptcha creates a captcha instance from driver and store
func NewCaptcha(driver Driver, store Store) *Captcha {
	return &Captcha{Driver: driver, Store: store}
}

//Generate 生成随机的ID，验证码图片的base64编码和错误
func (c *Captcha) Generate() (id, b64s string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}

	// 存储id和正确的验证码
	err = c.Store.Set(id, answer)

	if err != nil {
		return "", "", err
	}
	b64s = item.EncodeB64string()
	return
}

//Verify 通过给定的id键，并删除存储中的验证码值，返回一个布尔值
// 如果您有多个验证码实例共享同一个存储，可以调用 `store.Verify` 方法
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
	vv := c.Store.Get(id, clear)
	fmt.Println("=========", vv, answer)
	//fix issue for some redis key-value string value
	vv = strings.TrimSpace(vv)
	return vv == strings.TrimSpace(answer)
}
