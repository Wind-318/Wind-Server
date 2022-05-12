package mail

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"math/rand"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

// 发送验证码
func SendCode(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())

	result := map[string]interface{}{
		"msg": "已发送，若未接收到请检查垃圾箱。",
	}

	userEmail := ctx.PostForm("userEmail")

	// 检查是否存在
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer redisCli.Close()
	isExist, err := redis.Bool(redisCli.Do("EXISTS", userEmail))
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	} else if isExist {
		result["msg"] = "邮件已发送，请 5 分钟后再试"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 生产消息，投入 stream 中
	redisCli.Do("XADD", "SendCode", "*", "email", userEmail)
	ctx.JSON(http.StatusOK, result)
}

// 发送验证码
func ConsumeCode() {
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return
	}
	defer redisCli.Close()
	// 新建消费者队列
	redisCli.Do("XGROUP", "CREATE", "SendCode", "ReceiveCode", 0, "MKSTREAM")
	for {
		// 消费消息
		reply, err := redisCli.Do("XREADGROUP", "GROUP", "ReceiveCode", "consumer", "count", 1, "block", 0, "streams", "SendCode", ">")
		if err != nil {
			continue
		}
		// 解析内容
		emails := []string{}
		parseRedisStream(reply, &emails)
		// 获取本次接收方
		userEmail := emails[len(emails)-1]
		// 生成验证码
		verificationCode := generateCode()
		// 验证码持续时间 5 分钟，过期自动失效
		ok, err := redis.String(redisCli.Do("SET", userEmail, verificationCode, "ex", "300", "nx"))
		// 错误处理
		if err != nil || ok != "OK" {
			continue
		}

		// 接收者邮箱
		mails := GetNewMail(userEmail)

		// 发送
		err = mails.Send("验证码", "<h1>您的验证码为:"+verificationCode+"<h1>", gomail.NewMessage())
		if err != nil {
			continue
		}
	}
}

// 生成 6 位数验证码
func generateCode() string {
	// 从字符串中随机选择
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	verificationCode := ""
	// 生成长度 6 的字符串
	for i := 0; i < 6; i++ {
		verificationCode += string(letters[rand.Intn(len(letters))])
	}

	// 返回
	return verificationCode
}

// 解析
func parseRedisStream(reply interface{}, ans *[]string) {
	// 断言为 []interface{}
	temp := reply.([]interface{})
	for index := range temp {
		// 不能断言，可以转为 string
		if _, ok := temp[index].([]interface{}); !ok {
			str, err := redis.String(temp[index], nil)
			if err != nil {
				return
			}
			*ans = append(*ans, str)
		} else {
			parseRedisStream(temp[index], ans)
		}
	}
}
