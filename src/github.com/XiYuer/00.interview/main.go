package main

import "fmt"

func f1() (r int) {
	defer func() {
		// 返回之前执行 defer r 自增一
		r++
	}()
	// 返回 10 将 r 的值赋值为10
	return 10
}
func f2() (r int) {
	t := 5
	defer func() {
		// 执行 defer 修改的是变量 t
		t = t + 5
	}()
	// 返回值 r 赋值为了 5
	return t
}

func f3() (r int) {
	// 函数形参是赋值拷贝
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	n := f1()
	fmt.Println(n)
}
