package main

import "fmt"

func fn() {
	fmt.Println("start")
	defer fmt.Println("哈哈哈")  // defer 后面的语句将在函数执行完毕返回之前才会执行
	fmt.Println("end")
}

func main() {
	fn()
}


// panic
// recover