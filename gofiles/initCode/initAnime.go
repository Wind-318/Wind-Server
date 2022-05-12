package initCode

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitAnime() {
	YhdmCatch()
}

// 樱花动漫片源地址
func YhdmCatch() {
	rand.Seed(time.Now().UnixNano())
	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo+"spider")
	defer conn.Close()

	yhdmAnime := "https://www.yhdmp.cc/list/?year="
	for i := 2000; i <= time.Now().Year(); i++ {
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

				// 保存图片
				if data[2][:4] != "http" {
					data[2] = "http:" + data[2]
				}
				// 休息数秒
				time.Sleep(time.Second * time.Duration(algorithm.GetRandomNumberTime(54, 62)))
				picFile, _ := http.Get(data[2])
				picByte, err := ioutil.ReadAll(picFile.Body)
				picName := strconv.Itoa(int(time.Now().UnixNano()))
				if err == nil {
					ioutil.WriteFile("./picture/anime/yhdm/"+picName+".jpg", picByte, 0644)
				}

				// 检索后动漫名称如已存在，则更新资料；如不存在则插入
				isExist = 0
				conn.Get(&isExist, "SELECT count(*) FROM bangumi WHERE name = ?", data[3])
				if isExist > 0 {
					conn.Exec("UPDATE bangumi SET url = ?, description = ?, picurl = ? WHERE name = ?", "https://www.yhdmp.cc"+data[1], data[4], "../picture/anime/yhdm/"+picName+".jpg", data[3])
				} else {
					conn.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, data[3], "https://www.yhdmp.cc"+data[1], strconv.Itoa(i), data[4], "../picture/anime/yhdm/"+picName+".jpg", 0)
				}

				conn.Exec("INSERT INTO animesource VALUES(?, ?, ?, ?)", 0, data[3], "樱花动漫", "https://www.yhdmp.cc"+data[1])
			}

			res.Body.Close()
			time.Sleep(time.Second * time.Duration(algorithm.GetRandomNumberTime(54, 62)))
		}
	}
}
