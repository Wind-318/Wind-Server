package user

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callUser"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 验证是否处于登录状态
func IsLogin(ctx *gin.Context) {
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
	ok, err := callUser.CallUserIsLogin(ctx)
	if err != nil {
		result["msg"] = err.Error()
	} else if !ok {
		result["msg"] = "未登录"
	} else {
		result["msg"] = "success"
	}
	ctx.JSON(http.StatusOK, result)
}

// 获取邮箱
func GetUserEmail(ctx *gin.Context) {
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
	ok, err := callUser.CallUserGetUserEmail(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	result["email"] = ok
	ctx.JSON(http.StatusOK, result)
}

// 获取所有用户名称
func GetUsersName(ctx *gin.Context) {
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
	ok, err := callUser.CallUserGetUsersName(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	result["names"] = ok
	ctx.JSON(http.StatusOK, result)
}

// 获取用户名称
func GetName(ctx *gin.Context) {
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
	ok, err := callUser.CallUserGetName(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	result["name"] = ok
	ctx.JSON(http.StatusOK, result)
}

// 查看是否是管理员
func IsSystem(ctx *gin.Context) {
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
	ok, err := callUser.CallUserIsSystem(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
	} else if ok {
		result["msg"] = "success"
	}
	ctx.JSON(http.StatusOK, result)
}

// 查看是否是管理员或文章作者
func IsSystemOrAuthor(ctx *gin.Context) {
	result := map[string]interface{}{}
	ok, err := callUser.CallUserIsSystemOrAuthor(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
	} else if ok {
		result["msg"] = "success"
	}
	ctx.JSON(http.StatusOK, result)
}
