package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 模拟构造函数，返回一个结构体的指针
func newPerson(name string, age int) *Person {
	// & 根据值去地址并返回地址值
	return &Person{
		name,
		age,
	}
}

func main() {
	p := newPerson("Tom", 18)

	fmt.Println(*p)

	p.name = "Siiy"

	fmt.Println(*p)
}
