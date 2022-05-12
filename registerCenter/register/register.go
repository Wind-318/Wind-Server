package register

// 服务名
var RegisterServiceName = "RegisterService"

// 服务接口
type RegisterServiceInterface interface {
	// 接受注册请求
	ServiceRegisterRequest(request RegisterRequestInterface, reply *interface{}) error
	// 接受注销请求
	ServiceLogoutRequest(request RegisterRequestInterface, reply *interface{}) error
	// 服务发现
	ServiceSearch(request string, reply *[]string) error
	// 全服务搜索
	ServiceSearchAll(request string, reply *map[string][]string) error
	// 服务检测
	ServiceCheck(request string, reply *bool) error
}

// 注册请求信息
type RegisterRequestInterface struct {
	// 注册服务名
	RegisterRequestServiceName string
	// 注册实例 IP
	IP string
	// 注册实例端口
	Port string
}
