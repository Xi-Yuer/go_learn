package main

import "fmt"

type Person struct {
	name string
	age  int
	addr string
}

func fn(p Person) {
	p.name = "Sily"
}

func fn2(p *Person) {
	(*p).name = "Sily" // 传入是是一个指针，根据指针寻找变量并修改值
}

func main() {
	var p Person
	p.name = "Tom"
	p.age = 18
	p.addr = "北京市"
	// p.name = "Sily" // 可以更改

	fn(p) // 函数传入的参数永远都是拷贝赋值
	fmt.Println(p)
	fn2(&p)
	fmt.Println(p)
}
