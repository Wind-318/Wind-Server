package functions

import (
	"Project/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	userEmail := ctx.PostForm("userEmail")
	passWord := ctx.PostForm("userPassword")

	userInfo := &Users.User{
		MailAccount:  userEmail,
		MailPassword: passWord,
	}

	status := userInfo.Login()
	if status != "success" {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}

	ctx.SetCookie("cookie", userEmail, 86400, "/", "localhost/", false, true)

	ctx.HTML(http.StatusOK, "function.html", nil)
}
