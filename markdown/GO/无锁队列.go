package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

type LockFreeQueue struct {
	Container [100]int64
	Head      int64
	End       int64
}

func (q *LockFreeQueue) Push(val int64) {
	if q.Full() {
		fmt.Println("Full")
		return
	}
	if q.Empty() {
		atomic.StoreInt64(&q.Head, 0)
	}
	atomic.StoreInt64(&q.End, (q.End+1)%100)
	atomic.StoreInt64(&q.Container[q.End], val)
}

func (q *LockFreeQueue) Pop() int64 {
	if q.Empty() {
		fmt.Println("Empty!")
		return -1
	} else if q.Head == q.End {
		var ret int64
		atomic.StoreInt64(&ret, q.Container[q.Head])
		atomic.StoreInt64(&q.Head, -1)
		atomic.StoreInt64(&q.End, -1)
		return ret
	}
	var ret int64
	atomic.StoreInt64(&ret, q.Container[q.Head])
	atomic.StoreInt64(&q.Head, (q.Head+1)%100)
	return ret
}

func (q *LockFreeQueue) Empty() bool {
	return q.Head == -1
}

func (q *LockFreeQueue) Full() bool {
	return (q.End+1)%100 == q.Head
}

func main() {
	test := &LockFreeQueue{
		Head: -1,
		End:  -1,
	}
	for i := 100; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			test.Push(int64(i))
		}(i)
	}
	wg.Wait()
	fmt.Println(test)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(test.Pop())
		}()
	}
	wg.Wait()
}
