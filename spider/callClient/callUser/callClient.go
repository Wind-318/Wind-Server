package callUser

import (
	"net/rpc"
	"spider/callClient/data"
	"spider/serviceSearch"
)

// 操作客户端
type UserClient struct {
	*rpc.Client
}

// 拨号
func dialUserService(network, address string) (*UserClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &UserClient{Client: client}, nil
}

// 判断是否处于登录状态
func CallUserIsLogin(cookie string) (bool, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return false, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return false, err
	}

	var result bool
	// 调用
	err = client.Client.Call(data.UserServiceName+".IsLogin", cookie, &result)
	// 错误处理
	if err != nil {
		return false, err
	}

	return result, nil
}
