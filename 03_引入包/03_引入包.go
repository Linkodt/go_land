package main

import (
	"fmt"
	// 为了使用utils.go文件的变量或者函数，我们需要引入该包
	"../model_1"
)

func main(){

	var i int =  10
	i = 30
	i = 50

	fmt.Println("i=",i)

	// 包内首字母大写的变量为全局变量，可以被引用并使用
	fmt.Println(model_1.HeroName)
	num1 := 1.2
	num2 := 2.3
	var operator byte = '+'
	result := model_1.Cal(num1,num2,operator)
	fmt.Println(result)
}