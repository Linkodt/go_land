package main
// 基本类型反射
import (
	"fmt"
	"reflect"
)

// 专门演示反射
func reflectTest01(b interface{}){
	// 通过反射获取的传入的变量的type，kind，值
	rtyp := reflect.TypeOf(b)
	fmt.Println("rtype=",rtyp)

	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("rVal=",rVal)
	fmt.Println(n2)
	fmt.Printf("type=%T\n", rVal)


	// 下面将rVal转成interface{}
	iV := rVal.Interface()
	// interface{}断言成需要的
	num2:=iV.(int)
	fmt.Printf("num2=%v", num2)
}

func reflectTest02(b interface{}){
	// 对结构体反射
	rtyp := reflect.TypeOf(b)
	fmt.Println("rtype=",rtyp)

	rVal := reflect.ValueOf(b)
	// n2 := 2 + rVal.Int()
	fmt.Println("rVal=",rVal)
	// fmt.Println(n2)
	fmt.Printf("type=%T\n", rVal)

	fmt.Printf("kind1=%v kind2=%v\n", rVal.Kind(),rtyp.Kind())
	// 下面将rVal转成interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n",iV,iV)

	stu, ok :=iV.(Student)
	// 另一种方法断言----switch
	if ok{
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

type Student struct {
	Name string
	Age int
}

func main() {
	
	// var num int = 100
	// reflectTest01(num)
	var Stu Student = Student{"小王",10}
	reflectTest02(Stu)
}

/*
var student Stu
var num int

func test(b interface{}){
	// interface->reflect.Value
	rVal := relect.ValueOf(b)

	// relect.Value -> interface{}
	iVal := rVal.interface()

	// interface -> b
	v := iVal.(Stu)
}

*/