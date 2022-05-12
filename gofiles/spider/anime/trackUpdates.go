package anime

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

// 最后一次更新的月份
var lastestMonth = int(time.Now().Month())

// 是否为初次使用
var flag = true

// 更新策略：若当前时间为 1 月、4 月、7 月和 10 月，则每三天进行一次追踪更新，其余月份则每七天进行一次追踪更新
func ContinueGetNewAnime() {
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
		trackUpdateYhdm1()
	}
}

// 追踪更新新番
func trackUpdateYhdm1() {
	rand.Seed(time.Now().UnixNano())
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"spider")
	if err != nil {
		return
	}
	defer mysqlClient.Close()
	// 当前年
	nowYear := time.Now().Year()
	// 当前月
	nowMonth := time.Now().Month()
	// 如果最后更新月份处于临界月则清除数据库中新番 isNew 状态（临界月：最后更新月份为 3 月而当前月为 4 月，当前月为最后更新月的下一个季度）
	if (lastestMonth == 12 && nowMonth == 1) || (lastestMonth == 3 && nowMonth == 4) || (lastestMonth == 6 && nowMonth == 7) || (lastestMonth == 9 && nowMonth == 10) {
		mysqlClient.Exec("UPDATE bangumi SET isNew = 0 WHERE isNew = 1")
	}
	// 更新最后更新月份
	lastestMonth = int(nowMonth)
	if nowMonth >= 10 { // 4 季度
		nowMonth = 10
	} else if nowMonth >= 7 { // 3 季度
		nowMonth = 7
	} else if nowMonth >= 4 { // 2 季度
		nowMonth = 4
	} else { // 1 季度
		nowMonth = 1
	}
	// 更新新片源
	url := "https://www.yhdmp.cc/list/?year=" + strconv.Itoa(nowYear) + "&season=" + strconv.Itoa(int(nowMonth)) + "&pagesize=24&pageindex="
	for index := 0; ; index++ {
		// 随机等待一段时间
		time.Sleep(time.Duration(algorithm.GetRandomNumberTime(54, 62)) * time.Second)

		// 读取
		html, err := algorithm.GetRequestByte(url + strconv.Itoa(index))
		// 错误处理
		if err != nil {
			break
		}
		// 正则匹配新地址
		obj := regexp.MustCompile(`<a href="(/showp[\s\S]+?)">[\s\S]+?<img referrerpolicy="no-referrer" src="([\s\S]+?)" alt="([\s\S]+?)">[\s\S]+?<p>([\s\S]+?)</p>`)
		arr := obj.FindAllStringSubmatch(string(html), -1)
		// 无符合匹配则退出循环
		if len(arr) == 0 {
			break
		}
		for _, data := range arr {
			// 重复则跳过
			cnt := 0
			err = mysqlClient.Get(&cnt, "SELECT count(*) FROM animesource WHERE anime = ? AND source = ?", data[3], "樱花动漫")
			// 若存在，设为新番
			if err == nil && cnt > 0 {
				mysqlClient.Exec("UPDATE bangumi SET isNew = 1 WHERE name = ?", data[3])
				continue
			}

			// 休眠
			time.Sleep(time.Duration(rand.Intn(64)+59) * time.Second)
			// 保存图片
			if data[2][:4] != "http" {
				data[2] = "http:" + data[2]
			}

			// 随机起名
			picName := strconv.Itoa(int(time.Now().UnixNano()))
			bytes, err := algorithm.GetRequestByte(data[2])
			if err != nil {
				continue
			}
			// 保存
			ioutil.WriteFile("./picture/anime/yhdm/"+picName+".jpg", bytes, 0644)

			// 检索后动漫名称如已存在，则更新资料；如不存在则插入
			ans := 0
			isExist := false
			err = mysqlClient.Get(&ans, "SELECT count(*) FROM bangumi WHERE name = ?", data[3])
			if err == nil && ans > 0 {
				isExist = true
			}
			if isExist {
				// 更新资料
				_, err = mysqlClient.Exec("UPDATE bangumi SET url = ?, description = ?, picurl = ?, isNew = ? WHERE name = ?", "https://www.yhdmp.cc"+data[1], data[4], config.Addr+"picture/anime/yhdm/"+picName+".jpg", 1, data[3])
				if err != nil {
					continue
				}
			} else {
				// 新插入
				mysqlClient.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, data[3], "https://www.yhdmp.cc"+data[1], nowYear, data[4], config.Addr+"picture/anime/yhdm/"+picName+".jpg", 1)
				if err != nil {
					continue
				}
			}
			// 来源地址插入
			mysqlClient.Exec("INSERT INTO animesource VALUES(?, ?, ?, ?)", 0, data[3], "樱花动漫", "https://www.yhdmp.cc"+data[1])
		}
	}
}
