package sina

import (
	"Project/callClient/callDatabase"
	"Project/callClient/callMail"
	"Project/callClient/callSpider"
	"Project/config"
	"fmt"
	"math/rand"
	"time"
)

// 计时抓取存到数据库
func CountTime() {
	// 1.5 小时到 3.5 小时抓取一次
	rand.Seed(time.Now().UnixNano())
	for {
		_, err := callSpider.CallGenerateText()
		if err != nil {
			fmt.Println(err)
		}
		result := rand.Intn(7200) + 5400
		time.Sleep(time.Second * time.Duration(result))
	}
}

// 6 点和 18 点发送
func SendEveryUser() {
	for {
		// 得到现在时间
		nowHour, nowMinute := time.Now().Hour(), time.Now().Minute()
		// 等待时间
		waitSeconds := 0

		// 计算现在到 6 点或者到 18 点还有多少秒
		if nowHour < 18 && nowHour >= 6 {
			waitSeconds += (17-nowHour)*3600 + (60-nowMinute)*60
		} else if nowHour >= 18 {
			waitSeconds += (23-nowHour)*3600 + (60-nowMinute)*60 + 6*3600
		} else {
			waitSeconds += (5-nowHour)*3600 + (60-nowMinute)*60
		}

		// 等待到指定时间
		time.Sleep(time.Second * time.Duration(waitSeconds))
		// 得到订阅用户名单
		users, err := callDatabase.CallMySQLSelectUsersAccount(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user")
		if err != nil {
			return
		}
		// 发送邮件
		for _, user := range users {
			callMail.CallMailSendStockService(user)
		}
	}
}
