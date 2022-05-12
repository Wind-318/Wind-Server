package user

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callDatabase"
	"Project/callClient/callUser"
	"Project/config"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result := map[string]interface{}{
		"msg": "注册成功！",
	}
	result, err := callUser.CallUserRegister(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	} else if result["msg"] != "注册成功！" {
		ctx.JSON(http.StatusOK, result)
		return
	}

	// 选取 ID
	id, err := callDatabase.CallMySQLSelectUserIDByAccount(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", ctx.PostForm("userEmail"))
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 创建文件夹
	os.MkdirAll("./bbsFile/"+strconv.Itoa(id)+"/", 0644)
	// 读取模板
	bytes, err := ioutil.ReadFile("./blogTemplate.html")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 写入
	ioutil.WriteFile("./bbsFile/"+strconv.Itoa(id)+"/user.html", bytes, 0644)
	ctx.JSON(http.StatusOK, result)
}

// 修改密码
func ChangePassWord(ctx *gin.Context) {
	// 判断 IP 访问合法性
	if ok, err := callAlgorithm.CallAlgorithmIfRestricted(ctx); err == nil && ok {
		result := map[string]interface{}{
			"msg": "访问过于频繁，请 20 秒后再试！",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}
	callAlgorithm.CallAlgorithmAddIP(ctx)
	result, err := callUser.CallUserChangePassWord(ctx)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	} else if result["msg"] != "success" {
		ctx.JSON(http.StatusOK, result)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
