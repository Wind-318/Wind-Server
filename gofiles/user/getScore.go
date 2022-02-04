package user

import (
	"Project/gofiles/config"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// 每日签到奖励积分
func signAddScore(ctx *gin.Context) {
	result := map[string]interface{}{
		"msg": "创建成功",
	}
	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	if err != nil {
		result["msg"] = "请先登录"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	// 原积分
	score := 0
	conn.Get(&score, "SELECT score WHERE account = ?", cookie)
	score += 5

	// 签到奖励 5 积分
	conn.Exec("update user SET score = ? WHERE account = ?", score, cookie)

	result["msg"] = "success"
	ctx.JSON(http.StatusOK, result)
}

// 点赞奖励积分
func praiseAddScore(ctx *gin.Context) {
	result := map[string]interface{}{
		"msg": "创建成功",
	}
	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	if err != nil {
		result["msg"] = "请先登录"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	// 原积分
	score := 0
	conn.Get(&score, "SELECT score WHERE account = ?", cookie)
	score += 3

	// 点赞奖励 3 积分
	conn.Exec("update user SET score = ? WHERE account = ?", score, cookie)

	result["msg"] = "success"
	ctx.JSON(http.StatusOK, result)
}

// 回复奖励积分
func replyAddScore(ctx *gin.Context) {
	result := map[string]interface{}{
		"msg": "创建成功",
	}
	// 检查登录状态
	cookies, err := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	cookie, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	if err != nil {
		result["msg"] = "请先登录"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 连接数据库
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	// 原积分
	score := 0
	conn.Get(&score, "SELECT score WHERE account = ?", cookie)
	score += 1

	// 回复加 1 积分
	conn.Exec("update user SET score = ? WHERE account = ?", score, cookie)

	result["msg"] = "success"
	ctx.JSON(http.StatusOK, result)
}
