package main
//约瑟夫
import (
	"fmt"
)

type Boy struct {
	No int
	Next *Boy
}

func AddBoy(num int) *Boy {//构成单向的环形链表
	//判断
	first := &Boy{}
	curBoy:=&Boy{}
	if num<1{
		fmt.Println("num的值不对")
		return first
	}
	// 循环构建这个环形链表
	for i:=1;i<=num;i++{
		boy:=&Boy{
			No:i,
		}

		if i==1{
			first = boy // 不动
			curBoy = boy
			curBoy.Next = first
		}else{
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

// 显示环形链表
func ShowBoy(first *Boy){

	if first.Next == nil{
		fmt.Println("链表为空，没有小孩。。。")
		return 
	}


	curBoy := first

	for{
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		if curBoy.Next == first{
			break
		}else{
			curBoy = curBoy.Next
		}
	}
}

func PlayGame(first *Boy,startNo int,countNum int){
	// 定义辅助指针 帮助删除小孩
	// 1.empty
	if first.Next == nil{
		fmt.Println("empty link")
		return 
	}
	// startNo<小孩总数
	tail := first
	// tail->尾巴
	for{
		if tail.Next==first{
			break
		}
		tail = tail.Next
	}
	for i:=1;i<=startNo - 1;i++{
		first = first.Next
		tail = tail.Next
	}
	for{
		// 数countNum-1下
		for i:=1;i<=countNum -1;i++{
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈->",first.No)
		// 删除first的小孩
		first = first.Next
		tail.Next = first
		//tail==first,圈子只有一个小孩
		if tail==first{
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈", tail.No)
}


func main() {
	first := AddBoy(5)
	// 显示
	// ShowBoy(first)
	PlayGame(first, 2, 3)
}