package register

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 初始化数据库
func InitDatabase(account, password, ip, port string) error {
	mysqlInfo := account + ":" + password + "@tcp(" + ip + port + ")/"
	// 连接
	conn, err := sqlx.Connect("mysql", mysqlInfo+"mysql")
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	// 创建数据库
	drops, _ := os.Open("./sql/createDatabase.sql")
	defer drops.Close()
	buf := bufio.NewScanner(drops)
	for buf.Scan() {
		conn.Exec(buf.Text())
	}

	// 读取各表建表指令
	files, err := ioutil.ReadDir("./sql")
	if err != nil {
		log.Println(err)
		return err
	}
	for _, file := range files {
		if file.Name() == "createDatabase.sql" {
			continue
		}

		// 连接
		tempMySQLClient, err := sqlx.Connect("mysql", mysqlInfo+file.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		defer tempMySQLClient.Close()

		// 读取数据库各表内容
		database, err := ioutil.ReadDir("./sql/" + file.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		for index := range database {
			// 读取文件内容
			bytes, _ := ioutil.ReadFile("./sql/" + file.Name() + "/" + database[index].Name())
			// 执行
			tempMySQLClient.Exec(string(bytes))
		}
	}

	return nil
}
