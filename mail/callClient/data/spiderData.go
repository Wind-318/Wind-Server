package data

// AnimeSpider 服务名
var AnimeServiceName = "AnimeSpiderService"

// sina
var SinaServiceName = "SinaSpiderService"

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
