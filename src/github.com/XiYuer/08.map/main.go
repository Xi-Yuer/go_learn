package main

import (
	"fmt"
	"strings"
)

func main() {
	// 支持动态扩容，但应避免自动扩容，就好预估一下容量
	// 定义一个 map 类型的变量，他的 key 是 string 类型，值是 int 类型
	var m1 map[string]int
	//  map 必须初始化，需要使用 make 申请一块内存空间
	m1 = make(map[string]int, 10)
	m1["age"] = 18
	fmt.Println(m1)

	// 方式二：
	m2 := make(map[string]string, 5)
	m2["name"] = "Tom"
	fmt.Println(m2)
	fmt.Println(m2["name"])
	fmt.Println(m2["age"] == "") // 访问一个不存在的数值(会返回 map 对应类型的零值)

	// ============================================================================
	// map 遍历（for range）
	m3 := make(map[string]string, 5)
	m3["name"] = "Tom"
	m3["addr"] = "北京市"
	m3["phone"] = "184238213213"

	for k, v := range m3 {
		fmt.Println(k)
		fmt.Println(v)
	}

	// 删除 map 的默认键值(没有返回值)
	delete(m3, "addr")
	fmt.Println(m3)

	k, v := m3["name"] // 返回值值是 值 + 是否存在 (Tom true)
	n, m := m3["addr"] // 返回值值是 值 + 是否存在 (" " false)
	fmt.Println(k)
	fmt.Println(v)
	fmt.Println("==============")
	// fmt.Println(m3["name"])
	fmt.Println(n)
	fmt.Println(m)

	// 统计字符串中单词出现的次数
	str := "how do you do"
	m4 := make(map[string]int, 5)

	for _, v := range strings.Split(str, " ") {
		_, ok := m4[v]
		if ok {
			m4[v]++
		} else {
			m4[v] = 1
		}
	}

	fmt.Println(m4)
}
