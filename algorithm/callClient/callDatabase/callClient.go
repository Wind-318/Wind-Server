package callDatabase

import (
	"algorithm/callClient/data"
	"algorithm/serviceSearch"
	"net/rpc"
)

// 操作数据库客户端
type DatabaseClient struct {
	*rpc.Client
}

// 拨号
func dialRedisService(network, address string) (*DatabaseClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &DatabaseClient{Client: client}, nil
}

// 加入缓存
func (c *DatabaseClient) OperateRedisDatabase(request data.RedisDataInterface, reply *interface{}) error {
	return c.Client.Call(data.RedisServiceName+".OperateRedisDatabase", request, reply)
}

// 获取所有支持的运算
func (c *DatabaseClient) GetOperatorsForRedis(request struct{}, reply *[]string) error {
	return c.Client.Call(data.RedisServiceName+".GetOperatorsForRedis", request, reply)
}

// 调用微服务，传入指令和参数
func CallRedis(ip, port, operate string, parameters ...interface{}) (interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.RedisServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.RedisDataInterface{
		IP:         ip,
		Port:       port,
		Operate:    operate,
		Parameters: parameters,
		Operation:  "operate",
	}

	var result interface{}
	// 调用
	err = client.OperateRedisDatabase(dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 调用微服务，传入指令和参数
func CallMySQLSelectUsersAccount(account, password, ip, port, name string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
	}

	var result []string
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectUsersAccount", dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 选取剩余空间大小
func CallMySQLSelectRamainingSpace(userAccount, account, password, ip, port, name string) (int64, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return 0, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return 0, err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		UserAccount: userAccount,
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
	}

	var result int64
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectRemainingSpace", dataObj, &result)
	// 错误处理
	if err != nil {
		return 0, err
	}

	return result, nil
}

// 更新剩余空间大小
func CallMySQLUpdateRamainingSpace(userAccount, account, password, ip, port, name string, space int64) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		UserAccount: userAccount,
		Space:       space,
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".UpdateRemainingSpace", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 根据名称获取账号
func CallMySQLSelectUserAccountByName(userName, account, password, ip, port, name string) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		UserName: userName,
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
	}

	var result string
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectUserAccountByName", dataObj, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 根据账号和文件夹名获取图片
func CallMySQLSelectStoragePictureByFilename(userAccount, filename, account, password, ip, port, name string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		UserAccount: userAccount,
		Filename:    filename,
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
	}

	var result []string
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectStoragePictureByFilename", dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 根据账号和文件夹名获取图片
func CallMySQLSelectStorageSmallPictureByFilename(userAccount, filename, account, password, ip, port, name string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		UserAccount: userAccount,
		Filename:    filename,
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
	}

	var result []string
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectStorageSmallPictureByFilename", dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 根据账号和文件夹名获取文件总数
func CallMySQLSelectStorageFilesNumber(userAccount, filename, account, password, ip, port, name string) (int, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return 0, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return 0, err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		UserAccount: userAccount,
		Filename:    filename,
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
	}

	var result int
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectStorageFilesNumber", dataObj, &result)
	// 错误处理
	if err != nil {
		return 0, err
	}

	return result, nil
}

// 根据账号获取文件夹数量
func CallMySQLSelectStorageUserFilesNumber(userAccount, account, password, ip, port, name string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		UserAccount: userAccount,
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
	}

	var result []string
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectStorageUserFilesNumber", dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 初始化数据库
func CallMySQLInitDatabase(account, password, ip, port string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
	}

	var result []string
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".InitDatabase", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 根据名称获取动漫信息
func CallMySQLSelectAnimeInfoByName(account, password, ip, port, name, animeName string) (data.SelectAnimeDataInterface, error) {
	var result data.SelectAnimeDataInterface
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return result, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return result, err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		AnimeName: animeName,
		Account:   account,
		Password:  password,
		IP:        ip,
		Port:      port,
		Name:      name,
	}

	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectAnimeInfoByName", dataObj, &result)
	// 错误处理
	if err != nil {
		return result, err
	}

	return result, nil
}

// 根据名称获取来源
func CallMySQLSelectAnimeSourceByName(account, password, ip, port, name, animeName string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		AnimeName: animeName,
		Account:   account,
		Password:  password,
		IP:        ip,
		Port:      port,
		Name:      name,
	}

	var result []string
	// 调用
	err = client.Client.Call(data.MySQLServiceName+".SelectAnimeSourceByName", dataObj, &result)
	// 错误处理
	if err != nil {
		return result, err
	}

	return result, nil
}

// 根据名称获取 url
func CallMySQLSelectAnimeUrlByName(account, password, ip, port, name, animeName string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		AnimeName: animeName,
		Account:   account,
		Password:  password,
		IP:        ip,
		Port:      port,
		Name:      name,
	}

	// 调用
	var result []string
	err = client.Client.Call(data.MySQLServiceName+".SelectAnimeUrlByName", dataObj, &result)
	// 错误处理
	if err != nil {
		return result, err
	}

	return result, nil
}

// 选取动漫名
func CallMySQLSelectAnimeName(account, password, ip, port, name string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
	}

	// 调用
	var result []string
	err = client.Client.Call(data.MySQLServiceName+".SelectAnimeName", dataObj, &result)
	// 错误处理
	if err != nil {
		return result, err
	}

	return result, nil
}

// 选取新番
func CallMySQLSelectNewAnime(account, password, ip, port, name string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
	}

	// 调用
	var result []string
	err = client.Client.Call(data.MySQLServiceName+".SelectNewAnime", dataObj, &result)
	// 错误处理
	if err != nil {
		return result, err
	}

	return result, nil
}

// 根据年份选取动漫
func CallMySQLSelectAnimeByYear(account, password, ip, port, name string, year int) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
		Year:     year,
	}

	// 调用
	var result []string
	err = client.Client.Call(data.MySQLServiceName+".SelectAnimeByYear", dataObj, &result)
	// 错误处理
	if err != nil {
		return result, err
	}

	return result, nil
}

// 选取收藏网站
func CallMySQLSelectCollectionWebsite(account, password, ip, port, name string) ([]data.SelectCollectionDataInterface, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectCollectionDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
	}

	// 调用
	var result []data.SelectCollectionDataInterface
	err = client.Client.Call(data.MySQLServiceName+".SelectCollectionWebsite", dataObj, &result)
	// 错误处理
	if err != nil {
		return result, err
	}

	return result, nil
}

// 插入新网站
func CallMySQLInsertCollectionWebsite(account, password, ip, port, name, url, comment, picurl string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectCollectionDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
		Url:      url,
		Comment:  comment,
		Picurl:   picurl,
	}

	// 调用
	var result interface{}
	err = client.Client.Call(data.MySQLServiceName+".InsertCollectionWebsite", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 更新本季新番取消最新状态
func CallMySQLUpdateAnimeSetIsNew0(account, password, ip, port, name string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
	}

	// 调用
	var result interface{}
	err = client.Client.Call(data.MySQLServiceName+".UpdateAnimeSetIsNew0", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 查看某来源动漫是否已存在，已存在更新为新番
func CallMySQLSelectAnimeCountByNameAndSource(account, password, ip, port, name, animeName, source string) (bool, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return true, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return true, err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:   account,
		Password:  password,
		IP:        ip,
		Port:      port,
		Name:      name,
		AnimeName: animeName,
		Source:    []string{source},
	}

	// 调用
	var result bool
	err = client.Client.Call(data.MySQLServiceName+".SelectAnimeCountByNameAndSource", dataObj, &result)
	// 错误处理
	if err != nil {
		return true, err
	}

	return result, nil
}

// 根据名称查看动漫是否已存在
func CallMySQLSelectAnimeCountByName(account, password, ip, port, name, animeName string) (bool, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return true, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return true, err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:   account,
		Password:  password,
		IP:        ip,
		Port:      port,
		Name:      name,
		AnimeName: animeName,
	}

	// 调用
	var result = true
	err = client.Client.Call(data.MySQLServiceName+".SelectAnimeCountByName", dataObj, &result)
	// 错误处理
	if err != nil {
		return true, err
	}

	return result, nil
}

// 更新动漫资料
func CallMySQLUpdateAnimeInfo(account, password, ip, port, name, url, description, picurl, animeName string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
		Url:         url,
		Description: description,
		Picurl:      picurl,
		AnimeName:   animeName,
	}

	// 调用
	var result interface{}
	err = client.Client.Call(data.MySQLServiceName+".UpdateAnimeInfo", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 插入新动漫到 bangumi
func CallMySQLInsertAnimeForBangumi(account, password, ip, port, name, animeName, url, description, picurl string, year, isNew int) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
		AnimeName:   animeName,
		Url:         url,
		Description: description,
		Picurl:      picurl,
		Year:        year,
		IsNew:       isNew,
	}

	// 调用
	var result interface{}
	err = client.Client.Call(data.MySQLServiceName+".InsertAnimeForBangumi", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 插入新动漫到 source
func CallMySQLInsertAnimeForSource(account, password, ip, port, name, animeName, source, url string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectAnimeDataInterface{
		Account:   account,
		Password:  password,
		IP:        ip,
		Port:      port,
		Name:      name,
		AnimeName: animeName,
		Source:    []string{source},
		Url:       url,
	}

	// 调用
	var result interface{}
	err = client.Client.Call(data.MySQLServiceName+".InsertAnimeForSource", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 根据账号选取用户名
func CallMySQLSelectUserNameByAccount(account, password, ip, port, name, userAccount string) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
		UserAccount: userAccount,
	}

	// 调用
	var result string
	err = client.Client.Call(data.MySQLServiceName+".SelectUserNameByAccount", dataObj, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 选取所有用户名称
func CallMySQLSelectUserNames(account, password, ip, port, name string) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
	}

	// 调用
	var result []string
	err = client.Client.Call(data.MySQLServiceName+".SelectUserNames", dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 选取所有图片信息
func CallMySQLSelectStorageFilesInfo(account, password, ip, port, name string) ([]data.SelectStorageFilesInfo, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	// 创建数据对象
	dataObj := data.SelectStorageFilesInfo{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
	}

	// 调用
	var result []data.SelectStorageFilesInfo
	err = client.Client.Call(data.MySQLServiceName+".SelectStorageFilesInfo", dataObj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 更新缩略图路径
func CallMySQLUpdateStorageSmallPicture(account, password, ip, port, name, smallpath, path string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectStorageFilesInfo{
		Account:   account,
		Password:  password,
		IP:        ip,
		Port:      port,
		Name:      name,
		Smallpath: smallpath,
		Path:      path,
	}

	// 调用
	var result interface{}
	err = client.Client.Call(data.MySQLServiceName+".UpdateStorageSmallPicture", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 插入新用户
func CallMySQLInsertNewUser(account, password, ip, port, name, userAccount, userPassword, userName, pic string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:      account,
		Password:     password,
		IP:           ip,
		Port:         port,
		Name:         name,
		UserAccount:  userAccount,
		UserPassword: userPassword,
		UserName:     userName,
		Picture:      pic,
	}

	// 调用
	var result interface{}
	err = client.Client.Call(data.MySQLServiceName+".InsertNewUser", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 选取 ID
func CallMySQLSelectUserIDByAccount(account, password, ip, port, name, userAccount string) (int, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return 0, err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return 0, err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
		UserAccount: userAccount,
	}

	// 调用
	var result int
	err = client.Client.Call(data.MySQLServiceName+".SelectUserIDByAccount", dataObj, &result)
	// 错误处理
	if err != nil {
		return 0, err
	}

	return result, nil
}

// 选取用户密码
func CallMySQLSelectUserPasswordByAccount(account, password, ip, port, name, userAccount string) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:     account,
		Password:    password,
		IP:          ip,
		Port:        port,
		Name:        name,
		UserAccount: userAccount,
	}

	// 调用
	var result string
	err = client.Client.Call(data.MySQLServiceName+".SelectUserPasswordByAccount", dataObj, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}

// 修改密码
func CallMySQLUpdateUserPassword(account, password, ip, port, name, userAccount, userPassword string) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:      account,
		Password:     password,
		IP:           ip,
		Port:         port,
		Name:         name,
		UserAccount:  userAccount,
		UserPassword: userPassword,
	}

	// 调用
	var result interface{}
	err = client.Client.Call(data.MySQLServiceName+".UpdateUserPassword", dataObj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 根据 ID 获取账号
func CallMySQLSelectUserAccountByID(account, password, ip, port, name, id string) (string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.MySQLServiceName)
	if err != nil {
		return "", err
	}
	// 拨号
	client, err := dialRedisService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return "", err
	}

	// 创建数据对象
	dataObj := data.SelectDataInterface{
		Account:  account,
		Password: password,
		IP:       ip,
		Port:     port,
		Name:     name,
		UserName: id,
	}

	// 调用
	var result string
	err = client.Client.Call(data.MySQLServiceName+".SelectUserAccountByID", dataObj, &result)
	// 错误处理
	if err != nil {
		return "", err
	}

	return result, nil
}
