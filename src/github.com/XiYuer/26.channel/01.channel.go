package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个 channel 变量
var n chan int

var wg sync.WaitGroup

// 不带缓冲区的 channel
func nobufChannel() {
	n := make(chan int) // 不带缓冲区的通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("另一个线程")
		// 在另一个线程中接受通道值(只有接收到了通道发过来的值，代码才会继续往下执行)
		ret := <-n
		fmt.Println(ret)
	}()
	// 向通道发送值
	time.Sleep(time.Second * 2)
	n <- 10

	// 关闭通道
	close(n)
}

// 带缓冲区的channel

func hasBufChannel() {
	n := make(chan int, 10) // 缓冲区为 10 的 chaunnel
	n <- 10
	fmt.Println("10 发送到通道中去了")

	// 关闭通道
	close(n)
}

func main() {
	// 通道的初始化（内存申请分配）
	// m := make(chan int, 10) // 带缓冲区的通道初始化
	// 通道的操作
	// 发送 chan <- 1  将值发送给通道
	// 接受 ret := <-chan 从通道获取值
	// 关闭 close()
	nobufChannel()
	hasBufChannel()
	wg.Wait()
}
