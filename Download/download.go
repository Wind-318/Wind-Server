package Download

import (
	"Project/httpRequest"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"sync"
)

// 下载资源
func DownloadGet(url string, position ...string) {
	if len(position) > 1 {
		panic("too many args")
	}
	// 如果没有传入参数，则设置一个默认位置下载文件
	pos := "./files/" + strconv.Itoa(rand.Intn(100000))
	if len(position) == 1 {
		// 有传入位置则使用传入的位置
		pos = position[0]
	}

	// Windows 下文件名称不能包含以下字符，去除不合法字符
	obj := regexp.MustCompile(`[/\:*?"<>|]`)
	pos = obj.ReplaceAllString(pos, "")

	// 正则表达式搜索文件所需文件夹是否存在
	obj = regexp.MustCompile(`([\s\S]+)\\`)
	fileroot := obj.FindAllStringSubmatch(pos, -1)[0][1]

	// 递归创建所有所需文件夹
	os.MkdirAll(fileroot, 0644)

	data := httpRequest.GetRequestByte(url)

	err := ioutil.WriteFile(pos, data, 0644)
	if err != nil {
		return
	}
}

// 下载 html 页面中所有资源，参数为 url 地址和 保存的文件位置
func DownloadHtmlSource(url, rule string, position ...string) {
	// 参数检测，传入文件位置超过一个则 panic
	if len(position) > 1 {
		panic("too many args")
	}

	pos := "./files/" + strconv.Itoa(rand.Intn(100000))
	if len(position) == 1 {
		pos = position[0]
	}

	// 资源文件匹配规则
	sourceRule := rule
	sources := httpRequest.RegexpHtml(url, sourceRule)

	var wg sync.WaitGroup
	controlMaxNum := make(chan int, 5)

	// 并发下载资源，最大同时下载 5 条
	for _, urls := range sources {
		wg.Add(1)
		controlMaxNum <- 1
		go func(urls, pos string) {
			DownloadGet(urls, pos+urls)
			wg.Done()
			<-controlMaxNum
		}(urls, pos)
	}
	wg.Wait()
}
