package mainPackage

import (
	"database/registerCenter"
	"errors"

	"github.com/garyburd/redigo/redis"
)

// RedisServiceInterface 的接口实现
type RedisService struct{}

// 传入操作和参数
type Operation func(ip, port, operate string, parameters ...interface{}) (interface{}, error)

// 对缓存操作
func OperateDatabase(ip, port, operate string, parameters ...interface{}) (interface{}, error) {
	// 连接 redis
	var redisConnect, err = redis.Dial("tcp", ip+port)
	// 错误处理
	if err != nil {
		return nil, err
	}
	defer redisConnect.Close()
	// do
	ret, err := redisConnect.Do(operate, parameters...)
	// 错误处理
	if err != nil {
		return ret, err
	}
	// 返回
	return ret, nil
}

// 注册所有支持的运算
var Operators = map[string]Operation{
	"operate": OperateDatabase,
}

// 获取具体函数
func CreateOperation(operator string) (Operation, error) {
	var oper Operation
	if oper, ok := Operators[operator]; ok {
		return oper, nil
	}
	return oper, errors.New("illegal Operator")
}

// 调用函数，返回数据
func (c *RedisService) OperateRedisDatabase(request registerCenter.RedisDataInterface, reply *interface{}) error {
	// 获取具体函数
	oper, err := CreateOperation(request.Operation)
	if err != nil {
		return err
	}
	// 获取返回值
	*reply, err = oper(request.IP, request.Port, request.Operate, request.Parameters...)
	if err != nil {
		return err
	}
	return nil
}

// 获取所有支持的操作
func (c *RedisService) GetOperatorsForRedis(request interface{}, reply *[]string) error {
	arr := make([]string, 0, len(Operators))
	for key := range Operators {
		arr = append(arr, key)
	}
	*reply = arr
	return nil
}
