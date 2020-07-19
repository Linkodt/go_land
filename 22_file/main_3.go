package main

import (
	"fmt"
	_ "os"
	_ "bufio"
	_ "io"
	"io/ioutil"
)

func main() {
	file:="D:/文档储存/Go_project/22_file/test.txt"
	content,err := ioutil.ReadFile(file)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%v", string(content))
	// 因为没有open 所以也不需要close 已经被封装到readfile里面了
}