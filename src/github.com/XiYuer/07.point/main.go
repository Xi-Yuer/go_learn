package main

import "fmt"

func main() {

	// &: 根据变量取地址值
	// *：根据地址值获取变量
	n := 18

	p := &n

	m := *p // 根据地址值获取变量

	fmt.Println(&n) // 0xc00000a0b8(获取变量 n 的地址值)
	fmt.Println(p)  // 0xc00000a0b8
	fmt.Println(*p) // 18 (根据地址值获取值)
	fmt.Println(m)  // 18


	// 使用 new 可以申请一块内存地址
	var x = new(int)
	fmt.Println(x)  	// 0xc00000a100
	fmt.Println(*x)		// 0

	*x = 100			// 根据地址值 *x 获取变量并赋值
	fmt.Println(*x)
}
