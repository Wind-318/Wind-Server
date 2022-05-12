package mainPackage

import (
	"errors"
	"spider/anime"
	"spider/registerCenter"
)

// AnimeSpiderInterface 的接口实现
type AnimeSpiderService struct{}

// startYear：开始年份；rules：匹配规则；后几条为信息在数组中的下标；filename、webPre、source 分别为创建存储图片文件名和网站前缀和来源
type CatchAnimesFunc func(startYear, name, url, description, picurl int, rules, sourceurl, filename, webPre, source, picPre string) ([]string, []string, error)

// startYear：开始年份；rules：匹配规则；后几条为信息在数组中的下标；filename、webPre、source 分别为创建存储图片文件名和网站前缀和来源
func CatchAnimes(startYear, name, url, description, picurl int, rules, sourceurl, filename, webPre, source, picPre string) ([]string, []string, error) {
	// 调用
	urls, filenames, err := anime.CatchAnime(startYear, name, url, description, picurl, rules, sourceurl, filename, webPre, source, picPre)
	// 错误处理
	if err != nil {
		return nil, nil, err
	}
	return urls, filenames, nil
}

// 注册所有支持的运算
var AnimeOperators = map[string]CatchAnimesFunc{
	"catchAnime": CatchAnimes,
}

// 获取具体函数
func CreateAnimeOperation(operator string) (CatchAnimesFunc, error) {
	var oper CatchAnimesFunc
	if oper, ok := AnimeOperators[operator]; ok {
		return oper, nil
	}
	return oper, errors.New("illegal Operator")
}

// 调用函数，返回数据
func (c *AnimeSpiderService) StartGetAnime(request registerCenter.AnimeSpiderDataInterface, reply *registerCenter.GetAnimeData) error {
	// 获取具体函数
	oper, err := CreateAnimeOperation(request.Operation)
	if err != nil {
		return err
	}
	// 获取返回值
	urls, filnames, err := oper(request.StartYear, request.Name, request.Url, request.Description, request.Picurl, request.Rules, request.Sourceurl, request.Filename, request.WebPre, request.Source, request.PicPre)
	if err != nil {
		return err
	}
	reply.Urls = urls
	reply.Filenames = filnames
	return nil
}

// 搜索功能
func (c *AnimeSpiderService) SearchAnime(request registerCenter.AnimeData, reply *registerCenter.AnimeReplyData) error {
	(*reply).Ids, (*reply).Info = anime.Search(request)
	return nil
}

// 搜索新番
func (c *AnimeSpiderService) SearchNewAnime(request registerCenter.AnimeData, reply *registerCenter.AnimeReplyData) error {
	(*reply).Ids, (*reply).Info = anime.SearchNewAnime(request)
	return nil
}

// 选出指定年份番剧
func (c *AnimeSpiderService) SearchByYear(request registerCenter.AnimeData, reply *registerCenter.AnimeReplyData) error {
	(*reply).Ids, (*reply).Info = anime.SearchByYear(request)
	return nil
}

// 樱花动漫一号线追踪策略
func (c *AnimeSpiderService) TrackUpdateYhdm1(request string, reply *registerCenter.GetAnimeData) error {
	temp1, temp2, err := anime.TrackUpdateYhdm1()
	(*reply).Urls, (*reply).Filenames = temp1, temp2
	return err
}

// 获取所有支持的操作
func (c *AnimeSpiderService) GetOperatorsForAnime(request interface{}, reply *[]string) error {
	arr := make([]string, 0, len(AnimeOperators))
	for key := range AnimeOperators {
		arr = append(arr, key)
	}
	*reply = arr
	return nil
}
