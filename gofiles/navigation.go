package gofiles

import (
	"Project/gofiles/algorithm"
	"Project/gofiles/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 退出登录
func Exit(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 返回值
	result := map[string]interface{}{}
	// 删除 cookie
	err := user.DeleteCookie(ctx)
	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
	}
	ctx.HTML(http.StatusOK, "bbs.html", nil)
}

// 404 页面
func ToNotFound(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	ctx.HTML(http.StatusNotFound, "notFound.html", nil)
}

// 修改密码
func ToChangePassword(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	ctx.HTML(http.StatusOK, "resetPassword.html", nil)
}

// 登录页面
func ToLogin(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	ctx.HTML(http.StatusOK, "login.html", nil)
}

// 跳转到注册页面
func ToRegister(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	ctx.HTML(http.StatusOK, "register.html", nil)
}

// 根目录
func ToHome(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 重置 cookie 时间
	if cookie, err := ctx.Cookie("cookie"); err == nil {
		user.ResetCookieTime(ctx, cookie)
	}
	ctx.HTML(http.StatusOK, "home.html", nil)
}

// 讨论区页面
func ToBbs(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 重置 cookie 时间
	if cookie, err := ctx.Cookie("cookie"); err == nil {
		user.ResetCookieTime(ctx, cookie)
	}
	// 返回结果
	ctx.HTML(http.StatusOK, "bbs.html", nil)
}

// 新建文章
func ToCreateText(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 重置 cookie 时间
	if cookie, err := ctx.Cookie("cookie"); err == nil {
		user.ResetCookieTime(ctx, cookie)
	}
	ctx.HTML(http.StatusOK, "createText.html", nil)
}

// 收藏
func ToCollections(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 重置 cookie 时间
	if cookie, err := ctx.Cookie("cookie"); err == nil {
		user.ResetCookieTime(ctx, cookie)
	}
	ctx.HTML(http.StatusOK, "collections.html", nil)
}

// 资源页
func ToResources(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 重置 cookie 时间
	if cookie, err := ctx.Cookie("cookie"); err == nil {
		user.ResetCookieTime(ctx, cookie)
	}
	ctx.HTML(http.StatusOK, "resources.html", nil)
}

// 动漫列表
func ToAnime(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 重置 cookie 时间
	if cookie, err := ctx.Cookie("cookie"); err == nil {
		user.ResetCookieTime(ctx, cookie)
	}
	ctx.HTML(http.StatusOK, "anime.html", nil)
}

// 存储
func ToStorage(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	// 重置 cookie 时间
	if cookie, err := ctx.Cookie("cookie"); err == nil {
		user.ResetCookieTime(ctx, cookie)
	}
	ctx.HTML(http.StatusOK, "storage.html", nil)
}
