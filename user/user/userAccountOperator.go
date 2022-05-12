package user

import (
	"user/callClient/callAlgorithm"
	"user/callClient/callDatabase"
	"user/config"
	"user/registerCenter.go"
)

// 登录
func Login(request registerCenter.UserData) map[string]interface{} {
	// 返回结果
	result := map[string]interface{}{
		"msg": "success",
	}
	// 检查是否已登录
	if ok, err := IsLogin(request.Cookie); err != nil || ok {
		if err != nil {
			result["msg"] = err.Error()
		} else {
			result["msg"] = "已登录"
		}
		return result
	}

	// 检查字段合法性
	if request.UserAccount == "" || request.UserPassword == "" {
		result["msg"] = "账号或密码不能为空"
		return result
	}
	// 构造对象
	userInfo := &User{
		UserName: "",
		Account:  request.UserAccount,
		Password: request.UserPassword,
	}

	// 登录
	status := userInfo.Login()
	// 不成功返回结果
	if status != nil {
		result["msg"] = status.Error()
		return result
	}

	// 返回
	return result
}

// 注册
func Register(request registerCenter.UserData) map[string]interface{} {
	// 返回结果
	result := map[string]interface{}{
		"msg": "注册成功！",
	}

	// 构造对象
	userInfo := &User{
		UserName: request.UserName,
		Account:  request.UserAccount,
		Password: request.UserPassword,
	}

	// 检查是否含有特殊字符
	words := "~`!@#$%^&*()_+-=[]\\{}|'\";:,./<>?"

	// 遍历昵称
	for i := range request.UserName {
		for j := range words {
			if request.UserName[i] == words[j] {
				result["msg"] = "名称不能含有特殊字符"
				return result
			}
		}
	}

	// 是否允许注册
	if !config.AllowRegister && request.UserAccount != config.SystemAccount {
		result["msg"] = "暂不开放注册"
		return result
	}

	// 检查是否有字段为空或重复注册或验证码错误
	if request.UserName == "" || request.UserAccount == "" || request.UserPassword == "" || request.Code == "" {
		result["msg"] = "还有字段未填写"
		return result
	} else if userInfo.CheckUserExist() {
		result["msg"] = "用户已存在"
		return result
	} else if getCode, err := userInfo.GetVerificationCode(); err != nil || request.Code != getCode {
		if err != nil {
			result["msg"] = err.Error()
		} else {
			result["msg"] = "验证码错误"
		}
		return result
	} else if err := callAlgorithm.CallAlgorithmJudgePassword(request.UserPassword); err != nil {
		result["msg"] = err.Error()
		return result
	}

	// 注册
	err := userInfo.Register()
	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
		return result
	}

	// 返回结果
	return result
}

// 修改密码
func ChangePassWord(request registerCenter.UserData) map[string]interface{} {
	result := map[string]interface{}{
		"msg": "success",
	}

	// 创建对象
	userInfo := &User{
		Account: request.UserAccount,
	}

	// 验证码错误，直接返回
	if ret, err := userInfo.GetVerificationCode(); ret != request.Code || err != nil {
		if err != nil {
			result["msg"] = err.Error()
		} else {
			result["msg"] = "验证码错误！"
		}
		return result
	}
	// 判断密码合法性
	if err := callAlgorithm.CallAlgorithmJudgePassword(request.UserPassword); err != nil {
		result["msg"] = err.Error()
		return result
	}
	// 加密
	codeString, err := callAlgorithm.CallAlgorithmEncryption(request.UserPassword)
	if err != nil {
		result["msg"] = err.Error()
		return result
	}

	// 更改密码
	err = callDatabase.CallMySQLUpdateUserPassword(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", request.UserAccount, codeString)

	// 错误处理
	if err != nil {
		result["msg"] = err.Error()
	}

	return result
}
