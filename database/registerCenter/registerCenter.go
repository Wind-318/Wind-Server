package registerCenter

// 服务名
var RedisServiceName = "RedisService"

// 服务名
var MySQLServiceName = "MySQLService"

// RedisServiceInterface Redis 服务接口
type RedisServiceInterface interface {
	// 对 Redis 数据库进行操作
	OperateRedisDatabase(request RedisDataInterface, reply *interface{}) error
	// 获取 Redis 所有操作
	GetOperatorsForRedis(request interface{}, reply *[]string) error
}

// MySQLServiceInterface Redis 服务接口
type MySQLServiceInterface interface {
	// 初始化数据库
	InitDatabase(request SelectDataInterface, reply *interface{}) error
	// 选取订阅用户名单
	SelectUsersAccount(request SelectDataInterface, reply *[]string) error
	// 获取剩余可用空间
	SelectRemainingSpace(request SelectDataInterface, reply *int64) error
	// 获取剩余可用空间
	UpdateRemainingSpace(request SelectDataInterface, reply *interface{}) error
	// 获取账号
	SelectUserAccountByName(request SelectDataInterface, reply *string) error
	// 根据账号与文件名选取存储封面图片
	SelectStoragePictureByFilename(request SelectDataInterface, reply *[]string) error
	// 根据账号与文件名选取存储缩略图片
	SelectStorageSmallPictureByFilename(request SelectDataInterface, reply *[]string) error
	// 根据账号和文件夹获取文件总数
	SelectStorageFilesNumber(request SelectDataInterface, reply *int) error
	// 根据账号获取文件夹个数
	SelectStorageUserFilesNumber(request SelectDataInterface, reply *[]string) error
	// 根据动漫名查询信息
	SelectAnimeInfoByName(request SelectAnimeDataInterface, reply *SelectAnimeDataInterface) error
	// 根据动漫名查询播放来源
	SelectAnimeSourceByName(request SelectAnimeDataInterface, reply *[]string) error
	// 根据动漫名查询播放地址
	SelectAnimeUrlByName(request SelectAnimeDataInterface, reply *[]string) error
	// 获取所有动漫名
	SelectAnimeName(request SelectAnimeDataInterface, reply *[]string) error
	// 获取新番
	SelectNewAnime(request SelectAnimeDataInterface, reply *[]string) error
	// 根据年份选取动漫
	SelectAnimeByYear(request SelectAnimeDataInterface, reply *[]string) error
	// 选取收藏网站
	SelectCollectionWebsite(request SelectCollectionDataInterface, reply *[]SelectCollectionDataInterface) error
	// 插入新网站
	InsertCollectionWebsite(request SelectCollectionDataInterface, reply *interface{}) error
	// 更新本季新番取消最新状态
	UpdateAnimeSetIsNew0(request SelectAnimeDataInterface, reply *interface{}) error
	// 查看某来源动漫是否已存在，已存在更新为新番
	SelectAnimeCountByNameAndSource(request SelectAnimeDataInterface, reply *bool) error
	// 根据名称查看动漫是否已存在
	SelectAnimeCountByName(request SelectAnimeDataInterface, reply *bool) error
	// 更新动漫资料
	UpdateAnimeInfo(request SelectAnimeDataInterface, reply *interface{}) error
	// 插入新动漫到 bangumi
	InsertAnimeForBangumi(request SelectAnimeDataInterface, reply *interface{}) error
	// 插入新动漫到 source
	InsertAnimeForSource(request SelectAnimeDataInterface, reply *interface{}) error
	// 根据账号选取用户名
	SelectUserNameByAccount(request SelectDataInterface, reply *string) error
	// 选取所有用户名称
	SelectUserNames(request SelectDataInterface, reply *[]string) error
	// 选取所有图片信息
	SelectStorageFilesInfo(request SelectStorageFilesInfo, reply *[]SelectStorageFilesInfo) error
	// 更新缩略图路径
	UpdateStorageSmallPicture(request SelectStorageFilesInfo, reply *[]SelectStorageFilesInfo) error
	// 插入新用户
	InsertNewUser(request SelectDataInterface, reply *interface{}) error
	// 选取 ID
	SelectUserIDByAccount(request SelectDataInterface, reply *int) error
	// 选取用户密码
	SelectUserPasswordByAccount(request SelectDataInterface, reply *string) error
	// 修改密码
	UpdateUserPassword(request SelectDataInterface, reply *interface{}) error
	// 根据 ID 获取账号
	SelectUserAccountByID(request SelectDataInterface, reply *string) error
	// 更新积分数
	UpdateUserScore(request SelectDataInterface, reply *interface{}) error
	// 插入新文件夹
	InsertStorageNewFolder(request SelectStorageFilesInfo, reply *interface{}) error
	// 插入新文件
	InsertStorageNewFile(request SelectStorageFilesInfo, reply *interface{}) error
	// 查询文件夹是否存在
	SelectIsExistFolder(request SelectStorageFilesInfo, reply *interface{}) error
	// 减少可用空间大小
	UpdateStorageUnusedCapacity(request SelectStorageFilesInfo, reply *interface{}) error
}

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
