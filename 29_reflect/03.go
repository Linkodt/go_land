package main

import (
	"reflect"
	"fmt"
)

type user struct {
	Userld string
	Name string
}

func main() {
	var(
		model *user
		st reflect.Type
		elem reflect.Value
	)
	st = reflect.TypeOf(model)
	fmt.Printf("st=%v\n", st)
	st = st.Elem()
	fmt.Printf("st=%v\n", st)
	elem = reflect.New(st)
	fmt.Printf("elem=%v\n", elem)
	model = elem.Interface().(*user) // 类型断言 指向跟elem一致
	fmt.Printf("model=%v\n", model)
	elem = elem.Elem()
	elem.FieldByName("Userld").SetString("12345678")
	elem.FieldByName("Name").SetString("nickname")
	fmt.Printf("elem=%v elem=%T\n", elem,elem)

}