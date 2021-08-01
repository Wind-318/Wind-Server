package functions

import (
	"Project/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerificationFind(ctx *gin.Context) {
	userEmail := ctx.PostForm("userEmail")
	userInfo := &Users.User{
		MailAccount: userEmail,
	}

	if !userInfo.CheckUserExist() {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}

	err := userInfo.Verification()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "serverError.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "verificationFindPassword.html", nil)
}

func ConfirmFind(ctx *gin.Context) {
	code := ctx.PostForm("code")
	userEmail := ctx.PostForm("userEmail")
	userInfo := &Users.User{
		MailAccount: userEmail,
	}
	if code != userInfo.GetVerificationCode() {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "findPassword.html", nil)
}

func ModifyPassword(ctx *gin.Context) {
	newPassword := ctx.PostForm("newPassword")
	userEmail := ctx.PostForm("userEmail")
	userInfo := &Users.User{
		MailAccount: userEmail,
	}
	userInfo.ChangePassword(newPassword)
	ctx.HTML(http.StatusOK, "login.html", nil)
}
