package initCode

import "os"

func CreateFile() {
	// 新建文件夹
	os.MkdirAll("./bbsFile", 0644)
	os.MkdirAll("./userFile", 0644)
	os.MkdirAll("./picture/collections", 0644)
}
