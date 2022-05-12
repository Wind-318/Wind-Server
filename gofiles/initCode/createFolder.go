package initCode

import "os"

func CreateFolder() {
	// 创建文件夹
	os.MkdirAll("./picture/anime/yhdm", 0644)
	os.MkdirAll("./bbsFile", 0644)
	os.MkdirAll("./userFile", 0644)
	os.MkdirAll("./picture/collections", 0644)
}
