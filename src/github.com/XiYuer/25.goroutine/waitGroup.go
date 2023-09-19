package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义一个变量用于保存线程
var wg sync.WaitGroup

func routine(i int) {

	// 函数执行完毕线程池减一
	defer wg.Done()
	// 随机种子
	rand.Seed(time.Now().UnixNano())
	// 生成随机数
	n := rand.Intn(10) // 0 <= x < 10
	fmt.Println(i, n)
}

func main() {
	for i := 0; i < 5; i++ {
		//每次线程启动线程数量加一
		wg.Add(1)
		go routine(i)
	}
	// 主线程等待线程池执行完毕再结束
	wg.Wait()
}
