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
	"os"
	"project/Text"
	"project/functions"
>>>>>>> 9e17b98... /
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
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

	router.LoadHTMLGlob("HTML/*")
	router.GET("/", functions.ToLogin)
	router.GET("/toRegister", functions.ToRegister)
	router.GET("/toChangePassword", functions.ToChangePassword)
	router.GET("/toFindPassword", functions.FindPassword)
	router.GET("/toFunctions", functions.ToFuntions)
	router.GET("/toverificationCode", functions.ToVerificationCode)

	router.POST("/register", functions.Register)
	router.POST("/login", functions.Login)
	router.POST("/changePassword", functions.ChangePassWord)
	router.POST("/findPassword", functions.FindPassword)
	router.POST("/sendCode", functions.SendCode)
	router.GET("/sendStock", functions.SendStock)

	router.Run()
}

func countTime() {
	Text.GenerateText()
	for {
		// 1.5 小时到 3.5 小时抓取一次
		result, _ := rand.Int(rand.Reader, big.NewInt(7200))
		<-time.After(time.Hour * time.Duration(result.Int64()+5400))
		Text.GenerateText()
	}
}
