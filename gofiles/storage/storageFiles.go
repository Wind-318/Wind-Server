package storage

import (
	"Project/gofiles/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 存储文件
func StorageFiles(ctx *gin.Context) {
	// 返回数据
	result := map[string]interface{}{
		"msg": "上传成功",
	}
	// 获取 cookie
	cookies, err := ctx.Cookie("cookie")
	// 错误处理
	if err != nil {
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 连接 redis
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	// 延时关闭
	defer redisconn.Close()
	// 判断合法性
	isExist, err := redis.Bool(redisconn.Do("HEXISTS", cookies, "email"))
	// 若账号不存在
	if !isExist || err != nil {
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 重制 cookie
	ctx.SetCookie("cookie", cookies, 86400, "/", "localhost/", false, true)
	// 重设过期时间
	redisconn.Do("EXPIRE", cookies, 86400)
	// 获取账号
	email, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	// 获取文件夹名称
	filename := ctx.PostForm("filename")
	// 文件夹名称为空的处理
	if filename == "" {
		fmt.Println("文件夹名称不能为空!")
		result["msg"] = "文件夹名称不能为空!"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 存储路径
	userPath := "./userFile/" + email + "/" + filename + "/"
	// 渲染路径
	userPaths := "../userFile/" + email + "/" + filename + "/"
	// 创建文件夹
	os.MkdirAll(userPath, 0644)
	// 存储
	res, err := ctx.MultipartForm()
	// 错误处理
	if err != nil {
		fmt.Println(err)
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 文件
	files := res.File["file"]
	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLInfo)
	// 错误处理
	if err != nil {
		fmt.Println(err)
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 延时关闭数据库
	defer conn.Close()
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	// 获取该用户已用空间大小
	usedCapacity := int64(0)
	// 获取该用户被分配空间大小
	capacity := int64(0)

	// 加锁
	mutex := sync.Mutex{}
	mutex.Lock()
	// 获取数据
	err = conn.Get(&usedCapacity, "SELECT usedCapacity FROM user WHERE account = ?", email)
	// 错误处理
	if err != nil {
		fmt.Println(err)
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取被分配空间大小
	err = conn.Get(&capacity, "SELECT capacity FROM user WHERE account = ?", email)
	// 错误处理
	if err != nil {
		fmt.Println(err)
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 处理文件
	for _, file := range files {
		// 检查存储是否满，已满则直接返回
		if usedCapacity+file.Size/1024 > capacity {
			fmt.Println(usedCapacity+file.Size/1024, capacity, "存储已满")
			result["msg"] = "上传失败"
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
			// 插入数据库
			_, err = conn.Exec("INSERT INTO storage VALUES(?, ?, ?, ?, ?, ?, ?, ?)", 0, email, file.Size/1024, filename, randName, picType, userPaths+randName+"."+picType, userPaths+randName+"."+picType)
			// 出错返回
			if err != nil {
				fmt.Println(err)
				result["msg"] = "上传失败"
				ctx.JSON(http.StatusOK, result)
				return
			} else { // 未出错更新数据
				// 更新使用空间大小
				_, err = conn.Exec("UPDATE user SET usedCapacity = ? WHERE account = ?", usedCapacity+file.Size/1024, email)
				// 错误处理
				if err != nil {
					fmt.Println(err)
					result["msg"] = "上传失败"
					ctx.JSON(http.StatusOK, result)
					return
				}
			}
		} else { // 为其它格式时
			picType = file.Filename[length-4:]
			// 插入数据
			_, err = conn.Exec("INSERT INTO storage VALUES(?, ?, ?, ?, ?, ?, ?, ?)", 0, email, file.Size/1024, filename, randName, picType, userPaths+randName+"."+picType, userPaths+randName+"."+picType)
			// 出错返回
			if err != nil {
				fmt.Println(err)
				result["msg"] = "上传失败"
				ctx.JSON(http.StatusOK, result)
				return
			} else { // 未出错更新数据
				_, err = conn.Exec("UPDATE user SET usedCapacity = ? WHERE account = ?", usedCapacity+file.Size/1024, email)
				// 错误处理
				if err != nil {
					fmt.Println(err)
					result["msg"] = "上传失败"
					ctx.JSON(http.StatusOK, result)
					return
				}
			}
		}
		// 解锁
		mutex.Unlock()

		// 存储文件
		err = ctx.SaveUploadedFile(file, userPath+randName+"."+picType)
		// 错误处理
		if err != nil {
			fmt.Println(err)
			result["msg"] = "上传失败"
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
	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 延时关闭
	defer conn.Close()
	// 原图地址
	picPath := []string{}
	// 缩略图地址
	smallPicPath := []string{}
	// 文件类型
	types := []string{}
	// 账号
	accounts := []string{}
	// 文件夹路径
	filenames := []string{}
	// 文件名（不含后缀）
	names := []string{}
	// 选取数据
	err = conn.Select(&types, "SELECT type FROM storage")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Select(&picPath, "SELECT path FROM storage")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Select(&smallPicPath, "SELECT smallpath FROM storage")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Select(&accounts, "SELECT account FROM storage")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Select(&filenames, "SELECT filepath FROM storage")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Select(&names, "SELECT name FROM storage")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 循环
	for i := range types {
		userPath := "./userFile/" + accounts[i] + "/" + filenames[i] + "/"
		userPaths := "../userFile/" + accounts[i] + "/" + filenames[i] + "/"
		// 已有缩略图则跳过
		if picPath[i] != smallPicPath[i] {
			continue
		}
		// 随机起名
		randName := names[i]

		// 创建缩略图
		imgData, err := ioutil.ReadFile(picPath[i][1:])
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
		err = imaging.Save(image, userPath+randName+"small."+types[i])
		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 更新缩略图
		_, err = conn.Exec("UPDATE storage SET smallpath = ? WHERE path = ?", userPaths+randName+"small."+types[i], picPath[i])
		// 错误处理
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
