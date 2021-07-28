package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Queue struct {
	QueueContainer []interface{}
	mutex          sync.RWMutex
}

func (q *Queue) PushBack(val ...interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for _, value := range val {
		q.QueueContainer = append(q.QueueContainer, value)
	}
}

func (q *Queue) PushFront(val ...interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	temp := []interface{}{}
	for _, value := range val {
		temp = append(temp, value)
	}
	temp = append(temp, q.QueueContainer...)
	q.QueueContainer = temp
}

func (q *Queue) PopBack() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.QueueContainer) == 0 {
		panic("The queue is empty!")
	}
	data := q.QueueContainer[len(q.QueueContainer)-1]
	q.QueueContainer = q.QueueContainer[:len(q.QueueContainer)-1]
	return data
}

func (q *Queue) PopFront() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.QueueContainer) == 0 {
		panic("The queue is empty!")
	}
	data := q.QueueContainer[0]
	q.QueueContainer = q.QueueContainer[1:]
	return data
}

func (q *Queue) Back() interface{} {
	defer q.mutex.RUnlock()
	q.mutex.RLock()
	if q.Empty() {
		panic("The queue is empty!")
	}
	return q.QueueContainer[q.Size()-1]
}

func (q *Queue) Front() interface{} {
	defer q.mutex.RUnlock()
	q.mutex.RLock()
	if q.Empty() {
		panic("The queue is empty!")
	}
	return q.QueueContainer[0]
}

func (q *Queue) Empty() bool {
	defer q.mutex.RUnlock()
	q.mutex.RLock()
	return q.Size() == 0
}

func (q *Queue) Size() int {
	defer q.mutex.RUnlock()
	q.mutex.RLock()
	return len(q.QueueContainer)
}

func main() {
	q := &Queue{}
	for i := 1; i < 10001; i++ {
		wg.Add(1)
		go func(i int) {
			q.PushBack(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	data := q.Size()
	fmt.Println(q)

	fmt.Println("\n", data)
}
