package main

import (
	"Project/gofiles"
	"Project/gofiles/blogfunc"
	"Project/gofiles/collectionfunc"
	"Project/gofiles/ownmail"
	"Project/gofiles/spider/sina"
	"Project/gofiles/survive"
	"Project/gofiles/user"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/unrolled/secure"
)

func main() {
	// 新建路由
	router := gin.New()
	// 关闭控制台颜色
	gin.DisableConsoleColor()

	// 设置日志参数、路径
	f, _ := os.OpenFile("logs.log", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	// http 重定向 https
	router.Use(TlsHandler())

	// 计时抓取信息到数据库
	go countTime()

	// 发送邮件
	go sendEveryUser()

	// 加载静态资源
	router.LoadHTMLGlob("HTML/*")
	router.StaticFS("/blog", http.Dir("./blog"))
	router.StaticFS("/css", http.Dir("./css"))
	router.StaticFS("/files", http.Dir("./files"))
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.StaticFS("/js", http.Dir("./js"))
	router.StaticFS("/music", http.Dir("./music"))
	router.StaticFS("/picture", http.Dir("./picture"))
	router.StaticFile("/robots.txt", "./robots.txt")

	// 设置 404 界面
	router.NoRoute(gofiles.ToNotFound)

	// 设置根路由
	root := router.Group("")
	{
		root.GET("/", gofiles.ToHead)
		root.Any("/blog", gofiles.ToNotFound)
		root.Any("/css", gofiles.ToNotFound)
		root.Any("/files", gofiles.ToNotFound)
		root.Any("/js", gofiles.ToNotFound)
		root.Any("/music", gofiles.ToNotFound)
		root.Any("/picture", gofiles.ToNotFound)

		root.GET("/sendStock", sina.SendStock)
		root.GET("/serverError", gofiles.ToError)
	}

	// 设置博客路由
	blog := router.Group("/blogs")
	{
		blog.GET("/", gofiles.ToBlog)
		blog.GET("/CreateText", gofiles.ToCreateText)
		blog.GET("/InquireClassification", blogfunc.GetClassification)
		blog.POST("/AddComment", blogfunc.AddComment)
		blog.POST("/Author", blogfunc.Author)
		blog.POST("/CreateTexts", blogfunc.CreateText)
		blog.POST("/DeleteBlog", blogfunc.DeleteFromBlog)
		blog.POST("/DeleteComment", blogfunc.DeleteComment)
		blog.POST("/GetCommentsID", blogfunc.GetCommentsID)
		blog.POST("/GetUserText", blogfunc.GetUserText)
		blog.POST("/GetProfile", blogfunc.GetProfile)
		blog.POST("/GetLastModify", blogfunc.GetLastModify)
		blog.POST("/GetModifyBlog", blogfunc.GetModifyBlog)
		blog.POST("/Getpicurl", blogfunc.Getpicurl)
		blog.POST("/InquirePageNums", blogfunc.GetPageNums)
		blog.POST("/InquireText", blogfunc.GetText)
		blog.POST("/ModifyBlog", blogfunc.ModifyBlog)
		blog.POST("/Parise", blogfunc.Parise)
		blog.POST("/PariseNum", blogfunc.PariseNum)
		blog.POST("/Search", blogfunc.Search)
		blog.POST("/TextComment", blogfunc.TextComment)
		blog.POST("/Views", blogfunc.Views)
	}

	// 用户路由
	users := router.Group("/user")
	{
		users.GET("/Exit", gofiles.Exit)
		users.GET("/ToLogin", gofiles.ToLogin)
		users.GET("/TochangePassword", gofiles.ToChangePassword)

		users.POST("/changePassword", user.ChangePassWord)
		users.POST("/login", user.Login)
		users.POST("/register", user.Register)
		users.POST("/sendCode", sina.SendCode)
		users.POST("/verificationFind", user.VerificationFind)
		users.POST("/UploadProfile", blogfunc.UploadProfile)
	}

	// 收藏路由
	collections := router.Group("/collections")
	{
		collections.GET("/", gofiles.ToCollections)
		collections.GET("/GetWebs", collectionfunc.GetWebs)
		collections.GET("/IsSystem", collectionfunc.IsSystem)
		collections.POST("/IsSystems", collectionfunc.IsSystems)
		collections.POST("/PutWebs", collectionfunc.PutWebs)
		collections.POST("/PutPic", collectionfunc.PutPic)
	}

	// 资源路由
	Resources := router.Group("/resources")
	{
		Resources.GET("/", gofiles.ToResources)
	}

	Survive := router.Group("/survive")
	{
		Survive.GET("/getproperty", survive.Get)
		Survive.GET("/add", survive.Add)
	}

	// 监听 http
	go router.Run(":80")
	// 监听 https，自行选择 SSL 证书
	router.RunTLS(":443", "windserver.top.pem", "windserver.top.key")
}

// 重定向到 https
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "windserver.top:443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}

// 计时抓取存到数据库
func countTime() {
	for {
		sina.GenerateText()
		// 1.5 小时到 3.5 小时抓取一次
		rand.Seed(time.Now().UnixNano())
		result := rand.Intn(7200) + 5400
		time.Sleep(time.Second * time.Duration(result))
	}
}

// 6 点和 18 点发送给用户
func sendEveryUser() {
	for {
		// 得到现在时间
		nowHour, nowMinute := time.Now().Hour(), time.Now().Minute()
		// 等待时间
		waitSeconds := 0

		if nowHour < 18 && nowHour >= 6 {
			waitSeconds += (17-nowHour)*3600 + (60-nowMinute)*60
		} else if nowHour >= 18 {
			waitSeconds += (23-nowHour)*3600 + (60-nowMinute)*60 + 6*3600
		} else {
			waitSeconds += (5-nowHour)*3600 + (60-nowMinute)*60
		}

		time.Sleep(time.Second * time.Duration(waitSeconds))
		// 得到订阅用户名单
		users := user.SelectUsersAccount()
		// 发送邮件
		for _, user := range users {
			waitToSend := ownmail.GetNewMail(user)
			waitToSend.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", sina.SelectFirst10(), gomail.NewMessage())
		}
	}
}
