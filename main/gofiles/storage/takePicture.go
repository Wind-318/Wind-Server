package storage

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callDatabase"
	"Project/callClient/callStorage"
	"Project/callClient/callUser"
	"Project/config"
	"net/http"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// 批量下载
func DownloadFiles(ctx *gin.Context) {
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
	// 检验合法性
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}

	userAccount, err := callUser.CallUserGetUserEmail(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 加分布式锁
	if ok, err := redis.String(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "SET", userAccount+"DownloadFiles", "DownloadFiles", "EX", 30, "NX")); err != nil || ok != "OK" {
		return
	}
	// 设置检测，每 25 秒检测锁是否存在，存在就延长
	go func() {
		for {
			time.Sleep(time.Second * 25)
			// 锁存在
			if ok, err := redis.Bool(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "SETNX", userAccount+"DownloadFiles", "DownloadFiles")); err != nil || !ok {
				// 设置过期 30 秒
				callDatabase.CallRedis(config.RedisIP, config.RedisPort, "Expire", userAccount+"DownloadFiles", 30)
			} else {
				callDatabase.CallRedis(config.RedisIP, config.RedisPort, "del", userAccount+"DownloadFiles")
				return
			}
		}
	}()
	defer callDatabase.CallRedis(config.RedisIP, config.RedisPort, "del", userAccount+"DownloadFiles")

	result, err = callStorage.CallStorageDownloadFiles(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	} else if _, ok := result["msg"]; ok {
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 压缩
	err = Zip(result["zipfile1"].(string), result["zipfile2"].(string))
	if err != nil {
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}

// 获取指定用户上传的图片
func GetUserStoragePicture(ctx *gin.Context) {
	// 返回值
	result := map[string]interface{}{}
	// 检验合法性
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	var err error
	result["picPath"], result["smallPicPath"], result["num"], err = callStorage.CallStorageGetUserStoragePicture(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// 获取页数
func GetUserStoragePicturePage(ctx *gin.Context) {
	result := map[string]interface{}{}
	// 检验合法性
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	var err error
	result["page"], err = callStorage.CallStorageGetUserStoragePicturePage(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// 查看文件夹个数
func GetUserFileNums(ctx *gin.Context) {
	// 返回值
	result := map[string]interface{}{}
	// 检验合法性
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	var err error
	result["files"], err = callStorage.CallStorageGetUserFileNums(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// 新建文件夹
func MakeDirectory(ctx *gin.Context) {
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
	// 检验合法性
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 昵称
	name, err := callUser.CallUserGetName(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 账号
	email, err := callUser.CallUserGetUserEmail(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 文件夹名称
	path := ctx.PostForm("path")
	// 新建
	os.MkdirAll("./userFile/"+name+"/"+path, 0644)
	// 判断文件夹是否存在
	err = callDatabase.CallMySQLSelectIsExistFolder(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "storage", email, path)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 添加数据库
	err = callDatabase.CallMySQLInsertStorageNewFolder(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "storage", email, path)
	if err != nil {
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}
