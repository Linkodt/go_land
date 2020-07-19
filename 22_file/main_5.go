package main

// 复制文件

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func CopyFile(dstFileName string,srcFilename string)(written int64,err error){ //written拷贝的字节数
	srcFile, err := os.Open(srcFilename)
	if err!=nil{
		fmt.Printf("open file err=%v\n", err)
	}

	// 通过srcfile，获取到Reader
	reader := bufio.NewReader(srcFile)
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666) // 0666没什么用
	if err != nil {
		fmt.Printf("open file err=%v",err)
		return  
	}
	// 通过dstFile，获取到Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
	
}

func main() {
	// 将src拷贝到dst
	srcFile := "11223321.txt"
	dstFile := "test.txt"
	num , err := CopyFile(dstFile,srcFile)
	if err == nil{
		fmt.Println("success")
		fmt.Println(num)
	}else{
		fmt.Println("fault")
	}
// 没有拷贝成功 为什么 找不到原因
}