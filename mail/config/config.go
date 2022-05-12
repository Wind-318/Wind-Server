package config

// 邮件相关服务器 IP 地址
var MailListenAddress = "localhost"

// 邮件相关服务监听端口
var MailListenPort = ":8286"

// 发送方邮箱
var SenderAccount = "YourEmail@outlook.com"

// 认证码（不同邮件服务认证码不同，outlook 设置后可以直接使用密码）
var SenderPassword = "YourVerificationCode"

// 邮箱服务器地址，outlook 是 smtp.office365.com
var ServerAddr = "smtp.office365.com"

// 端口
var ServerPort = 587

// 服务注册发现中心服务器地址
var ServiceCenterAddress = "localhost"

// 端口
var ServiceCenterPort = ":8383"

var RedisIP = "localhost"

var RedisPort = ":6379"
