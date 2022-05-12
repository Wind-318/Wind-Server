package mysqlFunction

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 更新本季新番取消最新状态
func UpdateAnimeSetIsNew0(account, password, ip, port, name string) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	_, err = mysqlClient.Exec("UPDATE bangumi SET isNew = 0 WHERE isNew = 1")

	return err
}

// 更新动漫资料
func UpdateAnimeInfo(account, password, ip, port, name, url, description, picurl, animeName string, isNew int) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	_, err = mysqlClient.Exec("UPDATE bangumi SET url = ?, description = ?, picurl = ?, isNew = ? WHERE name = ?", url, description, picurl, isNew, animeName)

	return err
}

// 插入新动漫到 bangumi
func InsertAnimeForBangumi(account, password, ip, port, name, animeName, url, description, picurl string, year, isNew int) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	_, err = mysqlClient.Exec("INSERT INTO bangumi VALUES(?, ?, ?, ?, ?, ?, ?)", 0, animeName, url, year, description, picurl, isNew)

	return err
}

// 插入新动漫到 source
func InsertAnimeForSource(account, password, ip, port, name, animeName, sourcee, urls string) error {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return err
	}
	defer mysqlClient.Close()

	_, err = mysqlClient.Exec("INSERT INTO animesource VALUES(?, ?, ?, ?)", 0, animeName, sourcee, urls)

	return err
}
