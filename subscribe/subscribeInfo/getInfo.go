package subscribeInfo

import (
	"subscribe/callClient/callDatabase"
	"subscribe/config"
)

// 获取订阅名单
func SelectUsersAccount() []string {
	ret, err := callDatabase.CallMySQLSelectUsersAccount(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, config.MySQLName)
	if err != nil {
		return nil
	}

	return ret
}
