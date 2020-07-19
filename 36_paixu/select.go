package main

import (
	"fmt"
)

func SelectSort(arr *[5]int){
	// (*arr)[1] <==> arr[1]
	// 1. 最大和arr[0]交换
	for j:=0;j<len(arr)-1;j++{
	max := arr[j]
	maxIndex := j
	// 2. 遍历后面1-------[len(arr) - 1]比较
	for i:=j+1;i<len(arr);i++{
		if max<arr[i]{// 找到真正的最大值
			max = arr[i]
			maxIndex = i
		}
	}
	// 交换
	if maxIndex != j{
		arr[j],arr[maxIndex] = arr[maxIndex],arr[j]
	}

	fmt.Printf("第%d次 %v\n", j+1,*arr)

	}
}

func main() {
	arr := [5]int{6,88,32,1,15}	
	SelectSort(&arr)
}