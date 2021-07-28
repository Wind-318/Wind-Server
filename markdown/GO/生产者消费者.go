package main

import "fmt"

func producer(ch chan<- int) {
	//生产者一共生产200个物品
	for i := 0; i < 200; i++ {
		//如果缓冲区满则阻塞
		ch <- i
	}
	//关闭管道，通知消费者结束任务
	close(ch)
}

func consumer(ch <-chan int) {
	//开始消费，如果缓冲区空则阻塞
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	//缓冲区为100容量
	ch := make(chan int, 100)
	//生产者开始生产
	go producer(ch)
	//消费者开始消费
	consumer(ch)
}
