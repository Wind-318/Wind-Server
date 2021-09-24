package functions

import (
	"Project/infomation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// 上传文件
func Upload(ctx *gin.Context) {
	res, err := ctx.MultipartForm()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "serverError.html", nil)
		return
	}
	files := res.File["file"]

	for _, file := range files {
		err = ctx.SaveUploadedFile(file, "userFile"+"/"+file.Filename)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "serverError.html", nil)
			return
		}
	}

	ctx.HTML(http.StatusOK, "success.html", nil)
}

// 上传头像
func UploadProfile(ctx *gin.Context) {
	result := map[string]interface{}{}
	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	if err != nil {
		result["msg"] = "请先登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	pic, _ := ctx.FormFile("pic")
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()
	var id int
	conn.Get(&id, "SELECT id FROM user WHERE account = ?", cookie)
	// 保存文件
	ctx.SaveUploadedFile(pic, `blog/`+strconv.Itoa(id)+`/`+pic.Filename)

	conn.Exec("UPDATE user SET pic = ? WHERE ID = ?;", `https://windserver.top/blog/`+strconv.Itoa(id)+`/`+pic.Filename, id)
}
