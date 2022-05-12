package data

// Redis 服务名
var RedisServiceName = "RedisService"

// mysql 服务名
var MySQLServiceName = "MySQLService"

// 存储的图片数据
type SelectStorageFilesInfo struct {
	// 账号
	Account string
	// 密码
	Password string
	// 连接地址
	IP string
	// 连接端口
	Port string
	// 数据库名称
	Name string
	// ID
	Id int
	// 用户账号
	UserAccount string
	// 大小
	Size int64
	// 文件夹
	Filepath string
	// 文件名
	FileName string
	// 类型
	Type string
	// 路径
	Path string
	// 缩略图路径
	Smallpath string
}

// 收藏网站属性
type SelectCollectionDataInterface struct {
	// 账号
	Account string
	// 密码
	Password string
	// 连接地址
	IP string
	// 连接端口
	Port string
	// 数据库名称
	Name        string
	ID          int
	Url         string
	Comment     string
	Picurl      string
	Description string
}

// 动漫属性
type SelectAnimeDataInterface struct {
	// 账号
	Account string
	// 密码
	Password string
	// 连接地址
	IP string
	// 连接端口
	Port string
	// 数据库名称
	Name string
	// 动漫名
	AnimeName string
	// 链接
	Url string
	// 年份
	Year int
	// 简介
	Description string
	// 播放来源
	Source []string
	// 播放地址
	Urls []string
	// 封面地址
	Picurl string
	// 是否是当季番
	IsNew int
}

// 数据
type SelectDataInterface struct {
	// 账号
	Account string
	// 密码
	Password string
	// 连接地址
	IP string
	// 连接端口
	Port string
	// 数据库名称
	Name string
	// 用户账号
	UserAccount string
	// 用户名
	UserName string
	// 用户密码
	UserPassword string
	// 等级
	Level int
	// 权限
	Authority int
	// 头像路径
	Picture string
	// 积分
	Score int
	// 空间大小
	Capacity int64
	// 可用空间
	UnusedCapacity int64
	// 文件夹名称
	Filename string
	// 要更新的数据
	Space int64
}

// 数据
type RedisDataInterface struct {
	// 连接地址
	IP string
	// 连接端口
	Port string
	// 操作
	Operate string
	// 参数列表
	Parameters []interface{}
	// 进行的操作
	Operation string
}
