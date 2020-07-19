package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	// file <==> 叫法 file对象 file文件句柄 file指针
	file, err := os.Open("D:/文档储存/Go_project/22_file/test.txt")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%v\n", file) // 返回*File
	err = file.Close()
	if err!=nil{
		fmt.Println(err)
	}
}