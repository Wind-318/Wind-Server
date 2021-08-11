package functions

import (
	"Project/Mail"
	"Project/Text"
	"Project/Users"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

// 发送邮件
func SendStock(ctx *gin.Context) {
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.HTML(http.StatusNotAcceptable, "serverError.html", nil)
		return
	}
	users := Mail.GetNewMail(cookie)
	rand.Seed(time.Now().UnixNano())

	err = users.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", Text.SelectFirst10(), gomail.NewMessage())
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "serverError.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "success.html", nil)
}

// 发送验证码
func SendCode(ctx *gin.Context) {
	userEmail := ctx.PostForm("userEmail")
	result := map[string]interface{}{
		"code": 200,
		"msg":  "已发送，若未收到请检查垃圾箱",
	}
	userInfo := &Users.User{
		MailAccount: userEmail,
	}
	if userEmail == "" {
		result["msg"] = "邮箱不能为空"
		ctx.JSON(http.StatusOK, result)
		return
	} else if userInfo.CheckUserExist() {
		result["msg"] = "用户已存在"
		ctx.JSON(http.StatusOK, result)
		return
	}

	err := userInfo.Verification()
	if err != nil {
		result["msg"] = err.Error()
	}

	ctx.JSON(http.StatusOK, result)
}
