package serviceSearch

import (
	"errors"
	"net"

	"stathat.com/c/consistent"
)

// 服务合集
var Services = map[string][]string{}

// 一致性哈希环
var HashConsistent = map[string]*consistent.Consistent{}

// 获取节点
func GetNode(name string) (string, error) {
	// 双重判断是否有服务运行
	if _, ok := HashConsistent[name]; !ok {
		InitSearch()
	}
	// 刷入本地后仍未找到
	if _, ok := HashConsistent[name]; !ok {
		return "", errors.New(name + " 服务未找到可用节点！")
	}
	ips, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	ip := ""
	for _, address := range ips {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}

	// 获取节点
	remoteAddr, err := HashConsistent[name].Get(ip)
	if err != nil {
		return "", err
	}

	return remoteAddr, nil
}
