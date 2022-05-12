package register

import (
	"log"
	"registerCenter/config"

	"github.com/jmoiron/sqlx"
)

type instance struct {
	ID          int    `db:"id"`
	IP          string `db:"ip"`
	Port        string `db:"port"`
	ServiceName string `db:"serviceName"`
}

// 服务注册
func ServiceRegister(request RegisterRequestInterface) error {
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo)
	// 错误处理
	if err != nil {
		log.Println(err)
		return err
	}
	defer mysqlClient.Close()

	// 注册，插入到数据库
	_, err = mysqlClient.Exec("INSERT INTO register VALUES(0, ?, ?, ?)", request.IP, request.Port, request.RegisterRequestServiceName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 服务注销
func ServiceLogout(request RegisterRequestInterface) error {
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo)
	// 错误处理
	if err != nil {
		log.Println(err)
		return err
	}
	defer mysqlClient.Close()

	// 注销
	_, err = mysqlClient.Exec("DELETE FROM register WHERE ip = ? AND port = ? AND serviceName = ?", request.IP, request.Port, request.RegisterRequestServiceName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 服务发现
func ServiceSearch(request string) ([]string, error) {
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo)
	// 错误处理
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer mysqlClient.Close()

	// 查找
	ans := []instance{}
	err = mysqlClient.Select(&ans, "SELECT id, ip port FROM register WHERE serviceName = ?", request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 发现的服务实例组
	ret := []string{}
	for index := range ans {
		ret = append(ret, ans[index].IP+ans[index].Port)
	}

	return ret, nil
}

// 服务全发现
func ServiceSearchAll() (map[string][]string, error) {
	// 连接数据库
	mysqlClient, err := sqlx.Connect("mysql", config.MySQLInfo)
	// 错误处理
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer mysqlClient.Close()

	// 查找
	ans := []instance{}
	err = mysqlClient.Select(&ans, "SELECT * FROM register")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 发现的服务实例组
	ret := map[string][]string{}
	for index := range ans {
		ret[ans[index].ServiceName] = append(ret[ans[index].ServiceName], ans[index].IP+ans[index].Port)
	}

	return ret, nil
}
