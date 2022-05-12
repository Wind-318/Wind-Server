package storage

import (
	"Project/config"
	"Project/gofiles/info"
	"Project/gofiles/user"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var mutex sync.Mutex

// 存储文件
func StorageFiles(ctx *gin.Context) {
	// 返回数据
	result := map[string]interface{}{}
	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer redisCli.Close()
	// 连接 storage
	mysqlClient1, err := sqlx.Connect("mysql", config.MySQLInfo+"storage")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient1.Close()
	// 连接 user
	mysqlClient2, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient2.Close()
	// 判断登录
	if ok := user.IsLogin(ctx); !ok {
		result["msg"] = "尚未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取账号
	account, err := user.GetUserEmail(ctx)
	if err != nil {
		result["msg"] = "用户不存在"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取用户昵称
	userName, err := user.GetName(ctx)
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
		// 加锁
		mutex.Lock()
		// 从缓存读剩余可用空间，同时只有一个线程可以读取，若未访问到，刷入缓存中，其余线程就能直接读取到；
		remainSpace, err := redis.Int64(redisCli.Do("GET", account+"StorageSpace"))
		// 缓存未命中
		if err != nil {
			// 在数据库中查询
			var space int64
			err := mysqlClient2.Get(&space, "SELECT unusedCapacity FROM user WHERE account = ?", account)
			if err != nil {
				// 解锁
				mutex.Unlock()
				result["msg"] = err.Error()
				ctx.JSON(http.StatusOK, result)
				return
			}
			remainSpace = space
			// 写入缓存
			redisCli.Do("SET", account+"StorageSpace", remainSpace)
		}
		// 解锁
		mutex.Unlock()

		// 文件
		var nowSpace int64
		// 获取文件大小
		useSpace := -file.Size / 1024

		// 减少可用空间
		_, err = mysqlClient2.Exec("Update user Set unusedCapacity = unusedCapacity - ? WHERE account = ?", -useSpace, account)
		if err != nil {
			result["msg"] = err.Error()
			ctx.JSON(http.StatusOK, result)
			return
		}
		// 从缓存中减去文件大小
		nowSpace, err = redis.Int64(redisCli.Do("INCRBY", account+"StorageSpace", useSpace))
		if err != nil {
			result["msg"] = err.Error()
			// 补偿，退款！
			mysqlClient2.Exec("Update user Set unusedCapacity = unusedCapacity - ? WHERE account = ?", useSpace, account)
			ctx.JSON(http.StatusOK, result)
			return
		}
		// 空间不足时，淘汰缓存，等待刷入数据库的新数据
		if nowSpace < 0 {
			// 删去缓存
			redisCli.Do("del", account+"StorageSpace")
			// 补偿，退款！
			mysqlClient2.Exec("Update user Set unusedCapacity = unusedCapacity - ? WHERE account = ?", useSpace, account)
			// 延时双删，淘汰可能刷入的旧数据
			go func() {
				redisClis, err := redis.Dial("tcp", config.RedisInfo)
				if err != nil {
					return
				}
				defer redisCli.Close()
				time.Sleep(time.Millisecond * 500)
				redisClis.Do("del", account+"StorageSpace")
			}()
			result["msg"] = "剩余空间不足！"
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
		_, err = mysqlClient1.Exec("INSERT INTO storage VALUES(?, ?, ?, ?, ?, ?, ?, ?)", 0, account, file.Size, filename, randName, picType, userPaths+randName+"."+picType, userPaths+randName+"."+picType)
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

// 存储缩略图
func updatePicture() {
	// 加分布式锁
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return
	}
	defer redisCli.Close()
	if ok, err := redis.String(redisCli.Do("SET", "updatePicture", "updatePicture", "EX", 60, "NX")); err != nil || ok != "OK" {
		return
	}
	// 设置检测，每 50 秒检测锁是否存在，存在就延长
	go func() {
		for {
			time.Sleep(time.Second * 50)
			// 锁存在
			if ok, err := redis.Bool(redisCli.Do("SETNX", "updatePicture", "updatePicture")); err != nil || !ok {
				// 设置过期 60 秒
				redisCli.Do("Expire", "updatePicture", 60)
			} else {
				redisCli.Do("del", "updatePicture")
				return
			}
		}
	}()
	defer redisCli.Do("del", "updatePicture")

	mysqlCli, err := sqlx.Connect("mysql", config.MySQLInfo+"storage")
	if err != nil {
		return
	}
	defer mysqlCli.Close()
	// 选取数据
	arr := []info.StorageFileData{}
	err = mysqlCli.Select(&arr, "SELECT * FROM storage")
	if err != nil {
		return
	}

	// 循环
	for i := range arr {
		userName, err := user.SelectUserNameByAccount(arr[i].UserAccount)
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
		randName := arr[i].FileName

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
		_, err = mysqlCli.Exec("UPDATE storage SET smallpath = ? WHERE path = ?", userPaths+randName+"small."+arr[i].Type, arr[i].Path)

		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
