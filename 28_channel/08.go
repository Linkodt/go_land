package main
// 只读只写
import (
	"fmt"
)

func main() {
	var chan2 chan<- int // 只可以写
	chan2 = make(chan int, 3)
	chan2 <- 20
	// num:= <- chan2 // error
	fmt.Println("chan2=",chan2)

	var chan3 <-chan int // 只读
	num2:= <-chan3
	// chan3<- 30 //error
	fmt.Println("num2",num2)
}