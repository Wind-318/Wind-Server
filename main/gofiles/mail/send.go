package mail

import (
	"Project/callClient/callAlgorithm"
	"Project/callClient/callDatabase"
	"Project/config"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// 发送邮件
func SendStockToUser(ctx *gin.Context) {}

// 发送验证码
func SendCode(ctx *gin.Context) {
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
		"msg": "已发送，若未接收到请检查垃圾箱。",
	}

	userEmail := ctx.PostForm("userEmail")

	// 检查是否存在
	isExist, err := redis.Bool(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "EXISTS", userEmail))
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	} else if isExist {
		result["msg"] = "邮件已发送，请 5 分钟后再试"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 生产消息，投入 stream 中
	callDatabase.CallRedis(config.RedisIP, config.RedisPort, "XADD", "SendCode", "*", "email", userEmail)
	ctx.JSON(http.StatusOK, result)
}
