package main

import (
	"fmt"
)

func main() {
	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	totalWeight := 0.0
	for i := 0; i<len(hens);i++{
		totalWeight += hens[i]
	}

	avgWeight := fmt.Sprintf("%.2f", totalWeight / float64(len(hens))) //返回.2f值

	fmt.Printf("totalWeight=%v avgWeight=%v\n", totalWeight, avgWeight)

	for index,value := range hens{ // 遍历
		fmt.Println(index,value)
	}
}