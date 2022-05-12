package main

import (
	"encoding/gob"
	"mail/callClient/callRegister"
	"mail/config"
	"mail/registerCenter"
	"mail/sendEmail"
	"mail/serviceSearch"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
)

func main() {
	gob.Register([]interface{}{})
	gob.Register(map[string][]string{})
	gob.Register(map[string]interface{}{})
	// 服务发现
	go serviceSearch.ServiceSearch()
	go serviceSearch.InitSearch()
	rpc.RegisterName(registerCenter.MailSeviceName, new(MailService))
	rpc.HandleHTTP()

	go func() {
		// 注册
		callRegister.CallServiceRegister(registerCenter.MailSeviceName, config.MailListenAddress, config.MailListenPort)
		// 监听 ctrl + c
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan)
		<-sigChan
		// 注销服务
		callRegister.CallServiceLogout(registerCenter.MailSeviceName, config.MailListenAddress, config.MailListenPort)
		os.Exit(0)
	}()

	go sendEmail.SendCode()
	// 监听端口
	http.ListenAndServe(config.MailListenPort, nil)
}

// 邮件服务接口实例
type MailService struct{}

// 发送邮件
func (c *MailService) SendStockService(email string, reply *interface{}) error {
	return sendEmail.SendStock(email)
}

// 发送验证码
func (c *MailService) SendCodeService(userEmail string, reply *interface{}) error {
	return nil
}
