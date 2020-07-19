package main

import (
	"fmt"
)

// 声明定义一个接口
type Usb interface{  // 接口
	// 声明了两个没有实现的方法
	Start()
	Stop()
}


type Phone struct {
	
}

func (p Phone) Start(){
	fmt.Println("手机开始工作...")
}
func(p Phone) Stop(){
	fmt.Println("手机停止工作...")
}
func (p Phone) Call(){
	fmt.Println("手机打电话了打电话")
}

type Camera struct {
	
}
func (c Camera) Start(){
	fmt.Println("相机开始工作...")
}
func(c Camera) Stop(){
	fmt.Println("相机停止工作...")
}

type Computer struct{

}
// working方法接收了一个Usb接口类型的变量
// 只要是实现了Usb接口 (所谓实现Usb接口，就是指实现了 Usb接口声明的所有方法)
func (c Computer) Working(usb Usb){
	usb.Start()
	// 如果usb是指向phone结构体变量的，则还需要调用Call方法
	// 类型断言!!!!!!!
	if phone, ok := usb.(Phone);ok{
		phone.Call()
	}else{
		fmt.Println("errrrrrrr是相机")
	}
	usb.Stop()
}


func main() {
	// var相关解构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	camera_2 := Camera{}
	var b Usb = camera_2  // 方式2
	
	b.Start()
	b.Stop()

	// 关键
	computer.Working(phone)
	computer.Working(camera)
}