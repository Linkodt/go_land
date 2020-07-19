package main

import (
	"fmt"
)

func main() {
	// 演示管道的使用
	// 1. 创建一个存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int,3)

	// 看看intChan是什么
	fmt.Printf("intChan = %v     intChan本身的地址=%p \n",intChan,&intChan)

	// 向管道写入数据
	intChan <- 10
	NUM := 211
	intChan <- NUM
	intChan <- 51
	// 写入的数据不能超过容量
	// 看管道的长度和cap(容量)
	fmt.Printf("len=%v    cap=%v  \n",len(intChan),cap(intChan))

	// 读取数据
	var num2 int
	num2 = <- intChan
	fmt.Println("num2=",num2)
}