package user

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callDatabase"
	"Project/config"

	"github.com/gin-gonic/gin"
)

// 重置 cookie 时间
func ResetCookieTime(ctx *gin.Context, cookie string) error {
	// 一天 86400 秒
	oneDay := 86400
	// domain 是域名，path 是域名，合起来限制可以被哪些 url 访问，重设 cookie 过期时间
	ctx.SetCookie("cookie", cookie, oneDay*30, "/", "localhost/", false, true)

	// 重置
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "EXPIRE", cookie, oneDay*30)

	return nil
}

// 设置 cookie
func SetCookieTime(ctx *gin.Context, userEmail string) error {
	// domain 是域名，path 是域名，合起来限制可以被哪些 url 访问，重设 cookie 过期时间
	cookie, err := callAlgorithm.CallAlgorithmGenerateRandStr(50)
	if err != nil {
		return err
	}
	// 一天 86400 秒
	oneDay := 86400
	ctx.SetCookie("cookie", cookie, oneDay*30, "/", "localhost/", false, true)

	// 重设
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "HMSET", cookie, "email", userEmail)
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "EXPIRE", cookie, oneDay*30)

	return nil
}

// 删除 cookie
func DeleteCookie(ctx *gin.Context) error {
	// 获取 cookie
	cookie, err := ctx.Cookie("cookie")
	// 没有则直接返回
	if err != nil {
		return err
	}

	// 从 redis 中删除 cookie
	_, err = callDatabase.CallRedis(config.RedisIP, config.RedisPort, "DEL", cookie)
	// 错误处理
	if err != nil {
		return err
	}
	// 重置时间
	ctx.SetCookie("cookie", cookie, -1, "/", "localhost/", false, true)

	return nil
}
