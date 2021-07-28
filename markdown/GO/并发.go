//开启协程和并发控制
package main

import (
	"fmt"
	"sync"
	"time"
)

//声明一个等待组
var wg sync.WaitGroup

func hello(ch chan int) error {
	fmt.Println("开始执行")
	/*
		执行的操作
		有错误返回error
	*/
	//执行完任务后取出管道中的一个值，使得管道中有空位
	<-ch
	//等待组计数减一
	wg.Done()
	//没有错误，返回nil
	return nil
}

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 100; i++ {
		//等待组计数加一
		wg.Add(1)
		//并发控制，管道中最多有10个数，剩下的协程将被阻塞，等待进入管道
		ch <- 1
		//使用go关键字开启一个协程，主线程继续向后执行
		go hello(ch)
	}

	//声明一个互斥锁
	mutex := &sync.Mutex{}
	wg.Add(2)
	go mutexTest(mutex)
	go mutexTest(mutex)

	//声明一个读写锁
	rwMutex := &sync.RWMutex{}

	//等待所有任务执行完毕，即等待组计数为0
	wg.Wait()
}

func mutexTest(mutex *sync.Mutex) {
	defer wg.Done()
	//上互斥锁，此时其它线程进入会尝试加锁，如果已经是加锁状态，会堵塞在这里
	mutex.Lock()
	//互斥锁解锁
	defer mutex.Unlock()
	fmt.Println("测试互斥锁")
	time.Sleep(time.Second)
}

func RWmutexTest1(rwMutex *sync.RWMutex) {
	//上读锁，可以有多个线程进入加多个锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	//work...
}

func RWmutexTest2(rwMutex *sync.RWMutex) {
	//上写锁，只能由一个线程进入加锁，其它线程阻塞
	rwMutex.Lock()
	defer rwMutex.Unlock()

	//work...
}
