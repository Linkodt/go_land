package main

import (
	"fmt"
)

func sum(n1 int,n2 int) int{
	// 执行到defer时不执行，压入独立栈
	// 采用先入后出的方式出栈
	defer fmt.Println("ok1 n1=", n1) //3.  10
	defer fmt.Println("ok2 n2=", n2) //2.  20

	res := n1 + n2
	fmt.Println("ok3 res=", res)     //1.  30
	return res
	
}

func main() {
	res := sum(10,20)
	fmt.Println("res=",res)         //4.   30
}