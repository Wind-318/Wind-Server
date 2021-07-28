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
<<<<<<< HEAD:spiderUsers/user.go
<<<<<<< HEAD:spiderUsers/user.go
	"fmt"
	"log"
=======
>>>>>>> 9e17b98... /:Users/user.go
=======
	"errors"
>>>>>>> 0c71c88... 更新:Users/user.go
	"math/rand"
<<<<<<< HEAD:spiderUsers/user.go
=======
	"project/Mail"
	"project/infomation"
>>>>>>> d19263e... 更新:Users/user.go
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

<<<<<<< HEAD:spiderUsers/user.go
func (user *User) CheckUserExist(registerAccount string) bool {
=======
func SelectUsersAccount() []string {
	db := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer db.Close()

	useraccount := make([]userAcnt, 0)

	db.Select(&useraccount, "SELECT account FROM user")
	ret := make([]string, 0)
	for _, account := range useraccount {
		ret = append(ret, account.Accounts)
	}
	return ret
}

func (user *User) CheckUserExist() bool {
<<<<<<< HEAD:spiderUsers/user.go
>>>>>>> 0c71c88... 更新:Users/user.go
	db := sqlx.MustConnect("mysql", httpRequest.MySQLInfo)
=======
	db := sqlx.MustConnect("mysql", infomation.MySQLInfo)
>>>>>>> d19263e... 更新:Users/user.go
	defer db.Close()

	useraccount := userAcnt{}

	db.Get(&useraccount, "SELECT account FROM user WHERE account = ?", registerAccount)

	return useraccount.Accounts == registerAccount
}

// 注册功能
<<<<<<< HEAD:spiderUsers/user.go
func (user *User) Register(registerAccount, registerPassword string) string {
	db := sqlx.MustConnect("mysql", httpRequest.MySQLInfo)
=======
func (user *User) Register() string {
	db := sqlx.MustConnect("mysql", infomation.MySQLInfo)
>>>>>>> d19263e... 更新:Users/user.go
	defer db.Close()

	passwd := make([]byte, 0)
	// sha512 加密
	code := sha512.Sum512([]byte(registerPassword))
	passwd = append(passwd, code[:]...)

	// 将 16 进制转为字符串存储
	tx, err := db.Begin()
	if err != nil {
		return ""
	}

	_, err = tx.Exec("INSERT INTO user VALUES(?,?,?)", 0, registerAccount, hex.EncodeToString(passwd))
	if err != nil {
		return ""
	}
	err = tx.Commit()
	if err != nil {
		return ""
	}

	return "success"
}

// 登录功能
func (user *User) Login() string {
	userpasswd := userData{}
	db := sqlx.MustConnect("mysql", infomation.MySQLInfo)
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
	db := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		return "fail"
	}

	code := sha512.Sum512([]byte(newPassword))
	passwd := make([]byte, 0)
	passwd = append(passwd, code[:]...)

	_, err = tx.Exec("UPDATE user SET password = ? WHERE account = ?", hex.EncodeToString(passwd), user.MailAccount)
	if err != nil {
		return "fail"
	}
	err = tx.Commit()
	if err != nil {
		return "fail"
	}
	return "success"
}

// 发送验证码
<<<<<<< HEAD:spiderUsers/user.go
func (user *User) Verification(ReceiverAccount string) {
=======
func (user *User) Verification() error {
>>>>>>> 30f28d9... /:Users/user.go
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

<<<<<<< HEAD:spiderUsers/user.go
	// 验证码持续时间 5 分钟，过期自动失效
	_, err := connect.Do("SET", ReceiverAccount, verificationCode, "ex", "300")
=======
	if user.FindVerificationCode() {
		return errors.New("验证码已发送，请 1 分钟后再试")
	}
	// 验证码持续时间 1 分钟，过期自动失效
	_, err := connect.Do("SET", user.MailAccount, verificationCode, "ex", "60")
>>>>>>> 0c71c88... 更新:Users/user.go
	if err != nil {
<<<<<<< HEAD:spiderUsers/user.go
<<<<<<< HEAD:spiderUsers/user.go
		fmt.Println(err)
=======
		log.Println(err)
=======
>>>>>>> 9e17b98... /:Users/user.go
		return err
>>>>>>> 30f28d9... /:Users/user.go
	}
	return nil
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
