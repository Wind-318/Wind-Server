package algorithm

import (
	"algorithm/callClient/callDatabase"
	"algorithm/config"
	"errors"

	"github.com/garyburd/redigo/redis"
)

// 每个时间段时长（秒）
var allowTime = 5

// 当前 IP 访问次数加一
func AddOneForThisIP(ip string) error {
	// 次数加一
	_, err := callDatabase.CallRedis(config.RedisIP, config.RedisPort, "INCR", ip)
	// 错误处理
	if err != nil {
		return err
	}
	// 设置过期时间
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "EXPIRE", ip, allowTime)
	// 访问超限
	if IfShouldGetRestricted(ip) {
		// 设置过期时间
		callDatabase.CallRedis(config.RedisIP, config.RedisPort, "EXPIRE", ip, 20)
		return errors.New("访问过于频繁，请稍后再试！")
	}
	// 返回 nil
	return nil
}

// 判断是否访问过于频繁
func IfShouldGetRestricted(ip string) bool {
	// 获取时间段内访问次数
	number, err := redis.Int(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "GET", ip))
	if err != nil {
		return false
	}
	// 超过每秒 10 次则超限
	return number >= allowTime*10
}
