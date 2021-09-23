package main

import (
	"Project/Mail"
	"Project/Text"
	"Project/Users"
	"Project/functions"
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
	router := gin.New()
	gin.DisableConsoleColor()

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
	router.Use(TlsHandler())

	go countTime()
	go sendEveryUser()

	router.LoadHTMLGlob("HTML/*")
	router.StaticFS("/blog", http.Dir("./blog"))
	router.StaticFS("/css", http.Dir("./css"))
	router.StaticFS("/files", http.Dir("./files"))
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.StaticFS("/js", http.Dir("./js"))
	router.StaticFS("/music", http.Dir("./music"))
	router.StaticFS("/picture", http.Dir("./picture"))
	router.StaticFile("/robots.txt", "./robots.txt")

	router.NoRoute(functions.ToNotFound)
	root := router.Group("")
	{
		root.GET("/", functions.ToHead)
		root.Any("/blog", functions.ToNotFound)
		root.Any("/css", functions.ToNotFound)
		root.Any("/files", functions.ToNotFound)
		root.Any("/js", functions.ToNotFound)
		root.Any("/music", functions.ToNotFound)
		root.Any("/picture", functions.ToNotFound)

		root.GET("/sendStock", functions.SendStock)
		root.GET("/serverError", functions.ToError)
	}

	blog := router.Group("/blogs")
	{
		blog.GET("/", functions.ToBlog)
		blog.GET("/CreateText", functions.ToCreateText)
		blog.GET("/InquireClassification", functions.GetClassification)
		blog.POST("/AddComment", functions.AddComment)
		blog.POST("/Author", functions.Author)
		blog.POST("/CreateTexts", functions.CreateText)
		blog.POST("/DeleteBlog", functions.DeleteFromBlog)
		blog.POST("/DeleteComment", functions.DeleteComment)
		blog.POST("/GetCommentsID", functions.GetCommentsID)
		blog.POST("/GetUserText", functions.GetUserText)
		blog.POST("/GetProfile", functions.GetProfile)
		blog.POST("/GetLastModify", functions.GetLastModify)
		blog.POST("/GetModifyBlog", functions.GetModifyBlog)
		blog.POST("/Getpicurl", functions.Getpicurl)
		blog.POST("/InquirePageNums", functions.GetPageNums)
		blog.POST("/InquireText", functions.GetText)
		blog.POST("/ModifyBlog", functions.ModifyBlog)
		blog.POST("/Parise", functions.Parise)
		blog.POST("/PariseNum", functions.PariseNum)
		blog.POST("/Search", functions.Search)
		blog.POST("/TextComment", functions.TextComment)
		blog.POST("/Views", functions.Views)
	}

	user := router.Group("/user")
	{
		user.GET("/Exit", functions.Exit)
		user.GET("/ToLogin", functions.ToLogin)
		user.GET("/TochangePassword", functions.ToChangePassword)

		user.POST("/changePassword", functions.ChangePassWord)
		user.POST("/login", functions.Login)
		user.POST("/register", functions.Register)
		user.POST("/sendCode", functions.SendCode)
		user.POST("/verificationFind", functions.VerificationFind)
	}

	collections := router.Group("/collections")
	{
		collections.GET("/", functions.ToCollections)
		collections.GET("/GetWebs", functions.GetWebs)
		collections.GET("/IsSystem", functions.IsSystem)
		collections.POST("/IsSystems", functions.IsSystems)
		collections.POST("/PutWebs", functions.PutWebs)
		collections.POST("/PutPic", functions.PutPic)
	}

	Resources := router.Group("/resources")
	{
		Resources.GET("/", functions.ToResources)
	}

	router.Run(":80")
	// 自行选择 SSL 证书
	router.RunTLS(":443", "windserver.top.pem", "windserver.top.key")
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "https://windserver.top:443",
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
		Text.GenerateText()
		// 1.5 小时到 3.5 小时抓取一次
		rand.Seed(time.Now().UnixNano())
		result := rand.Intn(7200) + 5400
		time.Sleep(time.Second * time.Duration(result))
	}
}

// 6 点和 18 点发送给用户
func sendEveryUser() {
	for {
		nowHour, nowMinute := time.Now().Hour(), time.Now().Minute()
		waitSeconds := 0

		if nowHour < 18 && nowHour >= 6 {
			waitSeconds += (17-nowHour)*3600 + (60-nowMinute)*60
		} else if nowHour >= 18 {
			waitSeconds += (23-nowHour)*3600 + (60-nowMinute)*60 + 6*3600
		} else {
			waitSeconds += (5-nowHour)*3600 + (60-nowMinute)*60
		}

		time.Sleep(time.Second * time.Duration(waitSeconds))
		users := Users.SelectUsersAccount()
		for _, user := range users {
			waitToSend := Mail.GetNewMail(user)
			waitToSend.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", Text.SelectFirst10(), gomail.NewMessage())
		}
	}
}
