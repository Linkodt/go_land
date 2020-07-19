package main

import (
	"fmt"
	"reflect"
)
// monster struct
type Monster struct{
	Name string  `json:"name"`
	Age int `json:"monster_age"`
	Score float32
	Sex string
}
func (s Monster) Print(){
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end-----")
}
func (s Monster) GetSum(n1,n2 int) int{
	return n1+n2
}
func (s Monster) Set(name string,age int,score float32,sex string){
	s.Name = name
	s.Age = age
	s.Sex = sex
	s.Score = score
}

func TestStruct(a interface{}){
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)

	kd:=val.Kind()
	if kd != reflect.Struct{
		fmt.Println("expect struct")
		return 
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num) // 4
	// 遍历
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v 类型=%T\n",i, val.Field(i),val.Field(i))  // val.Field() 获取值
		tagVal := typ.Field(i).Tag.Get("json") // typ.Field() != val.Field()
		if tagVal != ""{
			fmt.Printf("Field %d: tag为=%v\n",i,tagVal)  // 获取tag标签
		}
		
	}
	// 获取方法数目
	numOfmethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfmethod) // 方法排序按照ascii码

	val.Method(1).Call(nil) // 获取第二个方法并调用 传了nil进去

	// 第一个方法
	var params []reflect.Value  // 切片
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) // 调用传值
	fmt.Printf("res=%v res=%T res=%v",res,res,res[0].Int()) // 返回[]reflect.Value 返回切片
}

func main() {
	var a Monster = Monster{
		Name:"小王",
		Age:400,
		Score:30.8,
	}
	TestStruct(a)

}