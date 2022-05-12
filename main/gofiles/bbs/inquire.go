package bbs

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callDatabase"
	"Project/callClient/callUser"
	"Project/config"
	"net/http"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 文章信息
type textInfo struct {
	Id          int    `db:"id"`
	Author      string `db:"author"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Clicknum    int    `db:"clicknum"`
	Great       int    `db:"great"`
	Types       string `db:"types"`
	Authority   int    `db:"authority"`
	Create_time string `db:"create_time"`
	Update_time string `db:"update_time"`
	Authorid    int    `db:"authorid"`
	Urls        string `db:"url"`
	Picurl      string `db:"picurl"`
	SmallPic    string `db:"smallpic"`
}

// 获取文章页数
func GetPageNums(ctx *gin.Context) {
	val, _ := strconv.Atoi(ctx.PostForm("num"))
	// 每页数量
	every := 10
	result := map[string]interface{}{
		"authority":   0,
		"author":      0,
		"num":         0,
		"start":       0,
		"end":         0,
		"status":      0,
		"id":          0,
		"urls":        0,
		"picurl":      0,
		"description": 0,
		"isSystem":    0,
	}

	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer conn.Close()

	//
	if cookie, err := ctx.Cookie("cookie"); err == nil {
		redisconn, _ := redis.Dial("tcp", "localhost:6379")
		defer redisconn.Close()
		email, _ := redis.String(redisconn.Do("HGET", cookie, "email"))
		if email == config.SystemUserAccount {
			result["isSystem"] = 1
		}

		result["status"] = 1

		userid, err := callDatabase.CallMySQLSelectUserIDByAccount(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", email)
		if err != nil {
			result["msg"] = err.Error()
			ctx.JSON(http.StatusOK, result)
			return
		}

		result["ids"] = userid
	}

	// 从数据库选取文章
	temp := []textInfo{}
	conn.Select(&temp, "SELECT id, author, title, description, types, clicknum, great, authority, create_time, update_time, authorid, url, picurl, smallpic FROM blog WHERE authority = 0")
	ids := []int{}
	urlsarr := []string{}
	titles := []string{}
	picurls := []string{}
	authoritys := []int{}
	descriptions := []string{}
	authors := []string{}
	create_time := []string{}
	update_time := []string{}
	for _, data := range temp {
		ids = append(ids, data.Id)
		urlsarr = append(urlsarr, data.Urls)
		titles = append(titles, data.Title)
		descriptions = append(descriptions, data.Description)
		authors = append(authors, data.Author)
		create_time = append(create_time, data.Create_time)
		update_time = append(update_time, data.Update_time)
		picurls = append(picurls, data.SmallPic)
		authoritys = append(authoritys, data.Authority)
	}
	result["urls"] = urlsarr
	result["titles"] = titles
	result["picurl"] = picurls
	result["description"] = descriptions
	result["author"] = authors
	result["create_time"] = create_time
	result["update_time"] = update_time
	result["id"] = ids
	result["authority"] = authoritys

	num := len(temp)
	if val >= num {
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 总文章数量
	result["num"] = num
	// 开始
	if num-val-every <= 0 {
		result["start"] = 1
	} else {
		result["start"] = num - val - every + 1
	}
	// 结束
	result["end"] = num - val + 1

	ctx.JSON(http.StatusOK, result)
}

// 获取分类的数量
func GetClassification(ctx *gin.Context) {
	result := map[string]interface{}{
		"num":   0,
		"types": 0,
		"Addr":  0,
		"pic":   0,
	}
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	result["Addr"] = config.Addr
	cookies, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.JSON(http.StatusOK, result)
		return
	}
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))

	conn, err := sqlx.Connect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	if err != nil {
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer conn.Close()

	arr := []textInfo{}
	conn.Select(&arr, "SELECT DISTINCT types FROM blog WHERE authoremail = ?", cookie)
	result["num"] = len(arr)

	tempStr := []string{}
	picarr := []string{}
	for _, textType := range arr {
		tempStr = append(tempStr, textType.Types)
		var picurl string
		conn.Get(&picurl, "SELECT smallpic FROM blog WHERE types = ?", textType.Types)
		picarr = append(picarr, picurl)
	}

	result["types"] = tempStr
	result["pic"] = picarr

	ctx.JSON(http.StatusOK, result)
}

// 获取某一类型的文章
func GetText(ctx *gin.Context) {
	result := map[string]interface{}{
		"num": 0,
	}
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}

	cookies, err := ctx.Cookie("cookie")
	types := ctx.PostForm("types")
	if err != nil {
		ctx.JSON(http.StatusOK, result)
		return
	}

	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))

	conn, _ := sqlx.Connect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	defer conn.Close()

	arr := []textInfo{}
	conn.Select(&arr, "SELECT id, author, title, description, clicknum, great, authority, create_time, update_time, authorid, url, picurl, smallpic FROM blog WHERE authoremail = ? AND types = ?", cookie, types)

	id := make([]int, len(arr))
	author := make([]string, len(arr))
	title := make([]string, len(arr))
	description := make([]string, len(arr))
	clicknum := make([]int, len(arr))
	great := make([]int, len(arr))
	authority := make([]int, len(arr))
	create_time := make([]string, len(arr))
	update_time := make([]string, len(arr))
	authorid := make([]int, len(arr))
	urls := make([]string, len(arr))
	picurls := make([]string, len(arr))

	//
	for index := range arr {
		id[index] = arr[index].Id
		author[index] = arr[index].Author
		title[index] = arr[index].Title
		description[index] = arr[index].Description
		clicknum[index] = arr[index].Clicknum
		great[index] = arr[index].Great
		authority[index] = arr[index].Authority
		create_time[index] = arr[index].Create_time
		update_time[index] = arr[index].Update_time
		authorid[index] = arr[index].Authorid
		urls[index] = arr[index].Urls
		picurls[index] = arr[index].SmallPic
	}

	// 添加到 json
	result["num"] = len(arr)
	result["id"] = id
	result["author"] = author
	result["titles"] = title
	result["description"] = description
	result["clicknum"] = clicknum
	result["great"] = great
	result["authority"] = authority
	result["create_time"] = create_time
	result["update_time"] = update_time
	result["authorid"] = authorid
	result["urls"] = urls
	result["picurl"] = picurls

	// 返回
	ctx.JSON(http.StatusOK, result)
}

// 搜索文章
func Search(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result := map[string]interface{}{}
	// 文章名
	text := strings.ToLower(ctx.PostForm("text"))
	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer conn.Close()

	// 选取文章信息
	temp := []textInfo{}
	conn.Select(&temp, "SELECT id, author, title, description, types, clicknum, great, authority, create_time, update_time, authorid, url, picurl, smallpic FROM blog WHERE authority = 0")

	// id
	ids := []int{}
	// url
	urlsarr := []string{}
	// 标题
	titles := []string{}
	// 背景图
	picurls := []string{}
	// 作者 id
	authoritys := []int{}
	// 简介
	descriptions := []string{}
	// 作者
	authors := []string{}
	// 创建时间
	create_time := []string{}
	// 更新时间
	update_time := []string{}

	// 匹配
	for index := range temp {
		ok, err := callAlgorithm.CallAlgorithmMatch(temp[index].Title, text)
		if err != nil {
			result["msg"] = err.Error()
			ctx.JSON(http.StatusOK, result)
			return
		}
		if text == "" || ok {
			ids = append(ids, temp[index].Id)
			urlsarr = append(urlsarr, temp[index].Urls)
			titles = append(titles, temp[index].Title)
			descriptions = append(descriptions, temp[index].Description)
			authors = append(authors, temp[index].Author)
			create_time = append(create_time, temp[index].Create_time)
			update_time = append(update_time, temp[index].Update_time)
			picurls = append(picurls, temp[index].SmallPic)
			authoritys = append(authoritys, temp[index].Authority)
		}
	}

	// 添加到 json 中
	result["urls"] = urlsarr
	result["titles"] = titles
	result["picurl"] = picurls
	result["description"] = descriptions
	result["author"] = authors
	result["create_time"] = create_time
	result["update_time"] = update_time
	result["id"] = ids
	result["num"] = len(urlsarr)
	result["authority"] = authoritys

	// 检查是否是管理员
	ok, err := callUser.CallUserIsSystem(ctx)
	if err != nil || !ok {
		result["isSystem"] = 0
	} else {
		result["isSystem"] = 1
	}

	ctx.JSON(http.StatusOK, result)
}
