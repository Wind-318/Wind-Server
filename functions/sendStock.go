package functions

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"1/Mail"
	"1/Text"
=======
=======
	"fmt"
>>>>>>> d19263e... 更新
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
	rand.Seed(time.Now().UnixNano())
	picNum := strconv.Itoa(rand.Intn(18) + 1)
	err = users.Send(time.Now().String()[:19]+" "+time.Now().Weekday().String()+"：每日要闻", Text.SelectFirst10WithPicture(picNum), gomail.NewMessage(), "./pic/"+picNum+".png")
	if err != nil {
		fmt.Println(err)
	}

	ctx.String(http.StatusOK, "已发送，如果没有收到请检查垃圾箱。")
	wg.Wait()
}
