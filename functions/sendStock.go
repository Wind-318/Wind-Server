package functions

import (
	"1/Mail"
	"1/Text"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

func ToFuntions(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "function.html", nil)
}

func SendStock(ctx *gin.Context) {
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		ctx.String(http.StatusBadRequest, "请先登录")
		return
	}
	users := Mail.GetNewMail(cookie)
	users.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", Text.SelectFirst20(), gomail.NewMessage())
	ctx.String(http.StatusOK, "已发送，如果没有收到请检查垃圾箱。")
}
