package storage

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callStorage"
	"Project/callClient/callUser"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取网站链接
func GetWebs(ctx *gin.Context) {
	// 判断身份
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result := map[string]interface{}{
			"msg": "尚未登陆！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}

	result, err := callStorage.CallStorageGetWebs(ctx)
	if err != nil {
		if result == nil {
			result = map[string]interface{}{}
		}
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}

// 增加收藏网站
func PutWebs(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	// 判断身份
	if ok, err := callUser.CallUserIsSystem(ctx); err != nil || !ok {
		result := map[string]interface{}{
			"msg": "非管理员，无权操作！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	result, err := callStorage.CallStoragePutWebs(ctx)
	if err != nil {
		if result == nil {
			result = map[string]interface{}{}
		}
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}

// 增加图片
func PutPic(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result := map[string]interface{}{
		"msg": "success",
	}
	// 判断管理员账户
	if ok, err := callUser.CallUserIsLogin(ctx); err != nil || !ok {
		result["msg"] = "非管理员账户"
		ctx.JSON(http.StatusOK, result)
		return
	}

	pic, _ := ctx.FormFile("pic")
	go func() {
		ctx.SaveUploadedFile(pic, `picture/collections/`+pic.Filename)
	}()
	ctx.JSON(http.StatusOK, result)
}
