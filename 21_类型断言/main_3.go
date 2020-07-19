package main

import (
	"fmt"
)

func TypeJugde(items ...interface{}){
	for i, x:=range items{
		switch x.(type){
			case int,int32,int64 :
				fmt.Printf("%v--int--%v\n",i, x)
			case float32,float64 :
				fmt.Printf("%v--float--%v\n",i, x)
			case nil:
				fmt.Printf("%v--nil--%v\n",i, x)
			case string:
				fmt.Printf("%v--string--%v\n",i, x)
			case bool:
				fmt.Printf("%v--bool--%v\n",i, x)
			default :
				fmt.Printf("%v--undefine???--%v\n",i, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 3.2
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300

	TypeJugde(n1,n2,n3,name,address,n4)
}