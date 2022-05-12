package user

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 注册
func Register(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())
	result := map[string]interface{}{
		"msg": "注册成功！",
	}

	// 验证码
	code := ctx.PostForm("code")
	// 构造对象
	userInfo := &User{
		UserName: ctx.PostForm("userName"),
		Account:  ctx.PostForm("userEmail"),
		Password: ctx.PostForm("userPassword"),
	}

	// 检查是否含有特殊字符
	words := "~`!@#$%^&*()_+-=[]\\{}|'\";:,./<>?"

	// 遍历昵称
	for i := range userInfo.UserName {
		for j := range words {
			if userInfo.UserName[i] == words[j] {
				result["msg"] = "名称不能含有特殊字符"
				ctx.JSON(http.StatusOK, result)
				return
			}
		}
	}

	// 是否允许注册
	if !config.AllowRegister && userInfo.Account != config.SystemAccount {
		result["msg"] = "暂不开放注册"
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 检查是否有字段为空或重复注册或验证码错误
	if userInfo.UserName == "" || userInfo.Account == "" || userInfo.Password == "" || code == "" {
		result["msg"] = "还有字段未填写"
		ctx.JSON(http.StatusOK, result)
		return
	} else if userInfo.CheckUserExist() {
		result["msg"] = "用户已存在"
		ctx.JSON(http.StatusOK, result)
		return
	} else if getCode, err := userInfo.GetVerificationCode(); err != nil || code != getCode {
		if err != nil {
			result["msg"] = err.Error()
		} else {
			result["msg"] = "验证码错误"
		}
		ctx.JSON(http.StatusOK, result)
		return
	} else if err := algorithm.JudgePasswordIllegal(userInfo.Password); err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 注册
	err := userInfo.Register()
	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 选取 ID
	id, err := SelectUserIDByAccount(userInfo.Account)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 创建文件夹
	os.MkdirAll("./bbsFile/"+id+"/", 0644)
	// 读取模板
	bytes, err := ioutil.ReadFile("./blogTemplate.html")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 写入
	ioutil.WriteFile("./bbsFile/"+id+"/user.html", bytes, 0644)
	ctx.JSON(http.StatusOK, result)
}

// 修改密码
func ChangePassWord(ctx *gin.Context) {

	// 判断 IP 访问合法性
	if ok := algorithm.IfShouldGetRestricted(ctx.ClientIP()); ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	algorithm.AddOneForThisIP(ctx.ClientIP())

	result := map[string]interface{}{
		"msg": "success",
	}

	// 验证码
	code := ctx.PostForm("code")
	// 密码
	passwd := ctx.PostForm("userPassword")
	// 创建对象
	userInfo := &User{
		Account: ctx.PostForm("userEmail"),
	}

	// 验证码错误，直接返回
	if ret, err := userInfo.GetVerificationCode(); ret != code || err != nil {
		if err != nil {
			result["msg"] = err.Error()
		} else {
			result["msg"] = "验证码错误！"
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 判断密码合法性
	if err := algorithm.JudgePasswordIllegal(passwd); err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 加密
	codeString := algorithm.Encryption(passwd)

	// 更改密码
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo+"user")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer mysqlClient.Close()

	// 选取
	_, err = mysqlClient.Exec("UPDATE user SET password = ? WHERE account = ?", codeString, userInfo.Account)

	if err != nil {
		result["msg"] = err.Error()
	}
	ctx.JSON(http.StatusOK, result)
}
