package main

import (
	"1/functions"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("HTML/*")
	router.GET("/toRegister", functions.ToRegister)
	router.POST("/register", functions.Register)
	router.Run()
}
