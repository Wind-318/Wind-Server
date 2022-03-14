package gofiles

import (
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// 退出登录
func Exit(ctx *gin.Context) {
	// 获得 cookie
	cookie, err := ctx.Cookie("cookie")
	// 没有则直接返回
	if err != nil {
		return
	}
	// 连接 redis
	redisconn, err := redis.Dial("tcp", "localhost:6379")
	// 错误处理
	if err != nil {
		fmt.Println(err)
		return
	}
	defer redisconn.Close()
	// 从 redis 中删除 cookie
	redisconn.Do("DEL", cookie)
	// 重置时间
	ctx.SetCookie("cookie", cookie, -1, "/", "localhost/", false, true)
	// 回到登录页面
	ctx.HTML(http.StatusOK, "login.html", nil)
}

// 404
func ToNotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "notFound.html", nil)
}

// 修改密码
func ToChangePassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "resetPassword.html", nil)
}

// 登录页面
func ToLogin(ctx *gin.Context) {
	_, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.HTML(http.StatusOK, "login.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "login.html", nil)
}

// 注册
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

func ToAnime(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "anime.html", nil)
}

// 存储
func ToStorage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "storage.html", nil)
}
