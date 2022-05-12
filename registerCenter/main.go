package main

import (
	"encoding/gob"
	"net/http"
	"net/rpc"
	"registerCenter/config"
	"registerCenter/register"
)

func main() {
	gob.Register(map[string][]string{})
	gob.Register(map[string]interface{}{})
	// 初始化数据库
	register.InitDatabase(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort)
	// 注册服务
	rpc.RegisterName(register.RegisterServiceName, new(RegisterService))
	rpc.HandleHTTP()
	// 监听端口
	http.ListenAndServe(config.ServiceCenterPort, nil)
}

// 注册服务实例
type RegisterService struct{}

// 服务注册
func (r *RegisterService) ServiceRegisterRequest(request register.RegisterRequestInterface, reply *interface{}) error {
	return register.ServiceRegister(request)
}

// 接受注销请求
func (r *RegisterService) ServiceLogoutRequest(request register.RegisterRequestInterface, reply *interface{}) error {
	return register.ServiceLogout(request)
}

// 服务发现
func (r *RegisterService) ServiceSearch(request string, reply *[]string) error {
	temp, err := register.ServiceSearch(request)
	*reply = temp
	return err
}

// 服务全发现
func (r *RegisterService) ServiceSearchAll(request string, reply *map[string][]string) error {
	temp, err := register.ServiceSearchAll()
	*reply = temp
	return err
}
