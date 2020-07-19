package main

import (
	"fmt"
	"strconv"
	"time"
)

func test3() {
	str := ""
	for i:=0;i<100000;i++{
		str += "hello" + strconv.Itoa(i)
	}
	// fmt.Printf(str)
}

func main() {
	start := time.Now().Unix()
	test3()
	end:= time.Now().Unix()
	fmt.Printf("%v秒", end-start)
}