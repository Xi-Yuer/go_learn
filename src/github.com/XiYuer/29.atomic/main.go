package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子性操作

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func fn1() {
	defer wg.Done()

	// 互斥锁加锁
	// lock.Lock()
	// x++
	// lock.Unlock()

	// 通过原子操作实现加锁
	atomic.AddInt64(&x, 1)
}

func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go fn1()
	}
	wg.Wait()
	fmt.Println(x)
}
