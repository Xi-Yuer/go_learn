package main

import (
	"errors"
	"fmt"
)

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

func foo() (i int) {

	i = 0
	defer func(i int) {
		fmt.Println(i)
	}(i)

	return 2
}

var run func() int = nil

func bar() {
	defer func() {
		// defer 捕获错误
		n := recover()
		fmt.Println("defer 延迟执行,捕获错误结果为：", n)
	}()
	panic("bar")

}

func chanFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误", err)
		}
	}()
	cha := make(chan int, 1)
	cha <- 1
}

type User struct {
	Name  string
	Email string
}

// 方法
func (u *User) Notify() {
	u.Name = "Sily"
	fmt.Println(u)
}

type Person struct {
	id   int
	name string
}

func (self *Person) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

// func getCircleArea(radius float32) (area float32) {
// 	if radius < 0 {
// 		// 自己抛
// 		panic("半径不能为负")
// 	}
// 	return 3.14 * radius * radius
// }

func test03() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束   （2）发生异常时
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}

func getCircleArea(radius float32) (area float32, err error) {
	if radius < 0 {
		// 构建个异常对象
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14 * radius * radius
	return
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	area, err := getCircleArea(-5)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(area)
	}
}
