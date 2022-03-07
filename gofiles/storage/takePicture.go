package storage

import (
	"Project/gofiles/config"
	"Project/gofiles/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// 获取指定用户上传的图片
func GetUserStoragePicture(ctx *gin.Context) {
	// 返回值
	result := map[string]interface{}{}
	// 检验合法性
	if !user.IsExist(ctx) {
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	// 获取邮箱
	name := ctx.PostForm("name")
	email := ""
	conn.Get(&email, "SELECT account FROM user WHERE username = ?", name)
	// 获取偏移量
	moveNum, _ := strconv.Atoi(ctx.PostForm("num"))
	// 一次选取的数量
	onceChoose, _ := strconv.Atoi(ctx.PostForm("onceChoose"))
	// 图片路径
	picPath := []string{}
	smallPicPath := []string{}
	conn.Select(&picPath, "SELECT path FROM storage WHERE account = ?", email)
	conn.Select(&smallPicPath, "SELECT smallpic FROM storage WHERE account = ?", email)
	if moveNum+onceChoose >= len(picPath) {
		picPath = picPath[moveNum:]
		smallPicPath = smallPicPath[moveNum:]
	} else {
		picPath = picPath[moveNum : moveNum+onceChoose]
		smallPicPath = smallPicPath[moveNum : moveNum+onceChoose]
	}
	result["picPath"] = picPath
	result["smallPicPath"] = smallPicPath
	result["num"] = len(picPath)
	ctx.JSON(http.StatusOK, result)
}

// 获取页数
func GetUserStoragePicturePage(ctx *gin.Context) {
	// 名称
	name := ctx.PostForm("name")
	// 每页数量
	pageNum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	email := ""
	conn.Get(&email, "SELECT account FROM user WHERE username = ?", name)
	// 总数
	sumNum := 0
	conn.Get(&sumNum, "SELECT count(*) FROM storage WHERE account = ?", email)
	// 计算页数
	page := 1
	if sumNum%pageNum == 0 {
		page = sumNum / pageNum
	} else {
		page = sumNum/pageNum + 1
	}
	if page == 0 {
		page = 1
	}
	result := map[string]interface{}{}
	result["page"] = page
	ctx.JSON(http.StatusOK, result)
}
