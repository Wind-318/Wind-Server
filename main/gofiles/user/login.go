package user

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callUser"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result, err := callUser.CallUserLogin(ctx)
	if err != nil {
		if result == nil {
			result = map[string]interface{}{}
		}
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	} else if result["msg"] != "success" {
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 设置 cookie
	err = SetCookieTime(ctx, ctx.PostForm("userEmail"))
	if err != nil {
		if result == nil {
			result = map[string]interface{}{}
		}
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}
