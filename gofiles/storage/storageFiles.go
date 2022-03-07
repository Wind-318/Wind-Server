package storage

import (
	"Project/gofiles/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
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
	result := map[string]interface{}{
		"msg": "上传成功",
	}
	cookies, err := ctx.Cookie("cookie")
	if err != nil {
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	isExist, err := redis.Bool(redisconn.Do("HEXISTS", cookies, "email"))
	if !isExist || err != nil {
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	ctx.SetCookie("cookie", cookies, 86400, "/", "localhost/", false, true)
	redisconn.Do("EXPIRE", cookies, 86400)
	// 获取账号
	email, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	// 存储路径
	userPath := "./userFile/" + email + "/img1/"
	// 渲染路径
	userPaths := "../userFile/" + email + "/img1/"
	// 创建文件夹
	os.MkdirAll(userPath, 0644)
	// 存储
	res, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println(err)
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	files := res.File["file"]
	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	// 获取该用户已用空间大小和可用空间大小
	usedCapacity := int64(0)
	capacity := int64(0)
	err = conn.Get(&usedCapacity, "SELECT usedCapacity FROM user WHERE account = ?", email)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Get(&capacity, "SELECT capacity FROM user WHERE account = ?", email)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 并发控制
	control := make(chan int, 5)
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			control <- 1
			defer func() {
				<-control
				wg.Done()
			}()
			// 加锁
			rwmutex := sync.RWMutex{}
			// 加锁
			rwmutex.Lock()
			if usedCapacity+file.Size/1024 > capacity {
				fmt.Println(usedCapacity+file.Size/1024, capacity, "存储已满")
				return
			}

			// 文件名长度
			length := len(file.Filename)
			// 文件类型
			picType := file.Filename[length-3:]
			// 随机起名
			randName := strconv.Itoa(int(time.Now().UnixNano()))
			if picType == "jpg" || picType == "png" || picType == "gif" || picType == "bmp" {
				_, err = conn.Exec("INSERT INTO storage VALUES(?, ?, ?, ?, ?, ?)", 0, email, picType, userPaths+randName+"."+picType, file.Size/1024, userPaths+randName+"small."+picType)
				if err != nil {
					fmt.Println(err)
				} else {
					usedCapacity += file.Size / 1024
					_, err = conn.Exec("UPDATE user SET usedCapacity = ? WHERE account = ?", usedCapacity, email)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			} else {
				picType = file.Filename[length-4:]
				_, err = conn.Exec("INSERT INTO storage VALUES(?, ?, ?, ?, ?, ?)", 0, email, picType, userPaths+randName+"."+picType, file.Size/1024, userPaths+randName+"small."+picType)
				if err != nil {
					fmt.Println(err)
				} else {
					usedCapacity += file.Size / 1024
					_, err = conn.Exec("UPDATE user SET usedCapacity = ? WHERE account = ?", usedCapacity, email)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}

			// 解锁
			rwmutex.Unlock()
			// 存储
			err = ctx.SaveUploadedFile(file, userPath+randName+"."+picType)
			if err != nil {
				fmt.Println(err)
				return
			}
			isError := false
			// 创建缩略图
			imgData, _ := ioutil.ReadFile(userPath + randName + "." + picType)
			buf := bytes.NewBuffer(imgData)
			image, err := imaging.Decode(buf)
			if err != nil {
				isError = true
				fmt.Println(err)
				return
			}
			// 图片缩略
			image = imaging.Resize(image, 0, 400, imaging.Lanczos)
			// 保存缩略图
			err = imaging.Save(image, userPath+randName+"small."+picType)
			if err != nil {
				isError = true
				fmt.Println(err)
				return
			}
			if isError {
				conn.Exec("UPDATE storage SET smallpic = ?", userPaths+randName+"."+picType)
			}
		}(file)
	}
	ctx.JSON(http.StatusOK, result)
	// 等待所有协程结束
	wg.Wait()
}
