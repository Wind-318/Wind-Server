package user

import "github.com/gomodule/redigo/redis"

func IsExist(cookies string) bool {
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	isExist, err := redis.Bool(redisconn.Do("HEXISTS", cookies, "email"))
	if !isExist || err != nil {
		return false
	}
	return true
}

func GetAccount(cookies string) string {
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	account, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	return account
}
