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

// 保存文件的属性
type StorageFileInfo struct {
	Id        int    `db:"id"`
	Account   string `db:"account"`
	Size      int64  `db:"size"`
	Filepath  string `db:"filepath"`
	Name      string `db:"name"`
	Type      string `db:"type"`
	Path      string `db:"path"`
	Smallpath string `db:"smallpath"`
}
