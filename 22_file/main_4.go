package main

import (
	"fmt"
	"os"
	"bufio"
)

func PathExists(path string)(bool, error){
	_ , err := os.Stat(path) // 返回nil说明文件存在,返回的错误为其他类型，说明不确定存不存在
	if err != nil {
		return true, nil
	}
	if os.IsNotExist(err){ // 返回false说明文件存在
		return false, nil
	}
	return false, err
}

func main() {
	// 写入5句 hello gardon
	// 打开文件
	filePath:="11223321.txt" // 当前目录下
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err!=nil{
		fmt.Println(err)
		return
	}
	// 关闭file句柄
	defer file.Close()
	str := "hello world\n"
	writer := bufio.NewWriter(file) // 带缓存
	for i:=0;i<5;i++ {
		writer.WriteString(str)
	}
	// 因为writer是带缓存的，所以调用方法时先写入缓存，所以需要调用Flush方法，将缓存的数据写入文件
	writer.Flush()
}