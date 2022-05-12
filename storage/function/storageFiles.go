package function

import (
	"errors"
	"storage/callClient/callDatabase"
	"storage/callClient/callUser"
	"storage/config"
	"storage/registerCenter"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 互斥锁
var mutex sync.Mutex

// 批量下载
func DownloadFiles(request registerCenter.StoragePictureData) map[string]interface{} {
	// 返回值
	result := map[string]interface{}{}
	// 检验合法性
	if ok, err := callUser.CallUserIsLogin(request.Cookie); err != nil || !ok {
		result["msg"] = "未登录"
		return result
	}

	// 获取邮箱
	name := request.UserName
	// 获取文件夹名称
	fileName := request.FolderName

	result["zipfile1"], result["zipfile2"] = "./userFile/"+name+"/"+fileName, "./userFile/"+name+"/"+fileName+".zip"
	result["url"] = "../userFile/" + name + "/" + fileName + ".zip"
	result["name"] = fileName + ".zip"
	return result
}

// 存储文件
func StorageFiles(request registerCenter.StoragePictureData) error {
	// 获取账号
	account, err := callUser.CallUserGetUserEmail(request.Cookie)
	if err != nil {
		return err
	}

	// 加锁
	mutex.Lock()
	// 从缓存读剩余可用空间，同时只有一个线程可以读取，若未访问到，刷入缓存中，其余线程就能直接读取到；
	remainSpace, err := redis.Int64(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "GET", account+"StorageSpace"))
	// 缓存未命中
	if err != nil {
		// 在数据库中查询
		space, err := callDatabase.CallMySQLSelectRamainingSpace(account, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user")
		if err != nil {
			// 解锁
			mutex.Unlock()
			return err
		}
		remainSpace = space
		// 写入缓存
		callDatabase.CallRedis(config.RedisIP, config.RedisPort, "SET", account+"StorageSpace", remainSpace)
	}
	// 解锁
	mutex.Unlock()

	// 文件
	var nowSpace int64
	// 获取文件大小
	useSpace := -request.Size / 1024

	// 减少可用空间
	err = callDatabase.CallMySQLUpdateStorageUnusedCapacity(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", account, -useSpace)
	if err != nil {
		return err
	}
	// 从缓存中减去文件大小
	nowSpace, err = redis.Int64(callDatabase.CallRedis(config.RedisIP, config.RedisPort, "INCRBY", account+"StorageSpace", useSpace))
	if err != nil {
		// 补偿，退款！
		callDatabase.CallMySQLUpdateStorageUnusedCapacity(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", account, useSpace)
		return err
	}
	// 空间不足时，淘汰缓存，等待刷入数据库的新数据
	if nowSpace < 0 {
		// 删去缓存
		callDatabase.CallRedis(config.RedisIP, config.RedisPort, "del", account+"StorageSpace")
		// 补偿，退款！
		callDatabase.CallMySQLUpdateStorageUnusedCapacity(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user", account, useSpace)
		// 延时双删，淘汰可能刷入的旧数据
		go func() {
			time.Sleep(time.Millisecond * 500)
			callDatabase.CallRedis(config.RedisIP, config.RedisPort, "del", account+"StorageSpace")
		}()
		return errors.New("剩余空间不足！")
	}

	return nil
}
