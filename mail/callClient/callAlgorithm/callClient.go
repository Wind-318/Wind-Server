package callAlgorithm

import (
	"mail/callClient/data"
	"mail/serviceSearch"
	"net/rpc"
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

// 解析 stream 数据
func CallAlgorithmParseStream(reply interface{}) ([]string, error) {
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

	var result []string
	// 调用
	err = client.Client.Call(data.AlgorithmServiceName+".ParseRedisStream", reply, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}
