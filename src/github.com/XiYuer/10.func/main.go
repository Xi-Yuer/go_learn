package main

import "fmt"

// 命名返回值
func f1() (ret int) {
	ret = 1
	return
}

// 多个返回值
func f2() (int, string) {
	return 1, "a"
}

// 类型一样可以简写
func f3(x, y int) int {
	return x + y
}

// 可变参数
func f4(x string, y ...int) {
	fmt.Println(y)
	fmt.Printf("%T", y) // []int 类型是一个切片
}

func main() {

	// 定义一个函数
	fn := func(v string, k string) string {
		return v + " " + k
	}
	str := fn("Hello", "Function")
	fmt.Println(str)

	fmt.Println(f1())

	v1, v2 := f2()
	_, v3 := f2()
	fmt.Println(v1, v2)
	fmt.Println(v3)

	f4("a", 1, 2, 3, 4)
}
