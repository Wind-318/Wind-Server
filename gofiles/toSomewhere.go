package gofiles

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func Exit(ctx *gin.Context) {
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	redisconn.Do("DEL", cookie)
	ctx.SetCookie("cookie", cookie, -1, "/", "localhost:80", false, true)
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func ToNotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "notFound.html", nil)
}

func ToChangePassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "resetPassword.html", nil)
}

func ToLogin(ctx *gin.Context) {
	_, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.HTML(http.StatusOK, "login.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func ToRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func ToVerificationCode(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "verificationCode.html", nil)
}

func ToHead(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ToSomewhere.html", nil)
}

func ToError(ctx *gin.Context) {
	ctx.HTML(http.StatusInternalServerError, "serverError.html", nil)
}

func ToBlog(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "blog.html", nil)
}

func ToCreateText(ctx *gin.Context) {
	_, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}
	ctx.HTML(http.StatusOK, "createText.html", nil)
}

func ToCollections(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "collections.html", nil)
}

func ToResources(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "resources.html", nil)
}
