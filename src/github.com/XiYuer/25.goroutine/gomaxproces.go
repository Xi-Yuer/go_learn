package main

import (
	"fmt"
	"runtime"
	"sync"
)

// gomaxprocs(n)  指定 cpu 采用几个核 执行代码默认为跑满整个 cpu

var wg sync.WaitGroup

func fn() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func fn2() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	// 指定 cpu 内核使用数量
	runtime.GOMAXPROCS(2)
	// 添加两个线程
	wg.Add(2)
	go fn()
	go fn2()
	// 等待线程执行完毕
	wg.Wait()
}
