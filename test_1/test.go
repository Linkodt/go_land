package main

import (
	"fmt"
)

func test(arr *[3]int){
	(*arr)[0] = 88
}

func test_2(arr [3]int){
	arr[0] = 20
	fmt.Println("test_2=",arr)  
}

func main() {
	var arr [3]int
	arr = [3]int{11,22,33}
	fmt.Println(arr)  // 11 22 33
	test(&arr)
	fmt.Println(arr)  // 88 22 33
	test_2(arr)  // 20 22 33
	fmt.Println("main=",arr)  // 88 22 33

}