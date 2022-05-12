package main

import (
	"encoding/gob"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"subscribe/callClient/callRegister"
	"subscribe/config"
	"subscribe/registerCenter"
	"subscribe/serviceSearch"
	"subscribe/subscribeInfo"
)

func main() {
	gob.Register([]interface{}{})
	gob.Register(map[string][]string{})
	gob.Register(map[string]interface{}{})
	// 服务发现
	go serviceSearch.ServiceSearch()
	go serviceSearch.InitSearch()
	// 注册服务
	rpc.RegisterName(registerCenter.SubscribeServiceName, new(SubscribeService))
	rpc.HandleHTTP()

	go func() {
		// 注册
		callRegister.CallServiceRegister(registerCenter.SubscribeServiceName, config.SubscribeListenAddress, config.SubscribeListenPort)
		// 监听 ctrl + c
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan)
		<-sigChan
		// 注销服务
		callRegister.CallServiceLogout(registerCenter.SubscribeServiceName, config.SubscribeListenAddress, config.SubscribeListenPort)
		os.Exit(0)
	}()
	// 监听端口
	http.ListenAndServe(config.SubscribeListenPort, nil)
}

// userServiceInterface 实现
type SubscribeService struct{}

// 选取用户名单
func (c *SubscribeService) SelectUsersAccount(request interface{}, reply *[]string) error {
	*reply = subscribeInfo.SelectUsersAccount()
	return nil
}
