package main

import (
	"fmt"
	"strings"
)
 // 累加器
func AddUpper() func(int) int{
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = x + n
		str += string(36)
		fmt.Println("str =", str)
		return n
	}
}

func makeSuffix(suffix string) func (string) string{

	return func (name string) string{
		// 如果name没有指定后缀则加上，否则返回原来的name
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}


func main() {
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16

	f2 := makeSuffix(".jpg") //f就是func(name string) string{ ... }
	fmt.Println("文件处理后=",f2("winter")) // winter.jpg
	fmt.Println("文件处理后=",f2("hello.jpg"))  // hello.jpg
}