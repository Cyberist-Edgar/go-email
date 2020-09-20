package main

import (
	"fmt"
	"project/email/email"
)

func main() {
	email := email.New("使用golang发送邮件", "text/plain")
	err := email.SetHost("smtp.126.com:25")
	if err != nil {
		fmt.Println(err)
	}
	email.SetTo("xxxx@qq.com", "xxx@qq.com")
	email.SetBody("可以收到邮件吗")
	// 这里的密码应该是授权密码，而非邮箱密码
	email.SetAuth("xxx@126.com", "xxx")
	err = email.SendMail()
	if err != nil {
		fmt.Println(err)
	}
}
