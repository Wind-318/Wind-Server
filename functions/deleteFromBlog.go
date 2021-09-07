package functions

import (
	"Project/infomation"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type modifyText struct {
	Title       string `db:"title"`
	Description string `db:"description"`
	Content     string `db:"content"`
	Picurl      string `db:"picurl"`
}

// 删除一篇文章
func DeleteFromBlog(ctx *gin.Context) {
	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))

	var ids []string
	jsonArr := ctx.PostForm("checked")
	json.Unmarshal([]byte(jsonArr), &ids)

	for _, id := range ids {
		if id != cookie && id != infomation.SystemUserAccount {
			return
		}
		conn.Exec("DELETE FROM blog WHERE id = ?", id)
	}
}

// 删除评论
func DeleteComment(ctx *gin.Context) {
	id := ctx.PostForm("id")

	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()

	conn.Exec("DELETE FROM comments WHERE id = ?", id)
}

// 获取修改文章的信息
func GetModifyBlog(ctx *gin.Context) {
	_, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}

	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()
	modifyT := modifyText{}
	conn.Get(&modifyT, "SELECT title, description, content, picurl FROM blog WHERE id = ?", id)
	result := map[string]interface{}{
		"id":          id,
		"title":       modifyT.Title,
		"description": modifyT.Description,
		"content":     modifyT.Content,
		"picurl":      modifyT.Picurl,
	}

	ctx.JSON(http.StatusOK, result)
}

// 修改文章
func ModifyBlog(ctx *gin.Context) {
	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	if err != nil {
		return
	}
	id := ctx.PostForm("id")
	text := ctx.PostForm("texts")
	titles := ctx.PostForm("titles")
	description := ctx.PostForm("description")
	pic, _ := ctx.FormFile("pic")
	attFile, _ := ctx.MultipartForm()
	attFiles := attFile.File["attFiles"]
	pictype := ctx.PostForm("picType")

	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()

	var ids int
	conn.Get(&ids, "SELECT id FROM user WHERE account = ?", cookie)
	types := ""
	conn.Get(&types, "SELECT types FROM blog WHERE id = ?", id)

	conn.Exec("UPDATE blog SET content = ? WHERE id = ?", text, id)
	if titles != "" {
		conn.Exec("UPDATE blog SET title = ? WHERE id = ?", titles, id)
	}
	if description != "" {
		conn.Exec("UPDATE blog SET description = ? WHERE id = ?", description, id)
	}
	if pictype != "" {
		var randTime = strconv.Itoa(int(time.Now().UnixNano()))
		var picAddr = infomation.Addr + `blog/` + strconv.Itoa(ids) + `/` + types + "/" + randTime + "." + pictype

		// 保存文件
		ctx.SaveUploadedFile(pic, `blog/`+strconv.Itoa(ids)+`/`+types+"/"+randTime+"."+pictype)

		conn.Exec("UPDATE blog SET picurl = ? WHERE id = ?", picAddr, id)
	}
	for i := range attFiles {
		ctx.SaveUploadedFile(attFiles[i], `blog/`+strconv.Itoa(ids)+`/`+types+"/"+attFiles[i].Filename)
	}
}
