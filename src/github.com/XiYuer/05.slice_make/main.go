package main

import (
	"fmt"
)

func main() {
	// 使用 make函数创建一个切片
	// 			类型，长度，容量（默认为长度）
	a1 := make([]int, 5, 10)  // make 用来申请一块内存空间，返回时的对应类型的指针

	fmt.Println(a1) // [0 0 0 0 0]
	fmt.Println(len(a1)) // 长度为：5
	fmt.Println(cap(a1)) // 容量为：10


	// 切片的遍历

	a2 := []string{"a", "b", "c"}

	for _, v := range a2 {
		fmt.Println(v)
	}

	// 向切片添加元素

	a3 := []string {"北京","上海","广州"}
	// a3[3] = "深圳"  // error: panic: runtime error: index out of range [3] with length 3 切片容量超出范围
	// fmt.Println(a3)

	// 正确地向切片中追加元素(调用 append 函数必须使用原切片的变量接受返回值)
	a3 = append(a3,"深圳")
	fmt.Println(a3)

	a4 := []string {"安徽","上海","杭州"}

	a3 = append(a3, a4...) // 将切片 a4 展开（...）添加到 a3 切片中

	fmt.Println(a3)
}