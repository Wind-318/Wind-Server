package callSpider

import (
	"mail/callClient/data"
	"mail/serviceSearch"
	"net/rpc"
)

// 操作客户端
type SpiderClient struct {
	*rpc.Client
}

// 拨号
func dialSpiderService(network, address string) (*SpiderClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &SpiderClient{Client: client}, nil
}

// 选取新闻
func (c *SpiderClient) SelectNews(request int, reply *string) error {
	return c.Client.Call(data.SinaServiceName+".SelectNews", request, reply)
}

// 生成文章
func (c *SpiderClient) GenerateText(request interface{}, reply *string) error {
	return c.Client.Call(data.SinaServiceName+".GenerateText", request, reply)
}

// 搜索
func (c *SpiderClient) Search(request interface{}, reply *[][]string) error {
	return c.Client.Call(data.SinaServiceName+".Search", request, reply)
}

// 调用服务，传入 n，选择前 n 条
func CallSelectNews(n int) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.SinaServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialSpiderService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	var result string
	// 调用
	err = client.SelectNews(n, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 调用服务，获取
func CallGenerateText() (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.SinaServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialSpiderService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	var result string
	// 调用
	err = client.GenerateText("", &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 调用服务，传入 n，选择前 n 条
func CallSearch() ([][]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.SinaServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialSpiderService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	var result [][]string
	// 调用
	err = client.Search("", &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}
