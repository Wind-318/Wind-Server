package sendEmail

import (
	"mail/config"

	"github.com/go-gomail/gomail"
)

// 邮件对象
type OwnMail struct {
	// 发送者邮箱
	SenderAccount string
	// 发送者邮箱密码（或授权码）
	SenderPassword string
	// 接收者邮箱
	Receiver string
	// 服务器地址
	ServerAddr string
	// 服务器端口
	ServerPort int
	// 可选附件
	Attchs []string
}

// 返回一个邮件对象
func GetNewMail(userMailAccount string) *OwnMail {
	return &OwnMail{
		SenderAccount:  config.SenderAccount,
		SenderPassword: config.SenderPassword,
		Receiver:       userMailAccount,
		ServerAddr:     config.ServerAddr,
		ServerPort:     config.ServerPort,
	}
}

// 发送邮件，需要标题和正文
func (mail *OwnMail) Send(title, text string, mess *gomail.Message, picturesAddr ...string) error {
	// 设置发送方
	mess.SetHeader("From", mail.SenderAccount)
	// 设置接收方
	mess.SetHeader("To", mail.Receiver)
	// 设置标题
	mess.SetHeader("Subject", title)
	// 设置正文
	mess.SetBody("text/html", text)

	// 如果有附件则添加附件
	if len(mail.Attchs) != 0 {
		for _, addr := range mail.Attchs {
			mess.Attach(addr)
		}
	}
	if len(picturesAddr) != 0 {
		for _, addr := range picturesAddr {
			mess.Embed(addr)
		}
	}
	// 发送
	dial := gomail.NewDialer(mail.ServerAddr, mail.ServerPort, mail.SenderAccount, mail.SenderPassword)
	err := dial.DialAndSend(mess)
	if err != nil {
		return err
	}
	return nil
}
