package bbs

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"Project/gofiles/user"
	"bytes"
	"io/ioutil"
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

// 创建一篇文章
func CreateText(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 返回信息
	result := map[string]interface{}{
		"msg": "success",
	}

	// 检查登录状态
	if !user.IsLogin(ctx) {
		result["msg"] = "未登录"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 获取账号
	cookie, err := user.GetUserEmail(ctx)
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
	defer redisCli.Close()
	if ok, err := redis.String(redisCli.Do("SET", cookie+"CreateText", "CreateText", "EX", 10, "NX")); err != nil || ok != "OK" {
		return
	}
	// 设置检测，每 8 秒检测锁是否存在，存在就延长
	go func() {
		for {
			time.Sleep(time.Second * 8)
			// 锁存在
			if ok, err := redis.Bool(redisCli.Do("SETNX", cookie+"CreateText", "CreateText")); err != nil || !ok {
				// 设置过期 10 秒
				redisCli.Do("Expire", cookie+"CreateText", 10)
			} else {
				redisCli.Do("del", cookie+"CreateText")
				return
			}
		}
	}()
	defer redisCli.Do("del", cookie+"CreateText")

	// 获取内容
	text := ctx.PostForm("texts")
	// 标题
	titles := ctx.PostForm("titles")
	// 简介
	description := ctx.PostForm("description")
	// 分类
	types := ctx.PostForm("types")
	// 权限
	authority := ctx.PostForm("authority")
	// 图片
	pic, _ := ctx.FormFile("pic")
	// 图片类型
	pictype := ctx.PostForm("picType")
	// 附件
	attFile, _ := ctx.MultipartForm()
	attFiles := attFile.File["attFiles"]

	// 数据检验
	val, err := strconv.Atoi(authority)
	if err != nil {
		result["msg"] = "权限等级只能为数字"
		ctx.JSON(http.StatusOK, result)
		return
	} else if text == "" || titles == "" || types == "" {
		result["msg"] = "文章或分类不能为空"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo+"bbs")
	defer conn.Close()

	tempid, err := user.SelectUserIDByAccount(cookie)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	id, _ := strconv.Atoi(tempid)
	name, err := user.SelectUserNameByAccount(cookie)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 创建文件夹
	os.Mkdir(`bbsFile/`+strconv.Itoa(id)+`/`+types, 0644)

	randtime := strconv.Itoa(int(time.Now().UnixNano()))
	go func() {
		// 保存文件
		err = ctx.SaveUploadedFile(pic, `bbsFile/`+strconv.Itoa(id)+`/`+types+"/"+randtime+"."+pictype)
		if err != nil {
			return
		}
		// 创建缩略图
		imgData, _ := ioutil.ReadFile(`bbsFile/` + strconv.Itoa(id) + `/` + types + "/" + randtime + "." + pictype)
		buf := bytes.NewBuffer(imgData)
		image, err := imaging.Decode(buf)
		if err != nil {
			return
		}
		// 图片缩略
		image = imaging.Resize(image, 0, 400, imaging.Lanczos)
		// 保存缩略图
		err = imaging.Save(image, `bbsFile/`+strconv.Itoa(id)+`/`+types+"/"+randtime+"small."+pictype)
		if err != nil {
			return
		}

		// 保存文件
		for i := range attFiles {
			err = ctx.SaveUploadedFile(attFiles[i], `bbsFile/`+strconv.Itoa(id)+`/`+types+"/"+attFiles[i].Filename)
			if err != nil {
				return
			}
		}
	}()

	// 加锁
	var mutex = &sync.Mutex{}
	mutex.Lock()
	var ids int
	// 起名
	num := strconv.Itoa(int(time.Now().UnixNano()))
	// 插入文章到数据库
	conn.Exec("INSERT INTO blog VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 0, name, cookie, titles, description, text, types, 0, 0, val, time.Now().String()[:19], time.Now().String()[:19], id, config.Addr+`bbsFile/`+strconv.Itoa(id)+`/`+types+`/`+num+`.html`, 0, config.Addr+`bbsFile/`+strconv.Itoa(id)+`/`+types+"/"+randtime+"."+pictype, config.Addr+`bbsFile/`+strconv.Itoa(id)+`/`+types+"/"+randtime+"small."+pictype)
	conn.Get(&ids, "select id from blog order by id DESC limit 1")
	mutex.Unlock()

	htmls := `<!DOCTYPE html>
	<html lang="en">
		
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title id="titles">` + titles + `</title>
			<script src="../../../js/sakura.js"></script>
			<script src="../../../js/marked.min.js"></script>
			<script src="../../../js/jquery.min.js"></script>
			<link rel="stylesheet" href="../../../css/markdowncss.css">
			<link rel="stylesheet" href="../../../css/content.css">
			</head>
			
			<body>
			<div id="root">
				<div style="flex-direction:row; display: flex;">
					<div style="display: flex; margin-left: 100px; width: 200px; flex-direction: column; background-color: rgba(122, 122, 122, 0.6); position: relative;">
						<img alt="" id="profile" style="border-radius: 50%; width: auto; height: auto;">
						<span id="author" style="color: rgb(255, 255, 255, 0.8); margin: 0 auto; text-align: center;"></span>
						<span id="views" style="color: rgb(255, 255, 255, 0.8); margin: 0 auto; text-align: center;"></span>
						<span id="lastmodify" style="color: rgb(255, 255, 255, 0.8); margin: 0 auto; text-align: center;"></span>

						<div style="bottom: 0; position: absolute; width: 100%; height: auto; flex-direction: row; display: flex;">
							<button id="praise" style="cursor: pointer; width: 60px; height: 60px; margin-left: 10px; flex-direction: row; display: flex; margin-bottom: 0;">
								<img src="../../../picture/praise.png" alt="">
								<span id="praiseNum" style="color: white; margin-bottom: 0; margin-left: 5px;"></span>
							</button>
							<button id="reply" style="cursor: pointer; width: 60px; height: 60px; border-radius: 50%; margin-left: 50px; background-color: rgb(50, 75, 150);">
								<span style="color: white;">回复</span>
							</button>
						</div>
					</div>

					<div class="divcontainer" id="` + strconv.Itoa(ids) + `" name="main">
						<div id="contentText" style="flex-direction: column; width: 1000px; margin: 0 auto; background-color: rgba(255, 255, 255, 0.7);" class="contents"></div>
					</div>
				</div>

				<div style="margin-left: 100px; width: 1200px; height: 20px; background-color: rgb(82, 60, 145);"></div>
			</div>

			<script src="../../../js/text.js"></script>

			<script>
				replyjs();
			</script>

			<script>
				adddelete();
			</script>
			</body>
			
			</html>`

	// 写入文件
	ioutil.WriteFile(`bbsFile/`+strconv.Itoa(id)+`/`+types+`/`+num+`.html`, []byte(htmls), 0644)

	// 返回
	ctx.JSON(http.StatusOK, result)
}

// 获取某一 id 的文章
func GetUserText(ctx *gin.Context) {
	// 返回值
	result := map[string]interface{}{
		"content": "",
	}
	content := ""
	id := ctx.PostForm("ids")
	conn := sqlx.MustConnect("mysql", config.MySQLInfo+"bbs")
	defer conn.Close()
	conn.Get(&content, "SELECT content FROM blog WHERE id = ?", id)
	title := ""
	conn.Get(&title, "SELECT title FROM blog WHERE id = ?", id)
	result["content"] = content
	result["title"] = title
	ctx.JSON(http.StatusOK, result)
}

// 获取头像
func GetProfile(ctx *gin.Context) {
	result := map[string]interface{}{
		"pic": "",
	}
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLInfo+"bbs")
	conn1 := sqlx.MustConnect("mysql", config.MySQLInfo+"user")
	defer conn.Close()
	defer conn1.Close()
	userid := ""
	conn.Get(&userid, "SELECT authorid FROM blog WHERE id = ?", id)
	pic := ""
	conn1.Get(&pic, "SELECT pic FROM user WHERE id = ?", userid)
	result["pic"] = pic
	ctx.JSON(http.StatusOK, result)
}

// 获取最后一次编辑时间
func GetLastModify(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", config.MySQLInfo+"bbs")
	defer conn.Close()
	lastmodify := ""
	conn.Get(&lastmodify, "SELECT update_time FROM blog WHERE id = ?", id)
	result := map[string]interface{}{
		"lastmodify": lastmodify,
	}
	ctx.JSON(http.StatusOK, result)
}

// 获取图片 url
func Getpicurl(ctx *gin.Context) {
	id := ctx.PostForm("id")

	conn := sqlx.MustConnect("mysql", config.MySQLInfo+"bbs")
	defer conn.Close()

	picurl := ""
	conn.Get(&picurl, "SELECT picurl FROM blog WHERE id = ?", id)

	result := map[string]interface{}{
		"picurl": picurl,
	}
	ctx.JSON(http.StatusOK, result)
}
