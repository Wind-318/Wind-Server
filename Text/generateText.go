package Text

import (
	"Project/httpRequest"
	"Project/infomation"
	"math/rand"
	"regexp"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var rwmutex *sync.RWMutex = &sync.RWMutex{}

// 返回文章链接和标题
func Search() [][]string {
	rand.Seed(time.Now().UnixNano())
	db, _ := sqlx.Open("mysql", infomation.MySQLInfo)
	defer db.Close()

	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	html := httpRequest.GetRequestByte("https://finance.sina.com.cn/stock/")

	// 匹配规则
	rule := `href="(https://finance.sina.com.cn/stock/[\S]+?html[\S]*?)">[\S]+?</a>`
	obj := regexp.MustCompile(rule)

	arr := obj.FindAllStringSubmatch(string(html), -1)

	for index := range arr {
		if FindFromCache("seturl", arr[index][1]) {
			continue
		}
		temp := httpRequest.RegexpHtml(arr[index][1], `<title>([\s\S]+?)</title>`)
		if len(temp) == 0 {
			continue
		}
		arr[index] = append(arr[index], temp[0])
		SaveRedis(arr[index][1], arr[index][2])
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}

	return arr
}

// 生成正文
func GenerateText() string {
	arr := Search()
	db, _ := sqlx.Open("mysql", infomation.MySQLInfo)
	defer db.Close()
	text := ``

	for _, data := range arr {
		if len(data) < 3 {
			continue
		}
		text += `<h2>
		<a target="_blank" href="` + data[1] + `">` + data[2] + `</a>
		<h2>`

		// 保存到数据库中
		tx, err := db.Begin()
		if err != nil {
			return ""
		}
		_, err = tx.Exec("INSERT INTO urlinfo values(?,?,?,?,?)", 0, "https://finance.sina.com.cn/stock/", data[1], 0, data[2])

		if err != nil {
			return ""
		}
		err = tx.Commit()
		if err != nil {
			return ""
		}
	}

	return text
}

// 查找缓存是否存在
func FindFromCache(key, member string) bool {
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()
	reply, err := redis.Bool(connect.Do("SISMEMBER", key, member))
	if err != nil {
		return false
	}
	return reply
}

// 存储到缓存中
func SaveRedis(member ...string) {
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	rwmutex.Lock()
	defer rwmutex.Unlock()

	llen, err := redis.Int(connect.Do("LLEN", "listurl"))
	if err != nil {
		return
	}

	// 链表中超过 100 条消息开始淘汰，保留 20 条最新消息
	if llen > 100 {
		connect.Do("LTRIM", "listurl", 0, 19)
		connect.Do("LTRIM", "listtitle", 0, 19)
	}

	connect.Do("SADD", "seturl", member[0])
	connect.Do("LPUSH", "listurl", member[0])
	connect.Do("LPUSH", "listtitle", member[1])
}

func SelectFirst10() string {
	rwmutex.RLock()

	con, _ := redis.Dial("tcp", "127.0.0.1:6379")
	titles, err := redis.Strings(con.Do("LRANGE", "listtitle", 0, 9))
	if err != nil {
		return ""
	}
	urls, err := redis.Strings(con.Do("LRANGE", "listurl", 0, 9))

	rwmutex.RUnlock()

	if err != nil {
		return ""
	}
	text := ``

	for index := range urls {
		text += `<h2>
		<a target="_blank" href="` + urls[index] + `">` + titles[index] + `</a>
		<h2><br>`
	}

	return text
}

func SelectFirst10WithPicture(picNum string) string {
	rwmutex.RLock()

	con, _ := redis.Dial("tcp", "127.0.0.1:6379")
	titles, err := redis.Strings(con.Do("LRANGE", "listtitle", 0, 9))
	if err != nil {
		return ""
	}
	urls, err := redis.Strings(con.Do("LRANGE", "listurl", 0, 9))

	rwmutex.RUnlock()

	if err != nil {
		return ""
	}
	text := ``

	for index := range urls {
		text += `<h2>
		<a target="_blank" href="` + urls[index] + `">` + titles[index] + `</a>
		<h2><br>`
	}

	text += `<br><img src="` + picNum + infomation.PicFormat + `" alt="My image" />`

	return text
}

func SelectFirst20() string {
	rwmutex.RLock()

	con, _ := redis.Dial("tcp", "127.0.0.1:6379")
	titles, err := redis.Strings(con.Do("LRANGE", "listtitle", 0, 19))
	if err != nil {
		return ""
	}
	urls, err := redis.Strings(con.Do("LRANGE", "listurl", 0, 19))

	rwmutex.RUnlock()

	if err != nil {
		return ""
	}
	text := ``

	for index := range urls {
		text += `<h2>
		<a target="_blank" href="` + urls[index] + `">` + titles[index] + `</a>
		<h2>`
	}

	return text
}
