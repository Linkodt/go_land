package main

import (
	"fmt"
	"strconv"
)

func main(){
	// 统计字符串长度 按字节
	// golang编码统一为utf-8 字符占1个字节，汉字占3个字节
	str := "hello北惊"
	fmt.Println("str len=", len(str))  // 11

	// 字符串遍历，同时解决有中文的问题
	r := []rune(str)
	for i := 0; i<len(r); i++{
		fmt.Printf("字符=%c\n", r[i])
	}
	// 字符=h
	// 字符=e
	// 字符=l
	// 字符=l
	// 字符=o
	// 字符=北
	// 字符=惊

	// 字符串转整数
	n_1 , err_1 := strconv.Atoi("12")
	n_2 , err_2 := strconv.Atoi("hello")
	if err_1 != nil{
		fmt.Println("转换错误", err_1)
	}else{
		fmt.Println("转换成果是", n_1)
	}
	if err_2 != nil{  // err!=nil 则没有错
		fmt.Println("转换错误", err_2)  // error: strconv.Atoi: parsing "hello": invalid syntax
	}else{
		fmt.Println("转换成果是", n_2)
	}

	// 整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v  str=%T\n", str, str)  // %v 打印字符串 %T 打印类型  12345 string

	// 字符串转 []byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)  // bytes=[104 101 108 108 111 32 103 111]

	// []byte 转 字符串
	str = string([]byte{97, 98, 99})  // abc
	fmt.Printf("str=%v\n", str)
}