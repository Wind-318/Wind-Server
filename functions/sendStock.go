package functions

import (
<<<<<<< HEAD
	"1/Mail"
	"1/Text"
=======
	"math/rand"
>>>>>>> e13850d... 更新图片发送
	"net/http"
<<<<<<< HEAD
=======
	"project/Mail"
	"project/Text"
	"strconv"
	"sync"
>>>>>>> 0c71c88... 更新
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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		rand.Seed(time.Now().UnixNano())
		picNum := strconv.Itoa(rand.Intn(248) + 1)
		users.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", Text.SelectFirst10WithPicture(picNum), gomail.NewMessage(), ".\\picture\\"+picNum+".jpg")
	}()

	ctx.String(http.StatusOK, "已发送，如果没有收到请检查垃圾箱。")
	wg.Wait()
}
