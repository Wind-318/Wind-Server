package user

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"errors"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 用户信息
type User struct {
	// 昵称
	UserName string
	// 账号
	Account string
	// 密码
	Password string
}

// 检查用户是否存在
func (user *User) CheckUserExist() bool {
	// 查看账号是否存在
	_, err := SelectUserAccountByName(user.UserName)
	if err == nil {
		return true
	}

	// 查看用户名是否存在
	_, err = SelectUserNameByAccount(user.Account)
	return err == nil
}

// 注册功能
func (user *User) Register() error {
	// 获取加密后密码
	code := algorithm.Encryption(user.Password)
	// 插入新用户
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 插入
	_, err = mysqlClient.Exec("INSERT INTO user VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 0, user.Account, code, user.UserName, 0, 0, config.Addr+"picture/defaultPic.jpg", 0, 0, 0)

	return err
}

// 登录功能
func (user *User) Login() error {
	// 加密
	code := algorithm.Encryption(user.Password)

	// 选取密码
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	var ans string
	// 选取
	err = mysqlClient.Get(&ans, "SELECT password FROM user WHERE account = ?", user.Account)
	if err != nil {
		return errors.New("账号不存在")
	} else if ans != code {
		return errors.New("密码错误")
	}

	// 返回
	return nil
}

// 获取验证码
func (user *User) GetVerificationCode() (string, error) {
	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return "", err
	}
	defer redisCli.Close()
	// 无验证码
	if !user.FindVerificationCode() {
		return "", errors.New("验证码不存在！")
	}

	// 获取验证码
	reply, err := redis.String(redisCli.Do("GET", user.Account))
	// 错误处理
	if err != nil {
		return "", err
	}
	return reply, nil
}

// 查找验证码是否过期
func (user *User) FindVerificationCode() bool {
	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return true
	}
	defer redisCli.Close()
	// 检查是否存在
	isExist, err := redis.Bool(redisCli.Do("EXISTS", user.Account))
	// 错误处理
	if err != nil {
		return false
	}
	return isExist
}

// 生成 6 位数验证码
func (user *User) GenerateCode() string {
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
