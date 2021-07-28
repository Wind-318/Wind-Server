package functions

import (
	"Project/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPassword(ctx *gin.Context) {
	userName := ctx.PostForm("userName")
	userInfo := &Users.User{
		MailAccount: userName,
	}

	if userName == "" {
		ctx.String(http.StatusOK, "用户名不能为空")
		return
	}

	code := ctx.PostForm("code")
	if code != userInfo.GetVerificationCode() {
		ctx.String(http.StatusOK, "验证码错误！")
		return
	}

	newPassword := ctx.PostForm("newPassword")
	if newPassword == "" {
		ctx.String(http.StatusOK, "密码不能为空")
		return
	}

	userInfo.ChangePassword(newPassword)

	ctx.HTML(http.StatusOK, "login.html", nil)
}
