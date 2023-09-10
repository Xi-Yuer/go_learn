package main

import (
	"fmt"
)

func main() {
	// 赋值切片
	a1 := []int{1, 2, 3}
	a2 := a1 // 这个是引用指针赋值
	a1[0] = 100
	fmt.Println(a2)

	var a3  = make([]int, 3)
	a4 := copy(a3, a1) 			// 将 切片 a1 拷贝到 切片 a3 中去, copy 函数的返回值是拷贝的数量长度
	fmt.Println(a3)
	
	a1[0] = 200				// 修改切片 a1 的值
	fmt.Println(a1)
	fmt.Println(a3)
	fmt.Println(a4)


	a6 := [...]int{1, 2, 3, 4, 5}
	a5 := a6[:]
	// 删除 a5 中的 2 3 元素
	a5 = append(a5[:1], a5[3:len(a5)]...)
	fmt.Println(a5)
	fmt.Println(a6)
}