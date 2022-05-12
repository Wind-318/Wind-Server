package data

// 服务注册和发现中心
var RegisterServiceName = "RegisterService"

// 注册请求信息
type RegisterRequestInterface struct {
	// 注册服务名
	RegisterRequestServiceName string
	// 注册实例 IP
	IP string
	// 注册实例端口
	Port string
}
