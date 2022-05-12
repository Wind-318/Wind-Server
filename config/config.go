package config

// 是否允许注册
var AllowRegister = false

// 发送方邮箱
var SenderAccount = "YourEmail@outlook.com"

// 认证码
var SenderPassword = "VerificationCode"

// 邮箱服务器地址，outlook 是 smtp.office365.com
var ServerAddr = "smtp.office365.com"

// 开启爬虫
var AllowSpider = false

// 开启 https
var TurnOnHttps = false

// 域名
var YourName = "127.0.0.1"

// pem 文件（可选）
var HttpsCertification1 = "https.pem"

// key 文件（可选）
var HttpsCertification2 = "https.key"

// 端口
var ServerPort = 587

// 管理员邮箱
var SystemAccount = "YourEmail@outlook.com"

// 域名
var Addr = "http://127.0.0.1/"

// MySQL 账号
var MySQLAccount = "root"

// 密码
var MySQLPassword = "root"

// IP 地址
var MySQLIP = "localhost"

// 监听端口
var MySQLPort = ":3306"

// 连接信息
var MySQLInfo = MySQLAccount + ":" + MySQLPassword + "@tcp(" + MySQLIP + MySQLPort + ")/"

// IP
var RedisIP = "localhost"

// Port
var RedisPort = ":6379"

// info
var RedisInfo = RedisIP + RedisPort
