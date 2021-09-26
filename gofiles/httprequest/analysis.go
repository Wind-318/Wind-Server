package httprequest

import (
	"Project/gofiles/config"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// 获取 robots 文件 disallow 内容
var Disallow []string

// 得到正则表达式分析的页面结果
func RegexpHtml(url string, regexpRule string) []string {
	// 如果是 disallow 的内容，则返回空
	if IsDisallow(url) {
		return nil
	}
	html := GetRequestByte(url)

	obj := regexp.MustCompile(regexpRule)
	arr := obj.FindAllStringSubmatch(string(html), -1)
	ret := make([]string, 0)
	for _, strs := range arr {
		ret = append(ret, strs[1])
	}
	return ret
}

// 分析 robots 文件
func analysisRobotsTxt(url string) []string {
	// 获得 robots.txt 文件内容
	html := helpToGetFirstHtml(url)

	// 查找 Disallow 的内容
	obj := regexp.MustCompile(`Disallow: ([\s\S]+?)(\n)`)
	arr := obj.FindAllStringSubmatch(string(html), -1)

	ret := make([]string, 0)
	for _, i := range arr {
		ret = append(ret, i[1])
	}

	return ret
}

// 判断是否允许访问
func IsDisallow(url string) bool {
	if Disallow == nil {
		Disallow = analysisRobotsTxt("https://finance.sina.com.cn/robots.txt")
	}
	for _, data := range Disallow {
		obj := regexp.MustCompile(data)
		ret := obj.FindAllString(url, -1)
		if len(ret) != 0 {
			return true
		}
	}
	return false
}

func helpToGetFirstHtml(url string) []byte {
	getHtmlClient := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	// 随机设置一个 useragent
	req.Header.Set("User-Agent", config.GetRandomUserAgent())

	res, err := getHtmlClient.Do(req)
	if err != nil {
		return nil
	}

	defer res.Body.Close()
	if err != nil {
		return nil
	} else if res.StatusCode == 404 {
		return nil
	}
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	return html
}
