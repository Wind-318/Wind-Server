package storage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// 存储文件
func StroageFiles(ctx *gin.Context) {
	result := map[string]interface{}{
		"msg": "success",
	}
	cookies, err := ctx.Cookie("cookie")
	if err != nil {
		result["msg"] = "fail"
	}
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	isExist, err := redis.Bool(redisconn.Do("HEXISTS", cookies, "email"))
	if !isExist || err != nil {
		result["msg"] = "fail"
	} else {
		email, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
		ctx.SetCookie("cookie", cookies, 86400, "/", "localhost/", false, true)
		redisconn.Do("EXPIRE", cookies, 86400)
	}
}
