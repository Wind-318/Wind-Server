package main

import (
<<<<<<< HEAD
	"1/Text"
	"1/functions"
=======
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"project/Mail"
	"project/Text"
	"project/Users"
	"project/functions"
<<<<<<< HEAD
>>>>>>> 9e17b98... /
=======
	"project/httpRequest"
>>>>>>> 0c71c88... 更新
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
	router.StaticFS("/TXT", http.Dir("./TXT"))
	router.GET("/", functions.ToLogin)
	router.GET("/toRegister", functions.ToRegister)
	router.GET("/toChangePassword", functions.ToChangePassword)
	router.GET("/toFindPassword", functions.FindPassword)
	router.GET("/toFunctions", functions.ToFuntions)
	router.GET("/toverificationCode", functions.ToVerificationCode)
	router.GET("/robots.txt", functions.ToRobots)

	router.POST("/register", functions.Register)
	router.POST("/login", functions.Login)
	router.POST("/changePassword", functions.ChangePassWord)
	router.POST("/findPassword", functions.FindPassword)
	router.POST("/sendCode", functions.SendCode)
	router.GET("/sendStock", functions.SendStock)

	router.Run()
}

func countTime() {
	for {
		Text.GenerateText()
		// 1.5 小时到 3.5 小时抓取一次
		result, _ := rand.Int(rand.Reader, big.NewInt(7200))
		time.Sleep(time.Second * time.Duration(result.Int64()+5400))
	}
}

// 6 点、18 点定时推送
func sendEveryUser() {
	db := sqlx.MustConnect("mysql", httpRequest.MySQLInfo)
	defer db.Close()

	for {
		users := Users.SelectUsersAccount()
		for _, user := range users {
			waitToSend := Mail.GetNewMail(user)
			waitToSend.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", Text.SelectFirst10(), gomail.NewMessage())
		}
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
	}
}
