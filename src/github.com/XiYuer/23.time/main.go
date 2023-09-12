package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 年
	fmt.Println(now.Year())
	// 月
	fmt.Println(now.Month())
	// 日
	fmt.Println(now.Date())
	// 第几周
	fmt.Println(now.Day())
	// 时
	fmt.Println(now.Hour())
	// 分
	fmt.Println(now.Minute())
	// 秒
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) // 毫秒级

	ret := time.Unix(1694529766, 0)
	fmt.Println(ret) // 2023-09-12 22:42:46 +0800 CST


	// 时间格式化(将时间格式化为字符串)  2006-01-02 03:04:05 === YYY-mm-dd MM:DD:ss 
	fmt.Println(now.Format("2006-01-02 03:04:05"))
}
