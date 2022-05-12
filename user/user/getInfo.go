package user

import (
	"user/callClient/callDatabase"
	"user/config"
	"user/registerCenter.go"

	"github.com/garyburd/redigo/redis"
)

// 获取邮箱
func GetUserEmail(cookie string) (string, error) {
	// 拿到存储的账号
	email, err := redis.String(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "HGET", cookie, "email"))
	if err != nil {
		return "", err
	}
	return email, nil
}

// 获取所有用户名称
func GetUsersName(request interface{}) ([]string, error) {
	// 检索名称
	names, err := callDatabase.CallMySQLSelectUserNames(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user")
	if err != nil {
		return nil, err
	}

	return names, nil
}

// 获取用户名称
func GetName(cookie string) (string, error) {
	// 检索名称
	email, err := GetUserEmail(cookie)
	if err != nil {
		return "", err
	}
	name, err := callDatabase.CallMySQLSelectUserNameByAccount(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", email)

	return name, err
}

// 查看是否是管理员
func IsSystem(cookie string) bool {
	// 验证
	if email, err := GetUserEmail(cookie); err != nil || email != config.SystemAccount {
		return false
	}

	return true
}

// 查看是否是管理员或文章作者
func IsSystemOrAuthor(request registerCenter.UserData) bool {
	// 获取 id
	id := request.AuthorID

	// 查询作者账号
	ids, err := callDatabase.CallMySQLSelectUserAccountByID(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "bbs", id)

	if err != nil {
		return false
	}
	account, err := GetUserEmail(request.Cookie)
	if err != nil {
		return false
	}
	// 验证
	if err != nil {
		return false
	} else if account != config.SystemAccount && account != ids {
		return false
	}

	return true
}
