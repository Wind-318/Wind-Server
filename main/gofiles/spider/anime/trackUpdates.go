package anime

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callSpider"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var flag = true

// 更新策略：若当前时间为 1 月、4 月、7 月和 10 月，则每三天进行一次追踪更新，其余月份则每七天进行一次追踪更新
func ContinueGetNewAnime() {
	os.MkdirAll("picture/anime/yhdm", 0644)
	rand.Seed(time.Now().UnixNano())
	// 一直轮询
	for {
		// 第一次进入不等待
		if !flag {
			// 更新时间
			everyUpdateDays := 7
			// 若为 1 月、4 月、7 月和 10 月则每三天更新
			if nowMonth := time.Now().Month(); nowMonth == 1 || nowMonth == 4 || nowMonth == 7 || nowMonth == 10 {
				everyUpdateDays = 3
			}
			// 等待小时数
			timeToWaitHour := 24 * everyUpdateDays
			// 等待轮询
			time.Sleep(time.Hour * time.Duration(timeToWaitHour))
		} else {
			flag = false
		}

		// 追踪新番
		urls, filenames, err := callSpider.CallAnimeTrackUpdateYhdm1()
		if err != nil {
			continue
		}
		// 保存封面图
		for index := range urls {
			// 读取图片内容
			bytes, err := callAlgorithm.CallAlgorithmGetBytes(urls[index])
			if err != nil {
				continue
			}
			// 下标越界 break
			if index >= len(filenames) {
				break
			}
			// 保存
			ioutil.WriteFile("./picture/anime/yhdm/"+filenames[index]+".jpg", bytes, 0644)
			time.Sleep(time.Duration(rand.Intn(64)+59) * time.Second)
		}
	}
}
