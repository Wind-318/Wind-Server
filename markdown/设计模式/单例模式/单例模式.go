package main

import "fmt"

var uniqueSingleton *singleton = &singleton{"Hello world!"}

type singleton struct {
	Hello string
}

func GetSingleton() *singleton {
	return uniqueSingleton
}

func (single *singleton) Output() {
	fmt.Println(single.Hello)
}

func main() {
	uniqueSingleton.Output()
}
