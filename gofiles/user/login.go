package user

import (
	"Project/gofiles/algorithm"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())

	// 返回结果
	result := map[string]interface{}{
		"msg": "success",
	}
	// 检查是否已登录
	if ok := IsLogin(ctx); ok {
		result["msg"] = "已登录"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 构造对象
	userInfo := &User{
		UserName: "",
		Account:  ctx.PostForm("userEmail"),
		Password: ctx.PostForm("userPassword"),
	}
	// 检查字段合法性
	if userInfo.Account == "" || userInfo.Password == "" {
		result["msg"] = "账号或密码不能为空"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 登录
	status := userInfo.Login()
	// 不成功返回结果
	if status != nil {
		result["msg"] = status.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 设置 cookie
	err := SetCookieTime(ctx, ctx.PostForm("userEmail"))
	if err != nil {
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}
