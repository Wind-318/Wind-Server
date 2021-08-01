package Users

import (
	"Project/Mail"
	"Project/infomation"
	"errors"
	"math/rand"
	"time"

	"github.com/go-gomail/gomail"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type User struct {
	UserName     string
	MailAccount  string
	MailPassword string
}

type userData struct {
	Passwd string `db:"password"`
}

type userAcnt struct {
	Accounts string `db:"account"`
}

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
	db := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer db.Close()

	useraccount := userAcnt{}

	db.Get(&useraccount, "SELECT account FROM user WHERE account = ?", user.MailAccount)

	return useraccount.Accounts == user.MailAccount
}

// 注册功能
func (user *User) Register() string {
	db := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer db.Close()

	code := infomation.Encryption(user.MailPassword)

	tx, err := db.Begin()
	if err != nil {
		return ""
	}

	_, err = tx.Exec("INSERT INTO user VALUES(?,?,?,?)", 0, user.MailAccount, code, user.UserName)
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

	code := infomation.Encryption(user.MailPassword)

	err := db.Get(&userpasswd, "SELECT password FROM user WHERE account = ?", user.MailAccount)
	if err != nil {
		return "No userAccount"
	} else if userpasswd.Passwd != code {
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

	code := infomation.Encryption(newPassword)

	_, err = tx.Exec("UPDATE user SET password = ? WHERE account = ?", code, user.MailAccount)
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
func (user *User) Verification() error {
	// 接收者邮箱
	mail := Mail.GetNewMail(user.MailAccount)

	verificationCode := user.sendCode()

	mail.Send("验证码", "<h1>您的验证码为："+verificationCode+"<h1>", gomail.NewMessage())

	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	if user.FindVerificationCode() {
		return errors.New("验证码已发送，请 1 分钟后再试")
	}
	// 验证码持续时间 1 分钟，过期自动失效
	_, err := connect.Do("SET", user.MailAccount, verificationCode, "ex", "60")
	if err != nil {
		return err
	}
	return nil
}

// 获取验证码
func (user *User) GetVerificationCode() string {
	if !user.FindVerificationCode() {
		return "not exist"
	}
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	reply, _ := redis.String(connect.Do("GET", user.MailAccount))
	return reply
}

// 查找验证码是否过期
func (user *User) FindVerificationCode() bool {
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	isExist, _ := redis.Bool(connect.Do("EXISTS", user.MailAccount))
	return isExist
}

// 生成 6 位数验证码
func (user *User) sendCode() string {
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	verificationCode := ""
	for i := 0; i < 6; i++ {
		verificationCode += string(letters[rand.Intn(len(letters))])
	}
	return verificationCode
}
