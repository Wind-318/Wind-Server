package user

import (
	"Project/gofiles/config"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// 每日签到奖励
func SignAddScore(ctx *gin.Context) {
	result := map[string]interface{}{
		"msg": "签到成功，获得 5 积分",
	}
	if !IsExist(ctx) {
		result["msg"] = "尚未登陆"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	if err != nil {
		result["msg"] = "尚未登陆"
		ctx.JSON(http.StatusOK, result)
		return
	}
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, err := redis.String(redisconn.Do("HGET", cookies, "email"))
	if err != nil {
		result["msg"] = "尚未登陆"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	// 判断是否已签到
	isSign, _ := redis.Bool(redisconn.Do("HEXISTS", cookie+"userSign", "userSign"))
	if isSign {
		signDay, _ := redis.String(redisconn.Do("HGET", cookie+"userSign", "userSign"))
		signDays, _ := strconv.Atoi(signDay)
		if (time.Now().Day() - signDays) < 1 {
			result["msg"] = "今日已签到"
			ctx.JSON(http.StatusOK, result)
			return
		}
	}

	// 签到日期为本日
	redisconn.Do("HMSET", cookie+"userSign", "userSign", time.Now().Day())

	// 签到奖励 5 积分
	conn.Exec("UPDATE user SET score = score + 5 WHERE account = ?", cookie)
	ctx.JSON(http.StatusOK, result)
}
