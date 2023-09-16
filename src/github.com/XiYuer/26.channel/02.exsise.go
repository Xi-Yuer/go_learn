package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 两个通道，一个 routine 放值 一个routine 去值，最后输入
// 单向通道，只能存值或者去值，多用于函数参数的限制
// channel <- : 只能存值
// <- channel :	只能去值

// 该线程负责发送数据
func fn1(c chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		c <- i
	}
	// 通道关闭之后不可以再往里面写数据，但是可以读
	close(c)
}

// 该线程负责读数据
func fn2(c1 <-chan int, c2 chan<- int) {
	defer wg.Done()
	for {
		val, ok := <-c1
		if !ok {
			break
		}
		c2 <- val * val
	}
	// 通道关闭之后不可以再往里面写数据，但是可以读
	close(c2)
}

func main() {
	m := make(chan int, 100)
	n := make(chan int, 100)

	wg.Add(2)
	go fn1(m)
	go fn2(m, n)

	// 等待上面两个线程执行完毕
	wg.Wait()

	// 等待两个线程运行完毕之后，从 channel 里面去数据
	for v := range n {
		fmt.Println(v)
	}
}
