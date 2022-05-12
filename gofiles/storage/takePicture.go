package storage

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"Project/gofiles/user"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// 批量下载
func DownloadFiles(ctx *gin.Context) {
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

	// 检查是否已登录
	if ok := !user.IsLogin(ctx); ok {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}

	userAccount, err := user.GetUserEmail(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 加分布式锁
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	if ok, err := redis.String(redisCli.Do("SET", userAccount+"DownloadFiles", "DownloadFiles", "EX", 30, "NX")); err != nil || ok != "OK" {
		return
	}
	// 设置检测，每 25 秒检测锁是否存在，存在就延长
	go func() {
		for {
			time.Sleep(time.Second * 25)
			// 锁存在
			if ok, err := redis.Bool(redisCli.Do("SETNX", userAccount+"DownloadFiles", "DownloadFiles")); err != nil || !ok {
				// 设置过期 30 秒
				redisCli.Do("Expire", userAccount+"DownloadFiles", 30)
			} else {
				redisCli.Do("del", userAccount+"DownloadFiles")
				return
			}
		}
	}()
	defer redisCli.Do("del", userAccount+"DownloadFiles")

	// 获取邮箱
	name := ctx.PostForm("name")
	// 获取文件夹名称
	fileName := ctx.PostForm("texts")
	if name == "" || fileName == "" {
		result["msg"] = "需进入文件夹后下载！"
		ctx.JSON(http.StatusOK, result)
		return
	}

	result["url"] = "../userFile/" + name + "/" + fileName + ".zip"
	result["name"] = fileName + ".zip"

	// 压缩
	err = algorithm.Zip("./userFile/"+name+"/"+fileName, "./userFile/"+name+"/"+fileName+".zip")
	if err != nil {
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}

// 获取指定用户上传的图片
func GetUserStoragePicture(ctx *gin.Context) {
	// 返回值
	result := map[string]interface{}{}
	// 检查是否已登录
	if ok := !user.IsLogin(ctx); ok {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}
	var err error

	// 获取邮箱
	name := ctx.PostForm("name")
	email, err := user.SelectUserAccountByName(name)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取偏移量
	moveNums := ctx.PostForm("num")
	moveNum, _ := strconv.Atoi(moveNums)
	// 获取分类
	folderName := ctx.PostForm("file")
	// 一次选取的数量
	onceChooses := ctx.PostForm("onceChoose")
	onceChoose, _ := strconv.Atoi(onceChooses)
	// 查询
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"storage")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient.Close()

	// 选取图片
	picpath := []string{}
	err = mysqlClient.Select(&picpath, "SELECT path FROM storage WHERE account = ? AND filepath = ?", email, folderName)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 选取缩略图
	smallpath := []string{}
	mysqlClient.Select(&smallpath, "SELECT smallpath FROM storage WHERE account = ? AND filepath = ?", email, folderName)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 根据偏移量选取显示的图片
	if moveNum+onceChoose >= len(picpath) {
		picpath = picpath[moveNum:]
		smallpath = smallpath[moveNum:]
	} else {
		picpath = picpath[moveNum : moveNum+onceChoose]
		smallpath = smallpath[moveNum : moveNum+onceChoose]
	}
	// 获取显示数量
	result["picPath"] = picpath
	result["smallPicPath"] = smallpath
	result["num"] = len(picpath)

	ctx.JSON(http.StatusOK, result)
}

// 获取页数
func GetUserStoragePicturePage(ctx *gin.Context) {
	result := map[string]interface{}{}
	// 检查是否已登录
	if ok := !user.IsLogin(ctx); ok {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 名称
	name := ctx.PostForm("name")
	// 分类
	folderName := ctx.PostForm("file")
	// 每页数量
	pageNums := ctx.PostForm("pageNum")
	pageNum, _ := strconv.Atoi(pageNums)

	// 账号
	email, err := user.SelectUserAccountByName(name)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"storage")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient.Close()
	// 总数
	sumNum := 0
	mysqlClient.Get(&sumNum, "SELECT count(*) FROM storage WHERE account = ? AND filepath = ?", email, folderName)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
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

	result["page"] = page
	ctx.JSON(http.StatusOK, result)
}

// 查看文件夹个数
func GetUserFileNums(ctx *gin.Context) {
	// 返回值
	result := map[string]interface{}{}
	// 检查是否已登录
	if ok := !user.IsLogin(ctx); ok {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 名称
	name := ctx.PostForm("name")
	// 账号
	email, err := user.SelectUserAccountByName(name)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"storage")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient.Close()
	// 获取该用户文件夹个数
	files := []string{}
	err = mysqlClient.Select(&files, "SELECT DISTINCT filename FROM storagefile WHERE account = ?", email)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	result["files"] = files
	ctx.JSON(http.StatusOK, result)
}

// 新建文件夹
func MakeDirectory(ctx *gin.Context) {
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
	// 检查是否已登录
	if ok := !user.IsLogin(ctx); ok {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 昵称
	name, err := user.GetName(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 账号
	email, err := user.GetUserEmail(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 文件夹名称
	path := ctx.PostForm("path")
	// 新建
	os.MkdirAll("./userFile/"+name+"/"+path, 0644)
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"storage")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient.Close()
	// 判断文件夹是否存在
	ret := ""
	err = mysqlClient.Get(&ret, "Select filename FROM storagefile WHERE account = ? AND filename = ?", email, path)
	if err == nil {
		result["msg"] = "文件夹已存在！"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 添加数据库
	_, err = mysqlClient.Exec("INSERT INTO storagefile VALUES(?, ?, ?)", 0, email, path)
	if err != nil {
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}
