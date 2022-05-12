package initCode

import (
	"Project/callClient/callDatabase"
	"Project/config"
)

// 建库建表
func InitDatabase() {
	CreateFile()
	callDatabase.CallMySQLInitDatabase(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort)
}
