package httpRequest

import (
	"io/ioutil"
	"net/http"
	"time"
)

// 通过 GET 请求获取请求 url 的 html 页面
func GetRequestByte(url string) []byte {
	if IsDisallow(url) {
		return nil
	}
	getHtmlClient := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	// 随即设置一个 useragent
	req.Header.Set("User-Agent", GetRandomUserAgent())

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
