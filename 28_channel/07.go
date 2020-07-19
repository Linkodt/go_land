package main
// 1-8000素数
import (
	"fmt"
	"time"
)
// 放数字进intChan
func putNum(intChan chan int){
	for i:=1;i<=8000;i++{
		intChan<-i
	}
	close(intChan)
	fmt.Println("putNum退出")
}
// 取数字 判断是否素数， 放入
func primeNum(intChan chan int,primeChan chan int, exitChan chan bool,i int){
	var flag bool
	for{
		time.Sleep(time.Millisecond*10)
		num,ok := <- intChan
		if!ok{
			// 已经取完
			break
		}
		// 假设是素数
		flag = true
		// 判断是否素数
		for i:=2;i<num;i++{
			if num%i==0{
				// 不是素数
				flag = false
				break
			}
		}
		if flag{
			primeChan <- num
		}

	}
	fmt.Println("primeNum协程取不到数据退出",i)
	exitChan<-true
}


func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4) // 4个协程
	go putNum(intChan)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan,i)
	}

	go func(){
		for i := 0; i < 4; i++ {
			<-exitChan // 没有四个的话就不会动，等到有了才执行
		}
		close(primeChan)
	}()

	for{
		res , ok := <- primeChan
		if !ok{
			break //取完
		}
		fmt.Println("素数=",res)
	}
	fmt.Println("main退出")
	
}