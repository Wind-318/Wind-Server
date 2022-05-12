package sendEmail

import (
	"mail/callClient/callDatabase"
	"mail/callClient/callSpider"
	"mail/config"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/go-gomail/gomail"
)

// 发送邮件
func SendStock(email string) error {
	// 获得邮件对象
	users := GetNewMail(email)
	rand.Seed(time.Now().UnixNano())
	temp, err := callSpider.CallSelectNews(15)
	if err != nil {
		return err
	}
	// 发送
	err = users.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", temp, gomail.NewMessage())
	if err != nil {
		return err
	}

	return nil
}

// 发送验证码
func SendCode() {
	// 新建消费者队列
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "XGROUP", "CREATE", "SendCode", "ReceiveCode", 0, "MKSTREAM")
	for {
		// 消费消息
		reply, err := callDatabase.CallRedis(config.RedisIP, config.RedisPort, "XREADGROUP", "GROUP", "ReceiveCode", "consumer"+config.MailListenPort, "count", 1, "block", 0, "streams", "SendCode", ">")
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
		ok, err := redis.String(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "SET", userEmail, verificationCode, "ex", "300", "nx"))
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
