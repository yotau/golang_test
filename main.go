package main

import (
	"fmt"
)

func main() {
	// 定义二维切片的长度
	length := 3

	// 初始化二维字符串切片
	res := make([][]string, length)

	// 为每个元素分配具体的字符串切片
	res[0] = []string{"a", "b", "c"}
	res[1] = []string{"d", "e", "f"}
	res[2] = []string{"g", "h", "i"}

	// 打印二维切片的内容
	for i, row := range res {
		fmt.Printf("Row %d: %v\n", i, row)
	}

	fmt.Printf("\n")
}
