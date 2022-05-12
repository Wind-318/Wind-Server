package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 检查权限
func CheckPermission(ctx *gin.Context) {
	result := map[string]interface{}{}
	if !IsLogin(ctx) {
		result["msg"] = "尚未登陆！"
	}

	ctx.JSON(http.StatusOK, result)
}
