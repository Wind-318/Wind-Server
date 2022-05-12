package registerCenter

import "mime/multipart"

// 服务名
var StorageServiceName = "StorageService"

// 存储服务接口
type StorageServiceInterface interface {
	// 删除文件
	DeleteFiles(request StoragePictureData, reply *interface{}) error
	// 存储文件
	StorageFiles(request StoragePictureData, reply *interface{}) error
	// 获取用户存储的图片合集
	GetUserStoragePicture(request StoragePictureData, reply *StoragePicture) error
	// 获取页数
	GetUserStoragePicturePage(request StoragePictureData, reply *int) error
	// 查看文件夹个数
	GetUserFileNums(request StoragePictureData, reply *[]string) error
	// 获取网站链接
	GetWebs(request StoragePictureData, reply *map[string]interface{}) error
	// 增加收藏网站
	PutWebs(request StoragePictureData, reply *map[string]interface{}) error
	// 批量下载
	DownloadFiles(request StoragePictureData, reply *map[string]interface{}) error
}

// 图片数据
type StoragePictureData struct {
	// cookie
	Cookie string
	// url
	Url string
	// comment
	Comment string
	// picurl
	Picurl string
	// 用户名
	UserName string
	// 文件夹名
	FolderName string
	// 文件数据
	Files *multipart.Form
	// 大小
	Size int64
	// 偏移量
	Offset int
	// 每次选择的量
	OnceChoose int
	// 每页数量
	PageNum int
}

// 存储图片的合集
type StoragePicture struct {
	// 图片链接
	PicPath []string
	// 缩略图链接
	SmallPicPath []string
	// 数量
	Number int
}
