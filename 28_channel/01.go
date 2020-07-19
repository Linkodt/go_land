package main

import (
	"fmt"
	"time"
)
// 计算1-200各个数的阶乘，放入map中
var(
	mymap = make(map[int]int,10)
)

func test(n int){
	// 计算阶乘
	num:=1
	for i:=1;i<=n;i++{
		num *= i
	}
	mymap[n] = num

}

func main() {
	for i:=1;i<=50;i++{
		go test(i)
	}
	time.Sleep(time.Second*10)
	for i,v :=range mymap{
		fmt.Printf("map[%v]=%v\n",i,v)
	}
}