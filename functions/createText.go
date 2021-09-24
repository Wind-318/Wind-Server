package functions

import (
	"Project/infomation"
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
	result := map[string]interface{}{
		"msg": "创建成功",
	}

	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	if err != nil {
		result["msg"] = "请先登录"
		ctx.JSON(http.StatusOK, result)
		return
	}

	text := ctx.PostForm("texts")
	titles := ctx.PostForm("titles")
	description := ctx.PostForm("description")
	types := ctx.PostForm("types")
	authority := ctx.PostForm("authority")
	pic, _ := ctx.FormFile("pic")
	attFile, _ := ctx.MultipartForm()
	attFiles := attFile.File["attFiles"]
	pictype := ctx.PostForm("picType")

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

	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()

	var name string
	var id int
	conn.Get(&id, "SELECT id FROM user WHERE account = ?", cookie)
	conn.Get(&name, "SELECT username FROM user WHERE account = ?", cookie)

	// 创建文件夹
	os.Mkdir(`blog/`+strconv.Itoa(id)+`/`+types, 0644)

	// 文件名中加入当前时间
	var randTime = strconv.Itoa(int(time.Now().UnixNano()))
	var picAddr = infomation.Addr + `blog/` + strconv.Itoa(id) + `/` + types + "/" + randTime + "." + pictype

	// 保存文件
	ctx.SaveUploadedFile(pic, `blog/`+strconv.Itoa(id)+`/`+types+"/"+randTime+"."+pictype)
	for i := range attFiles {
		ctx.SaveUploadedFile(attFiles[i], `blog/`+strconv.Itoa(id)+`/`+types+"/"+attFiles[i].Filename)
	}

	// 创建缩略图
	imgData, _ := ioutil.ReadFile(`blog/` + strconv.Itoa(id) + `/` + types + "/" + randTime + "." + pictype)
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		result["msg"] = "保存文件失败"
		return
	}
	image = imaging.Resize(image, 0, 400, imaging.Lanczos)
	err = imaging.Save(image, `blog/`+strconv.Itoa(id)+`/`+types+"/"+randTime+"small."+pictype)
	if err != nil {
		result["msg"] = "保存文件失败"
		return
	}

	// 插入文章到数据库
	var mutex = &sync.Mutex{}
	mutex.Lock()
	var ids int
	var num int
	conn.Get(&num, "SELECT count(id) FROM blog WHERE authoremail = ? AND types = ?", cookie, types)
	conn.Exec("INSERT INTO blog VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 0, name, cookie, titles, description, text, types, 0, 0, val, time.Now().String()[:19], time.Now().String()[:19], id, infomation.Addr+`blog/`+strconv.Itoa(id)+`/`+types+`/`+strconv.Itoa(num+1)+`.html`, 0, picAddr, infomation.Addr+`blog/`+strconv.Itoa(id)+`/`+types+"/"+randTime+"small."+pictype)
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
	ioutil.WriteFile(`blog/`+strconv.Itoa(id)+`/`+types+`/`+strconv.Itoa(num+1)+`.html`, []byte(htmls), 0644)

	ctx.JSON(http.StatusOK, result)
}

// 获取某一 id 的文章
func GetUserText(ctx *gin.Context) {
	result := map[string]interface{}{
		"content": "",
	}
	content := ""
	id := ctx.PostForm("ids")
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()
	conn.Get(&content, "SELECT content FROM blog WHERE id = ?", id)
	result["content"] = content
	ctx.JSON(http.StatusOK, result)
}

// 获取头像
func GetProfile(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()
	userid := ""
	conn.Get(&userid, "SELECT authorid FROM blog WHERE id = ?", id)
	pic := ""
	conn.Get(&pic, "SELECT pic FROM user WHERE id = ?", userid)
	result := map[string]interface{}{
		"pic": pic,
	}
	ctx.JSON(http.StatusOK, result)
}

// 获取最后一次编辑时间
func GetLastModify(ctx *gin.Context) {
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
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

	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()

	picurl := ""
	conn.Get(&picurl, "SELECT picurl FROM blog WHERE id = ?", id)

	result := map[string]interface{}{
		"picurl": picurl,
	}
	ctx.JSON(http.StatusOK, result)
}
