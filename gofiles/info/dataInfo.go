package info

// 动漫属性
type AnimeInfo struct {
	// 名称
	Name string `db:"name"`
	// bangumi 链接
	Url string `db:"url"`
	// 年份
	Year int `db:"year"`
	// 简介
	Description string `db:"description"`
	// 播放来源
	Source []string `db:"source"`
	// 播放地址
	Urls []string `db:"urls"`
	// 封面地址
	Picurl string `db:"picurl"`
}

// 收藏网站属性
type CollectionData struct {
	ID          int    `db:"id"`
	Url         string `db:"url"`
	Description string `db:"description"`
	Picurl      string `db:"picurl"`
}

type StorageFileData struct {
	// ID
	Id int `db:"id"`
	// 用户账号
	UserAccount string `db:"account"`
	// 大小
	Size int64 `db:"size"`
	// 文件夹
	Filepath string `db:"filepath"`
	// 文件名
	FileName string `db:"name"`
	// 类型
	Type string `db:"type"`
	// 路径
	Path string `db:"path"`
	// 缩略图路径
	Smallpath string `db:"smallpath"`
}
