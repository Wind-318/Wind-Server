package callUser

import (
	"Project/callClient/data"
	"Project/serviceSearch"
	"errors"
	"net/rpc"

	"github.com/gin-gonic/gin"
)

// 操作客户端
type UserClient struct {
	*rpc.Client
}

// 拨号
func dialUserService(network, address string) (*UserClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &UserClient{Client: client}, nil
}

// 判断是否处于登录状态
func CallUserIsLogin(ctx *gin.Context) (bool, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return false, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return false, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return false, errors.New("尚未登陆！")
	}
	var result bool
	// 调用
	err = client.Client.Call(data.UserServiceName+".IsLogin", cookie, &result)
	// 错误处理
	if err != nil {
		return false, err
	}

	return result, nil
}

// 获取邮箱
func CallUserGetUserEmail(ctx *gin.Context) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return "", errors.New("尚未登陆！")
	}
	var result string
	// 调用
	err = client.Client.Call(data.UserServiceName+".GetUserEmail", cookie, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 获取所有用户名称
func CallUserGetUsersName(ctx *gin.Context) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	_, err = ctx.Cookie("cookie")
	if err != nil {
		return nil, errors.New("尚未登陆！")
	}
	var result []string
	// 调用
	err = client.Client.Call(data.UserServiceName+".GetUsersName", "", &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 检查是否是管理员
func CallUserIsSystem(ctx *gin.Context) (bool, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return false, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return false, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return false, errors.New("尚未登陆！")
	}
	var result bool
	// 调用
	err = client.Client.Call(data.UserServiceName+".IsSystem", cookie, &result)
	// 错误处理
	if err != nil {
		return false, err
	}

	return result, nil
}

// 检查是否是管理员或文章作者
func CallUserIsSystemOrAuthor(ctx *gin.Context) (bool, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return false, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return false, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return false, errors.New("尚未登陆！")
	}
	obj := data.UserData{
		AuthorID: ctx.PostForm("id"),
		Cookie:   cookie,
	}
	var result bool
	// 调用
	err = client.Client.Call(data.UserServiceName+".IsSystemOrAuthor", obj, &result)
	// 错误处理
	if err != nil {
		return false, err
	}

	return result, nil
}

// 获取用户名
func CallUserGetName(ctx *gin.Context) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return "", errors.New("尚未登陆！")
	}
	var result string
	// 调用
	err = client.Client.Call(data.UserServiceName+".GetName", cookie, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 签到积分
func CallUserSignGetScore(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	// 调用
	err = client.Client.Call(data.UserServiceName+".SignGetScore", cookie, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 登录
func CallUserLogin(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	_, err = ctx.Cookie("cookie")
	if err == nil {
		return nil, errors.New("已登录！")
	}
	obj := data.UserData{
		Cookie:       " ",
		UserAccount:  ctx.PostForm("userEmail"),
		UserPassword: ctx.PostForm("userPassword"),
	}
	var result map[string]interface{}
	// 调用
	err = client.Client.Call(data.UserServiceName+".Login", obj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 注册
func CallUserRegister(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	obj := data.UserData{
		UserName:     ctx.PostForm("userName"),
		UserAccount:  ctx.PostForm("userEmail"),
		UserPassword: ctx.PostForm("userPassword"),
		Code:         ctx.PostForm("code"),
	}
	var result map[string]interface{}
	// 调用
	err = client.Client.Call(data.UserServiceName+".Register", obj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 修改密码
func CallUserChangePassWord(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.UserServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialUserService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	obj := data.UserData{
		UserAccount:  ctx.PostForm("userEmail"),
		UserPassword: ctx.PostForm("userPassword"),
		Code:         ctx.PostForm("code"),
	}
	var result map[string]interface{}
	// 调用
	err = client.Client.Call(data.UserServiceName+".ChangePassWord", obj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}
