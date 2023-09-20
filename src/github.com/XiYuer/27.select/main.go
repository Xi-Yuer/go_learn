package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fn1() {
	n := make(chan int, 1)

	// select 语句会随机的选择一条语句执行，如果执行不同就会执行另外一条语句
	for i := 0; i < 10; i++ {
		select {
		case val := <-n:
			fmt.Println(val)
		case n <- i:
		}
	}
}

func main() {
	n := make(chan int, 1)
	m := make(chan int, 1)
	x := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()

		// 多路复用
		for {
			select {
			case i := <-n:
				fmt.Println("从 n 中获得数据:", i)
			case i := <-m:
				fmt.Println("从 m 中获得数据:", i)
			default:
				fmt.Println("暂无数据")
				<-x
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second * 3)
			n <- rand.Intn(100)
			x <- 1
		}
	}()
	go func() {
		defer wg.Done()
		for {
			x <- 1
			time.Sleep(time.Second * 2)
			m <- rand.Intn(100)
		}
	}()

	wg.Wait()
}
