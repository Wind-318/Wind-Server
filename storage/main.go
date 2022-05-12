package main

import (
	"encoding/gob"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"storage/callClient/callRegister"
	"storage/config"
	"storage/function"
	"storage/registerCenter"
	"storage/serviceSearch"
)

func main() {
	// 注册类型
	gob.Register([]interface{}{})
	gob.Register(map[string][]string{})
	gob.Register(map[string]string{})
	// 服务发现
	go serviceSearch.ServiceSearch()
	go serviceSearch.InitSearch()
	// 注册服务
	rpc.RegisterName(registerCenter.StorageServiceName, new(StorageService))
	rpc.HandleHTTP()

	go func() {
		// 注册
		callRegister.CallServiceRegister(registerCenter.StorageServiceName, config.StorageServiceAddress, config.StorageServicePort)
		// 监听 ctrl + c
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan)
		<-sigChan
		// 注销服务
		callRegister.CallServiceLogout(registerCenter.StorageServiceName, config.StorageServiceAddress, config.StorageServicePort)
		os.Exit(0)
	}()
	// 监听端口
	http.ListenAndServe(config.StorageServicePort, nil)
}

// 存储服务
type StorageService struct{}

// 删除文件
func (s *StorageService) DeleteFiles(request registerCenter.StoragePictureData, reply *interface{}) error {
	return function.DeleteFiles(request)
}

// 存储文件
func (s *StorageService) StorageFiles(request registerCenter.StoragePictureData, reply *interface{}) error {
	return function.StorageFiles(request)
}

// 获取用户存储的图片合集
func (s *StorageService) GetUserStoragePicture(request registerCenter.StoragePictureData, reply *registerCenter.StoragePicture) error {
	temp, err := function.GetUserStoragePicture(request)
	if err != nil {
		return err
	}
	*reply = temp
	return nil
}

// 获取页数
func (s *StorageService) GetUserStoragePicturePage(request registerCenter.StoragePictureData, reply *int) error {
	temp, err := function.GetUserStoragePicturePage(request)
	if err != nil {
		return err
	}
	*reply = temp
	return nil
}

// 查看文件夹个数
func (s *StorageService) GetUserFileNums(request registerCenter.StoragePictureData, reply *[]string) error {
	temp, err := function.GetUserFileNums(request)
	if err != nil {
		return err
	}
	*reply = temp
	return nil
}

// 获取网站链接
func (s *StorageService) GetWebs(request registerCenter.StoragePictureData, reply *map[string]interface{}) error {
	temp := function.GetWebs(request)
	*reply = temp
	return nil
}

// 增加收藏网站
func (s *StorageService) PutWebs(request registerCenter.StoragePictureData, reply *map[string]interface{}) error {
	temp := function.PutWebs(request)
	*reply = temp
	return nil
}

// 批量下载
func (s *StorageService) DownloadFiles(request registerCenter.StoragePictureData, reply *map[string]interface{}) error {
	temp := function.DownloadFiles(request)
	*reply = temp
	return nil
}
