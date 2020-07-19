package main

import (
	"fmt"
)

func test(arr *[5]int) {
	fmt.Println(*arr)
	temp := 0

	for i:=0;i<len(arr) - 1;i++{
		for j:=0;j<len(arr) - 1 - i;j++{
			if (*arr)[j]>(*arr)[j+1]{

				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
				
			}
		}
	}
	fmt.Println(*arr)
}

func main() {
	arr := [5]int{12,2,55,1,9}
	test(&arr)

	fmt.Println(arr)

}