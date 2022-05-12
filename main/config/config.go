package config

// 管理员账户
var SystemUserAccount = "Administrator@outlook.com"

// 是否开放注册，开启填 true
var AllowRegister = false

// 是否启用爬虫，开启填 true
var AllowSpider = false

// 域名
var YourName = "127.0.0.1"

// 域名
var Addr = "http://127.0.0.1/"

// 默认关闭 https，开启填 true
var TurnOnHttps = false

// https 证书 1 名（.pem 文件，放置在根目录下）
var HttpsCertification1 = "https.pem"

// https 证书 2 名（.key 文件，放置在根目录下）
var HttpsCertification2 = "https.key"

// 数据库账号
var MySQLAccount = "root"

// 密码
var MySQLPassword = "root"

// IP
var MySQLIP = "localhost"

// 端口
var MySQLPort = ":3306"

// Redis IP
var RedisIP = "localhost"

// 端口
var RedisPort = ":6379"

// 服务注册发现中心服务器地址
var ServiceCenterAddress = "localhost"

// 端口
var ServiceCenterPort = ":8383"
