package initCode

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callSpider"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// 初始化拉取动漫
func InitAnime() {
	// go Yhdm2Catch()
	go FcdmCatch()
	YhdmCatch()
}

// 樱花动漫 1 号线
func YhdmCatch() {
	rand.Seed(time.Now().UnixNano())
	os.MkdirAll("./picture/anime/yhdm", 0644)
	// 匹配规则
	regexpRule := `<a href="(/showp[\s\S]+?)">[\s\S]+?<img referrerpolicy="no-referrer" src="([\s\S]+?)" alt="([\s\S]+?)">[\s\S]+?<p>([\s\S]+?)</p>`
	urls, filenames, err := callSpider.CallAnimeSpider(2000, 3, 1, 4, 2, regexpRule, "https://www.yhdmp.cc/list/?year=", "yhdm", "https://www.yhdmp.cc", "樱花动漫", "http:", "catchAnime")
	if err != nil {
		fmt.Println(err)
		return
	}
	for index := range urls {
		// 保存图片
		bytes, err := callAlgorithm.CallAlgorithmGetBytes(urls[index])
		if err != nil {
			continue
		} else if index >= len(filenames) {
			break
		}
		// 文件名
		picName := filenames[index]
		// 写入文件
		ioutil.WriteFile("./picture/anime/yhdm/"+picName+".jpg", bytes, 0644)
		time.Sleep(time.Duration(rand.Intn(54)+62) * time.Second)
	}
}

// 樱花动漫 2 号线
func Yhdm2Catch() {
	rand.Seed(time.Now().UnixNano())
	os.MkdirAll("./picture/anime/yhdm2", 0644)
	// 匹配规则
	regexpRule := `<li><a href="(/view/[\S]+?.html)"[\s\S]+?<img src="([\s\S]+?)" alt="([\s\S]+?)"[\s\S]+?<p>([\s\S]+?)</p>`
	urls, filenames, err := callSpider.CallAnimeSpider(1926, 3, 1, 4, 2, regexpRule, "http://www.imomoe.live/so.asp?page=1&fl=0&pl=time&nf=", "yhdm2", "http://www.imomoe.live", "樱花动漫 2", "https:", "catchAnime")
	if err != nil {
		fmt.Println(err)
		return
	}
	for index := range urls {
		// 保存图片
		bytes, err := callAlgorithm.CallAlgorithmGetBytes(urls[index])
		if err != nil {
			continue
		} else if index >= len(filenames) {
			break
		}
		// 文件名
		picName := filenames[index]
		// 写入文件
		ioutil.WriteFile("./picture/anime/yhdm2/"+picName+".jpg", bytes, 0644)
		time.Sleep(time.Duration(rand.Intn(54)+62) * time.Second)
	}
}

// 风车动漫线路
func FcdmCatch() {
	rand.Seed(time.Now().UnixNano())
	os.MkdirAll("./picture/anime/fcdm", 0644)
	// 匹配规则
	regexpRule := `<a href="(/showp[\s\S]+?)">[\s\S]+?<img referrerpolicy="no-referrer" src="([\s\S]+?)" alt="([\s\S]+?)">[\s\S]+?<p>([\s\S]+?)</p>`
	urls, filenames, err := callSpider.CallAnimeSpider(2000, 3, 1, 4, 2, regexpRule, "https://www.dm530p.net/list/?year=", "fcdm", "https://www.dm530p.net", "风车动漫", "https:", "catchAnime")
	if err != nil {
		fmt.Println(err)
		return
	}
	for index := range urls {
		// 保存图片
		bytes, err := callAlgorithm.CallAlgorithmGetBytes(urls[index])
		if err != nil {
			continue
		} else if index >= len(filenames) {
			break
		}
		// 文件名
		picName := filenames[index]
		// 写入文件
		ioutil.WriteFile("./picture/anime/fcdm/"+picName+".jpg", bytes, 0644)
		time.Sleep(time.Duration(rand.Intn(54)+62) * time.Second)
	}
}
