package user

import "github.com/gomodule/redigo/redis"

// 获取账号
func GetAccount(cookies string) string {
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	account, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	return account
}
