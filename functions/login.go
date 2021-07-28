package functions

import (
	"1/Users"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func ToLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func Login(ctx *gin.Context) {
	userName := ctx.PostForm("userName")
	passWord := ctx.PostForm("passWord")

	userInfo := &Users.User{
		MailAccount:  userName,
		MailPassword: passWord,
	}

	statu := userInfo.Login()

	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if statu != "success" {
		ctx.String(http.StatusOK, "登陆失败")
		return
	}

	ctx.SetCookie("cookie", userName, 86400, "/", "localhost:8080", false, true)
	ctx.HTML(http.StatusOK, "function.html", nil)
}
