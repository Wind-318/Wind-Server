package main

import (
	"encoding/gob"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"user/callClient/callRegister"
	"user/config"
	"user/registerCenter.go"
	"user/serviceSearch"
	"user/user"
)

func main() {
	// 注册类型
	gob.Register([]interface{}{})
	gob.Register(map[string][]string{})
	gob.Register(map[string]interface{}{})
	// 服务发现
	go serviceSearch.ServiceSearch()
	go serviceSearch.InitSearch()
	// 注册服务
	rpc.RegisterName(registerCenter.UserServiceName, new(UserService))
	rpc.HandleHTTP()

	go func() {
		// 注册
		callRegister.CallServiceRegister(registerCenter.UserServiceName, config.UserListenAddress, config.UserListenPort)
		// 监听 ctrl + c
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan)
		<-sigChan
		// 注销服务
		callRegister.CallServiceLogout(registerCenter.UserServiceName, config.UserListenAddress, config.UserListenPort)
		os.Exit(0)
	}()
	// 监听端口
	http.ListenAndServe(config.UserListenPort, nil)
}

// userServiceInterface 实现
type UserService struct{}

// 验证是否处于登录状态
func (c *UserService) IsLogin(cookie string, reply *bool) error {
	temp, err := user.IsLogin(cookie)
	*reply = temp
	return err
}

// 获取邮箱
func (c *UserService) GetUserEmail(cookie string, reply *string) error {
	temp, err := user.GetUserEmail(cookie)
	*reply = temp
	return err
}

// 获取所有用户名称
func (c *UserService) GetUsersName(cookie string, reply *[]string) error {
	temp, err := user.GetUsersName(cookie)
	*reply = temp
	return err
}

// 查看是否是管理员
func (c *UserService) IsSystem(cookie string, reply *bool) error {
	*reply = user.IsSystem(cookie)
	return nil
}

// 查看是否是管理员或文章作者
func (c *UserService) IsSystemOrAuthor(request registerCenter.UserData, reply *bool) error {
	*reply = user.IsSystemOrAuthor(request)
	return nil
}

// 获取用户名
func (c *UserService) GetName(cookie string, reply *string) error {
	temp, err := user.GetName(cookie)
	*reply = temp
	return err
}

// 登录
func (c *UserService) Login(request registerCenter.UserData, reply *map[string]interface{}) error {
	*reply = user.Login(request)
	return nil
}

// 登录
func (c *UserService) Register(request registerCenter.UserData, reply *map[string]interface{}) error {
	*reply = user.Register(request)
	return nil
}

// 签到积分
func (c *UserService) SignGetScore(cookie string, reply *map[string]interface{}) error {
	*reply = user.SignGetScore(cookie)
	return nil
}

// 修改密码
func (c *UserService) ChangePassWord(request registerCenter.UserData, reply *map[string]interface{}) error {
	*reply = user.ChangePassWord(request)
	return nil
}
