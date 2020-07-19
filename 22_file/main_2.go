package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	// 打开文件
	// file <==> 叫法 file对象 file文件句柄 file指针
	file, err := os.Open("D:/文档储存/Go_project/22_file/test.txt")
	if err!=nil{
		fmt.Println(err)
	}
	// 当函数退出时需要及时关闭file
	defer file.Close() // 用defer更好,否则会有内存泄漏

	// 创建一个*Reader , 带缓冲的
	/*
	默认缓冲区为4096
	*/
	reader := bufio.NewReader(file)  // 创建缓冲区
	// 循环的读取文件的内容
	for{
		str, err := reader.ReadString('\n')  // 读到一个换行停一次
		fmt.Print(str)
		if err == io.EOF{ // 文件结束符
			break
		}
	}
	fmt.Println("文件读取结束")
}