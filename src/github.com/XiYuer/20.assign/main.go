package main

import "fmt"

// 空接口函数
func assign(n interface{}) {
	// val, ok := n.(int)
	// if !ok {
	// 	fmt.Printf("%T", val)
	// 	fmt.Printf("类型不正确")
	// } else {
	// 	fmt.Printf("%T", val)
	// 	fmt.Println(val)
	// }

	switch t := n.(type) {
	case int:
		fmt.Println("数字", t)
	case string:
		fmt.Println("字符串", t)
	case bool:
		fmt.Println("布尔值", t)
	default:
		fmt.Println("其他类型")
	}

}

func main() {
	// 类型断言
	assign(10)
	assign("10")
}
