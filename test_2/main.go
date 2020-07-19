package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func(m *Monster)Store() bool{
 	data,err := json.Marshal(m)
 	if err != nil {
 		fmt.Println("marshal error!",err)
 		return false
 	}
 	// 保存文件
 	err = ioutil.WriteFile("test_2.txt", data, 0666)
 	if err !=nil{
 		fmt.Println("WriteFile error!", err)
 		return false
 	}
 	return true

}

func(m *Monster)Restore()bool{
	content,err := ioutil.ReadFile("test_2.txt")
	if err != nil {
	 	fmt.Println("ReadFile error!",err)
	 	return false
	}
	err = json.Unmarshal(content, m)
	if err != nil {
	 	fmt.Println("Unmarshal error!",err)
	 	return false
	}
	return true
}
