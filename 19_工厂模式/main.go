package main

import (
	"fmt"
	"../model_2"
)

func main() {
	var stu = model_2.Student_2{
		Name : "tom",
	}
	fmt.Println(stu)
	stu.SetCore(15)
	fmt.Println(stu.Getscore())

	var stu_2 = model_2.NewStudent("tom~", 88.8)
	fmt.Println(stu_2)
	fmt.Println(*stu_2)
}