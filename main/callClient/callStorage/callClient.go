package callStorage

import (
	"Project/callClient/data"
	"Project/serviceSearch"
	"net/rpc"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 操作客户端
type StorageClient struct {
	*rpc.Client
}

// 拨号
func dialStorageService(network, address string) (*StorageClient, error) {
	client, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &StorageClient{Client: client}, nil
}

// 删除文件
func CallStorageDeleteFiles() error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.StorageServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialStorageService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.StorageServiceName+".DeleteFiles", "", &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 存储文件
func CallStorageStorageFiles(ctx *gin.Context, size int64) error {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.StorageServiceName)
	if err != nil {
		return err
	}
	// 拨号
	client, err := dialStorageService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return err
	}

	// 检查数据
	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return err
	}

	obj := data.StoragePictureData{
		Cookie: cookie,
		Size:   size,
	}

	var result interface{}
	// 调用
	err = client.Client.Call(data.StorageServiceName+".StorageFiles", obj, &result)
	// 错误处理
	if err != nil {
		return err
	}

	return nil
}

// 获取用户存储的图片合集
func CallStorageGetUserStoragePicture(ctx *gin.Context) ([]string, []string, int, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.StorageServiceName)
	if err != nil {
		return nil, nil, 0, err
	}
	// 拨号
	client, err := dialStorageService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, nil, 0, err
	}

	num, _ := strconv.Atoi(ctx.PostForm("num"))
	choose, _ := strconv.Atoi(ctx.PostForm("onceChoose"))
	obj := data.StoragePictureData{
		UserName:   ctx.PostForm("name"),
		Offset:     num,
		FolderName: ctx.PostForm("file"),
		OnceChoose: choose,
	}

	var result data.StoragePicture
	// 调用
	err = client.Client.Call(data.StorageServiceName+".GetUserStoragePicture", obj, &result)
	// 错误处理
	if err != nil {
		return nil, nil, 0, err
	}

	return result.PicPath, result.SmallPicPath, result.Number, nil
}

// 获取页数
func CallStorageGetUserStoragePicturePage(ctx *gin.Context) (int, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.StorageServiceName)
	if err != nil {
		return 0, err
	}
	// 拨号
	client, err := dialStorageService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return 0, err
	}

	pagenum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
	obj := data.StoragePictureData{
		UserName:   ctx.PostForm("name"),
		FolderName: ctx.PostForm("file"),
		PageNum:    pagenum,
	}

	var result int
	// 调用
	err = client.Client.Call(data.StorageServiceName+".GetUserStoragePicturePage", obj, &result)
	// 错误处理
	if err != nil {
		return 0, err
	}

	return result, nil
}

// 查看文件夹个数
func CallStorageGetUserFileNums(ctx *gin.Context) ([]string, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.StorageServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialStorageService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	obj := data.StoragePictureData{
		UserName: ctx.PostForm("name"),
	}

	var result []string
	// 调用
	err = client.Client.Call(data.StorageServiceName+".GetUserFileNums", obj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 获取网站链接
func CallStorageGetWebs(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.StorageServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialStorageService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return nil, err
	}
	obj := data.StoragePictureData{
		Cookie: cookie,
	}
	var result map[string]interface{}
	// 调用
	err = client.Client.Call(data.StorageServiceName+".GetWebs", obj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 增加收藏网站
func CallStoragePutWebs(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.StorageServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialStorageService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return nil, err
	}
	obj := data.StoragePictureData{
		Cookie:  cookie,
		Url:     ctx.PostForm("url"),
		Comment: ctx.PostForm("comment"),
		Picurl:  ctx.PostForm("picurl"),
	}
	var result map[string]interface{}
	// 调用
	err = client.Client.Call(data.StorageServiceName+".PutWebs", obj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 批量下载
func CallStorageDownloadFiles(ctx *gin.Context) (map[string]interface{}, error) {
	// 获取节点
	remoteAddr, err := serviceSearch.GetNode(data.StorageServiceName)
	if err != nil {
		return nil, err
	}
	// 拨号
	client, err := dialStorageService("tcp", remoteAddr)
	// 错误处理
	if err != nil {
		return nil, err
	}

	cookie, err := ctx.Cookie("cookie")
	if err != nil {
		return nil, err
	}
	obj := data.StoragePictureData{
		Cookie:     cookie,
		UserName:   ctx.PostForm("name"),
		FolderName: ctx.PostForm("texts"),
	}
	var result map[string]interface{}
	// 调用
	err = client.Client.Call(data.StorageServiceName+".DownloadFiles", obj, &result)
	// 错误处理
	if err != nil {
		return nil, err
	}

	return result, nil
}
