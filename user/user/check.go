package user

import (
	"user/callClient/callDatabase"
	"user/config"

	"github.com/garyburd/redigo/redis"
)

// 验证是否处于登录状态
func IsLogin(cookie string) (bool, error) {
	// 检查存在
	isExist, err := redis.Bool(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "HEXISTS", cookie, "email"))
	// 不存在或者出错，返回 false
	if err != nil || !isExist {
		return false, err
	}

	// 检验完成，返回 true
	return true, nil
}
