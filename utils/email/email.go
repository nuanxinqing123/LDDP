// -*- coding: utf-8 -*-

package email

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"regexp"
	"strconv"
)

// VerifyEmailFormat 正则验证邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// SendEmailCode 发送邮箱验证码
func SendEmailCode(Title, UserEmail, Content string) error {
	//定义收件人
	mailTo := []string{UserEmail}
	//邮件主题为"Hello"
	subject := Title

	err := SendMail(mailTo, subject, Content)
	if err != nil {
		return err
	}

	return nil
}

// SendMail 发送邮件
func SendMail(mailTo []string, subject, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": viper.GetString("email.user"),
		"pass": viper.GetString("email.pass"),
		"host": viper.GetString("email.host"),
		"port": viper.GetString("email.port"),
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "LDDP官方<system@yurenxianbao.com>")) //这种方式可以添加别名，即“XX官方”
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	m.SetHeader("To", mailTo...)    // 发送给多个用户
	m.SetHeader("Subject", subject) // 设置邮件主题
	m.SetBody("text/html", body)    // 设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err != nil {
		zap.L().Error("发送邮件失败，原因：" + err.Error())
		return err
	}
	return nil
}
