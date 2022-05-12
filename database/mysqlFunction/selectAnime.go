package mysqlFunction

import (
	"database/registerCenter"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 根据动漫名查询信息
func SelectAnimeInfoByName(account, password, ip, port, name, animeName string) (registerCenter.SelectAnimeDataInterface, error) {
	ans := registerCenter.SelectAnimeDataInterface{}
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return ans, err
	}
	defer mysqlClient.Close()

	// 选取
	err = mysqlClient.Get(&ans, "SELECT name, url, year, description, picurl FROM bangumi WHERE name = ?", animeName)

	return ans, err
}

// 根据动漫名查询播放来源
func SelectAnimeSourceByName(account, password, ip, port, name, animeName string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	// 选取
	ans := []string{}
	err = mysqlClient.Select(&ans, "SELECT source FROM animesource WHERE anime = ?", animeName)

	return ans, err
}

// 根据动漫名查询播放地址
func SelectAnimeUrlByName(account, password, ip, port, name, animeName string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	// 选取
	ans := []string{}
	err = mysqlClient.Select(&ans, "SELECT urls FROM animesource WHERE anime = ?", animeName)

	return ans, err
}

// 获取所有动漫名
func SelectAnimeName(account, password, ip, port, name string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	// 选取
	ans := []string{}
	err = mysqlClient.Select(&ans, "SELECT name FROM bangumi")

	return ans, err
}

// 选取新番
func SelectNewAnime(account, password, ip, port, name string) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	// 选取
	ans := []string{}
	err = mysqlClient.Select(&ans, "SELECT name FROM bangumi WHERE isNew = 1")

	return ans, err
}

// 根据年份选取动漫
func SelectAnimeByYear(account, password, ip, port, name string, year int) ([]string, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return nil, err
	}
	defer mysqlClient.Close()

	// 选取
	ans := []string{}
	if year > 2000 {
		err = mysqlClient.Select(&ans, "SELECT name FROM bangumi WHERE year = ?", year)
	} else {
		err = mysqlClient.Select(&ans, "SELECT name FROM bangumi WHERE year <= ?", year)
	}

	return ans, err
}

// 查看某来源动漫是否存在
func SelectAnimeCountByNameAndSource(account, password, ip, port, name, animeName, source string) (bool, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return true, err
	}
	defer mysqlClient.Close()

	// 选取
	ans := false
	cnt := 0
	err = mysqlClient.Get(&cnt, "SELECT count(*) FROM animesource WHERE anime = ? AND source = ?", animeName, source)
	if err != nil {
		return ans, err
	}
	if cnt > 0 {
		ans = true
	}
	// 若存在，设为新番
	if ans {
		_, err = mysqlClient.Exec("UPDATE bangumi SET isNew = 1 WHERE name = ?", animeName)
	}

	return ans, err
}

// 根据名称查看动漫是否存在
func SelectAnimeCountByName(account, password, ip, port, name, animeName string) (bool, error) {
	mysqlClient, err := sqlx.Connect("mysql", account+":"+password+"@tcp("+ip+port+")/"+name)
	if err != nil {
		return true, err
	}
	defer mysqlClient.Close()

	// 选取
	ans := 0
	err = mysqlClient.Get(&ans, "SELECT count(*) FROM bangumi WHERE name = ?", animeName)
	if ans > 0 {
		return true, nil
	}

	return false, err
}
