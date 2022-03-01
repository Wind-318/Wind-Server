package initCode

import (
	"Project/gofiles/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 从 bangumi 获取基本信息
func InitAnime() {
	bangumi := "https://bgm.tv/anime/browser/?sort=date&page="
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	for i := 1; ; i++ {
		res, err := http.Get(bangumi + strconv.Itoa(i))
		if err != nil {
			return
		}
		defer res.Body.Close()
		html, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}

		curPage := string(html)
		obj1 := regexp.MustCompile(`<li id="item[\s\S]+?<a href="(/subject[\S]+?)" class="subjectCover cover ll[\s\S]+?<img src="([\s\S]+?)" class="cover"[\s\S]+?<a href="/subject/[\S]+?" class="l">([\s\S]+?)</a>[\s\S]+?<p class="info tip">([\s\S]+?)</p>`)
		obj2 := regexp.MustCompile(`<li id="item[\s\S]+?<a href="(/subject[\S]+?)" class="subjectCover cover ll[\s\S]+?<a href="/subject/[\S]+?" class="l">([\s\S]+?)</a>[\s\S]+?<p class="info tip">([\s\S]+?)</p>`)
		ans1 := obj1.FindAllStringSubmatch(curPage, -1)
		ans2 := obj2.FindAllStringSubmatch(curPage, -1)
		if len(ans1)+len(ans2) == 0 {
			break
		}
		// 将基本信息中的日期提取出来并插入数据库
		for index := range ans1 {
			// 已存在则跳过
			isExist := 0
			conn.Get(&isExist, "SELECT COUNT(*) FROM bangumi WHERE name = ?", ans1[index][3])
			if isExist > 0 {
				continue
			}
			// 不存在则加入
			tempObj := regexp.MustCompile(`\d{4}`)
			arr := tempObj.FindAllStringSubmatch(ans1[index][4], -1)
			if len(arr) == 0 {
				continue
			}
			ans1[index][4] = arr[0][0]

			ans1[index][1] = "https://bgm.tv" + ans1[index][1]
			ans1[index][2] = "https:" + ans1[index][2]

			_, err := conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, ans1[index][3], ans1[index][1], ans1[index][4], "", ans1[index][2], 0)
			if err != nil {
				fmt.Println(err)
			}
			_, err = conn.Exec("INSERT INTO animesource VALUES(?, ?, ?, ?)", 0, ans1[index][3], "bangumi", ans1[index][1])
			if err != nil {
				fmt.Println(err)
			}
		}
		for index := range ans2 {
			// 已存在则跳过
			isExist := 0
			conn.Get(&isExist, "SELECT COUNT(*) FROM bangumi WHERE name = ?", ans2[index][2])
			if isExist > 0 {
				continue
			}
			// 不存在则加入
			tempObj := regexp.MustCompile(`\d{4}`)
			arr := tempObj.FindAllStringSubmatch(ans2[index][3], -1)
			if len(arr) == 0 {
				continue
			}

			ans2[index][3] = arr[0][0]
			ans2[index][1] = "https://bgm.tv" + ans2[index][1]

			_, err := conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, ans2[index][2], ans2[index][1], ans2[index][3], "", "", 0)
			if err != nil {
				fmt.Println(err)
			}
			_, err = conn.Exec("INSERT INTO animesource VALUES(?, ?, ?, ?)", 0, ans2[index][2], "bangumi", ans2[index][1])
			if err != nil {
				fmt.Println(err)
			}
		}
		time.Sleep(5000 * time.Millisecond)
	}

	ysjdmCatch()
	yhdmCatch()
}

// 更新策略：若当前时间为 1 月、4 月、7 月和 10 月，则每天进行一次追踪更新，其余月份则每七天进行一次追踪更新
func ContinueGetNewAnime() {
	everyUpdateDays := 7
	if nowMonth := time.Now().Month(); nowMonth == 1 || nowMonth == 4 || nowMonth == 7 || nowMonth == 10 {
		everyUpdateDays = 1
	}

}

// 樱花动漫片源地址
func yhdmCatch() {
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	yhdmAnime := "https://www.yhdmp.cc/list/?year="
	for i := 2000; i <= time.Now().Year()+1; i++ {
		for j := 0; ; j++ {
			res, err := http.Get(yhdmAnime + strconv.Itoa(i) + "&pagesize=24&pageindex=" + strconv.Itoa(j))
			if err != nil {
				continue
			}
			html, err := ioutil.ReadAll(res.Body)
			if err != nil {
				continue
			}
			obj := regexp.MustCompile(`<a href="(/showp[\s\S]+?)">[\s\S]+?<img referrerpolicy="no-referrer" src="([\s\S]+?)" alt="([\s\S]+?)">[\s\S]+?<p>([\s\S]+?)</p>`)
			arr := obj.FindAllStringSubmatch(string(html), -1)
			if len(arr) == 0 {
				break
			}
			for _, data := range arr {
				// 重复的跳过
				isExist := 0
				conn.Get(&isExist, "SELECT count(*) FROM animesource WHERE anime = ? AND source = ?", data[3], "樱花动漫")
				if isExist > 0 {
					continue
				}

				// 检索后动漫名称如已存在，则更新资料；如不存在则插入
				isExist = 0
				conn.Get(&isExist, "SELECT count(*) FROM bangumi WHERE name = ?", data[3])
				if isExist > 0 {
					conn.Exec("UPDATE bangumi SET url = ?, description = ?, picurl = ? WHERE name = ?", "https://www.yhdmp.cc"+data[1], data[4], data[2], data[3])
				} else {
					conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, data[3], "https://www.yhdmp.cc"+data[1], strconv.Itoa(i), data[4], data[2], 0)
				}

				conn.Exec("INSERT INTO animesource VALUES(?, ?, ?, ?)", 0, data[3], "樱花动漫", "https://www.yhdmp.cc"+data[1])
			}

			res.Body.Close()
			time.Sleep(5000 * time.Millisecond)
		}
	}
}

// 异世界动漫片源地址
func ysjdmCatch() {
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	ysjdmAnime := "http://121.4.190.96:9991/getsortdata_all_z.php?action=acg&page="
	for i := 1; ; i++ {
		res, err := http.Get(ysjdmAnime + strconv.Itoa(i) + "&year=0&area=all&class=0&dect=&id=")
		if err != nil {
			continue
		}
		html, err := ioutil.ReadAll(res.Body)
		if err != nil {
			continue
		}
		text := string(html)
		obj := regexp.MustCompile(`<a class="li-hv" href="([\s\S]+?)" title="([\s\S]+?)">[\s\S]+?data-original="([\s\S]+?)"`)
		arr := obj.FindAllStringSubmatch(text, -1)
		if len(arr) == 0 {
			break
		}
		for _, data := range arr {
			// 重复的跳过
			isExist := 0
			conn.Get(&isExist, "SELECT count(*) FROM animesource WHERE anime = ? AND source = ?", data[2], "异世界动漫 1")
			if isExist > 0 {
				continue
			}

			time.Sleep(time.Millisecond * 5000)

			// 检查是否有简介可更新
			tempRes, err := http.Get("http://ysjdm8.com" + data[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			defer tempRes.Body.Close()
			description, err := ioutil.ReadAll(tempRes.Body)
			if err != nil {
				fmt.Println(err)
				continue
			}
			tempObj := regexp.MustCompile(`<div class="des2"><b>([\s\S]+?)</div>`)
			tempArr := tempObj.FindAllStringSubmatch(string(description), -1)

			// 检索后动漫名称如已存在，则更新资料；如不存在则插入
			isExist = 0
			conn.Get(&isExist, "SELECT count(*) FROM bangumi WHERE name = ?", data[2])
			if isExist > 0 {
				if len(tempArr) > 0 && len(tempArr[0]) > 1 {
					conn.Exec("UPDATE bangumi SET url = ?, description = ?, picurl = ? WHERE name = ?", "http://ysjdm8.com"+data[1], tempArr[0][1], data[3], data[2])
				} else {
					conn.Exec("UPDATE bangumi SET url = ?, picurl = ? WHERE name = ?", "http://ysjdm8.com"+data[1], data[3], data[2])
				}
			} else {
				yearObj := regexp.MustCompile(`<b>年代：</b>(\d{4})</dd>`)
				yearArr := yearObj.FindAllStringSubmatch(string(description), -1)
				animeYear := "2000"
				if len(yearArr) > 0 && len(yearArr[0]) > 1 {
					animeYear = yearArr[0][1]
				}

				if len(tempArr) > 0 && len(tempArr[0]) > 1 {
					conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, data[2], "http://ysjdm8.com"+data[1], animeYear, tempArr[0][1], data[3], 0)
				} else {
					conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, data[2], "http://ysjdm8.com"+data[1], animeYear, "", data[3], 0)
				}
			}

			conn.Exec("INSERT INTO animesource VALUES(?, ?, ?, ?)", 0, data[2], "异世界动漫 1", "http://ysjdm8.com"+data[1])
		}

		res.Body.Close()
		time.Sleep(5000 * time.Millisecond)
	}
}
