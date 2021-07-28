package Users

import (
	"Project/Mail"
	"Project/infomation"
	"crypto/sha512"
	"encoding/hex"
	"errors"
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

	passwd := make([]byte, 0)
	// sha512 加密
	code := sha512.Sum512([]byte(user.MailPassword))
	passwd = append(passwd, code[:]...)

	// 将 16 进制转为字符串存储
	tx, err := db.Begin()
	if err != nil {
		return ""
	}

	_, err = tx.Exec("INSERT INTO user VALUES(?,?,?)", 0, user.MailAccount, hex.EncodeToString(passwd))
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
	rand.Seed(time.Now().UnixNano())
	verificationCode := ""
	for i := 0; i < 6; i++ {
		verificationCode += strconv.Itoa(rand.Intn(10))
	}
	return verificationCode
}
