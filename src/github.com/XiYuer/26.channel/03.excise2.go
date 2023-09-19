package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	x int64
}
type result struct {
	job *job
	sum int64
}

// 不断的生产任务向任务池中推送任务
func fn1(j chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		j <- &job{
			x,
		}
		time.Sleep(time.Second)
	}
}

// 不断的向任务池中去任务并执行
func fn2(j <-chan *job, res chan<- *result) {
	defer wg.Done()
	// 不断的从池之中读取值
	for {
		job := <-j
		sum := int64(0)
		value := job.x
		for value > 0 {
			sum += value % 10
			value = value / 10
		}
		// 造一个新的结果
		res <- &result{
			job,
			sum,
		}
	}
}

var wg sync.WaitGroup
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func main() {
	wg.Add(1)
	go fn1(jobChan)
	// 开启 24 个线程池执行
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go fn2(jobChan, resultChan)
	}
	// 取出结果并打印输出
	for val := range resultChan {
		fmt.Println(val.job.x, val.sum)
	}
	wg.Wait()
}
