package sina

import (
	"log"
	"math/rand"
	"regexp"
	"spider/callClient/callAlgorithm"
	"spider/callClient/callDatabase"
	"spider/config"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 返回文章链接和标题
func Search() ([][]string, error) {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	// 连接数据库
	db, err := sqlx.Open("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/"+config.MySQLName)
	// 错误处理
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// 获取内容
	html, err := callAlgorithm.CallAlgorithmGetBytes("https://finance.sina.com.cn/stock/")
	if err != nil {
		return nil, err
	}
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 匹配规则
	rule := `href="(https://finance.sina.com.cn/stock/[\S]+?html[\S]*?)">[\S]+?</a>`

	// 新建正则表达式对象
	obj := regexp.MustCompile(rule)

	// 解析
	arr := obj.FindAllStringSubmatch(string(html), -1)

	// 循环
	for index := range arr {
		// 等待
		waitTime, err := callAlgorithm.CallAlgorithmRandomTime(54, 62)
		if err != nil {
			log.Println(err)
			continue
		}
		time.Sleep(time.Duration(waitTime))
		// 在缓存中查找是否存在
		if findFromCache("seturl", arr[index][1]) {
			continue
		}
		bytes, err := callAlgorithm.CallAlgorithmGetBytes(arr[index][1])
		if err != nil {
			continue
		}
		// 匹配
		temps, err := callAlgorithm.CallAlgorithmStringRegexp(string(bytes), `<title>([\s\S]+?)</title>`)
		if err != nil {
			log.Println(err)
			continue
		}
		// 错误处理
		if err != nil {
			return nil, err
		}
		// 未找到，返回
		if len(temps) == 0 || len(temps[0]) == 0 {
			continue
		}
		arr[index] = append(arr[index], temps[0][1])
		// 加入缓存中
		saveRedis(100, 19, arr[index][1], arr[index][2])
	}

	return arr, nil
}

// 生成正文
func GenerateText() (string, error) {
	// 获取标题和内容
	arr, err := Search()
	if err != nil {
		return "", err
	}
	// 连接数据库
	db, err := sqlx.Open("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/"+config.MySQLName)
	// 错误处理
	if err != nil {
		return "", err
	}
	defer db.Close()

	// 邮件内容
	text := ``
	// 按条生成格式
	for _, data := range arr {
		// 数据错误，跳过
		if len(data) < 3 {
			continue
		}
		// 链接
		text += `<h2>
		<a target="_blank" href="` + data[1] + `">` + data[2] + `</a>
		<h2>`

		// 保存到数据库中
		db.Exec("INSERT INTO urlinfo values(?,?,?,?,?)", 0, "https://finance.sina.com.cn/stock/", data[1], 0, data[2])
	}

	// 返回邮件正文
	return text, nil
}

// 选择前 n 条
func SelectNews(n int) (string, error) {
	// 获取 title
	titles, err := redis.Strings(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "LRANGE", "listtitle", 0, n))
	// 错误处理
	if err != nil {
		return "", err
	}
	// 获取链接
	urls, err := redis.Strings(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "LRANGE", "listurl", 0, n))

	// 错误处理
	if err != nil {
		return "", err
	}
	// 正文
	text := ``

	// 添加到正文
	for index := range urls {
		text += `<h2>
		<a target="_blank" href="` + urls[index] + `">` + titles[index] + `</a>
		<h2><br>`
	}

	return text, nil
}

// 查找缓存是否存在
func findFromCache(key, member string) bool {
	// 验证是否已存在
	reply, err := redis.Bool(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "SISMEMBER", key, member))
	if err != nil {
		return false
	}
	return reply
}

// 存储到缓存中，超过 n 条开始淘汰，保留最新 m 条消息
func saveRedis(n, m int, member ...string) error {
	// 获取长度
	llen, err := redis.Int(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "LLEN", "listurl"))
	// 错误处理
	if err != nil {
		return err
	}

	// 链表中超过 n 条消息开始淘汰，保留 m 条最新消息
	if llen > n {
		callDatabase.CallRedis(config.RedisIP, config.RedisPort, "LTRIM", "listurl", 0, m-1)
		callDatabase.CallRedis(config.RedisIP, config.RedisPort, "LTRIM", "listtitle", 0, m-1)
	}

	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "SADD", "seturl", member[0])
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "LPUSH", "listurl", member[0])
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "LPUSH", "listtitle", member[1])

	return nil
}
