package main

import (
	"fmt"
)


func main() {
	var intArr [5]int = [...]int{1, 22, 33, 44, 99}
	var slice = intArr[1:3]  // 下标 从1开始到3结束 不包括3
	fmt.Println("intArr=",intArr)
	fmt.Println("slice的元素=",slice)
	fmt.Println("slice的元素个数=",len(slice))
	fmt.Println("slice的容量=",cap(slice)) // 目前可存放最多个数 可动态变化 容量不够时自动增长
	// intArr= [1 22 33 44 99]
	// slice的元素= [22 33]
	// slice的元素个数 = 2
	// slice的容量= 4
	slice[0] = 1234
	fmt.Println("intArr=",intArr)
	fmt.Println("slice的元素=",slice)
	// intArr= [1 1234 33 44 99]
	// slice的元素= [1234 33]
	// 说明slice内是指针 指向原数组被引用的地址

	// 切片必须make后使用
	var slice_2 []int = make([]int, 5, 10) // 5是大小 10是容量，选填
	fmt.Println(slice_2)  // [0 0 0 0 0]

	slice_3 := append(slice_2, slice...) // 切片追加
	fmt.Println(slice_3) // [0 0 0 0 0 1234 33]

	for i:=0; i<len(slice_2);i++{
		slice_2[i] = i
	}
	fmt.Println()
	slice_2 = append(slice_2, 2)
	fmt.Println(slice_2) // [0 1 2 3 4 2]
	slice_4 := copy(slice_2, slice_3) // slice_3:[0 0 0 0 0 1234 33]
	fmt.Println(slice_2)
	fmt.Println(slice_3)
	fmt.Println(slice_4) // [0 1 2 3 4 2 33]
}