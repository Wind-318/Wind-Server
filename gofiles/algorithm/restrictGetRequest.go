package algorithm

import (
	"Project/config"
	"errors"

	"github.com/garyburd/redigo/redis"
)

// 每个时间段时长（秒）
var allowTime = 5

// 当前 IP 访问次数加一
func AddOneForThisIP(ip string) error {
	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return err
	}
	defer redisCli.Close()
	// 次数加一
	_, err = redisCli.Do("INCR", ip)
	// 错误处理
	if err != nil {
		return err
	}
	// 设置过期时间
	redisCli.Do("EXPIRE", ip, allowTime)
	// 访问超限
	if IfShouldGetRestricted(ip) {
		// 设置过期时间
		redisCli.Do("EXPIRE", ip, 20)
		return errors.New("访问过于频繁，请稍后再试！")
	}
	// 返回 nil
	return nil
}

// 判断是否访问过于频繁
func IfShouldGetRestricted(ip string) bool {
	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return false
	}
	defer redisCli.Close()
	// 获取时间段内访问次数
	number, err := redis.Int(redisCli.Do("GET", ip))
	if err != nil {
		return false
	}
	// 超过每秒 10 次则超限
	return number >= allowTime*10
}
