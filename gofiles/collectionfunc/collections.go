package collectionfunc

import (
	"Project/gofiles/config"
	"Project/gofiles/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// 收藏网站信息
type collections struct {
	ID      int    `db:"id"`
	Url     string `db:"url"`
	Comment string `db:"description"`
	Picurl  string `db:"picurl"`
}

// 查看是否是管理员
func IsSystem(ctx *gin.Context) {
	result := map[string]interface{}{
		"msg": "fail",
	}
	// 获取登录信息
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.JSON(http.StatusOK, result)
		return
	}

	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()

	// 获取账号
	account, _ := redis.String(redisconn.Do("HGET", cookie, "email"))
	if account == config.SystemUserAccount {
		result["msg"] = "success"
	}

	ctx.JSON(http.StatusOK, result)
}

// 查看是否是管理员或作者
func IsSystems(ctx *gin.Context) {
	result := map[string]interface{}{
		"msg": "fail",
	}
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	ids := ""
	conn.Get(&ids, "SELECT authoremail FROM blog WHERE id = ?", id)

	// 获取登录信息
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.JSON(http.StatusOK, result)
		return
	}

	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()

	account, _ := redis.String(redisconn.Do("HGET", cookie, "email"))
	if account == config.SystemUserAccount || account == ids {
		result["msg"] = "success"
	}

	ctx.JSON(http.StatusOK, result)
}

// 获取网站链接
func GetWebs(ctx *gin.Context) {
	if !user.IsExist(ctx) {
		return
	}
	result := map[string]interface{}{
		"ids":      0,
		"urls":     0,
		"comments": 0,
		"picurls":  0,
	}

	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	// 查找收藏网站信息
	var arr = []collections{}
	conn.Select(&arr, "SELECT * FROM collections")

	ids := make([]int, len(arr))
	urls := make([]string, len(arr))
	comments := make([]string, len(arr))
	picurls := make([]string, len(arr))

	for index := range arr {
		ids[index] = arr[index].ID
		urls[index] = arr[index].Url
		comments[index] = arr[index].Comment
		picurls[index] = arr[index].Picurl
	}

	result["ids"] = ids
	result["urls"] = urls
	result["comments"] = comments
	result["picurls"] = picurls
	ctx.JSON(http.StatusOK, result)
}

// 增加收藏网站
func PutWebs(ctx *gin.Context) {
	url := ctx.PostForm("url")
	comment := ctx.PostForm("comment")
	pic := ctx.PostForm("picurl")
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	conn.Exec("INSERT INTO collections VALUES(?, ?, ?, ?)", 0, url, comment, pic)
}

// 增加图片
func PutPic(ctx *gin.Context) {
	pic, _ := ctx.FormFile("pic")
	go func() {
		ctx.SaveUploadedFile(pic, `picture/collections/`+pic.Filename)
	}()
	result := map[string]interface{}{
		"msg": "success",
	}
	ctx.JSON(http.StatusOK, result)
}
