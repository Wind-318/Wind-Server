package main

import (
	"encoding/gob"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"spider/callClient/callRegister"
	"spider/config"
	"spider/mainPackage"
	"spider/registerCenter"
	"spider/serviceSearch"
)

func main() {
	gob.Register([]interface{}{})
	gob.Register(map[string][]string{})
	gob.Register(map[string]interface{}{})
	// 服务发现
	go serviceSearch.ServiceSearch()
	go serviceSearch.InitSearch()
	// 注册服务
	rpc.RegisterName(registerCenter.AnimeServiceName, new(mainPackage.AnimeSpiderService))
	rpc.RegisterName(registerCenter.SinaServiceName, new(mainPackage.SinaSpiderService))
	rpc.HandleHTTP()

	go func() {
		// 注册
		callRegister.CallServiceRegister(registerCenter.AnimeServiceName, config.SpiderListenAddress, config.SpiderListenPort)
		callRegister.CallServiceRegister(registerCenter.SinaServiceName, config.SpiderListenAddress, config.SpiderListenPort)
		// 监听 ctrl + c
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan)
		<-sigChan
		// 注销服务
		callRegister.CallServiceLogout(registerCenter.AnimeServiceName, config.SpiderListenAddress, config.SpiderListenPort)
		callRegister.CallServiceLogout(registerCenter.SinaServiceName, config.SpiderListenAddress, config.SpiderListenPort)
		os.Exit(0)
	}()
	// 监听端口
	http.ListenAndServe(config.SpiderListenPort, nil)
}
