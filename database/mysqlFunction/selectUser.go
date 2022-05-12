package mysqlFunction

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 查询订阅股票信息用户
func SelectUsersAccount(account, password, ip, port, name string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	ans := []string{}
	// 选取
	err = mysqlClient.Select(&ans, "SELECT account FROM subscribe WHERE stock = ?", 1)

	return ans, err
}

// 获取剩余可用空间
func SelectRemainingSpace(account, password, ip, port, name, userAccount string) (int64, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return 0, err
	}
	defer mysqlClient.Close()

	var ans int64
	// 选取
	err = mysqlClient.Get(&ans, "SELECT unusedCapacity FROM user WHERE account = ?", userAccount)

	return ans, err
}

// 获取用户名选取账号
func SelectUserAccountByName(account, password, ip, port, name, userName string) (string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return "", err
	}
	defer mysqlClient.Close()

	var ans string
	// 选取
	err = mysqlClient.Get(&ans, "SELECT account FROM user WHERE username = ?", userName)

	return ans, err
}

// 根据账号选取用户名
func SelectUserNameByAccount(account, password, ip, port, name, userAccount string) (string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return "", err
	}
	defer mysqlClient.Close()

	var ans string
	// 选取
	err = mysqlClient.Get(&ans, "SELECT username FROM user WHERE account = ?", userAccount)

	return ans, err
}

// 根据文件名选取存储封面图片
func SelectStoragePictureByFilename(account, password, ip, port, name, userAccount, filename string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	var ans []string
	// 选取
	err = mysqlClient.Select(&ans, "SELECT path FROM storage WHERE account = ? AND filepath = ?", userAccount, filename)

	return ans, err
}

// 根据文件名选取存储缩略图片
func SelectStorageSmallPictureByFilename(account, password, ip, port, name, userAccount, filename string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	var ans []string
	// 选取
	err = mysqlClient.Select(&ans, "SELECT smallpath FROM storage WHERE account = ? AND filepath = ?", userAccount, filename)

	return ans, err
}

// 根据账号和文件夹获取总数
func SelectStorageFilesNumber(account, password, ip, port, name, userAccount, filename string) (int, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return 0, err
	}
	defer mysqlClient.Close()

	var ans int
	// 选取
	err = mysqlClient.Get(&ans, "SELECT count(*) FROM storage WHERE account = ? AND filepath = ?", userAccount, filename)

	return ans, err
}

// 根据账号获取文件夹个数
func SelectStorageUserFilesNumber(account, password, ip, port, name, userAccount string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	var ans []string
	// 选取
	err = mysqlClient.Select(&ans, "SELECT DISTINCT filename FROM storagefile WHERE account = ?", userAccount)

	return ans, err
}

// 选取所有用户名称
func SelectUserNames(account, password, ip, port, name string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	var ans []string
	// 选取
	err = mysqlClient.Select(&ans, "SELECT username FROM user")

	return ans, err
}

// 插入新用户
func InsertNewUser(account, password, ip, port, name, userAccount, userPassword, userName, pic string) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 选取
	_, err = mysqlClient.Exec("INSERT INTO user VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 0, userAccount, userPassword, userName, 0, 0, pic, 0, 0, 0)

	return err
}

// 选取 ID
func SelectUserIDByAccount(account, password, ip, port, name, userAccount string) (int, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return 0, err
	}
	defer mysqlClient.Close()

	var ans int
	// 选取
	err = mysqlClient.Get(&ans, "SELECT id FROM user WHERE account = ?", userAccount)

	return ans, err
}

// 选取用户密码
func SelectUserPasswordByAccount(account, password, ip, port, name, userAccount string) (string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return "", err
	}
	defer mysqlClient.Close()

	var ans string
	// 选取
	err = mysqlClient.Get(&ans, "SELECT password FROM user WHERE account = ?", userAccount)

	return ans, err
}

// 修改密码
func UpdateUserPassword(account, password, ip, port, name, userPassword, userAccount string) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 选取
	_, err = mysqlClient.Exec("UPDATE user SET password = ? WHERE account = ?", userPassword, userAccount)

	return err
}

// 根据 ID 获取账号
func SelectUserAccountByID(account, password, ip, port, name, id string) (string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return "", err
	}
	defer mysqlClient.Close()

	// 选取
	ans := ""
	err = mysqlClient.Get(&ans, "SELECT authoremail FROM blog WHERE id = ?", id)

	return ans, err
}

// 根据 ID 获取账号
func UpdateUserScore(account, password, ip, port, name, userAccount string, score int) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	_, err = mysqlClient.Exec("UPDATE user SET score = score + ? WHERE account = ?", score, userAccount)

	return err
}
