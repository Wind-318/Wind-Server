package algorithm

import (
	"Project/config"
	"math"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 生成随机数字用于 rand.Intn()，start 为最少时间，range 为波动范围
func GetRandomNumberTime(start, rangeNumber int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(rangeNumber)) + start
}

// 生成符合正态分布（平均值为 offest，标准差为 1）的整数，传入偏移量，允许的下限和上限
func GetNormalDistributionNumber(offset, minimum, maximum int) int {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	// 四舍五入取整生成
	ans := int(math.Floor(rand.NormFloat64()+0.5) + float64(offset))

	// 若不满足上下界条件则递归生成
	if ans > maximum || ans < minimum {
		return GetNormalDistributionNumber(offset, minimum, maximum)
	}

	return ans
}

// 获取一串长度为 length 的随机字符串
func GenerateRandStr(length int) (string, error) {
	redisCli, err := redis.Dial("tcp", config.RedisInfo)
	if err != nil {
		return "", err
	}
	defer redisCli.Close()
	code := ""
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	rand.Seed(time.Now().UnixNano())

	// 长度为 50
	for i := 0; i < length; i++ {
		code += string(letters[rand.Intn(len(letters))])
	}

	// 已存在，则再生成
	if ok, _ := redis.Bool(redisCli.Do("EXISTS", code)); ok {
		code, err = GenerateRandStr(length)
		if err != nil {
			return "", err
		}
	}

	return code, nil
}
