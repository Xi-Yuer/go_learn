package main

import "fmt"

// 接口
type person interface {
	run()
	eat(string)
}

// 空接口(任何类型都满足：当不知道函数需要什么类型时，可以使用空接口)
// 任意的变量都满足空接口
func emtypeFn(n interface{}) {
	fmt.Println(n)
}

func bar(p person) {
	p.run()
	p.eat("零食")
}

type human struct {
	name string
	age  int
}

func (h *human) run() {
	fmt.Println(h.name + "在跑步")
}
func (h *human) eat(food string) {
	fmt.Println(h.name + "在吃东西")
}

func main() {
	p := &human{
		name: "Tom",
		age:  18,
	}
	p.eat("苹果")
	bar(p)

	// 定义一个空接口的 map ,可以存储任何值
	m := map[string]interface{}{}
	m["name"] = "Tom"
	m["age"] = 18
	m["isWuman"] = true
	m["hobby"] = [...]string{"唱", "跳", "Rap"}
	fmt.Println(m) // map[age:18 isWuman:true name:Tom]
}
