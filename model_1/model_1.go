package model_1

import (
	"fmt"
)
// 定义一个变量
var HeroName string = "宋江"

func Cal(i float64,n float64,operator byte) float64 {  // 大写

	var res float64
	switch operator{
	case '+':
		res = i + n
	case '-':
		res = i - n
	case '*':
		res = i * n
	case '/':
		res = i / n
	default:
		fmt.Println("操作符号错误")
	}
	return res
}