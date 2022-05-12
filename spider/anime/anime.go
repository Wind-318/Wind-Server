package anime

import (
	"spider/callClient/callAlgorithm"
	"spider/callClient/callDatabase"
	"spider/config"
	"strconv"
	"time"
)

// 抓取模板   startYear：开始年份；rules：匹配规则；后几条为信息在数组中的下标；filename、webPre、source 分别为创建存储图片文件名和网站前缀和来源，返回封面图 url 和文件名
func CatchAnime(startYear, name, url, description, picurl int, rules, sourceurl, filename, webPre, source, picPre string) ([]string, []string, error) {
	urls, filenames := []string{}, []string{}

	// 源 url
	sourceUrl := sourceurl
	// 从 startYear 年开始到当前年份
	for i := startYear; i <= time.Now().Year(); i++ {
		for j := 1; ; j++ {
			// 随机等待一段时间
			waitTime, err := callAlgorithm.CallAlgorithmRandomTime(54, 62)
			if err != nil {
				return urls, filenames, err
			}
			time.Sleep(time.Duration(waitTime) * time.Second)
			// 匹配规则
			rule := rules
			// 获取信息
			arr, err := callAlgorithm.CallAlgorithmUrlRegexp(sourceUrl+strconv.Itoa(i)+"&pagesize=24&pageindex="+strconv.Itoa(j), rule)
			// 错误则返回
			if err != nil {
				return urls, filenames, err
			} else if len(arr) == 0 {
				break
			}
			// 循环
			for index := range arr {
				// 重复的跳过
				isExist, err := callDatabase.CallMySQLSelectAnimeCountByNameAndSource(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", arr[index][name], source)
				if err == nil && isExist {
					continue
				}

				// 收录封面图
				urls = append(urls, picPre+arr[index][picurl])
				// 文件名
				picName := strconv.Itoa(int(time.Now().UnixNano()))
				filenames = append(filenames, picName)

				// 检索后动漫名称如已存在，则更新资料；如不存在则插入
				isExist, err = callDatabase.CallMySQLSelectAnimeCountByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", arr[index][name])
				if err != nil {
					isExist = false
				}
				// 已存在，更新数据
				if isExist {
					err = callDatabase.CallMySQLUpdateAnimeInfo(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", webPre+arr[index][url], arr[index][description], config.YourName+"picture/anime/"+filename+"/"+picName+".jpg", arr[index][name], 0)
					if err != nil {
						continue
					}
				} else { // 不存在，插入
					err = callDatabase.CallMySQLInsertAnimeForBangumi(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", arr[index][name], webPre+arr[index][url], arr[index][description], config.YourName+"picture/anime/"+filename+"/"+picName+".jpg", i, 0)
					if err != nil {
						continue
					}
				}
				// 来源中插入
				callDatabase.CallMySQLInsertAnimeForSource(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", arr[index][name], source, webPre+arr[index][url])
				if err != nil {
					continue
				}
			}
		}
	}

	return urls, filenames, nil
}
