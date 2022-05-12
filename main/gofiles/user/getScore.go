package user

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callUser"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 每日签到奖励
func SignAddScore(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result := map[string]interface{}{}

	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}

	result, err := callUser.CallUserSignGetScore(ctx)
	if err != nil {
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}
