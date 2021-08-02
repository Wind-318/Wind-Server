package functions

import (
	"Project/WindCount"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func Exit(ctx *gin.Context) {
	ctx.SetCookie("cookie", "", -1, "/", "localhost:80", false, true)
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func ToNotFound(ctx *gin.Context) {
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()

	if ok, _ := redis.Bool(conn.Do("HEXISTS", "banned", ctx.ClientIP())); ok {
		return
	}
	num, _ := redis.Int64(conn.Do("INCR", ctx.ClientIP()))
	conn.Do("EXPIRE", ctx.ClientIP(), 5)
	if num > 20 {
		conn.Do("HSET", "banned", ctx.ClientIP(), "-1")
		return
	}

	ctx.HTML(http.StatusNotFound, "notFound.html", nil)
}

func ToChangePassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "findPassword.html", nil)
}

func ToFindPassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "verificationFindPassword.html", nil)
}

func ToLogin(ctx *gin.Context) {
	_, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.HTML(http.StatusOK, "login.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "function.html", nil)
}

func ToFunction(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "function.html", nil)
}

func ToRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func ToVerificationCode(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "verificationCode.html", nil)
}

func ToHead(ctx *gin.Context) {
	WindCount.GetCount().AddNum()
	ctx.HTML(http.StatusOK, "ToSomewhere.html", nil)
}
