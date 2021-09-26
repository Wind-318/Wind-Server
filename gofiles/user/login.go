package user

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// 登录
func Login(ctx *gin.Context) {
	userEmail := ctx.PostForm("userEmail")
	passWord := ctx.PostForm("userPassword")

	userInfo := &User{
		Account:  userEmail,
		Password: passWord,
	}

	result := map[string]interface{}{
		"msg": "success",
	}

	if userEmail == "" || passWord == "" {
		result["msg"] = "账号或密码不能为空"
		ctx.JSON(http.StatusOK, result)
		return
	}

	status := userInfo.Login()
	if status != "success" {
		result["msg"] = status
		ctx.JSON(http.StatusOK, result)
		return
	}

	// domain 是域名，path 是域名，合起来限制可以被哪些 url 访问
	cookie := generateRandStr()
	ctx.SetCookie("cookie", cookie, 86400, "/", "localhost/", false, true)
	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()
	redisconn.Do("HMSET", cookie, "email", userEmail)
	redisconn.Do("EXPIRE", cookie, 86400)

	ctx.JSON(http.StatusOK, result)
}

// 获取一串随机字符串
func generateRandStr() string {
	code := ""
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	redisconn, _ := redis.Dial("tcp", "localhost:6379")
	defer redisconn.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		code += string(letters[rand.Intn(len(letters))])
	}

	if ok, _ := redis.Bool(redisconn.Do("EXISTS", code)); ok {
		code = generateRandStr()
	}
	return code
}
