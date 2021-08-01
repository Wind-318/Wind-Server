package functions

import (
	"Project/Mail"
	"Project/Text"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

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
