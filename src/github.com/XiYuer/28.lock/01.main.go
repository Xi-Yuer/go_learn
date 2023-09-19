package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
var lock sync.Mutex


// 互斥锁
func fn1() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		// 对操作的资源进行加锁
		lock.Lock()
		x += 1
		lock.Unlock()
	}

}

func fn2() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go fn1()
	go fn2()
	wg.Wait()
	fmt.Println(x)
}
