package functions

import (
	"Project/infomation"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

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

	var mutex = &sync.Mutex{}
	mutex.Lock()
	var ids int
	var num int
	conn.Get(&num, "SELECT count(id) FROM blog WHERE authoremail = ? AND types = ?", cookie, types)
	conn.Exec("INSERT INTO blog VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 0, name, cookie, titles, description, text, types, 0, 0, val, time.Now().String()[:19], time.Now().String()[:19], id, infomation.Addr+`blog/`+strconv.Itoa(id)+`/`+types+`/`+strconv.Itoa(num+1)+`.html`, 0, picAddr)
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
				<div class="imgs">
					<img src="` + picAddr + `" alt="">
				</div>
				<div class="divcontainer">
					<div id="content" style="flex-direction: column; width: 800px; margin: 0 auto; background-color: rgba(255, 255, 255, 0.6);" class="contents"></div>
				</div>
				
				<script>
					$.ajax({
						url:"/blogs/GetUserText",
						type:"POST",
						data:{
							id:` + strconv.Itoa(ids) + `
						},
						success:function(data) {
							document.getElementById('content').innerHTML = marked(data["content"]);
						},
					})
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
	id := ctx.PostForm("id")
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()
	conn.Get(&content, "SELECT content FROM blog WHERE id = ?", id)
	result["content"] = content
	ctx.JSON(http.StatusOK, result)
}
