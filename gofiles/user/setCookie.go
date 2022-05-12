package user

import (
	"Project/config"
	"Project/gofiles/algorithm"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// 重置 cookie 时间
func ResetCookieTime(ctx *gin.Context, cookie string) error {
	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return err
	}
	defer redisCli.Close()
	// 一天 86400 秒
	oneDay := 86400
	// domain 是域名，path 是域名，合起来限制可以被哪些 url 访问，重设 cookie 过期时间
	ctx.SetCookie("cookie", cookie, oneDay*30, "/", "localhost/", false, true)

	// 重置
	redisCli.Do("EXPIRE", cookie, oneDay*30)

	return nil
}

// 设置 cookie
func SetCookieTime(ctx *gin.Context, userEmail string) error {
	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return err
	}
	defer redisCli.Close()
	// domain 是域名，path 是域名，合起来限制可以被哪些 url 访问，重设 cookie 过期时间
	cookie, err := algorithm.GenerateRandStr(50)
	if err != nil {
		return err
	}
	// 一天 86400 秒
	oneDay := 86400
	ctx.SetCookie("cookie", cookie, oneDay*30, "/", "localhost/", false, true)

	// 重设
	redisCli.Do("HMSET", cookie, "email", userEmail)
	redisCli.Do("EXPIRE", cookie, oneDay*30)

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

	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return err
	}
	defer redisCli.Close()
	// 从 redis 中删除 cookie
	_, err = redisCli.Do("DEL", cookie)
	// 错误处理
	if err != nil {
		return err
	}
	// 重置时间
	ctx.SetCookie("cookie", cookie, -1, "/", "localhost/", false, true)

	return nil
}
