package main

import "fmt"

type father interface {
	show()
}

type fac1 struct{}

type fac2 struct{}

func (fac fac1) show() {
	fmt.Println("produce fac1")
}

func (fac fac2) show() {
	fmt.Println("produce fac2")
}

type factory struct{}

func (fac factory) selects(sel string) father {
	if sel == "fac1" {
		return fac1{}
	} else if sel == "fac2" {
		return fac2{}
	} else {
		fmt.Println("error")
		return nil
	}
}

func main() {
	sel := ""
	fmt.Scanln(&sel)
	fac := &factory{}
	fac.selects(sel).show()
}
