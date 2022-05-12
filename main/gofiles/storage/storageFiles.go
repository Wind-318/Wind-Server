package storage

import (
	"Project/callClient/callDatabase"
	"Project/callClient/callStorage"
	"Project/callClient/callUser"
	"Project/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

// 并发控制，最大量为 50
var chanInt50 chan int = make(chan int, 50)

// 存储文件
func StorageFiles(ctx *gin.Context) {
	chanInt50 <- 1
	defer func() {
		<-chanInt50
	}()
	// 返回数据
	result := map[string]interface{}{}
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "尚未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取账号
	account, err := callUser.CallUserGetUserEmail(ctx)
	if err != nil {
		result["msg"] = "用户不存在"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取用户昵称
	userName, err := callUser.CallUserGetName(ctx)
	if err != nil {
		result["msg"] = "用户不存在"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取文件夹名称
	filename := ctx.PostForm("filename")
	// 文件夹名称为空的处理
	if filename == "" {
		result["msg"] = "需先进入文件夹上传！"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 存储路径
	userPath := "./userFile/" + userName + "/" + filename + "/"
	// 渲染路径
	userPaths := config.Addr + "userFile/" + userName + "/" + filename + "/"
	// 创建文件夹
	os.MkdirAll(userPath, 0644)
	// 存储
	res, err := ctx.MultipartForm()
	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 文件
	files := res.File["file"]

	// 处理文件
	for _, file := range files {
		// 检查存储是否满，已满则直接返回
		err = callStorage.CallStorageStorageFiles(ctx, file.Size)
		if err != nil {
			result["msg"] = err.Error()
			ctx.JSON(http.StatusOK, result)
			return
		}

		// 文件名长度
		length := len(file.Filename)
		// 文件类型
		picType := file.Filename[length-3:]
		// 随机起名
		randName := strconv.Itoa(int(time.Now().UnixNano()))
		// 文件格式为以下四种时
		if picType == "jpg" || picType == "png" || picType == "gif" || picType == "bmp" {
			picType = file.Filename[length-3:]
		} else { // 为其它格式时
			picType = file.Filename[length-4:]
		}

		// 存储文件
		err = ctx.SaveUploadedFile(file, userPath+randName+"."+picType)
		// 错误处理
		if err != nil {
			result["msg"] = err.Error()
			ctx.JSON(http.StatusOK, result)
			return
		}
		// 插入数据库
		err = callDatabase.CallMySQLInsertStorageNewFile(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "storage", account, filename, randName, picType, userPaths+randName+"."+picType, userPaths+randName+"."+picType, file.Size/1024)
		// 出错返回
		if err != nil {
			result["msg"] = err.Error()
			ctx.JSON(http.StatusOK, result)
			return
		}
	}
	// 返回 json 数据
	ctx.JSON(http.StatusOK, result)

	// 更新缩略图
	go updatePicture()
}

// 并发控制，同时准入一次
var ch chan int = make(chan int, 1)

// 存储缩略图
func updatePicture() {
	ch <- 1
	defer func() {
		<-ch
	}()

	// 选取数据
	arr, err := callDatabase.CallMySQLSelectStorageFilesInfo(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "storage")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 循环
	for i := range arr {
		userName, err := callDatabase.CallMySQLSelectUserNameByAccount(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", arr[i].Account)
		if err != nil {
			fmt.Println(err)
			continue
		}
		userPath := "./userFile/" + userName + "/" + arr[i].Filepath + "/"
		userPaths := config.Addr + "userFile/" + userName + "/" + arr[i].Filepath + "/"
		// 已有缩略图则跳过
		if arr[i].Path != arr[i].Smallpath {
			continue
		}
		// 起名
		randName := arr[i].Name

		// 创建缩略图
		imgData, err := ioutil.ReadFile("./" + arr[i].Path[len(config.Addr):])
		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
		}
		buf := bytes.NewBuffer(imgData)
		image, err := imaging.Decode(buf)
		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 图片缩略
		image = imaging.Resize(image, 0, 400, imaging.Lanczos)
		// 保存缩略图
		err = imaging.Save(image, userPath+randName+"small."+arr[i].Type)
		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 更新缩略图
		err = callDatabase.CallMySQLUpdateStorageSmallPicture(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "storage", userPaths+randName+"small."+arr[i].Type, arr[i].Path)

		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
