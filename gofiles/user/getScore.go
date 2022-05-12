package user

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"net/http"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 每日签到奖励
func SignAddScore(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())

	result := map[string]interface{}{}
	if !IsLogin(ctx) {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 随机积分数
	score := algorithm.GetNormalDistributionNumber(7, 3, 15)

	// 检查登录状态
	userAccount, err := GetUserEmail(ctx)
	if err != nil {
		result["msg"] = "尚未登陆"
		ctx.JSON(http.StatusOK, result)
		return
	}

	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer redisCli.Close()
	// 判断是否已签到
	isSign, _ := redis.Bool(redisCli.Do("HEXISTS", userAccount+"userSign", "userSign"))
	if isSign {
		signDay, _ := redis.String(redisCli.Do("HGET", userAccount+"userSign", "userSign"))
		signDays, _ := strconv.Atoi(signDay)
		if (time.Now().Day() - signDays) < 1 {
			result["msg"] = "今日已签到！"
			ctx.JSON(http.StatusOK, result)
			return
		}
	}

	// 设置签到日期为本日
	redisCli.Do("HMSET", userAccount+"userSign", "userSign", time.Now().Day())
	redisCli.Do("expire", userAccount+"userSign", 86400)

	// 签到奖励积分

	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient.Close()

	mysqlClient.Exec("UPDATE user SET score = score + ? WHERE account = ?", score, userAccount)

	result["msg"] = "签到成功，获得 " + strconv.Itoa(score) + " 积分！"

	ctx.JSON(http.StatusOK, result)
}
