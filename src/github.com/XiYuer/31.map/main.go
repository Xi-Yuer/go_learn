package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 线程安全的 Map
var wg sync.WaitGroup

func main() {
	m := sync.Map{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(m)
}
