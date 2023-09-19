package main

import "fmt"

func main() {
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
