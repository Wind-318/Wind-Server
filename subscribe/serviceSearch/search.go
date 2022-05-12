package serviceSearch

import (
	"fmt"
	"subscribe/callClient/callRegister"
	"subscribe/config"
	"subscribe/registerCenter"
	"time"

	"stathat.com/c/consistent"
)

// 定时扫描
func ServiceSearch() {
	for {
		time.Sleep(time.Hour)
		HashConsistent = map[string]*consistent.Consistent{}
	}
}

// 初始化，进行服务发现，拉取服务缓存到本地
func InitSearch() {
	// 搜索
	rec, err := callRegister.CallServiceSearchAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 加入本地
	Services = rec
	// 构造哈希环
	for i := range rec {
		HashConsistent[i] = consistent.New()
		for j := range rec[i] {
			if i == registerCenter.SubscribeServiceName && rec[i][j] == config.SubscribeListenAddress+config.SubscribeListenPort {
				continue
			}
			HashConsistent[i].Add(rec[i][j])
		}
	}
}

// 发送心跳包检测服务是否正常运行
func CheckHealth() {}