package main

import (
	"fmt"
	"os"
)

func main() {

	// 文件读写
	file, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("文件读取失败")
		return
	} else {
		fmt.Println(string(file))
	}

	fmt.Println("===========================")

	// 方式二：打开文件
	fileObj, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	} else {
		for {
			// 指定文件读取的长度(每次读取 128 个字节)
			content := make([]byte, 128)
			// 读取文件
			n, err := fileObj.Read(content)
			if err != nil {
				fmt.Println("文件读取失败")
			} else {
				fmt.Println(string(content))
			}
			if n < 128 {
				break
			}

		}
	}

	defer fileObj.Close() // 关闭文件

}
