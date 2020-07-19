package main
// 空心金字塔
import (
	"fmt"
)

func main(){

	// 分析编程思路:
	/*
		1. 打印一个矩形
		***
		***
		***
	
	for i := 1; i<=3; i++{
		for j:=1; j<=3; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
	*/
	/*
		2. 打印半个金字塔
		*
		**
		***

		for i:=1;i<=3;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
	*/
	
	/*
		3. 打印整个金字塔
		*     2*1 - 1  空格2   总层数-当前层数
	   ***    2*2 - 1  空格1
	  *****   2*3 - 1  空格0

	var totalLevel int = 9

	for i:=1;i<=totalLevel;i++{
		for k:=1;k<=totalLevel - i;k++{
			fmt.Printf(" ")
		}
		for j:=1;j<=2*i-1;j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
	*/
	/*
		4.打印空心金字塔
		分析 每行打印星号时考虑是打印*还是打印空格
		每层 第一跟最后打星号
		最后一层打星
	*/
	var totalLevel int = 9

	for i:=1;i<=totalLevel;i++{
		for k:=1;k<=totalLevel - i;k++{
			fmt.Printf(" ")
		}
		for j:=1;j<=2*i-1;j++{
			if j==1 || j == 2*i - 1 || i==totalLevel{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
			
		}
		fmt.Println()
	}
}

