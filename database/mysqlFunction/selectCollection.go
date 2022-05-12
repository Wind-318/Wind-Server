package mysqlFunction

import (
	"database/registerCenter"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 选取收藏网站
func SelectCollectionWebsite(account, password, ip, port, name string) ([]registerCenter.SelectCollectionDataInterface, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	ans := []registerCenter.SelectCollectionDataInterface{}
	// 选取
	err = mysqlClient.Select(&ans, "SELECT * FROM collections")

	return ans, err
}

// 插入新网站
func InsertCollectionWebsite(account, password, ip, port, name, url, comment, pic string) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	// 选取
	_, err = mysqlClient.Exec("INSERT INTO collections VALUES(?, ?, ?, ?)", 0, url, comment, pic)

	return err
}
