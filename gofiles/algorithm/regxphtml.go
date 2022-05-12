package algorithm

import (
	"crypto/rand"
	"io/ioutil"
	"math/big"
	"net/http"
	"regexp"
)

// User-Agent 列表，每次随机选择一个使用
var userAgents []string = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) ",
	"Chrome/45.0.2454.85 Safari/537.36 115Browser/6.0.3",
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
	"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SE 2.X MetaSr 1.0; SE 2.X MetaSr 1.0; .NET CLR 2.0.50727; SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0",
	"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1"}

// 随机选取
func GetRandomUserAgent() string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(userAgents))))
	if err != nil {
		return ""
	}
	return userAgents[n.Int64()]
}

// 通过 GET 请求获取请求 url 的 html 页面
func GetRequestByte(url string) ([]byte, error) {
	// 新建一个 client
	getHtmlClient := &http.Client{}

	// 创建 request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 随机设置一个 useragent
	req.Header.Set("User-Agent", GetRandomUserAgent())

	// 请求
	res, err := getHtmlClient.Do(req)
	// 错误处理
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 读取内容
	html, err := ioutil.ReadAll(res.Body)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 返回
	return html, nil
}

// 得到正则表达式分析的 url 页面结果
func RegexpUrlHtml(url string, regexpRule string) ([][]string, error) {
	// 获取内容
	html, err := GetRequestByte(url)
	if err != nil {
		return nil, err
	}

	// 新建对象
	obj := regexp.MustCompile(regexpRule)
	// 解析
	arr := obj.FindAllStringSubmatch(string(html), -1)
	return arr, nil
}

// 得到正则表达式分析的结果
func RegexpHtml(html string, regexpRule string) ([][]string, error) {
	// 新建对象
	obj := regexp.MustCompile(regexpRule)
	// 解析
	arr := obj.FindAllStringSubmatch(string(html), -1)
	return arr, nil
}
