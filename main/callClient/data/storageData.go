package data

import "mime/multipart"

// 服务名
var StorageServiceName = "StorageService"

// 存储图片的合集
type StoragePicture struct {
	// 图片链接
	PicPath []string
	// 缩略图链接
	SmallPicPath []string
	// 数量
	Number int
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
