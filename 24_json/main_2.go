package main
// 反序列化
import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string
	Age int
	Brithday string
	Sal float64
	Skill string
}
// 1. 反序列化成结构体
func unmarshalStruct(){
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Brithday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"喷火\"}"

	// 定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster) // 字节切片，改变哪一个实例
	if err != nil {
		fmt.Println(err) 
	}
	fmt.Println("反序列化=",monster)
}

func unmarshalMap(){
	str := "{\"address\":\"宝莲山\",\"age\":30,\"name\":\"红孩儿\"}"

	//定义一个map
	var a map[string]interface{}
	// a = make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &a) // 字节切片，改变哪一个实例
	if err != nil {
		fmt.Println(err) 
	}
	fmt.Println("反序列化=",a)
}

func unmarshalSlice(){
	str := "[{\"address\":[\"北京\",\"天津\"],\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"广州\",\"上海\"],\"age\":\"9\",\"name\":\"tom\"},{\"address\":[\"北京\",\"天津\"],\"age\":\"20\",\"name\":\"alsy\"}]"
	var slice []map[string]interface{}
	// 反序列化，不需要make
	err := json.Unmarshal([]byte(str), &slice) // 字节切片，改变哪一个实例
	if err != nil {
		fmt.Println(err) 
	}
	fmt.Println("反序列化=",slice)
}


func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}