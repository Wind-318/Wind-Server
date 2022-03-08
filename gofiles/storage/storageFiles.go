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
	// 连接 redis
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	// 判断合法性
	isExist, err := redis.Bool(redisconn.Do("HEXISTS", cookies, "email"))
	if !isExist || err != nil {
		result["msg"] = "上传失败"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 重制 cookie
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
	conn, err := sqlx.Connect("mysql", config.MySQLInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
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

	for _, file := range files {
		// 检查存储是否满
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
			_, err = conn.Exec("INSERT INTO storage VALUES(?, ?, ?, ?, ?, ?)", 0, email, picType, userPaths+randName+"."+picType, file.Size/1024, userPaths+randName+"."+picType)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				_, err = conn.Exec("UPDATE user SET usedCapacity = ? WHERE account = ?", usedCapacity+file.Size/1024, email)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		} else {
			picType = file.Filename[length-4:]
			_, err = conn.Exec("INSERT INTO storage VALUES(?, ?, ?, ?, ?, ?)", 0, email, picType, userPaths+randName+"."+picType, file.Size/1024, userPaths+randName+"."+picType)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				_, err = conn.Exec("UPDATE user SET usedCapacity = ? WHERE account = ?", usedCapacity+file.Size/1024, email)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}

		conn, err := sqlx.Connect("mysql", config.MySQLInfo)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		// 存储
		err = ctx.SaveUploadedFile(file, userPath+randName+"."+picType)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, result)

	updatePicture()
}

var ch chan int = make(chan int, 1)

// 存储缩略图
func updatePicture() {
	ch <- 1
	defer func() {
		<-ch
	}()
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	conn, err := sqlx.Connect("mysql", config.MySQLInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	picPath := []string{}
	smallPicPath := []string{}
	types := []string{}
	accounts := []string{}
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
	err = conn.Select(&smallPicPath, "SELECT smallpic FROM storage")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Select(&accounts, "SELECT account FROM storage")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(types))
	// 循环
	for i := range types {
		userPath := "./userFile/" + accounts[i] + "/img1/"
		userPaths := "../userFile/" + accounts[i] + "/img1/"
		if picPath[i] != smallPicPath[i] {
			continue
		}
		// 随机起名
		randName := strconv.Itoa(int(time.Now().UnixNano()))

		// 创建缩略图
		imgData, err := ioutil.ReadFile(picPath[i][1:])
		if err != nil {
			fmt.Println(err)
			return
		}
		buf := bytes.NewBuffer(imgData)
		image, err := imaging.Decode(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 图片缩略
		image = imaging.Resize(image, 0, 400, imaging.Lanczos)
		// 保存缩略图
		err = imaging.Save(image, userPath+randName+"small."+types[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		// 更新缩略图
		_, err = conn.Exec("UPDATE storage SET smallpic = ? WHERE path = ?", userPaths+randName+"small."+types[i], picPath[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
