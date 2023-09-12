package main

import (
	"fmt"

	format "github.com/XiYuer/21.format/format"
)

// package 组织代码结构的一种方式

func main() {
	price := format.Format(100)
	fmt.Println(price)
}
