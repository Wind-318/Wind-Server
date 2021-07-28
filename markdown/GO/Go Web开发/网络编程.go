package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func get(urls string) {
	res, err := http.Get(urls)
	defer res.Body.Close()
	if err != nil {
		// 处理错误
		return
	}

	strs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// 处理错误
		return
	}
	fmt.Println(string(strs))
}

func post(urls string) {
	val := &url.Values{
		"name": {"213", "asd"},
	}
	body := val.Encode()
	res, err := http.Post(urls, "text/html", strings.NewReader(body))
	defer res.Body.Close()
	if err != nil {
		// 处理错误
		return
	}

	strs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// 处理错误
		return
	}
	fmt.Println(string(strs))
}

func own(urls string) {
	clients := &http.Client{
		// 设置超时时间为 3 秒
		Timeout: 3 * time.Second,
	}
	req, _ := http.NewRequest("GET", urls, nil)
	req.Header.Set("name", "123")
	res, _ := clients.Do(req)
	defer res.Body.Close()
	strs, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(strs))
}

func main() {
	get("https://cn.bing.com/")
	post("https://cn.bing.com/")
	own("https://cn.bing.com/")
}
