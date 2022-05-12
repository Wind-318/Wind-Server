package mysqlFunction

import (
	"database/registerCenter"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 选取所有图片信息
func SelectStorageFilesInfo(account, password, ip, port, name string) ([]registerCenter.SelectStorageFilesInfo, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	ans := []registerCenter.SelectStorageFilesInfo{}
	// 选取
	err = mysqlClient.Select(&ans, "SELECT * FROM storage")

	return ans, err
}

// 更新缩略图路径
func UpdateStorageSmallPicture(account, password, ip, port, name, smallPath, path string) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 选取
	_, err = mysqlClient.Exec("UPDATE storage SET smallpath = ? WHERE path = ?", smallPath, path)

	return err
}

// 插入新文件夹
func InsertStorageNewFolder(account, password, ip, port, name, userAccount, filename string) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 选取
	_, err = mysqlClient.Exec("INSERT INTO storagefile VALUES(?, ?, ?)", 0, userAccount, filename)

	return err
}

// 插入新文件
func InsertStorageNewFile(account, password, ip, port, name, userAccount, filepath, filename, types, paths, smallpath string, size int64) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 选取
	_, err = mysqlClient.Exec("INSERT INTO storage VALUES(?, ?, ?, ?, ?, ?, ?, ?)", 0, userAccount, size, filepath, filename, types, paths, smallpath)

	return err
}

// 查询文件夹是否存在
func SelectIsExistFolder(account, password, ip, port, name, userAccount, filename string) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	ret := ""
	// 选取
	err = mysqlClient.Get(&ret, "Select filename FROM storagefile WHERE account = ? AND filename = ?", userAccount, filename)
	// err != nil 则未找到文件夹，可以创建
	if err != nil {
		return nil
	}
	// 否则返回 err
	return errors.New("文件夹已存在！")
}

// 减少可用空间大小
func UpdateStorageUnusedCapacity(account, password, ip, port, name, userAccount string, fileSize int64) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 选取
	_, err = mysqlClient.Exec("Update user Set unusedCapacity = unusedCapacity - ? WHERE account = ?", fileSize, userAccount)

	return err
}
