package spiderMail

import (
	"github.com/go-gomail/gomail"
)

type mail struct {
	// 发送者邮箱
	SenderAccount string

	// 发送者邮箱密码（或授权码）
	SenderPassword string

	// 接收者邮箱
	Receiver string

	// 服务器地址，outlook 为 smtp.office365.com
	ServerAddr string

	// 服务器端口，outlook 为 587
	ServerPort int

	// 可选附件
	Attchs []string
}

func GetNewMail(userMailAccount string) *mail {
	return &mail{
		SenderAccount:  "",
		SenderPassword: "",
		Receiver:       userMailAccount,
		ServerAddr:     "",
		ServerPort:     1,
	}
}

// 发送邮件，需要标题和正文
func (mail *mail) Send(title, text string, mess *gomail.Message) error {
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
	// 发送
	dial := gomail.NewDialer(mail.ServerAddr, mail.ServerPort, mail.SenderAccount, mail.SenderPassword)
	err := dial.DialAndSend(mess)
	if err != nil {
		return err
	}
	return nil
}
