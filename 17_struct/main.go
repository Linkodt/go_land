package main

import (
	"fmt"
	"encoding/json"
)
// json.Marshal 函数中使用反射

type Cat struct{
	Name string
	Age int
	Color string
	Hobby string
}

type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}
// 小写不可被包引用
type Monster struct{
	Name string   `json:"name"`
	Age int       `json:"age"`
	Skill string  `json:"skill"`
}

func main() {
	var cat1 Cat
	cat1.Name = "bai"
	cat1.Age = 3
	cat1.Color = "white"
	cat1.Hobby = "eat"
	fmt.Println(cat1)

	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}
	// 使用slice时需要用make分配内存！！！
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	// 使用map时需要用make分配内存！！！
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom~"

	fmt.Println(p1)

	var p3 *Person = new(Person)
	p3.Name = "smith" // <==> (*p3).Name = "smith"
	p3.Age = 30 // <==>(*p3).Age = 30

	fmt.Println((*p3))

	var p4 *Person = &Person{}
	p4.Name = "smith" // <==> (*p3).Name = "smith"
	p4.Age = 30 // <==>(*p3).Age = 30
	p4.Scores = [...]float64{11,22,33,44,55}

	fmt.Println((*p4))


	var p5 Person
	p5.Age = 10
	p5.Name = "小明"
	var p6 *Person = &p5
	fmt.Println((*p6).Name) // 小明
	p6.Name = "tom"
	fmt.Println(p5) // tom



	// monster
	var nmw = Monster{"牛魔王", 500, "芭蕉扇~"}
	// 序列化为json格式字符串
	jsonStr,err := json.Marshal(nmw)
	if err != nil{

	}else{
		fmt.Println("error")
	}
	fmt.Println(nmw)
	fmt.Println("jsonStr")
	fmt.Println(string(jsonStr))

}