package blogfunc

import (
	"Project/gofiles/config"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
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
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))

	// 获取删除文章 id
	var ids []string
	jsonArr := ctx.PostForm("checked")
	json.Unmarshal([]byte(jsonArr), &ids)

	// 删除
	for _, id := range ids {
		conf := ""
		conn.Get(&conf, "SELECT authoremail FROM blog WHERE id = ?", id)
		if conf != cookie && conf != config.SystemUserAccount {
			return
		}
		conn.Exec("DELETE FROM blog WHERE id = ?", id)
	}
}

// 删除评论
func DeleteComment(ctx *gin.Context) {
	id := ctx.PostForm("id")

	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
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
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
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
	_, err := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	if err != nil {
		return
	}

	// 文章 ID
	id := ctx.PostForm("id")
	// 文章正文
	text := ctx.PostForm("texts")
	// 文章标题
	titles := ctx.PostForm("titles")
	// 文章简介
	description := ctx.PostForm("description")
	// 背景
	pic, _ := ctx.FormFile("pic")
	// 背景类型
	pictype := ctx.PostForm("picType")

	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	// 作者 ID
	authorid := 0
	conn.Get(&authorid, "SELECT authorid FROM blog WHERE id = ?", id)
	// 文章类型
	types := ""
	conn.Get(&types, "SELECT types FROM blog WHERE id = ?", id)

	// 修改正文、标题、简介
	conn.Exec("UPDATE blog SET content = ? WHERE id = ?", text, id)
	if titles != "" && titles != "undefined" {
		conn.Exec("UPDATE blog SET title = ? WHERE id = ?", titles, id)
	}
	if description != "" {
		conn.Exec("UPDATE blog SET description = ? WHERE id = ?", description, id)
	}

	// 更新修改时间
	conn.Exec("UPDATE blog SET update_time = ? WHERE id = ?", time.Now().String()[:19], id)
	if pictype != "" {
		var picAddr = config.Addr + `blog/` + strconv.Itoa(authorid) + `/` + types + "/" + pic.Filename

		conn.Exec("UPDATE blog SET picurl = ? WHERE id = ?", picAddr, id)
		conn.Exec("UPDATE blog SET smallpic = ? WHERE id = ?", config.Addr+`blog/`+strconv.Itoa(authorid)+`/`+types+"/small"+pic.Filename, id)
	}
}

// 修改文章上传的文件
func ModifyBlogFiles(ctx *gin.Context) {
	// 检查登录状态
	_, err := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	if err != nil {
		return
	}
	id := ctx.PostForm("id")
	pic, _ := ctx.FormFile("pic")
	pictype := ctx.PostForm("picType")
	attFile, _ := ctx.MultipartForm()
	attFiles := attFile.File["attFiles"]

	// 获取文章分类
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	types := ""
	conn.Get(&types, "SELECT types FROM blog WHERE id = ?", id)
	authorid := 0
	conn.Get(&authorid, "SELECT authorid FROM blog WHERE id = ?", id)

	if pictype != "" {
		ctx.SaveUploadedFile(pic, `blog/`+strconv.Itoa(authorid)+`/`+types+"/"+pic.Filename)

		// 创建缩略图
		imgData, _ := ioutil.ReadFile(`blog/` + strconv.Itoa(authorid) + `/` + types + "/" + pic.Filename)
		buf := bytes.NewBuffer(imgData)
		image, err := imaging.Decode(buf)
		if err != nil {
			return
		}
		image = imaging.Resize(image, 0, 400, imaging.Lanczos)
		err = imaging.Save(image, `blog/`+strconv.Itoa(authorid)+`/`+types+"/small"+pic.Filename)
		if err != nil {
			return
		}
	}

	// 保存文件
	for i := range attFiles {
		err = ctx.SaveUploadedFile(attFiles[i], `blog/`+strconv.Itoa(authorid)+`/`+types+"/"+attFiles[i].Filename)
		if err != nil {
			return
		}
	}
}
