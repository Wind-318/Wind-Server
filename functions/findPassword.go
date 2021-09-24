package functions

import (
	"Project/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 修改密码
func VerificationFind(ctx *gin.Context) {
	// 账号
	userEmail := ctx.PostForm("userEmail")
	// 新密码
	newPassword := ctx.PostForm("userPassword")
	// 验证码
	code := ctx.PostForm("code")
	userInfo := &Users.User{
		MailAccount: userEmail,
	}
	result := map[string]interface{}{
		"msg": "success",
	}

	// 检查信息
	if userEmail == "" || newPassword == "" || code == "" {
		result["msg"] = "还有字段未填写"
		ctx.JSON(http.StatusOK, result)
		return
	} else if !userInfo.CheckUserExist() {
		result["msg"] = "用户不存在"
		ctx.JSON(http.StatusOK, result)
		return
	} else if userInfo.GetVerificationCode() != code {
		result["msg"] = "验证码错误"
		ctx.JSON(http.StatusOK, result)
		return
	}

	err := userInfo.ChangePassword(newPassword)
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
