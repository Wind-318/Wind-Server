package registerCenter

// 服务名
var MailSeviceName = "MailService"

// 服务接口
type MailServiceInterface interface {
	// 发送邮件
	SendStockService(email string, reply *interface{}) error
	// 发送验证码
	SendCodeService(userEmail string, reply *interface{}) error
}

// 数据接口
type MailInterface struct{}
