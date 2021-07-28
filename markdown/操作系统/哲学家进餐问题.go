package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

//定义哲学家对象
type philosopher struct {
	ID int
}

//哲学家的吃饭行为
func (p *philosopher) eating(mutex []sync.Mutex, i int) {
	//哲学家左手边筷子加锁，如果还未解锁，会等待其它哲学家吃完解锁
	mutex[i%5].Lock()
	//哲学家左手筷子延迟解锁
	defer mutex[i%5].Unlock()
	//哲学家右手边筷子加锁，如果还未解锁，会等待其它哲学家吃完解锁
	mutex[(i+1)%5].Lock()
	//哲学家右手边筷子延迟解锁
	defer mutex[(i+1)%5].Unlock()
	fmt.Println(strconv.Itoa(p.ID) + "正在干饭...")
	//干饭中
	time.Sleep(time.Second)
	fmt.Println(strconv.Itoa(p.ID) + "干饭完毕")
}

func main() {
	//5个筷子对应5个互斥锁
	mutex := make([]sync.Mutex, 5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			//创建哲学家对象
			pher := &philosopher{
				ID: i,
			}
			//哲学家开始干饭
			pher.eating(mutex, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
