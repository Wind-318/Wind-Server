package callRegister

import (
	"net/rpc"
	"storage/callClient/data"
	"storage/config"
)

// 操作客户端
type RegisterClient struct {
	*rpc.Client
}

// 拨号
func dialRegisterService(network, address string) (*RegisterClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &RegisterClient{Client: client}, nil
}

// 注册服务
func CallServiceRegister(serviceName, ip, port string) error {
	// 拨号
	client, err := dialRegisterService("tcp", config.ServiceCenterAddress+config.ServiceCenterPort)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.RegisterRequestInterface{
		RegisterRequestServiceName: serviceName,
		IP:                         ip,
		Port:                       port,
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.RegisterServiceName+".ServiceRegisterRequest", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 注销服务
func CallServiceLogout(serviceName, ip, port string) error {
	// 拨号
	client, err := dialRegisterService("tcp", config.ServiceCenterAddress+config.ServiceCenterPort)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.RegisterRequestInterface{
		RegisterRequestServiceName: serviceName,
		IP:                         ip,
		Port:                       port,
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.RegisterServiceName+".ServiceLogoutRequest", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 服务发现
func CallServiceSearch(serviceName string) ([]string, error) {
	// 拨号
	client, err := dialRegisterService("tcp", config.ServiceCenterAddress+config.ServiceCenterPort)
	// 错误处理
	if err != nil {
		return nil, err
	}

	var result []string
	// 调用
	err = client.Client.Call(data.RegisterServiceName+".ServiceSearch", serviceName, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 服务全发现
func CallServiceSearchAll() (map[string][]string, error) {
	// 拨号
	client, err := dialRegisterService("tcp", config.ServiceCenterAddress+config.ServiceCenterPort)
	// 错误处理
	if err != nil {
		return nil, err
	}

	var result map[string][]string
	// 调用
	err = client.Client.Call(data.RegisterServiceName+".ServiceSearchAll", "", &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}
