package httpRequest

import (
<<<<<<< HEAD
	"GoProject/spider/spiderMail"
=======
>>>>>>> 589e205... 简易web版
	"crypto/rand"
	"log"
	"math/big"
)

// User-Agent 列表，每次随机选择一个使用
var userAgents []string = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) ",
	"Chrome/45.0.2454.85 Safari/537.36 115Browser/6.0.3",
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
	"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SE 2.X MetaSr 1.0; SE 2.X MetaSr 1.0; .NET CLR 2.0.50727; SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0",
	"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1"}

<<<<<<< HEAD
<<<<<<< HEAD
var MySQLInfo string = 数据库账号密码

var CliConnect string = 服务器 IP 
=======
var MySQLInfo string = "root:密码@tcp(localhost:3306)/数据库名称"
>>>>>>> 995f18e... 更新
=======
var MySQLInfo string = "root:密码@tcp(localhost:3306)/表名"
>>>>>>> 589e205... 简易web版

func GetRandomUserAgent() string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(userAgents))))
	if err != nil {
		log.Println(err)
	}
	return userAgents[n.Int64()]
}
<<<<<<< HEAD

func GetNewMail(userMailAccount string) *spiderMail.Mail {
	return &spiderMail.Mail{
		SenderAccount:  发送账号,
		SenderPassword: 发送密码,
		Receiver:       userMailAccount,
		ServerAddr:     "smtp.office365.com",
		ServerPort:     587,
	}
}
=======
>>>>>>> 589e205... 简易web版
