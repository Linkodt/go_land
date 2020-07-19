package main

import "fmt"

func testPtr(num *int) {

	*num = 20
}

func test(num_1 int, num_2 int) (int, int) {
	// 返回多个值
	sum := num_2 + num_1
	sub := num_1 - num_2
	return sum, sub
}

func main() {

}
