package anime

import (
	"Project/gofiles/config"
	"Project/gofiles/initCode"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 最后一次更新的月份
var lastestMonth = int(time.Now().Month())

// 更新策略：若当前时间为 1 月、4 月、7 月和 10 月，则每三天进行一次追踪更新，其余月份则每七天进行一次追踪更新
func ContinueGetNewAnime() {
	// 一直轮询
	for {
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

		// 追踪新番
		trackUpdates()
		// 抓取片源一地址
		initCode.YhdmCatch()
	}
}

// 追踪更新新番
func trackUpdates() {
	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	// 当前年
	nowYear := time.Now().Year()
	// 当前月
	nowMonth := time.Now().Month()
	// 如果最后更新月份处于临界月则清除数据库中新番 isNew 状态（临界月：最后更新月份为 3 月而当前月为 4 月，当前月为最后更新月的下一个季度）
	if (lastestMonth == 12 && nowMonth == 1) || (lastestMonth == 3 && nowMonth == 4) || (lastestMonth == 6 && nowMonth == 7) || (lastestMonth == 9 && nowMonth == 10) {
		conn.Exec("UPDATE bangumi SET isNew = 0")
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
	for index := 1; ; index++ {
		// 请求获取页面字符串
		res, err := http.Get(url + strconv.Itoa(index))
		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer res.Body.Close()
		// 读取
		html, err := ioutil.ReadAll(res.Body)
		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
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
			isExist := 0
			conn.Get(&isExist, "SELECT count(*) FROM animesource WHERE anime = ? AND source = ?", data[3], "樱花动漫")
			if isExist > 0 {
				continue
			}

			// 保存图片
			if data[2][:4] != "http" {
				data[2] = "http:" + data[2]
			}
			// 休息数秒
			time.Sleep(10 * time.Second)
			picFile, err := http.Get(data[2])
			// 错误处理
			if err != nil {
				fmt.Println(err)
				continue
			}
			// 图片数据
			picByte, err := ioutil.ReadAll(picFile.Body)
			// 随机起名
			picName := strconv.Itoa(int(time.Now().UnixNano()))
			// 无错误则保存到本地
			if err == nil {
				ioutil.WriteFile("./picture/anime/yhdm/"+picName+".jpg", picByte, 0644)
			}

			// 检索后动漫名称如已存在，则更新资料；如不存在则插入
			isExist = 0
			conn.Get(&isExist, "SELECT count(*) FROM bangumi WHERE name = ?", data[3])
			if isExist > 0 {
				// 更新资料
				conn.Exec("UPDATE bangumi SET url = ?, description = ?, picurl = ? WHERE name = ?", "https://www.yhdmp.cc"+data[1], data[4], "../picture/anime/yhdm/"+picName+".jpg", data[3])
			} else {
				// 新插入
				conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, data[3], "https://www.yhdmp.cc"+data[1], strconv.Itoa(nowYear), data[4], "../picture/anime/yhdm/"+picName+".jpg", 1)
			}

			// 来源地址插入
			conn.Exec("INSERT INTO animesource VALUES(?, ?, ?, ?)", 0, data[3], "樱花动漫", "https://www.yhdmp.cc"+data[1])
		}
		// 休眠 45 秒
		time.Sleep(45 * time.Second)
	}
}
