package main

import (
	"strings"
	"fmt"
)

func main() {
	
	str := "Hello World"

	// 获取字符串的长度
	fmt.Println(len(str))

	// 字符串拼接
	fmt.Println(str + " " +  "Let's" + " "+ "Go") // Hello World Let's Go

	// 字符串拼接并返回接收
	s1 :=  fmt.Sprintf("%s%s","Hi", " Hello") // Hi Hello
	fmt.Println(s1)

	// 字符串拼接直接打印
	fmt.Printf("%s","hello\n")  // hello

	// 字符串分割
	fmt.Println(strings.Split(str, " ")) // [Hello World] 

	// 包含
	s2 := "Hello"
	isContains := strings.Contains(s2,"H")
	fmt.Println(isContains) // true

	// 前缀/后缀
	s3 := "Hello World"
	isPadStart := strings.HasPrefix(s3,"He")
	isPadEnd := strings.HasSuffix(s3,"ld")
	fmt.Println(isPadStart) // true
	fmt.Println(isPadEnd)   // true

	// 判断字符出现在字符串中的位置
	s4 := "Hello World"
	findIndex := strings.Index(s4,"o")
	findIndex2 := strings.Index(s4,"b")
	findLastIndex := strings.LastIndex(s4,"o")
	fmt.Println(findIndex) // 4
	fmt.Println(findIndex2) // -1
	fmt.Println(findLastIndex) // 7

	// 字符串替换
	s5 := "Hello World"
	s6 := strings.Replace(s5,"World","Go",1)
	fmt.Println(s6) // Hello Go

}