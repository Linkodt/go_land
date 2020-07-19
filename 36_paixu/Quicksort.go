package main

import (
	"fmt"
)
// left表示数组左边的下标
// right表示数组右边的下标
// array表示要排序的数组
func quickSort(left int,right int,array *[6]int) {
	l :=left
	r :=right
	// pivot中轴
	pivot :=array[(left + right)/2]
	temp := 0
	// for循环将数组以pivot分隔开
	for ; l<r;{
		for ;array[1]<pivot;{
			l++
		}
		for ;array[r]>pivot;{
			r--
		}
		if l>=r{
			break
		}
		array[l],array[r] = array[r],array[l]
		//优化
		if array[l]==pivot{
			r--
		}
		if array[r]==pivot{
			l--
		}
	}
	if l==r{
		l++
		r--
	}

	if left<r{
		quickSort(left, r, array)
	}
	if right>l{
		quickSort(l, right, array)
	}

}

func main() {
	arr := [6]int{-9,87,78,66,12}
	quickSort(0, len(arr) - 1, &arr)
}