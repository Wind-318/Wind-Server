package main

import (
	"algorithm/algorithm"
	"algorithm/callClient/callRegister"
	"algorithm/config"
	"algorithm/registerCenter"
	"algorithm/serviceSearch"
	"encoding/gob"
	"errors"
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
	rpc.RegisterName(registerCenter.AlgorithmServiceName, new(AlgorithmService))
	rpc.HandleHTTP()

	go func() {
		// 注册
		callRegister.CallServiceRegister(registerCenter.AlgorithmServiceName, config.AlgorithmListenAddress, config.AlgorithmListenPort)
		// 监听 ctrl + c
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan)
		<-sigChan
		// 注销服务
		callRegister.CallServiceLogout(registerCenter.AlgorithmServiceName, config.AlgorithmListenAddress, config.AlgorithmListenPort)
		os.Exit(0)
	}()
	// 监听端口
	http.ListenAndServe(config.AlgorithmListenPort, nil)
}

// AlgorithmInterface 的接口实现
type AlgorithmService struct{}

// 调用函数，返回数据
func (c *AlgorithmService) ZipAlgorithm(request registerCenter.AlgorithmZipDataInterface, reply *interface{}) error {
	return algorithm.Zip(request.Source, request.Target)
}

// randomTime 算法
func (c *AlgorithmService) RandomTimeAlgorithm(request registerCenter.AlgorithmRandomTimeInterface, reply *int) error {
	if request.RangeNumber < 0 || request.Start < 0 {
		return errors.New("时间不能小于 0")
	}
	*reply = algorithm.GetRandomNumberTime(request.Start, request.RangeNumber)
	return nil
}

// match 算法
func (c *AlgorithmService) MatchAlgorithm(request registerCenter.AlgorithmMatchInterface, reply *bool) error {
	*reply = algorithm.Match(request.Str, request.Target)
	return nil
}

// url 页面分析
func (c *AlgorithmService) UrlRegexpAlgorithm(request registerCenter.AlgorithmUrlHtmlInterface, reply *[][]string) error {
	arr, err := algorithm.RegexpUrlHtml(request.Url, request.RegexpRule)
	*reply = arr
	return err
}

// 字符串分析
func (c *AlgorithmService) StringRegexpAlgorithm(request registerCenter.AlgorithmHtmlInterface, reply *[][]string) error {
	arr, err := algorithm.RegexpHtml(request.Html, request.RegexpRule)
	*reply = arr
	return err
}

// 判断密码合法性
func (c *AlgorithmService) JudgePasswordAlgorithm(password string, reply *interface{}) error {
	return algorithm.JudgePasswordIllegal(password)
}

// 获取 url 页面内容
func (c *AlgorithmService) GetBytesAlgorithm(url string, reply *[]byte) error {
	bytes, err := algorithm.GetRequestByte(url)
	*reply = bytes
	return err
}

// IP 访问次数加一
func (c *AlgorithmService) AddIPAlgorithm(ip string, reply *interface{}) error {
	return algorithm.AddOneForThisIP(ip)
}

// 判断是否频繁访问
func (c *AlgorithmService) IfRestrictedAlgorithm(ip string, reply *bool) error {
	*reply = algorithm.IfShouldGetRestricted(ip)
	return nil
}

// 生成随机 user-agent
func (c *AlgorithmService) GetRandomUserAgentAlgorithm(request interface{}, reply *string) error {
	*reply = algorithm.GetRandomUserAgent()
	return nil
}

// 加密
func (c *AlgorithmService) EncryptionAlgorithm(password string, reply *string) error {
	*reply = algorithm.Encryption(password)
	return nil
}

// 生成符合正态分布的整数，传入偏移量，允许的下限和上限
func (c *AlgorithmService) GetNormalDistributionNumberAlgorithm(request registerCenter.AlgorithmNormalDistributeInterface, reply *int) error {
	if request.Maximum < request.Minimum {
		return errors.New("上限不能低于下限！")
	}
	*reply = algorithm.GetNormalDistributionNumber(request.Offset, request.Minimum, request.Maximum)
	return nil
}

// 获取一串长度为 length 的随机字符串
func (c *AlgorithmService) GenerateRandStrAlgorithm(length int, reply *string) error {
	*reply = algorithm.GenerateRandStr(length)
	return nil
}
