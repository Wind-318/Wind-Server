package storage

import (
	"Project/gofiles/config"
	"Project/gofiles/user"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 获取某文件夹内容
func DownloadFiles(ctx *gin.Context) {
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
	// 获取文件夹名称
	fileName := ctx.PostForm("texts")
	urls := []string{}
	names := []string{}
	types := []string{}
	conn.Select(&urls, "SELECT path FROM storage WHERE filepath = ? AND account = ?", fileName, email)
	conn.Select(&names, "SELECT name FROM storage WHERE filepath = ? AND account = ?", fileName, email)
	conn.Select(&types, "SELECT type FROM storage WHERE filepath = ? AND account = ?", fileName, email)
	result["urls"] = urls
	result["names"] = names
	result["types"] = types
	ctx.JSON(http.StatusOK, result)
}

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
	// 获取分类
	filename := ctx.PostForm("file")
	// 一次选取的数量
	onceChoose, _ := strconv.Atoi(ctx.PostForm("onceChoose"))
	// 图片路径
	picPath := []string{}
	smallPicPath := []string{}
	conn.Select(&picPath, "SELECT path FROM storage WHERE account = ? AND filepath = ?", email, filename)
	conn.Select(&smallPicPath, "SELECT smallpath FROM storage WHERE account = ? AND filepath = ?", email, filename)
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
	// 分类
	filename := ctx.PostForm("file")
	// 每页数量
	pageNum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	email := ""
	conn.Get(&email, "SELECT account FROM user WHERE username = ?", name)
	// 总数
	sumNum := 0
	conn.Get(&sumNum, "SELECT count(*) FROM storage WHERE account = ? AND filepath = ?", email, filename)
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

// 查看文件夹个数
func GetUserFileNums(ctx *gin.Context) {
	// 返回值
	result := map[string]interface{}{}
	// 检验合法性
	if !user.IsExist(ctx) {
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLInfo)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 延时关闭
	defer conn.Close()
	// 名称
	name := ctx.PostForm("name")
	// 邮箱
	email := ""
	err = conn.Get(&email, "SELECT account FROM user WHERE username = ?", name)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取该用户文件夹个数
	files := []string{}
	err = conn.Select(&files, "SELECT DISTINCT filename FROM storagefile WHERE account = ?", email)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, result)
		return
	}
	result["files"] = files
	ctx.JSON(http.StatusOK, result)
}

// 新建文件夹
func MakeDirectory(ctx *gin.Context) {
	// 检验合法性
	if !user.IsExist(ctx) {
		return
	}
	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 延时关闭
	defer conn.Close()
	// 账号
	email := user.GetUserEmail(ctx)
	// 文件夹名称
	path := ctx.PostForm("path")
	// 新建
	os.MkdirAll("./userFile/"+email+"/"+path, 0644)
	// 添加数据库
	conn.Exec("INSERT INTO storagefile VALUES(?, ?, ?)", 0, email, path)
}
