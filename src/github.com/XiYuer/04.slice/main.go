package main

import (
	"fmt"
)

func main() {
	// 定义一个切片
	var a1 []int
	var a2 []string

	fmt.Println(a1 == nil) // true （nil 表示空，未分配内存空间）

	// 赋值
	a1 = []int{1, 2, 3}
	a2 = []string{"a", "b", "c"}

	// 
	fmt.Println("长度：",len(a1))
	fmt.Println("容量：",cap(a1))
	fmt.Println(a1)
	fmt.Println(a2)


	// 由数组得到切片
	a3 := []int{1, 2, 3, 4, 5, 6}
	a4 := a3[0:3] // [1 2 3]  基于一个数组切割，左闭右开（包含，不包含）
	a5 := a3[:4]
	a6 := a3[2:]
	a7 := a3[:]
	fmt.Println("a4[0:3]", a4)  // [0, 3]
	fmt.Println("a5[:4]",  a5)  // [0, 4]
	fmt.Println("a6[2:]",  a6)  // [2, len(a3)]
	fmt.Println("a7[:]",   a7)  // [0, len(a3)]  全部切割

	// 切片的容量(切片的容量是底层数组切片之后的的剩余长度)
	fmt.Println("切片的容量", cap(a6)) // 4
	fmt.Println("切片的长度", len(a5)) // 4

	// 切片是一个引用类型，当修改了底层数组的值，切片的值也会被修改
	fmt.Println(a4)  	// [1 2 3]
	a3[0] = 100
	fmt.Println(a4) 	// [100 2 3]
}