package registerCenter

// 服务名
var SubscribeServiceName = "SubscribeService"

// 服务接口
type SubscribeServiceInterface interface {
	// 获取订阅名单
	SelectUsersAccount(request interface{}, reply *[]string) error
}
