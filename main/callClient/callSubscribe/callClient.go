package callSubscribe

import (
	"Project/serviceSearch"
	"net/rpc"
)

// 操作客户端
type SubscribeClient struct {
	*rpc.Client
}

// 拨号
func dialSubscribeService(network, address string) (*SubscribeClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &SubscribeClient{Client: client}, nil
}

// 构造用户对象，传入用户名，账号和密码
func CallSubscribeSelectUsersAccount() ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode("SubscribeService")
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialSubscribeService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	var result []string
	// 调用
	err = client.Client.Call("SubscribeService.SelectUsersAccount", "", &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}
