package main

import "fmt"

func hello() {
	fmt.Println("hello")
}


// 程序启动会创建一个 mian groutine 去执行
func main() {
	// 开启一个单独的线程 routine 去执行 hello 这个函数
	go hello()
	fmt.Println("main")
	// main 函数执行完毕之后 又 mian 函数启动的 groutine 都会结束
}
