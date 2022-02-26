package initCode

import (
	"Project/gofiles/config"
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
		defer res.Body.Close()
		if err != nil {
			return
		}
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
		recordName := map[string]bool{}
		for index := range ans1 {
			tempObj := regexp.MustCompile(`\d{4}`)
			arr := tempObj.FindAllStringSubmatch(ans1[index][4], -1)
			if len(arr) == 0 {
				continue
			}
			ans1[index][4] = arr[0][0]

			ans1[index][1] = "https://bgm.tv" + ans1[index][1]
			ans1[index][2] = "https:" + ans1[index][2]
			recordName[ans1[index][3]] = true

			conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, ans1[index][3], ans1[index][1], ans1[index][4], "", ans1[index][2], 0)
		}
		for index := range ans2 {
			if _, ok := recordName[ans2[index][2]]; ok {
				continue
			}
			tempObj := regexp.MustCompile(`\d{4}`)
			arr := tempObj.FindAllStringSubmatch(ans2[index][3], -1)
			if len(arr) == 0 {
				continue
			}

			ans2[index][3] = arr[0][0]
			ans2[index][1] = "https://bgm.tv" + ans2[index][1]

			conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, ans2[index][2], ans2[index][1], ans2[index][3], "", "", 0)
		}
		time.Sleep(2500 * time.Millisecond)
	}
}

// 持续追踪
func ContinueGetNewAnime() {
	// 若当前时间为 1 月、4 月、7 月和 10 月，则每天进行一次追踪更新，其余月份则每七天进行一次追踪更新
	//everyUpdateDays := 7
	//if nowMonth := time.Now().Month(); nowMonth == 1 || nowMonth == 4 || nowMonth == 7 || nowMonth == 10 {
	//	everyUpdateDays = 1
	//}

}
