package callAlgorithm

import (
	"net/rpc"
	"spider/callClient/data"
	"spider/serviceSearch"
)

// 操作客户端
type AlgorithmClient struct {
	*rpc.Client
}

// 拨号
func dialAlgorithmService(network, address string) (*AlgorithmClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &AlgorithmClient{Client: client}, nil
}

// 调用微服务 zip，传入参数：源文件夹和目标文件夹路径
func CallAlgorithmZip(source, target string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.AlgorithmZipDataInterface{
		Source: source,
		Target: target,
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".ZipAlgorithm", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 调用微服务 RandomTime，传入参数：最短等待时间和增加随机等待时间
func CallAlgorithmRandomTime(start, randTime int) (int, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return 0, err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return 0, err
	}

	// 创建数据对象
	dataObj := data.AlgorithmRandomTimeInterface{
		Start:       start,
		RangeNumber: randTime,
	}

	var result int
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".RandomTimeAlgorithm", dataObj, &result)
	// 错误处理
	if err != nil {
		return 0, err
	}

	return result, nil
}

// 调用微服务 match，传入参数：要匹配字符串和待匹配字符串
func CallAlgorithmMatch(str, target string) (bool, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return false, err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return false, err
	}

	// 创建数据对象
	dataObj := data.AlgorithmMatchInterface{
		Str:    str,
		Target: target,
	}

	var result bool
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".MatchAlgorithm", dataObj, &result)
	// 错误处理
	if err != nil {
		return false, err
	}

	return result, nil
}

// 调用服务，传入参数：要获取的页面 url 和正则匹配规则
func CallAlgorithmUrlRegexp(url, regexpRule string) ([][]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.AlgorithmUrlHtmlInterface{
		Url:        url,
		RegexpRule: regexpRule,
	}

	var result [][]string
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".UrlRegexpAlgorithm", dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 调用服务，传入参数：待匹配字符串和正则匹配规则
func CallAlgorithmStringRegexp(html, regexpRule string) ([][]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.AlgorithmHtmlInterface{
		Html:       html,
		RegexpRule: regexpRule,
	}

	var result [][]string
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".StringRegexpAlgorithm", dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 判断密码合法性
func CallAlgorithmJudgePassword(password string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".JudgePasswordAlgorithm", password, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 获取 url 页面内容
func CallAlgorithmGetBytes(url string) ([]byte, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	var result []byte
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".GetBytesAlgorithm", url, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 获取随机 user-agent
func CallAlgorithmGetRandomUserAgent() (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	// 创建数据对象
	var dataObj interface{}

	var result string
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".GetRandomUserAgentAlgorithm", dataObj, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 加密
func CallAlgorithmEncryption(password string) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	var result string
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".EncryptionAlgorithm", password, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 获取随机正态分布数字，传入偏移量、下界和上界
func CallAlgorithmGetNormalDistributionNumber(offset, minimum, maximum int) (int, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return 0, err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return 0, err
	}

	// 创建数据对象
	dataObj := data.AlgorithmNormalDistributeInterface{
		Offset:  offset,
		Minimum: minimum,
		Maximum: maximum,
	}

	var result int
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".GetNormalDistributionNumberAlgorithm", dataObj, &result)
	// 错误处理
	if err != nil {
		return 0, err
	}

	return result, nil
}

// 获取随机字符串
func CallAlgorithmGenerateRandStr(length int) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AlgorithmServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialAlgorithmService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	var result string
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".GenerateRandStrAlgorithm", length, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}
