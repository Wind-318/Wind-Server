package user

import (
	"Project/gofiles/config"
	"Project/gofiles/ownmail"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-gomail/gomail"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// 用户信息
type User struct {
	UserName string
	Account  string
	Password string
}

type userData struct {
	Passwd string `db:"password"`
}

type userAcnt struct {
	Accounts string `db:"account"`
}

// 得到订阅用户名单
func SelectUsersAccount() []string {
	db := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer db.Close()

	useraccount := make([]userAcnt, 0)

	db.Select(&useraccount, "SELECT account FROM subscribe WHERE stock = ?", 1)
	ret := make([]string, 0)
	for _, account := range useraccount {
		ret = append(ret, account.Accounts)
	}
	return ret
}

// 检查用户是否存在
func (user *User) CheckUserExist() bool {
	db := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer db.Close()

	useraccount := userAcnt{}

	db.Get(&useraccount, "SELECT account FROM user WHERE account = ?", user.Account)

	return useraccount.Accounts == user.Account
}

// 注册功能
func (user *User) Register() error {
	db := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer db.Close()

	code := config.Encryption(user.Password)

	db.Exec("INSERT INTO user VALUES(?, ?, ?, ?, ?, ?, ?)", 0, user.Account, code, user.UserName, 0, 0, config.Addr+"picture/defaultPic.jpg")

	var id int
	db.Get(&id, "SELECT id FROM user WHERE account = ?", user.Account)

	os.Mkdir(`blog/`+strconv.Itoa(id), 0644)

	content, _ := ioutil.ReadFile(`./blogTemplate.html`)
	ioutil.WriteFile(`blog/`+strconv.Itoa(id)+`/user.html`, content, 0644)
	db.Exec("INSERT INTO blog VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 0, user.UserName, user.Account, "这是你的第一篇文章", "", "这是你的第一篇文章", "未分类", 0, 0, 0, time.Now().String()[:19], time.Now().String()[:19])

	return nil
}

// 登录功能
func (user *User) Login() string {
	userpasswd := userData{}
	db := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer db.Close()

	code := config.Encryption(user.Password)

	err := db.Get(&userpasswd, "SELECT password FROM user WHERE account = ?", user.Account)
	if err != nil {
		return "账号不存在"
	} else if userpasswd.Passwd != code {
		return "密码错误"
	}
	return "success"
}

// 修改密码
func (user *User) ChangePassword(newPassword string) error {
	db := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	code := config.Encryption(newPassword)

	_, err = tx.Exec("UPDATE user SET password = ? WHERE account = ?", code, user.Account)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// 发送验证码
func (user *User) Verification() error {
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	if user.FindVerificationCode() {
		return errors.New("验证码已发送，请 2 分钟后再试")
	}

	verificationCode := user.sendCode()
	// 验证码持续时间 2 分钟，过期自动失效
	_, err := connect.Do("SET", user.Account, verificationCode, "ex", "120")
	if err != nil {
		return err
	}

	// 接收者邮箱
	mails := ownmail.GetNewMail(user.Account)

	mails.Send("验证码", "<h1>您的验证码为:"+verificationCode+"<h1>", gomail.NewMessage())

	return nil
}

// 获取验证码
func (user *User) GetVerificationCode() string {
	if !user.FindVerificationCode() {
		return "not exist"
	}
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	reply, _ := redis.String(connect.Do("GET", user.Account))
	return reply
}

// 查找验证码是否过期
func (user *User) FindVerificationCode() bool {
	connect, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer connect.Close()

	isExist, _ := redis.Bool(connect.Do("EXISTS", user.Account))
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
