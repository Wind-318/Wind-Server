package callSpider

import (
	"Project/callClient/data"
	"Project/serviceSearch"
	"net/rpc"

	"github.com/gin-gonic/gin"
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

// 开始初始化过程
func (c *SpiderClient) StartGetAnimeDatabase(request data.AnimeSpiderDataInterface, reply *data.GetAnimeData) error {
	return c.Client.Call(data.AnimeServiceName+".StartGetAnime", request, reply)
}

// 获取所有支持的运算
func (c *SpiderClient) GetOperatorsForAnime(request struct{}, reply *[]string) error {
	return c.Client.Call(data.AnimeServiceName+".GetOperatorsForAnime", request, reply)
}

// 选取新闻
func (c *SpiderClient) SelectNews(request int, reply *string) error {
	return c.Client.Call(data.SinaServiceName+".SelectNews", request, reply)
}

// 调用微服务，传入 startYear：开始年份；rules：匹配规则；后几条为信息在数组中的下标；filename、webPre、source 分别为创建存储图片文件名和网站前缀和来源
func CallAnimeSpider(startYear, name, url, description, picurl int, rules, sourceurl, filename, webPre, source, picPre, operation string) ([]string, []string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AnimeServiceName)
	if err != nil {
		return nil, nil, err
	}
	// 拨号
	client, err := dialSpiderService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, nil, err
	}

	// 创建数据对象
	dataObj := data.AnimeSpiderDataInterface{
		StartYear:   startYear,
		Name:        name,
		Url:         url,
		Description: description,
		Picurl:      picurl,
		Rules:       rules,
		Sourceurl:   sourceurl,
		Filename:    filename,
		WebPre:      webPre,
		Source:      source,
		PicPre:      picPre,
		Operation:   operation,
	}

	var result data.GetAnimeData
	// 调用
	err = client.StartGetAnimeDatabase(dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, nil, err
	}

	return result.Urls, result.Filenames, nil
}

// 搜索
func CallAnimeSearchAnime(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AnimeServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialSpiderService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return nil, err
	}
	obj := data.AnimeData{
		Cookie: cookie,
		Text:   ctx.PostForm("text"),
	}
	reply := data.AnimeReplyData{}
	// 调用
	err = client.Call(data.AnimeServiceName+".SearchAnime", obj, &reply)

	// 错误处理
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"count": reply.Ids,
	}
	for index := range reply.Ids {
		result[reply.Ids[index]] = reply.Info[index]
	}

	return result, nil
}

// 搜索新番
func CallAnimeSearchNewAnime(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AnimeServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialSpiderService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return nil, err
	}
	obj := data.AnimeData{
		Cookie: cookie,
	}
	reply := data.AnimeReplyData{}
	// 调用
	err = client.Client.Call(data.AnimeServiceName+".SearchNewAnime", obj, &reply)

	// 错误处理
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"count": reply.Ids,
	}
	for index := range reply.Ids {
		result[reply.Ids[index]] = reply.Info[index]
	}

	return result, nil
}

// 选出指定年份番剧
func CallAnimeSearchByYear(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AnimeServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialSpiderService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return nil, err
	}
	obj := data.AnimeData{
		Cookie: cookie,
		Year:   ctx.PostForm("year"),
	}
	reply := data.AnimeReplyData{}
	// 调用
	err = client.Call(data.AnimeServiceName+".SearchByYear", obj, &reply)

	// 错误处理
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"count": reply.Ids,
	}
	for index := range reply.Ids {
		result[reply.Ids[index]] = reply.Info[index]
	}

	return result, nil
}

// 樱花动漫一号线追踪策略
func CallAnimeTrackUpdateYhdm1() ([]string, []string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.AnimeServiceName)
	if err != nil {
		return nil, nil, err
	}
	// 拨号
	client, err := dialSpiderService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, nil, err
	}

	result := data.GetAnimeData{}
	// 调用
	err = client.Call(data.AnimeServiceName+".TrackUpdateYhdm1", "", &result)

	// 错误处理
	if err != nil {
		return nil, nil, err
	}

	return result.Urls, result.Filenames, nil
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
	err = client.Client.Call(data.SinaServiceName+".GenerateText", "", &result)
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
	err = client.Client.Call(data.SinaServiceName+".Search", "", &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}
