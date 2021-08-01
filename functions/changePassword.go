package functions

import (
	"Project/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangePassWord(ctx *gin.Context) {
	userEmail, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}
	newPassword := ctx.PostForm("newPassword")

	userInfo := &Users.User{
		MailAccount: userEmail,
	}
	userInfo.ChangePassword(newPassword)
	ctx.HTML(http.StatusOK, "login.html", nil)
}
