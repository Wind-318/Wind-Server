package initCode

import (
	"Project/gofiles/config"
	"bufio"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 建库建表
func InitDatabase() {
	// 连接
	conn := sqlx.MustConnect("mysql", config.MySQLInit)
	defer conn.Close()

	// 如果存在同名表则丢弃
	drops, _ := os.Open("./sql/drop.sql")
	defer drops.Close()
	buf := bufio.NewScanner(drops)
	for buf.Scan() {
		conn.Exec(buf.Text())
	}

	// 读取各表建表指令
	files, _ := ioutil.ReadDir("./sql")
	for _, file := range files {
		if file.Name() == "drop.sql" {
			continue
		}
		bytes, _ := ioutil.ReadFile("./sql/" + file.Name())
		// 执行
		conn.Exec(string(bytes))
	}
}
