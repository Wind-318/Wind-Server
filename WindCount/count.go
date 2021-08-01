package WindCount

import "sync/atomic"

type count struct {
	todayNum int64
	allNum   int64
}

var onlyCount *count = &count{
	todayNum: 0,
	allNum:   0,
}

func GetCount() *count {
	return onlyCount
}

func (c *count) GetTodayNum() int64 {
	return c.todayNum
}

func (c *count) GetAllNum() int64 {
	return c.allNum
}

func (c *count) AddNum() {
	atomic.AddInt64(&c.todayNum, 1)
	atomic.AddInt64(&c.allNum, 1)
}
