package functions

import (
	"Project/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	userName := ctx.PostForm("userName")
	userEmail := ctx.PostForm("userEmail")
	passWord := ctx.PostForm("passWord")

	userInfo := &Users.User{
		UserName:     userName,
		MailAccount:  userEmail,
		MailPassword: passWord,
	}

	if userInfo.CheckUserExist() {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}

	userInfo.Verification()

	// domain 是域名，path 是域名，合起来限制可以被哪些 url 访问
	ctx.SetCookie("userName", userName, 60, "/", "localhost:80", false, true)
	ctx.SetCookie("userEmail", userEmail, 60, "/", "localhost:80", false, true)
	ctx.SetCookie("passWord", passWord, 60, "/", "localhost:80", false, true)
	ctx.HTML(http.StatusOK, "verificationCode.html", nil)
}

func SendCode(ctx *gin.Context) {
	userName, err := ctx.Cookie("userName")
	if err != nil {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}
	userEmail, err := ctx.Cookie("userEmail")
	if err != nil {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}
	passWord, err := ctx.Cookie("passWord")
	if err != nil {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}

	code := ctx.PostForm("code")
	userInfo := &Users.User{
		UserName:     userName,
		MailAccount:  userEmail,
		MailPassword: passWord,
	}

	if code != userInfo.GetVerificationCode() {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}

	userInfo.Register()

	ctx.HTML(http.StatusOK, "login.html", nil)
}
