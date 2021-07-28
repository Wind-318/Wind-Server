package main

import "fmt"

//定义接口
type inter interface {
	//方法一
	way1() int
	//方法二
	way2(type1 int, type2 string) error
}

//返回error类型
func test() error {
	/*
		若发生错误，则返回错误信息
		if ... {
			return errors.New("This is an error!")
		}
	*/
	//否则返回nil
	return nil
}

func main() {
	//////////////////////////切片
	//声明一个int类型切片，此时切片为nil
	var slice []int
	//使用make创建一个容量为0的切片，此时切片不为nil，而只是容量为0
	slice1 := make([]int, 0)
	fmt.Println(slice1)
	//初始化一个切片
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	//切下slice2的1到最后一个元素，给另一个切片赋值
	slice3 := slice2[1:]
	fmt.Println(slice3)

	//求出切片的长度
	length := len(slice)
	fmt.Println(length)
	//求出切片的容量
	capacity := cap(slice)
	fmt.Println(capacity)

	//追加元素
	slice = append(slice, 1)
	//拷贝slice2的元素给slice
	copy(slice, slice2)
	///////////////////////////////映射
	//定义一个map，类似于切片
	m := make(map[int]int)
	m[5] = 1
	m[2] = 3
	fmt.Println(m)
	//删除键值5
	delete(m, 5)
	fmt.Println(m)
	///////////////////////////////管道
	//创建一个有5个容量的管道
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		//向管道中写入数据,若管道已满，则阻塞
		ch <- i
	}
	for i := range ch {
		//读取管道中的元素( <-ch )，若已空，则阻塞
		fmt.Println(i)
	}
	//关闭管道，通知读取的协程不用继续阻塞，若为空则继续向后执行
	close(ch)
}
