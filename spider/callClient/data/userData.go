package data

// 用户模块服务名
var UserServiceName = "UserService"

// 用户模块数据接口
type UserData struct {
	// 用户名
	UserName string
	// 账户
	UserAccount string
	// 密码
	UserPassword string
	// 作者 ID
	AuthorID string
	// cookie
	Cookie string
	// 验证码
	Code string
}
