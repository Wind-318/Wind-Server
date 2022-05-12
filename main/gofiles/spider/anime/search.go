package anime

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callSpider"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 搜索功能
func Search(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result, err := callSpider.CallAnimeSearchAnime(ctx)
	if err != nil {
		if result == nil {
			result = map[string]interface{}{}
		}
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}

// 选出新番
func SearchNewAnime(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result, err := callSpider.CallAnimeSearchNewAnime(ctx)
	if err != nil {
		if result == nil {
			result = map[string]interface{}{}
		}
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}

// 选出指定年份番剧
func SearchByYear(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result, err := callSpider.CallAnimeSearchByYear(ctx)
	if err != nil {
		if result == nil {
			result = map[string]interface{}{}
		}
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}
