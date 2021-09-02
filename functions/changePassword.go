package functions

import (
	"Project/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangePassWord(ctx *gin.Context) {
	newPassword := ctx.PostForm("userPassword")
	useremail := ctx.PostForm("userEmail")
	code := ctx.PostForm("code")

	userInfo := &Users.User{
		MailAccount: useremail,
	}

	if code != userInfo.GetVerificationCode() {
		return
	}

	userInfo.ChangePassword(newPassword)
	ctx.HTML(http.StatusOK, "login.html", nil)
}
