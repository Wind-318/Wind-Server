package user

import (
	"strconv"
	"time"
	"user/callClient/callAlgorithm"
	"user/callClient/callDatabase"
	"user/config"

	"github.com/garyburd/redigo/redis"
)

// 签到积分
func SignGetScore(cookie string) map[string]interface{} {
	// 返回结果
	result := map[string]interface{}{
		"msg": "success",
	}
	// 随机积分数
	score, err := callAlgorithm.CallAlgorithmGetNormalDistributionNumber(7, 3, 15)
	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
		return result
	}
	// 检查登录状态
	userAccount, err := GetUserEmail(cookie)
	if err != nil {
		result["msg"] = "尚未登陆"
		return result
	}

	// 判断是否已签到
	isSign, _ := redis.Bool(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "HEXISTS", userAccount+"userSign", "userSign"))
	if isSign {
		signDay, _ := redis.String(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "HGET", userAccount+"userSign", "userSign"))
		signDays, _ := strconv.Atoi(signDay)
		if (time.Now().Day() - signDays) < 1 {
			result["msg"] = "今日已签到！"
			return result
		}
	}

	// 设置签到日期为本日
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "HMSET", userAccount+"userSign", "userSign", time.Now().Day())
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "expire", userAccount+"userSign", 86400)

	// 签到奖励积分
	callDatabase.CallMySQLUpdateUserScore(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", userAccount, score)

	result["msg"] = "签到成功，获得 " + strconv.Itoa(score) + " 积分！"
	return result
}
