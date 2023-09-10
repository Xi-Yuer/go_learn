package main

import "fmt"

func main() {
	// 定义一个元素类型为 map 切片
	s1 := make([]map[string]string, 5)

	// 给 map 初始化
	s1[0] = make(map[string]string, 1)
	// 给 map 赋值
	s1[0]["name"] = "Tom"

	fmt.Println(s1) // [map[name:Tom] map[] map[] map[] map[]]

	// 定义一个元素类型为切片类型的 map
	m1 := make(map[string][]string, 10) // 申请 map 内存空间 		{string:[]string,10}
	m1["A"] = make([]string, 3)         // 申请 slice 内存空间		[string,string,string]
	m1["A"] = []string{"a", "b", "c"}   // 为 slice 内存空间赋值

	fmt.Println(m1)
}
