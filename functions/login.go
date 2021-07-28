package functions

import (
	"1/Users"
	"net/http"
<<<<<<< HEAD
=======
	"project/Users"
	"project/infomation"
>>>>>>> d19263e... 更新

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

	if statu != "success" {
		ctx.String(http.StatusBadRequest, "登陆失败")
		return
	}

	ctx.SetCookie("cookie", userName, 86400, "/", "localhost:8080", false, true)
	if userName == infomation.SystemUserAccount {
		ctx.HTML(http.StatusOK, "systemUser.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "function.html", nil)
}
