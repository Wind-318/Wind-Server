package anime

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"spider/callClient/callAlgorithm"
	"spider/callClient/callDatabase"
	"spider/config"
	"strconv"
	"time"
)

// 最后一次更新的月份
var lastestMonth = int(time.Now().Month())

// 是否为初次使用
var isNewUse = true

// 追踪更新新番
func TrackUpdateYhdm1() ([]string, []string, error) {
	ret := []string{}
	names := []string{}
	// 当前年
	nowYear := time.Now().Year()
	// 当前月
	nowMonth := time.Now().Month()
	// 如果最后更新月份处于临界月则清除数据库中新番 isNew 状态（临界月：最后更新月份为 3 月而当前月为 4 月，当前月为最后更新月的下一个季度）
	if isNewUse || (lastestMonth == 12 && nowMonth == 1) || (lastestMonth == 3 && nowMonth == 4) || (lastestMonth == 6 && nowMonth == 7) || (lastestMonth == 9 && nowMonth == 10) {
		isNewUse = false
		// 新番置旧
		callDatabase.CallMySQLUpdateAnimeSetIsNew0(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider")
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
		waitTime, err := callAlgorithm.CallAlgorithmRandomTime(54, 62)
		if err != nil {
			log.Println(err)
			continue
		}
		time.Sleep(time.Duration(waitTime) * time.Second)
		// 请求获取页面字符串
		res, err := http.Get(url + strconv.Itoa(index))
		// 错误处理
		if err != nil {
			break
		}
		defer res.Body.Close()
		// 读取
		html, err := ioutil.ReadAll(res.Body)
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
			isExist, err := callDatabase.CallMySQLSelectAnimeCountByNameAndSource(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", data[3], "樱花动漫")

			if err == nil && isExist {
				continue
			}

			// 保存图片
			if data[2][:4] != "http" {
				data[2] = "http:" + data[2]
			}
			ret = append(ret, data[2])
			// 随机起名
			picName := strconv.Itoa(int(time.Now().UnixNano()))
			names = append(names, picName)

			// 检索后动漫名称如已存在，则更新资料；如不存在则插入
			isExist, _ = callDatabase.CallMySQLSelectAnimeCountByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", data[3])
			if err != nil {
				isExist = false
			}
			if isExist {
				// 更新资料
				err = callDatabase.CallMySQLUpdateAnimeInfo(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", "https://www.yhdmp.cc"+data[1], data[4], config.YourName+"picture/anime/yhdm/"+picName+".jpg", data[3], 1)
				if err != nil {
					continue
				}
			} else {
				// 新插入
				err = callDatabase.CallMySQLInsertAnimeForBangumi(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", data[3], "https://www.yhdmp.cc"+data[1], data[4], config.YourName+"picture/anime/yhdm/"+picName+".jpg", nowYear, 1)
				if err != nil {
					continue
				}
			}
			// 来源地址插入
			callDatabase.CallMySQLInsertAnimeForSource(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", data[3], "樱花动漫", "https://www.yhdmp.cc"+data[1])
		}
	}

	return ret, names, nil
}
