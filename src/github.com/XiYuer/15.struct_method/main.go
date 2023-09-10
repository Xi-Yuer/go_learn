package main

import "fmt"

type Person struct {
	name string
}

func newPerson(name string) *Person {
	return &Person{
		name,
	}
}

// 这个函数是一个方法，他指定了谁可以调用它
func (p Person) run() {
	fmt.Println(p.name, "在跑步")
}

func main() {
	p := newPerson("Tom")
	p.run()
	fmt.Println(*p)
}
