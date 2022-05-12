package user

import (
	"Project/callClient/callUser"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 检查权限
func CheckPermission(ctx *gin.Context) {
	result := map[string]interface{}{}
	ok, err := callUser.CallUserIsLogin(ctx)
	if err != nil || !ok {
		result["msg"] = "尚未登陆！"
	}
	ctx.JSON(http.StatusOK, result)
}
