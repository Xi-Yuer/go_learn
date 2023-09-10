package main

import "fmt"

// 模拟实现继承

type animal struct {
	name string
}

// 结构体的方法
func (a animal) run() {
	fmt.Println(a.name + "在跑")
}

// 结构体 Dog 继承了Animal 结构体
type dog struct {
	animal
}

// 结构体的方法
func (d dog) wolf() {
	fmt.Println(d.name, "汪汪汪")
}

// 返回一个狗的指针
func newDog(name string) *dog {
	return &dog{
		animal{
			name,
		},
	}
}

func main() {
	d := newDog("小黑")
	fmt.Println(*d)
	d.animal.run()
	d.run()
	d.wolf()
}
