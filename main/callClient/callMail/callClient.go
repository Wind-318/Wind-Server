package callMail

import (
	"Project/callClient/data"
	"Project/serviceSearch"
	"net/rpc"

	"github.com/gin-gonic/gin"
)

// 操作数据库客户端
type MailClient struct {
	*rpc.Client
}

// 拨号
func dialMailService(network, address string) (*MailClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &MailClient{Client: client}, nil
}

// 发送邮件
func CallMailSendStockService(email string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MailServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialMailService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.MailServiceName+".SendStockService", email, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 发送验证码
func CallMailSendCodeService(ctx *gin.Context) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MailServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialMailService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.MailServiceName+".SendCodeService", ctx.PostForm("userEmail"), &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}
