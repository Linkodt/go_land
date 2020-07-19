package main

import (
	"fmt"
)

type Stu struct {
	num int
}

func (s Stu) say(){
	fmt.Println("-------",s.num,"-------")
}
type Ainterface interface{
	say()
}

func main() {
	var stu Stu = Stu{num:3}
	var a Ainterface = stu
	stu.num = 4
	fmt.Println(stu)  // 4
	fmt.Println(a)
	a.say()  // 打印的还是3
}
