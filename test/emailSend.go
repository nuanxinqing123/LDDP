package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

func main() {
	// "system@yurenxianbao.com"
	// "system@yurenxianbao.com"
	// "lala/yurenxianbao-com"
	// "Ox1IWi9Fc4NRZqSrOB"
	// "smtpdm.aliyun.com:465"
	err := SendEmailCode("nuanxinqing@gmail.com", "邮件测试")
	if err != nil {
		fmt.Println(err)
	}
}

// SendEmailCode 发送邮箱验证码
func SendEmailCode(em, body string) error {
	//定义收件人
	mailTo := []string{em}
	//邮件主题为"Hello"
	subject := "LDDP"

	err := SendMail(mailTo, subject, body)
	if err != nil {
		return err
	}

	return nil
}

// SendMail 发送邮件
func SendMail(mailTo []string, subject, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "system@yurenxianbao.com",
		"pass": "Ox1IWi9Fc4NRZqSrOB",
		"host": "smtpdm.aliyun.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "LDDP官方")) //这种方式可以添加别名，即“XX官方”
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	m.SetHeader("To", mailTo...)    // 发送给多个用户
	m.SetHeader("Subject", subject) // 设置邮件主题
	m.SetBody("text/html", body)    // 设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
