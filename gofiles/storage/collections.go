package storage

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"Project/gofiles/info"
	"Project/gofiles/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// 获取网站链接
func GetWebs(ctx *gin.Context) {
	// 判断身份
	if ok := user.IsLogin(ctx); !ok {
		result := map[string]interface{}{
			"msg": "尚未登陆！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 返回信息
	result := map[string]interface{}{
		"ids":      0,
		"urls":     0,
		"comments": 0,
		"picurls":  0,
	}
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"storage")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient.Close()
	arr := []info.CollectionData{}
	// 查找收藏网站信息
	err = mysqlClient.Select(&arr, "SELECT * FROM collections")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	ids := make([]int, len(arr))
	urls := make([]string, len(arr))
	comments := make([]string, len(arr))
	picurls := make([]string, len(arr))

	for index := range arr {
		ids[index] = arr[index].ID
		urls[index] = arr[index].Url
		comments[index] = arr[index].Description
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
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())

	result := map[string]interface{}{}

	// 验证
	if email, err := user.GetUserEmail(ctx); err != nil || email != config.SystemAccount {
		result["msg"] = "非管理员账户"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取 url
	url := ctx.PostForm("url")
	// 获取名称
	comment := ctx.PostForm("comment")
	// 获取网站图标
	pic := ctx.PostForm("picurl")

	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"storage")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient.Close()
	mysqlClient.Exec("INSERT INTO collections VALUES(?, ?, ?, ?)", 0, url, comment, pic)

	ctx.JSON(http.StatusOK, result)
}

// 增加图片
func PutPic(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())

	result := map[string]interface{}{
		"msg": "success",
	}
	// 判断管理员账户
	if email, err := user.GetUserEmail(ctx); err != nil || email != config.SystemAccount {
		result["msg"] = "非管理员账户"
		ctx.JSON(http.StatusOK, result)
		return
	}

	pic, _ := ctx.FormFile("pic")
	go func() {
		ctx.SaveUploadedFile(pic, `picture/collections/`+pic.Filename)
	}()
	ctx.JSON(http.StatusOK, result)
}
