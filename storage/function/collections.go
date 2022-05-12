package function

import (
	"storage/callClient/callDatabase"
	"storage/callClient/callUser"
	"storage/config"
	"storage/registerCenter"
)

// 获取网站链接
func GetWebs(request registerCenter.StoragePictureData) map[string]interface{} {
	// 返回信息
	result := map[string]interface{}{
		"ids":      0,
		"urls":     0,
		"comments": 0,
		"picurls":  0,
	}
	// 非登录用户直接返回
	if ok, err := callUser.CallUserIsLogin(request.Cookie); err != nil || !ok {
		result["msg"] = "尚未登陆"
		return result
	}

	// 查找收藏网站信息
	arr, err := callDatabase.CallMySQLSelectCollectionWebsite(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "storage")
	if err != nil {
		result["msg"] = err.Error()
		return result
	}

	ids := make([]int, len(arr))
	urls := make([]string, len(arr))
	comments := make([]string, len(arr))
	picurls := make([]string, len(arr))

	for index := range arr {
		ids[index] = arr[index].ID
		urls[index] = arr[index].Url
		comments[index] = arr[index].Description
		picurls[index] = arr[index].Picurl
	}

	result["ids"] = ids
	result["urls"] = urls
	result["comments"] = comments
	result["picurls"] = picurls

	return result
}

// 增加收藏网站
func PutWebs(request registerCenter.StoragePictureData) map[string]interface{} {
	result := map[string]interface{}{}
	// 判断管理员账户
	if ok, err := callUser.CallUserIsLogin(request.Cookie); err != nil || !ok {
		result["msg"] = "非管理员账户"
		return result
	}

	// 获取 url
	url := request.Url
	// 获取名称
	comment := request.Comment
	// 获取网站图标
	pic := request.Picurl

	err := callDatabase.CallMySQLInsertCollectionWebsite(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "storage", url, comment, pic)
	if err != nil {
		result["msg"] = err.Error()
	}
	return result
}
