package mainPackage

import (
	"database/mysqlFunction"
	"database/registerCenter"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库服务
type MySQLService struct{}

// 初始化数据库
func (m *MySQLService) InitDatabase(request registerCenter.SelectDataInterface, reply *interface{}) error {
	return mysqlFunction.InitDatabase(request.Account, request.Password, request.IP, request.Port)
}

// 查询订阅股票信息用户
func (m *MySQLService) SelectUsersAccount(request registerCenter.SelectDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectUsersAccount(request.Account, request.Password, request.IP, request.Port, request.Name)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 获取剩余可用空间
func (m *MySQLService) SelectRemainingSpace(request registerCenter.SelectDataInterface, reply *int64) error {
	temp, err := mysqlFunction.SelectRemainingSpace(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 更新数据库剩余可用空间
func (m *MySQLService) UpdateRemainingSpace(request registerCenter.SelectDataInterface, reply *int64) error {
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", request.Account+":"+request.Password+"@tcp("+request.IP+request.Port+")/"+request.Name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 更新数据
	_, err = mysqlClient.Exec("UPDATE user SET unusedCapacity = ? WHERE account = ?", request.Space, request.UserAccount)

	return err
}

// 获取账号
func (m *MySQLService) SelectUserAccountByName(request registerCenter.SelectDataInterface, reply *string) error {
	temp, err := mysqlFunction.SelectUserAccountByName(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserName)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据文件名选取存储封面图片
func (m *MySQLService) SelectStoragePictureByFilename(request registerCenter.SelectDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectStoragePictureByFilename(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.Filename)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据文件名选取存储缩略图片
func (m *MySQLService) SelectStorageSmallPictureByFilename(request registerCenter.SelectDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectStorageSmallPictureByFilename(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.Filename)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据账号和文件夹获取总数
func (m *MySQLService) SelectStorageFilesNumber(request registerCenter.SelectDataInterface, reply *int) error {
	temp, err := mysqlFunction.SelectStorageFilesNumber(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.Filename)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据账号获取文件夹个数
func (m *MySQLService) SelectStorageUserFilesNumber(request registerCenter.SelectDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectStorageUserFilesNumber(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据动漫名查询信息
func (m *MySQLService) SelectAnimeInfoByName(request registerCenter.SelectAnimeDataInterface, reply *registerCenter.SelectAnimeDataInterface) error {
	temp, err := mysqlFunction.SelectAnimeInfoByName(request.Account, request.Password, request.IP, request.Port, request.Name, request.AnimeName)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据动漫名查询播放来源
func (m *MySQLService) SelectAnimeSourceByName(request registerCenter.SelectAnimeDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectAnimeSourceByName(request.Account, request.Password, request.IP, request.Port, request.Name, request.AnimeName)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据动漫名查询播放地址
func (m *MySQLService) SelectAnimeUrlByName(request registerCenter.SelectAnimeDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectAnimeUrlByName(request.Account, request.Password, request.IP, request.Port, request.Name, request.AnimeName)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 获取所有动漫名
func (m *MySQLService) SelectAnimeName(request registerCenter.SelectAnimeDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectAnimeName(request.Account, request.Password, request.IP, request.Port, request.Name)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 获取新番
func (m *MySQLService) SelectNewAnime(request registerCenter.SelectAnimeDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectNewAnime(request.Account, request.Password, request.IP, request.Port, request.Name)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据年份选取动漫
func (m *MySQLService) SelectAnimeByYear(request registerCenter.SelectAnimeDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectAnimeByYear(request.Account, request.Password, request.IP, request.Port, request.Name, request.Year)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 选取收藏网站
func (m *MySQLService) SelectCollectionWebsite(request registerCenter.SelectCollectionDataInterface, reply *[]registerCenter.SelectCollectionDataInterface) error {
	temp, err := mysqlFunction.SelectCollectionWebsite(request.Account, request.Password, request.IP, request.Port, request.Name)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 插入新网站
func (m *MySQLService) InsertCollectionWebsite(request registerCenter.SelectCollectionDataInterface, reply *interface{}) error {
	err := mysqlFunction.InsertCollectionWebsite(request.Account, request.Password, request.IP, request.Port, request.Name, request.Url, request.Comment, request.Picurl)

	return err
}

// 更新本季新番取消最新状态
func (m *MySQLService) UpdateAnimeSetIsNew0(request registerCenter.SelectAnimeDataInterface, reply *interface{}) error {
	err := mysqlFunction.UpdateAnimeSetIsNew0(request.Account, request.Password, request.IP, request.Port, request.Name)

	return err
}

// 查看某来源动漫是否已存在
func (m *MySQLService) SelectAnimeCountByNameAndSource(request registerCenter.SelectAnimeDataInterface, reply *bool) error {
	temp, err := mysqlFunction.SelectAnimeCountByNameAndSource(request.Account, request.Password, request.IP, request.Port, request.Name, request.AnimeName, request.Source[0])
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 根据名称查看动漫是否已存在
func (m *MySQLService) SelectAnimeCountByName(request registerCenter.SelectAnimeDataInterface, reply *bool) error {
	temp, err := mysqlFunction.SelectAnimeCountByName(request.Account, request.Password, request.IP, request.Port, request.Name, request.AnimeName)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 更新动漫资料
func (m *MySQLService) UpdateAnimeInfo(request registerCenter.SelectAnimeDataInterface, reply *interface{}) error {
	err := mysqlFunction.UpdateAnimeInfo(request.Account, request.Password, request.IP, request.Port, request.Name, request.Url, request.Description, request.Picurl, request.AnimeName, request.IsNew)

	return err
}

// 插入新动漫到 bangumi
func (m *MySQLService) InsertAnimeForBangumi(request registerCenter.SelectAnimeDataInterface, reply *interface{}) error {
	err := mysqlFunction.InsertAnimeForBangumi(request.Account, request.Password, request.IP, request.Port, request.Name, request.AnimeName, request.Url, request.Description, request.Picurl, request.Year, request.IsNew)

	return err
}

// 插入新动漫到 source
func (m *MySQLService) InsertAnimeForSource(request registerCenter.SelectAnimeDataInterface, reply *interface{}) error {
	err := mysqlFunction.InsertAnimeForSource(request.Account, request.Password, request.IP, request.Port, request.Name, request.AnimeName, request.Source[0], request.Url)

	return err
}

// 根据账号选取用户名
func (m *MySQLService) SelectUserNameByAccount(request registerCenter.SelectDataInterface, reply *string) error {
	temp, err := mysqlFunction.SelectUserNameByAccount(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 选取所有用户名称
func (m *MySQLService) SelectUserNames(request registerCenter.SelectDataInterface, reply *[]string) error {
	temp, err := mysqlFunction.SelectUserNames(request.Account, request.Password, request.IP, request.Port, request.Name)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 选取所有图片信息
func (m *MySQLService) SelectStorageFilesInfo(request registerCenter.SelectStorageFilesInfo, reply *[]registerCenter.SelectStorageFilesInfo) error {
	temp, err := mysqlFunction.SelectStorageFilesInfo(request.Account, request.Password, request.IP, request.Port, request.Name)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 更新缩略图路径
func (m *MySQLService) UpdateStorageSmallPicture(request registerCenter.SelectStorageFilesInfo, reply *interface{}) error {
	err := mysqlFunction.UpdateStorageSmallPicture(request.Account, request.Password, request.IP, request.Port, request.Name, request.Smallpath, request.Path)

	return err
}

// 插入新用户
func (m *MySQLService) InsertNewUser(request registerCenter.SelectDataInterface, reply *interface{}) error {
	err := mysqlFunction.InsertNewUser(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.UserPassword, request.UserName, request.Picture)

	return err
}

// 选取 ID
func (m *MySQLService) SelectUserIDByAccount(request registerCenter.SelectDataInterface, reply *int) error {
	temp, err := mysqlFunction.SelectUserIDByAccount(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 选取用户密码
func (m *MySQLService) SelectUserPasswordByAccount(request registerCenter.SelectDataInterface, reply *string) error {
	temp, err := mysqlFunction.SelectUserPasswordByAccount(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 修改密码
func (m *MySQLService) UpdateUserPassword(request registerCenter.SelectDataInterface, reply *interface{}) error {
	err := mysqlFunction.UpdateUserPassword(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserPassword, request.UserAccount)

	return err
}

// 根据 ID 获取账号
func (m *MySQLService) SelectUserAccountByID(request registerCenter.SelectDataInterface, reply *string) error {
	temp, err := mysqlFunction.SelectUserAccountByID(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserName)
	if err != nil {
		return err
	}
	*reply = temp

	return nil
}

// 更新积分数
func (m *MySQLService) UpdateUserScore(request registerCenter.SelectDataInterface, reply *interface{}) error {
	err := mysqlFunction.UpdateUserScore(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.Score)

	return err
}

// 插入新文件夹
func (m *MySQLService) InsertStorageNewFolder(request registerCenter.SelectStorageFilesInfo, reply *interface{}) error {
	err := mysqlFunction.InsertStorageNewFolder(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.FileName)

	return err
}

// 插入新文件
func (m *MySQLService) InsertStorageNewFile(request registerCenter.SelectStorageFilesInfo, reply *interface{}) error {
	err := mysqlFunction.InsertStorageNewFile(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.Filepath, request.FileName, request.Type, request.Path, request.Smallpath, request.Size)

	return err
}

// 查询文件夹是否存在
func (m *MySQLService) SelectIsExistFolder(request registerCenter.SelectStorageFilesInfo, reply *interface{}) error {
	err := mysqlFunction.SelectIsExistFolder(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.FileName)

	return err
}

// 减少可用空间大小
func (m *MySQLService) UpdateStorageUnusedCapacity(request registerCenter.SelectStorageFilesInfo, reply *interface{}) error {
	err := mysqlFunction.UpdateStorageUnusedCapacity(request.Account, request.Password, request.IP, request.Port, request.Name, request.UserAccount, request.Size)

	return err
}
