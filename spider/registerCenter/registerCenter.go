package registerCenter

import "spider/callClient/data"

// AnimeSpider 服务名
var AnimeServiceName = "AnimeSpiderService"

// sina
var SinaServiceName = "SinaSpiderService"

// sinaSpider 服务接口
type SinaSpiderServiceInterface interface {
	// 选取新闻
	SelectNews(n int, reply *string) error
	// 生成文章
	GenerateText(request string, reply *string) error
	// 搜索
	Search(request string, reply *[][]string) error
}

// AnimeSpider 服务接口
type AnimeSpiderServiceInterface interface {
	// 搜索功能
	SearchAnime(request AnimeData, reply *AnimeReplyData) error
	// 获取动漫
	StartGetAnime(request AnimeSpiderDataInterface, reply *GetAnimeData) error
	// 搜索新番
	SearchNewAnime(request AnimeData, reply *AnimeReplyData) error
	// 选出指定年份番剧
	SearchByYear(request AnimeData, reply *AnimeReplyData) error
	// 樱花动漫一号线追踪策略
	TrackUpdateYhdm1(request string, reply *GetAnimeData) error
	// 获取所有操作
	GetOperatorsForAnime(request interface{}, reply *[]string) error
}

// 返回值
type AnimeReplyData struct {
	// id 集合
	Ids []string
	// 信息
	Info []data.SelectAnimeDataInterface
}

type AnimeData struct {
	// cookie
	Cookie string
	// 年份
	Year string
	// 查询字段
	Text string
}

type GetAnimeData struct {
	Urls      []string
	Filenames []string
}

// AnimeSpider 数据
type AnimeSpiderDataInterface struct {
	// 开始年份
	StartYear int
	// 名称对应的下标
	Name int
	// url 对应的下标
	Url int
	// 简介对应下标
	Description int
	// 图片对应的下标
	Picurl int
	// 正则匹配规则
	Rules string
	// 源 url 前缀（去除年份）
	Sourceurl string
	// 文件名（英文缩写）
	Filename string
	// 访问 url 前缀
	WebPre string
	// 来源网站（如：樱花动漫）
	Source string
	// 图片 url 前缀
	PicPre string
	// 进行的操作
	Operation string
}
