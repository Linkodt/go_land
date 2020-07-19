package main

import (
	"fmt"
)

func main() {
	// 获取输入！！！！记得地址符号
	var b int
	fmt.Println("请输入：")
	fmt.Scanln(&b)
	// fmt.Scanf("%d",b)
	fmt.Println(b)
}