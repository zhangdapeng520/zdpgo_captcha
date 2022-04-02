go get -u github.com/jteeuwen/go-bindata/...

go-bindata fonts
sed -i "s/package main/package base64captcha/g" bindata.go


https://github.com/jteeuwen/go-bindata