package bbs

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callDatabase"
	"Project/callClient/callUser"
	"Project/config"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 评论信息
type comments struct {
	Id          int    `db:"id"`
	Blog        int    `db:"blog"`
	Content     string `db:"content"`
	Create_time string `db:"create_time"`
	Update_time string `db:"update_time"`
	Parent      int    `db:"parent"`
	Pic         string `db:"pic"`
	Author      string `db:"author"`
}

// 查找评论
func TextComment(ctx *gin.Context) {
	conn := sqlx.MustConnect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	defer conn.Close()

	id := ctx.PostForm("id")
	comment := []comments{}
	conn.Select(&comment, "SELECT * FROM comments WHERE blog = ?", id)
	ids := make([]int, len(comment))
	blogs := make([]int, len(comment))
	contents := make([]string, len(comment))
	create_time := make([]string, len(comment))
	update_time := make([]string, len(comment))
	pics := make([]string, len(comment))
	authors := make([]string, len(comment))
	parents := make([]string, len(comment))

	for i := range comment {
		father := ""
		conn.Get(&father, "SELECT author FROM comments WHERE id = ?", comment[i].Parent)
		ids[i] = comment[i].Id
		blogs[i] = comment[i].Blog
		contents[i] = comment[i].Content
		create_time[i] = comment[i].Create_time
		update_time[i] = comment[i].Update_time
		pics[i] = comment[i].Pic
		authors[i] = comment[i].Author
		parents[i] = father
	}

	result := map[string]interface{}{
		"nums":        len(ids),
		"ids":         ids,
		"blogs":       blogs,
		"contents":    contents,
		"create_time": create_time,
		"update_time": update_time,
		"parents":     parents,
		"pics":        pics,
		"authors":     authors,
	}

	ctx.JSON(http.StatusOK, result)
}

// 添加评论
func AddComment(ctx *gin.Context) {
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
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	id := ctx.PostForm("id")
	parent := ctx.PostForm("parent")
	content := ctx.PostForm("content")
	conn := sqlx.MustConnect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	conn1 := sqlx.MustConnect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/user")
	defer conn.Close()
	defer conn1.Close()
	pic := ""
	author := ""
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}
	// 获取账号
	email, _ := redis.String(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "HGET", cookie, "email"))

	// 加分布式锁
	if ok, err := redis.String(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "SET", email+"Comment", "comment", "EX", 10, "NX")); err != nil || ok != "OK" {
		return
	}
	// 设置检测，每 8 秒检测锁是否存在，存在就延长
	go func() {
		for {
			// 锁存在
			if ok, err := redis.Bool(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "SETNX", email+"Comment", "comment")); err != nil || !ok {
				// 设置过期 10 秒
				callDatabase.CallRedis(config.RedisIP, config.RedisPort, "Expire", email+"Comment", 10)
			} else {
				callDatabase.CallRedis(config.RedisIP, config.RedisPort, "del", email+"Comment")
				return
			}
			time.Sleep(time.Second * 8)
		}
	}()
	defer callDatabase.CallRedis(config.RedisIP, config.RedisPort, "del", email+"Comment")

	// 选出头像
	conn1.Get(&pic, "SELECT pic FROM user WHERE account = ?", email)
	// 选出作者
	conn1.Get(&author, "SELECT username FROM user WHERE account = ?", email)

	conn.Exec("INSERT INTO comments VALUES(?, ?, ?, ?, ?, ?, ?, ?)", 0, id, content, time.Now().String()[:19], time.Now().String()[:19], parent, pic, author)
}

// 点赞
func Parise(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	// 连接数据库
	conn1 := sqlx.MustConnect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")

	defer conn1.Close()
	id := ctx.PostForm("id")
	conn1.Exec("UPDATE blog SET great = great + 1 WHERE id = ?", id)
}

// 点赞数
func PariseNum(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	defer conn.Close()

	num := 0
	conn.Get(&num, "SELECT great FROM blog WHERE id = ?", id)

	result := map[string]interface{}{
		"num": num,
	}

	ctx.JSON(http.StatusOK, result)
}

// 浏览量
func Views(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	defer conn.Close()

	conn.Exec("UPDATE blog SET clicknum = clicknum + 1 WHERE id = ?", id)
	num := 0
	conn.Get(&num, "SELECT clicknum FROM blog WHERE id = ?", id)
	result := map[string]interface{}{
		"num": num,
	}

	ctx.JSON(http.StatusOK, result)
}

func Author(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	defer conn.Close()

	author := ""
	conn.Get(&author, "SELECT author FROM blog WHERE id = ?", id)
	result := map[string]interface{}{
		"author": author,
	}

	ctx.JSON(http.StatusOK, result)
}

// 获取所有评论id
func GetCommentsID(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLAccount+":"+config.MySQLPassword+"@tcp("+config.MySQLIP+config.MySQLPort+")/bbs")
	defer conn.Close()

	comment := []comments{}
	conn.Select(&comment, "SELECT id FROM comments WHERE blog = ?", id)

	ids := make([]int, len(comment))
	for i := range comment {
		ids[i] = comment[i].Id
	}
	result := map[string]interface{}{
		"ids": ids,
	}

	ctx.JSON(http.StatusOK, result)
}
