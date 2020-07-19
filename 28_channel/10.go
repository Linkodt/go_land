package main

import (
	"fmt"
)

func main() {
	// 使用select可以解决从管道取数据的阻塞问题
	// 1. 定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i:= 0;i<10;i++{
		intChan <- i
	}
	// 2. 定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i:= 0;i<5;i++{
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// 遍历管道时 不关闭则堵塞 导致deadlock
	// select方式
	for {
		select {
		case v:= <-intChan: // 如果管道一直没有关闭，不会堵塞而导致deadlock
			// 自动到下一个case匹配
			fmt.Println("从intchan读取的数据",v)
		case v:= <-stringChan:
			fmt.Println("从stringchan读取的数据",v)
		default:
			fmt.Println("都取不到")
			return 
		}
	}
}