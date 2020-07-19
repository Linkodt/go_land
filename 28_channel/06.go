package main

import (
	"fmt"
	_"time"
)

func writeData(intChan chan int){
	for i:=1;i<=50;i++{
		intChan <- i
		fmt.Println("writeData=",i)
		// time.Sleep(time.Second)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool){
	for{
		v, ok := <- intChan
		// time.Sleep(time.Second)
		if !ok{
			break
		}
		fmt.Println("数据=",v)
	}
	// readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int,15)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for{
		i , _ := <- exitChan
		if i == true{
			break
		}	
	}
	
}