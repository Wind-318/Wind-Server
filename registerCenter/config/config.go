package config

// MySQL 账号
var MySQLAccount = "root"

// 密码
var MySQLPassword = "root"

// IP 地址
var MySQLIP = "localhost"

// 监听端口
var MySQLPort = ":3306"

// 连接信息
var MySQLInfo = MySQLAccount + ":" + MySQLPassword + "@tcp(" + MySQLIP + MySQLPort + ")/registercenter"

// 服务注册发现中心服务器地址
var ServiceCenterAddress = "localhost"

// 端口
var ServiceCenterPort = ":8383"
