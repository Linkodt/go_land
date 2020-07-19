package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	b = a.(Point)  // b = a ==>报错
	fmt.Println(b)

	// 如何在进行断言时带上检测机制
	var x interface{}
	var b2 float32 = 1.1
	x = b2
	// float32 才是对的
	if y, flag := x.(float64); flag {
		fmt.Printf("y的类型%T 值是=%v", y, y)
	}else{
		fmt.Println("fffffffffff")
	}
	fmt.Println("continue")
}