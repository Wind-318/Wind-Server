package functions

import (
	"1/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToChangePassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "changePassword.html", nil)
}

func ChangePassWord(ctx *gin.Context) {
	userName := ctx.PostForm("userName")
	newPassword := ctx.PostForm("newPassword")
	newPasswordConfirm := ctx.PostForm("newPasswordConfirm")

	if newPasswordConfirm != newPassword {
		ctx.String(http.StatusOK, "两次密码输入不一致！")
		return
	} else if userName == "" {
		ctx.String(http.StatusOK, "请输入用户名")
		return
	}

	userInfo := &Users.User{
		MailAccount: userName,
	}
	userInfo.ChangePassword(newPassword)
	ctx.HTML(http.StatusOK, "login.html", nil)
}
