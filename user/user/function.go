package user

import (
	"errors"
	"math/rand"
	"time"
	"user/callClient/callAlgorithm"
	"user/callClient/callDatabase"
	"user/config"

	"github.com/garyburd/redigo/redis"
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
	// 账号
	_, err := callDatabase.CallMySQLSelectUserAccountByName(user.UserName, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user")
	if err == nil {
		return true
	}
	// 用户名
	_, err = callDatabase.CallMySQLSelectUserNameByAccount(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", user.Account)

	return err == nil
}

// 注册功能
func (user *User) Register() error {
	// 获取加密后密码
	code, err := callAlgorithm.CallAlgorithmEncryption(user.Password)
	if err != nil {
		return err
	}
	// 插入新用户
	err = callDatabase.CallMySQLInsertNewUser(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", user.Account, code, user.UserName, config.YourName+"picture/defaultPic.jpg")
	if err != nil {
		return err
	}

	return nil
}

// 登录功能
func (user *User) Login() error {
	// 加密
	code, err := callAlgorithm.CallAlgorithmEncryption(user.Password)
	// 错误处理
	if err != nil {
		return err
	}

	// 选取密码
	userpasswd, err := callDatabase.CallMySQLSelectUserPasswordByAccount(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", user.Account)
	// 错误处理
	if err != nil {
		return errors.New("账号不存在")
	} else if userpasswd != code {
		return errors.New("密码错误")
	}

	// 返回
	return nil
}

// 获取验证码
func (user *User) GetVerificationCode() (string, error) {
	// 无验证码
	if !user.FindVerificationCode() {
		return "", errors.New("验证码不存在！")
	}

	// 获取验证码
	reply, err := redis.String(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "GET", user.Account))
	// 错误处理
	if err != nil {
		return "", err
	}
	return reply, nil
}

// 查找验证码是否过期
func (user *User) FindVerificationCode() bool {
	// 检查是否存在
	isExist, err := redis.Bool(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "EXISTS", user.Account))
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
