package spiderUsers

import (
<<<<<<< HEAD:spiderUsers/user.go
	"GoProject/spider/httpRequest"
=======
	"1/Mail"
	"1/httpRequest"
>>>>>>> 589e205... 简易web版:Users/user.go
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-gomail/gomail"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type User struct {
	MailAccount  string
	MailPassword string
}

type userData struct {
	Passwd string `db:"password"`
}

type userAcnt struct {
	Accounts string `db:"account"`
}

func (user *User) CheckUserExist(registerAccount string) bool {
	db := sqlx.MustConnect("mysql", httpRequest.MySQLInfo)
	defer db.Close()

	useraccount := userAcnt{}

	db.Get(&useraccount, "SELECT account FROM user WHERE account = ?", registerAccount)

	return useraccount.Accounts == registerAccount
}

// 注册功能
func (user *User) Register(registerAccount, registerPassword string) string {
	db := sqlx.MustConnect("mysql", httpRequest.MySQLInfo)
	defer db.Close()

	passwd := make([]byte, 0)
	// sha512 加密
	code := sha512.Sum512([]byte(registerPassword))
	passwd = append(passwd, code[:]...)

	// 将 16 进制转为字符串存储
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return ""
	}

	_, err = tx.Exec("INSERT INTO user VALUES(?,?,?)", 0, registerAccount, hex.EncodeToString(passwd))
	if err != nil {
		log.Println(err)
		return ""
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return ""
	}

	return "success"
}

// 登录功能
func (user *User) Login() string {
	userpasswd := userData{}
	db := sqlx.MustConnect("mysql", httpRequest.MySQLInfo)
	defer db.Close()

	code := sha512.Sum512([]byte(user.MailPassword))
	passwd := make([]byte, 0)
	passwd = append(passwd, code[:]...)

	err := db.Get(&userpasswd, "SELECT password FROM user WHERE account = ?", user.MailAccount)
	if err != nil {
		return "No userAccount"
	} else if userpasswd.Passwd != hex.EncodeToString(passwd) {
		return "password wrong"
	}
	return "success"
}

// 修改密码
func (user *User) ChangePassword(newPassword string) string {
	db := sqlx.MustConnect("mysql", httpRequest.MySQLInfo)
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
		return "fail"
	}

	code := sha512.Sum512([]byte(newPassword))
	passwd := make([]byte, 0)
	passwd = append(passwd, code[:]...)

	_, err = tx.Exec("UPDATE user SET password = ? WHERE account = ?", hex.EncodeToString(passwd), user.MailAccount)
	if err != nil {
		log.Println(err)
		return "fail"
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "fail"
	}
	return "success"
}

// 发送验证码
func (user *User) Verification(ReceiverAccount string) {
	// 接收者邮箱
<<<<<<< HEAD:spiderUsers/user.go
	mail := httpRequest.GetNewMail(ReceiverAccount)
=======
	mail := Mail.GetNewMail(user.MailAccount)
>>>>>>> 589e205... 简易web版:Users/user.go

	verificationCode := user.sendCode()

	mail.Send("验证码", "<h1>您的验证码为："+verificationCode+"<h1>", gomail.NewMessage())

	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	// 验证码持续时间 5 分钟，过期自动失效
	_, err := connect.Do("SET", ReceiverAccount, verificationCode, "ex", "300")
	if err != nil {
		fmt.Println(err)
	}
}

// 获取验证码
func (user *User) GetVerificationCode(ReceiverAccount string) string {
	if !user.FindVerificationCode(ReceiverAccount) {
		return "not exist"
	}
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	reply, _ := redis.String(connect.Do("GET", ReceiverAccount))
	return reply
}

// 查找验证码是否过期
func (user *User) FindVerificationCode(ReceiverAccount string) bool {
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	isExist, _ := redis.Bool(connect.Do("EXISTS", ReceiverAccount))
	return isExist
}

// 生成 6 位数验证码
func (user *User) sendCode() string {
	rand.Seed(time.Now().UnixNano())
	verificationCode := ""
	for i := 0; i < 6; i++ {
		verificationCode += strconv.Itoa(rand.Intn(10))
	}
	return verificationCode
}
