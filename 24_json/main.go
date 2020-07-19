package main
// 序列化
import (
	"fmt"
	"encoding/json"
)
type Monster struct {
	Name string `json:"monster_name"`
	Age int
	Brithday string
	Sal float64
	Skill string
}

func testStruct(){
	// 演示 结构体序列化
	monster := Monster{
		Name:"牛魔王",
		Age :500,
		Brithday : "2011-11-11",
		Sal:8000.0,
		Skill:"喷火",
	}

	data,err := json.Marshal(monster)
	if err !=nil{
		fmt.Println("序列化失败",err)
	}
	// 输出序列化后的结果
	fmt.Println("monster序列化后=",string(data))
}

func testMap(){
	var a map[string]interface{}
	// 使用map之前需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "宝莲山"

	data,err := json.Marshal(a)
	if err !=nil{
		fmt.Println("序列化失败",err)
	}
	// 输出序列化后的结果
	fmt.Println("a map[string]interface{}序列化后=",string(data))
}

// 切片序列化
func testSlice(){
	/*
		slice []map[string]interface{}序列化后= [{"address":["北京","天津"],"age":"7","name":"jack"},{"address":["广州","上海"],"age":"9","name":"tom"},{"address":["北京","天津"],"age":"20","name":"alsy"}]
	*/
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = [2]string{"北京","天津"}
	m2["name"] = "tom"
	m2["age"] = "9"
	m2["address"] = [2]string{"广州","上海"}
	m3["name"] = "alsy"
	m3["age"] = "20"
	m3["address"] = [2]string{"北京","天津"}
	slice = append(slice,m1,m2,m3)

	data,err := json.Marshal(slice)
	if err !=nil{
		fmt.Println("序列化失败",err)
	}
	// 输出序列化后的结果
	fmt.Println("slice []map[string]interface{}序列化后=",string(data))
}

func testFloat64(){ // 23.64
	var num1 float64 = 23.64
	data,err := json.Marshal(num1)
	if err !=nil{
		fmt.Println("序列化失败",err)
	}
	// 输出序列化后的结果
	fmt.Println("num1序列化后=",string(data))
}

func main() {
	// 序列化结构体、map、切片
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}