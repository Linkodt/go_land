package main
// os.Args 是一个string的切片，用来存储所有的命令行参数
import (
	"fmt"
	"os"
)


func main() {
	fmt.Println("命令行的参数有",len(os.Args))
	for i,v := range os.Args{
		fmt.Printf("args[%v]=%v\n",i,v)
	}
	/*
		D:\文档储存\Go_project\23_cmdline>main.exe 11 string www.baidu.com ok
		命令行的参数有 5
		args[0]=main.exe
		args[1]=11
		args[2]=string
		args[3]=www.baidu.com
		args[4]=ok
	*/
}