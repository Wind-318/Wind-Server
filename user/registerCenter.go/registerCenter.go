package registerCenter

// 用户模块服务名
var UserServiceName = "UserService"

// 用户模块服务接口
type UserServiceInterface interface {
	// 验证是否处于登录状态
	IsLogin(cookie string, reply *bool) error
	// 获取邮箱
	GetUserEmail(cookie string, reply *string) error
	// 获取所有用户名称
	GetUsersName(request interface{}, reply *[]string) error
	// 查看是否是管理员
	IsSystem(cookie string, reply *bool) error
	// 查看是否是管理员或文章作者
	IsSystemOrAuthor(request UserData, reply *bool) error
	// 获取用户名
	GetName(cookie string, reply *string) error
	// 签到积分
	SignGetScore(cookie string, reply *map[string]interface{}) error
	// 登录
	Login(request UserData, reply *map[string]interface{}) error
	// 注册
	Register(request UserData, reply *map[string]interface{}) error
	// 修改密码
	ChangePassWord(request UserData, reply *map[string]interface{}) error
}

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
