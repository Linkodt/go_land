package main
// channel遍历 关闭
import (
	 "fmt"
)

func main() {
	intChan := make(chan int,3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// Not write
	// intChan <- 300
	num := <- intChan
	fmt.Println("num=",num)

	intChan2 := make(chan int, 100)
	for i:=0;i<100;i++{
		intChan2 <- i*2
	}
	// 遍历时未关闭则会死锁
	close(intChan2)
	for v := range intChan2{
		fmt.Println("v=",v)
	}
}