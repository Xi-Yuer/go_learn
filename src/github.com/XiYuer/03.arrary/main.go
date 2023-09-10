package main

import (
	"fmt"
)

func main() {
	// a1 := [2]int{1,2}
	// fmt.Print(a1)
	

	// 数组求和
	a2 := [...]int{1,2,3,4,5,6,7,8,9}

	sum := 0
	for _,v := range a2 {
		sum += v
	}

	fmt.Println(sum)


	// 数组赋值
	a3 := [2]int{1,2}
	a4 := a3
	a4[0] = 100

	fmt.Println("a3",a3) 	//  [1 2]  
	fmt.Println("a4",a4)	//	[100 2]

	// 练习

	a5 := [...]int{1,3,5,7,8}
	// var a6 [2][2]int

	for i := 0; i < len(a5); i++{
		for j := i+1; j< len(a5); j++ {
			if a5[i] + a5[j] == 8 {
				fmt.Println(a5[i],a5[j])
			}
		}
	} 
}