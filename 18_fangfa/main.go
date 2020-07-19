package main

import (
	"fmt"
)

type A struct {
	Name string  `json:"name"`
}

func (a A) test(){
	fmt.Println(a.Name)
}

type Person struct {
	Name string `json:"name"`
	OtherName string
}

func (p Person) speak(){   // 记住格式 func (<name> struct) speak(){}
	fmt.Printf("%v say %v is a good man~\n",p.Name,p.OtherName)
	p.Name = p.OtherName
	fmt.Println(p)  //真实的不变 可用结构体指针
}

func (p *Person) test(){  // 结构体指针！ 可以改变真实值
	(*p).Name = (*p).OtherName
	fmt.Println(*p)
}

type intergra int
func (i *intergra)change() { // 方法首字母大小写也很重要
	*i = 10
}
func (i *intergra) String() string{
	str := fmt.Sprintf("这是我们自定义的fmt.Println的打印模型,%v",*i)
	return str
}

func main() {
	var a A
	a.Name = "tom~"
	a.test()
	var p = Person{"tom","jack~"}
	p.speak()
	fmt.Println(p)
	p.test()
	fmt.Println(p)
	var i intergra = 5
	fmt.Println(i) //5
	i.change()
	fmt.Println(i)  //10
	fmt.Println(&i)  // 打印了自定义语句 String
}