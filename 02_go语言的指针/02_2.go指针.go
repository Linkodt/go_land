package main

import "fmt"


func main(){
    var i int = 10

    fmt.Println("i的地址=", &i)
    var ptr *int = &i
    *ptr = 20 // 修改了i的值
    fmt.Println("p的地址=", ptr)
    fmt.Println("p的值=", *ptr)
    fmt.Println("i的值=", i)
}