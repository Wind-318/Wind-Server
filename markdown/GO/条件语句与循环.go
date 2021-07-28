package main

import "fmt"

func main() {
	//使用make创建一个切片
	arr := make([]int, 0)
	//循环10次，给切片中添加数字
	for i := 0; i < 10; i++ {
		arr = append(arr, i*100)
	}
	//range循环中，i为下标，j为值，如果不需要某一个，可以使用"_"来替代
	for i, j := range arr {
		fmt.Printf("第%d个数是%d\n", i, j)
	}
	/*
		输出为：
		    	第0个数是0
				第1个数是100
				第2个数是200
				第3个数是300
				第4个数是400
				第5个数是500
				第6个数是600
				第7个数是700
				第8个数是800
				第9个数是900
	*/

	//判断arr长度是否为0，输出相应结果
	if len(arr) != 0 {
		fmt.Println("arr不为空")
	} else {
		fmt.Println("arr为空")
	}

	num := 5
	switch num {
	case 1:
		//操作
	case 2:
		//操作
	// case ...
	case 5:
		fmt.Println(5)
		//fallthrough关键字可以让switch继续执行下一个case
		fallthrough
	case 6:
		fmt.Println(6)
	}
}
