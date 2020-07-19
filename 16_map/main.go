package main

import (
	"fmt"
)

func main(){
	var a map[string]string = make(map[string]string, 2)
	a["no1"] = "大大"
	a["no2"] = "小小"
	a["no4"] = "small" // 自动扩容
	fmt.Println(a)
	a["no1"] = "种种" //修改
	fmt.Println(a)
	delete(a, "no2") // 删除 不存在时不报错
	fmt.Println(a)
	a["no3"] = "test"
	for k, v := range a{
		fmt.Println(k,v)
	}
	fmt.Println(a)
}