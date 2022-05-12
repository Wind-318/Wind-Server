package mainPackage

import (
	"spider/sina"
)

// AnimeSpiderInterface 的接口实现
type SinaSpiderService struct{}

// 调用函数，返回数据
func (c *SinaSpiderService) SelectNews(n int, reply *string) error {
	temp, err := sina.SelectNews(n)
	*reply = temp
	return err
}

// 获取所有支持的操作
func (c *SinaSpiderService) GenerateText(request string, reply *string) error {
	temp, err := sina.GenerateText()
	*reply = temp
	return err
}

// 搜索操作
func (c *SinaSpiderService) Search(request string, reply *[][]string) error {
	temp, err := sina.Search()
	*reply = temp
	return err
}
