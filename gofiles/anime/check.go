package anime

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func CheckPermission(ctx *gin.Context) {
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
	}
	ctx.JSON(http.StatusOK, result)
}

func GetYear(ctx *gin.Context) {
	result := map[string]interface{}{}
	result["year"] = time.Now().Year()
	ctx.JSON(http.StatusOK, result)
}
