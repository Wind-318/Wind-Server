package callUser

import (
	"net/rpc"
	"storage/callClient/data"
	"storage/serviceSearch"
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

// 获取邮箱
func CallUserGetUserEmail(cookie string) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	var result string
	// 调用
	err = client.Client.Call(data.UserServiceName+".GetUserEmail", cookie, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 获取所有用户名称
func CallUserGetUsersName(cookie string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	var result []string
	// 调用
	err = client.Client.Call(data.UserServiceName+".GetUsersName", "", &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 检查是否是管理员
func CallUserIsSystem(cookie string) (bool, error) {
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
	err = client.Client.Call(data.UserServiceName+".IsSystem", cookie, &result)
	// 错误处理
	if err != nil {
		return false, err
	}

	return result, nil
}

// 获取用户名
func CallUserGetName(cookie string) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	var result string
	// 调用
	err = client.Client.Call(data.UserServiceName+".GetName", cookie, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}
