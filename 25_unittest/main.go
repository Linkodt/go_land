package main

import (
	_ "fmt"
)

func addUpper(n int) int{
	res := 0
	for i:=1;i<=n-1;i++{
		res += i
	}
	return res
}

func addUpper2(n int) int{
	res := 0
	for i:= 1;i<=n-1;i++{
		res += i
	}
	return res
}

// testing 框架

func main() {
	
}