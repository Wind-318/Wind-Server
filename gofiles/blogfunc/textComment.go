package blogfunc

import (
	"Project/gofiles/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

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
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
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
	id := ctx.PostForm("id")
	parent := ctx.PostForm("parent")
	content := ctx.PostForm("content")
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	pic := ""
	author := ""
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	email, _ := redis.String(redisconn.Do("HGET", cookie, "email"))
	conn.Get(&pic, "SELECT pic FROM user WHERE account = ?", email)
	conn.Get(&author, "SELECT username FROM user WHERE account = ?", email)

	conn.Exec("INSERT INTO comments VALUES(?, ?, ?, ?, ?, ?, ?, ?)", 0, id, content, time.Now().String()[:19], time.Now().String()[:19], parent, pic, author)
}

// 点赞
func Parise(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	conn.Exec("UPDATE blog SET great = great + 1 WHERE id = ?", id)
}

// 点赞数
func PariseNum(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
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
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
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
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
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
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
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
