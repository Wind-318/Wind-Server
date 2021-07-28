package flowcontroller

import (
	"errors"
	"sync/atomic"
	"time"
)

var Controller *flowControl = &flowControl{
	limit:     10000,
	remainNum: 10000,
	container: make(chan int, 10000),
}

type flowControl struct {
	limit     int64
	remainNum int64
	container chan int
}

func (t *flowControl) PopToken() int64 {
	<-t.container
	return atomic.AddInt64(&t.remainNum, 1)
}

func (t *flowControl) GetToken() (int64, error) {
	// 设置 15 秒超时时间
	select {
	case t.container <- 1:
	case <-time.After(time.Second * 15):
		return -1, errors.New("time out")
	}

	return atomic.AddInt64(&t.remainNum, -1), nil
}

func (t *flowControl) GetRemainNum() int64 {
	return atomic.LoadInt64(&t.remainNum)
}
