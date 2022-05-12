package function

import (
	"storage/callClient/callDatabase"
	"storage/config"
	"storage/registerCenter"
)

// 获取指定用户上传的图片
func GetUserStoragePicture(request registerCenter.StoragePictureData) (registerCenter.StoragePicture, error) {
	// 返回值
	ans := registerCenter.StoragePicture{}
	// 获取邮箱
	name := request.UserName
	email, err := callDatabase.CallMySQLSelectUserAccountByName(name, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user")
	if err != nil {
		return ans, err
	}
	// 获取偏移量
	moveNum := request.Offset
	// 获取分类
	folderName := request.FolderName
	// 一次选取的数量
	onceChoose := request.OnceChoose
	// 查询
	ans.PicPath, err = callDatabase.CallMySQLSelectStoragePictureByFilename(email, folderName, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, config.MySQLName)
	if err != nil {
		return ans, err
	}
	ans.SmallPicPath, err = callDatabase.CallMySQLSelectStoragePictureByFilename(email, folderName, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, config.MySQLName)
	if err != nil {
		return ans, err
	}

	// 根据偏移量选取显示的图片
	if moveNum+onceChoose >= len(ans.PicPath) {
		ans.PicPath = ans.PicPath[moveNum:]
		ans.SmallPicPath = ans.SmallPicPath[moveNum:]
	} else {
		ans.PicPath = ans.PicPath[moveNum : moveNum+onceChoose]
		ans.SmallPicPath = ans.SmallPicPath[moveNum : moveNum+onceChoose]
	}
	// 获取显示数量
	ans.Number = len(ans.PicPath)

	return ans, nil
}

// 获取页数
func GetUserStoragePicturePage(request registerCenter.StoragePictureData) (int, error) {
	// 名称
	name := request.UserName
	// 分类
	folderName := request.FolderName
	// 每页数量
	pageNum := request.PageNum

	// 账号
	email, err := callDatabase.CallMySQLSelectUserAccountByName(name, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user")
	if err != nil {
		return 0, nil
	}
	// 总数
	sumNum, err := callDatabase.CallMySQLSelectStorageFilesNumber(email, folderName, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, config.MySQLName)
	if err != nil {
		return 0, nil
	}
	// 计算页数
	page := 1
	if sumNum%pageNum == 0 {
		page = sumNum / pageNum
	} else {
		page = sumNum/pageNum + 1
	}
	if page == 0 {
		page = 1
	}

	return page, nil
}

// 查看文件夹个数
func GetUserFileNums(request registerCenter.StoragePictureData) ([]string, error) {
	// 名称
	name := request.UserName
	// 账号
	email, err := callDatabase.CallMySQLSelectUserAccountByName(name, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "user")
	if err != nil {
		return nil, err
	}
	// 获取该用户文件夹个数
	files, err := callDatabase.CallMySQLSelectStorageUserFilesNumber(email, config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, config.MySQLName)

	return files, err
}
