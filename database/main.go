package main

import (
	"database/callClient/callRegister"
	"database/config"
	"database/mainPackage"
	"database/registerCenter"
	"database/serviceSearch"
	"encoding/gob"
	"log"
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
	// 注册服务
	rpc.RegisterName(registerCenter.RedisServiceName, new(mainPackage.RedisService))
	// 注册服务
	rpc.RegisterName(registerCenter.MySQLServiceName, new(mainPackage.MySQLService))
	rpc.HandleHTTP()

	go func() {
		// 注册
		err := callRegister.CallServiceRegister(registerCenter.MySQLServiceName, config.DatabaseListenAddress, config.DatabaseListenPort)
		if err != nil {
			log.Println(err)
		}
		err = callRegister.CallServiceRegister(registerCenter.RedisServiceName, config.DatabaseListenAddress, config.DatabaseListenPort)
		if err != nil {
			log.Println(err)
		}
		// 监听 ctrl + c
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan)
		<-sigChan
		// 注销服务
		err = callRegister.CallServiceLogout(registerCenter.MySQLServiceName, config.DatabaseListenAddress, config.DatabaseListenPort)
		if err != nil {
			log.Println(err)
		}
		err = callRegister.CallServiceLogout(registerCenter.RedisServiceName, config.DatabaseListenAddress, config.DatabaseListenPort)
		if err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}()
	// 监听端口
	http.ListenAndServe(config.DatabaseListenPort, nil)
}
