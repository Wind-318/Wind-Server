package functions

import (
	"Project/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 更改密码
func ChangePassWord(ctx *gin.Context) {
	// 新密码
	newPassword := ctx.PostForm("userPassword")
	// 账号
	useremail := ctx.PostForm("userEmail")
	// 验证码
	code := ctx.PostForm("code")

	userInfo := &Users.User{
		MailAccount: useremail,
	}

	// 验证码错误，直接返回
	if code != userInfo.GetVerificationCode() {
		return
	}

	// 更改密码
	userInfo.ChangePassword(newPassword)
	ctx.HTML(http.StatusOK, "login.html", nil)
}
