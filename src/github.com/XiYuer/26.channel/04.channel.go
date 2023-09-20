package main

import (
	"fmt"
	"sync"
	"time"
)

func fn1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("发生错误")
		}
	}()
	cha := make(chan int, 1)
	cha <- 1
	close(cha)
	// 关闭通道之后不能够再发送值
	cha <- 1
	n := <-cha
	// 关闭后的通道有以下特点：
	/*
		1.对一个关闭的通道再发送值就会导致panic。
		2.对一个关闭的通道进行接收会一直获取值直到通道为空。
		3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
		4.关闭一个已经关闭的通道会导致panic。
	*/
	fmt.Println(n)
}

func fn2() {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	cha <- 1
}

func fn3() {
	defer wg.Done()
	fmt.Println("等待中...")
	n, ok := <-cha
	if !ok {
		fmt.Println("错误")
	}
	fmt.Println(n)
}

func fn4() {
	cha1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		cha1 <- i
	}
	close(cha1)

	//  for range 遍历通道会使得通道里面的内容自动被取出
	fmt.Println("遍历之前的长度", len(cha1))
	for v := range cha1 {
		fmt.Println(v)
	}
	fmt.Println("遍历之后的长度", len(cha1))
}

var cha = make(chan int)
var wg sync.WaitGroup

func main() {
	// wg.Add(2)
	// go fn3()
	// go fn2()
	// wg.Wait()

}
