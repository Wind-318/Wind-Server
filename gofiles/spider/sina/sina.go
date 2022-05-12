package sina

import (
	"Project/gofiles/mail"
	"Project/gofiles/user"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-gomail/gomail"
)

// 计时抓取存到数据库
func CountTime() {
	// 1.5 小时到 3.5 小时抓取一次
	rand.Seed(time.Now().UnixNano())
	for {
		_, err := GenerateText()
		if err != nil {
			fmt.Println(err)
		}
		result := rand.Intn(7200) + 5400
		time.Sleep(time.Second * time.Duration(result))
	}
}

var f = true

// 6 点和 18 点发送
func SendEveryUser() {
	for {
		if !f {
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
		} else {
			f = false
		}
		// 得到订阅用户名单
		users, err := user.SelectUsersAccount()
		if err != nil {
			return
		}

		temp, err := SelectNews(15)
		if err != nil {
			continue
		}
		// 发送邮件
		for _, user := range users {
			SToMail := mail.GetNewMail(user)
			SToMail.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", temp, gomail.NewMessage())
		}
	}
}
