package functions

import (
	"Project/Users"
	"Project/infomation"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(ctx *gin.Context) {
	userName := ctx.PostForm("userName")
	code := ctx.PostForm("code")
	userEmail := ctx.PostForm("userEmail")
	passWord := ctx.PostForm("userPassword")

	userInfo := &Users.User{
		UserName:     userName,
		MailAccount:  userEmail,
		MailPassword: passWord,
	}

	result := map[string]interface{}{
		"msg": "注册成功",
	}

	if userEmail != infomation.SystemUserAccount {
		result["msg"] = "暂不开放注册"
		ctx.JSON(http.StatusOK, result)
		return
	}

	if userName == "" || userEmail == "" || passWord == "" || code == "" {
		result["msg"] = "还有字段未填写"
		ctx.JSON(http.StatusOK, result)
		return
	} else if userInfo.CheckUserExist() {
		result["msg"] = "用户已存在"
		ctx.JSON(http.StatusOK, result)
		return
	} else if code != userInfo.GetVerificationCode() {
		result["msg"] = "验证码错误"
		ctx.JSON(http.StatusOK, result)
		return
	}

	err := userInfo.Register()
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
