package main
import (
	"fmt"
	"time"
	"errors"
)

func test(){
	defer func() {
		err := recover()
		if err != nil{
			fmt.Println(err)
			fmt.Println("发邮件发邮件")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)	
}

func readConf(name string)(err error){
	if name == "config.ini"{
		return nil
	}else{
		return errors.New("读取文件错误")
	}
}

func test02(){
	err :=readConf("onfig.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02()....")
}



func main(){
	// test()

	// for i:=0;i<10;i++{
	// 	fmt.Println("main()______")
	// 	time.Sleep(time.Second)
	// }
	test02()
	for i:=0;i<10;i++{
		fmt.Println("main()______")
		time.Sleep(time.Second)
	}
}