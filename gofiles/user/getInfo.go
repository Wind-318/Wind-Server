package user

import (
	"Project/gofiles/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// 验证是否存在该用户
func IsExist(ctx *gin.Context) bool {
	cookies, err := ctx.Cookie("cookie")
	if err != nil {
		return false
	}
	redisconn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer redisconn.Close()
	isExist, err := redis.Bool(redisconn.Do("HEXISTS", cookies, "email"))
	if !isExist || err != nil {
		return false
	}
	return true
}

// 获取邮箱
func GetUserEmail(ctx *gin.Context) string {
	cookies, _ := ctx.Cookie("cookie")
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	email, _ := redis.String(redisconn.Do("HGET", cookies, "email"))
	return email
}

// 获取所有用户名称
func GetUsersName(ctx *gin.Context) {
	result := map[string]interface{}{}
	// 验证存在
	if !IsExist(ctx) {
		ctx.JSON(http.StatusOK, result)
		return
	}
	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()
	// 检索名称
	names := []string{}
	conn.Select(&names, "SELECT username FROM user")
	result["names"] = names
	ctx.JSON(http.StatusOK, result)
}
