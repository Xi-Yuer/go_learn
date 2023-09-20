package main

import (
	"fmt"
	"sync"
)

// 线程安全的 Map
var wg sync.WaitGroup

func main() {
	m := sync.Map{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		for i := 0; i < 10; i++ {
			m.Store("A", i)
		}
	}()
	go func() {
		defer wg.Done()
		wg.Add(1)
		for i := 0; i < 10; i++ {
			m.Store("B", i)
		}
	}()

	wg.Wait()

	fmt.Println(m)
}
