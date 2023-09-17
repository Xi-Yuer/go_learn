package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁

var (
	x = 0
	// 互斥锁
	lock sync.Mutex
	// 读写锁
	rwLock sync.RWMutex
	wg     sync.WaitGroup
)

// 读
func read() {
	defer wg.Done()
	// 读数据时加锁
	// lock.Lock()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Nanosecond)
	// lock.Unlock()
	rwLock.RUnlock()
}

// 写
func write() {
	defer wg.Done()
	// 写数据时也要加锁
	// lock.Lock()
	rwLock.Lock()
	x += 1
	time.Sleep(time.Nanosecond * 2)
	rwLock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()

	fmt.Println(x)
}
