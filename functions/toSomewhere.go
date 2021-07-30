package functions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToNotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "notFound.html", nil)
}

func ToChangePassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "changePassword.html", nil)
}

func ToFindPassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "findPassword.html", nil)
}

func ToLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
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
	ctx.HTML(http.StatusOK, "ToSomewhere.html", nil)
}
