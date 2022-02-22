package user

import (
	"Project/gofiles/config"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// 每日签到奖励
func SignAddScore(ctx *gin.Context) {
	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	// 判断是否已签到
	isSign, _ := redis.Bool(redisconn.Do("HEXISTS", cookie, "userSign"))
	if isSign {
		signDay, _ := redis.Int(redisconn.Do("HGET", cookie, "userSign"))
		if (time.Now().Day() - signDay) < 1 {
			return
		}
	}

	// 签到日期为本日
	redisconn.Do("HMSET", cookie, "userSign", time.Now().Day())
	// 原积分
	score := 0
	conn.Get(&score, "SELECT score WHERE account = ?", cookie)
	score += 5

	// 签到奖励 5 积分
	conn.Exec("udate user SET score = ? WHERE account = ?", score, cookie)
}
