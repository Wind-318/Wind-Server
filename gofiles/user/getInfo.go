package user

import (
	"Project/config"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 验证是否处于登录状态
func IsLogin(ctx *gin.Context) bool {
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return false
	}

	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return false
	}
	defer redisCli.Close()
	// 检查存在
	isExist, err := redis.Bool(redisCli.Do("HEXISTS", cookie, "email"))
	// 不存在或者出错，返回 false
	if err != nil || !isExist {
		return false
	}

	return isExist
}

// 查看是否是管理员
func IsSystem(ctx *gin.Context) {
	result := map[string]interface{}{}

	ok := true
	// 验证
	if email, err := GetUserEmail(ctx); err != nil || email != config.SystemAccount {
		ok = false
	}

	if ok {
		result["msg"] = "success"
	} else {
		result["msg"] = "false"
	}
	ctx.JSON(http.StatusOK, result)
}

// 查看是否是管理员或文章作者
func IsSystemOrAuthor(ctx *gin.Context) {
	result := map[string]interface{}{}

	// 获取 id
	id := ctx.PostForm("id")

	ok := true
	// 查询作者账号
	ids, err := SelectUserAccountByID(id)
	if err != nil {
		ok = false
	}
	account, err := GetUserEmail(ctx)
	if err != nil {
		ok = false
	}
	// 验证
	if err != nil {
		ok = false
	} else if account != config.SystemAccount && account != ids {
		ok = false
	}

	if ok {
		result["msg"] = "success"
	} else {
		result["msg"] = "false"
	}
	ctx.JSON(http.StatusOK, result)
}

// 根据账号选取 ID
func SelectUserIDByAccount(userAccount string) (string, error) {
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		return "", err
	}
	defer mysqlClient.Close()

	var ans string
	// 选取
	err = mysqlClient.Get(&ans, "SELECT id FROM user WHERE account = ?", userAccount)

	return ans, err
}

// 根据文章 ID 选择账号
func SelectUserAccountByID(id string) (string, error) {
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"bbs")
	if err != nil {
		return "", err
	}
	defer mysqlClient.Close()

	// 选取
	ans := ""
	err = mysqlClient.Get(&ans, "SELECT authoremail FROM blog WHERE id = ?", id)

	return ans, err
}

// 根据用户名选取账号
func SelectUserAccountByName(userName string) (string, error) {
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		return "", err
	}
	defer mysqlClient.Close()

	var ans string
	// 选取
	err = mysqlClient.Get(&ans, "SELECT account FROM user WHERE username = ?", userName)

	return ans, err
}

// 根据账号选取用户名
func SelectUserNameByAccount(userAccount string) (string, error) {
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		return "", err
	}
	defer mysqlClient.Close()

	var ans string
	// 选取
	err = mysqlClient.Get(&ans, "SELECT username FROM user WHERE account = ?", userAccount)

	return ans, err
}

// 查询订阅股票信息用户
func SelectUsersAccount() ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	ans := []string{}
	// 选取
	err = mysqlClient.Select(&ans, "SELECT account FROM subscribe WHERE stock = ?", 1)

	return ans, err
}

// 获取邮箱
func GetUserEmail(ctx *gin.Context) (string, error) {
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return "", nil
	}
	// 连接 redis
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return "", err
	}
	defer redisCli.Close()
	// 拿到存储的账号
	email, err := redis.String(redisCli.Do("HGET", cookie, "email"))
	if err != nil {
		return "", err
	}
	return email, nil
}

func GetUsersNames(ctx *gin.Context) {
	result := map[string]interface{}{}
	result["names"], result["msg"] = GetUsersName()
	ctx.JSON(http.StatusOK, result)
}

// 获取所有用户名称
func GetUsersName() ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	var ans []string
	// 选取
	err = mysqlClient.Select(&ans, "SELECT username FROM user")
	if err != nil {
		return nil, err
	}

	return ans, nil
}

// 获取用户名称
func GetName(ctx *gin.Context) (string, error) {
	// 检索名称
	email, err := GetUserEmail(ctx)
	if err != nil {
		return "", err
	}
	name, err := SelectUserNameByAccount(email)

	return name, err
}
