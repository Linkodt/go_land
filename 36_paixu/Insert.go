package main

import(
	"fmt"
)

func InsertSort(arr *[5]int){
	// 第一次
	for i:=1;i<len(arr);i++{
		InsertVal := arr[i]
		insertIndex := i-1
		for insertIndex>=0 && arr[insertIndex]<InsertVal{
			arr[insertIndex + 1] = arr[insertIndex]	// 后移..将大于要排列的数往后移动
			insertIndex--
		}
		// 插入
		if insertIndex + 1 != i{
		arr[insertIndex + 1] = InsertVal
		}
		fmt.Printf("第%d次插入后 %v\n", i,*arr)
	}
}

func main(){
	arr := [5]int{23, 0, 12, 56, 34}
	InsertSort(&arr)
}