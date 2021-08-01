package main

import (
	"Project/Mail"
	"Project/Text"
	"Project/Users"
	"Project/functions"
	"Project/infomation"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/jmoiron/sqlx"
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

	go countTime()
	go sendEveryUser()

	router.LoadHTMLGlob("HTML/*")
	router.StaticFS("/js", http.Dir("./js"))
	router.StaticFS("/css", http.Dir("./css"))
	router.StaticFS("/music", http.Dir("./music"))
	router.StaticFS("/picture", http.Dir("./picture"))
	router.StaticFS("/markdown", http.Dir("./markdown"))
	router.StaticFile("/robots.txt", "./robots.txt")

	router.NoRoute(functions.ToNotFound)
	router.GET("/", functions.ToHead)
	router.GET("/ToLogin", functions.ToLogin)
	router.GET("/TochangePassword", functions.ToChangePassword)
	router.GET("/ToFindPassword", functions.ToFindPassword)
	router.GET("/ToFunction", functions.ToFunction)
	router.GET("/ToVerificationCode", functions.ToVerificationCode)
	router.GET("/Exit", functions.Exit)

	router.POST("/modifyPassword", functions.ModifyPassword)
	router.POST("/ConfirmFind", functions.ConfirmFind)
	router.POST("/verificationFind", functions.VerificationFind)
	router.POST("/register", functions.Register)
	router.POST("/login", functions.Login)
	router.POST("/changePassword", functions.ChangePassWord)
	router.POST("/sendCode", functions.SendCode)
	router.POST("/upload", functions.Upload)
	router.GET("/sendStock", functions.SendStock)

	router.Run(":80")
}

func countTime() {
	for {
		Text.GenerateText()
		// 1.5 小时到 3.5 小时抓取一次
		rand.Seed(time.Now().UnixNano())
		result := rand.Intn(7200) + 5400
		time.Sleep(time.Second * time.Duration(result))
	}
}

// 6 点、18 点定时推送
func sendEveryUser() {
	db := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer db.Close()

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
