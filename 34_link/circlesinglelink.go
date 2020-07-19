package main

import (
	"fmt"
)

type CatNode struct {
	no int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode){

	// 判断是不是添加第一只猫
	if head.next == nil{
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println("第一只猫")
		return 
	}
	// 不是第一只猫
	temp := head
	for{
		if temp.next == head{
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	temp.next.next = head
}

// 输出
func ListCircleLink(head *CatNode){
	temp := head
	if temp.next == nil{
		fmt.Println("empty")
		return
	}
	for{
		fmt.Println("猫的信息为=",temp.name,temp.no,"->")
		if temp.next == head{
			break
		}
		temp = temp.next
	}
}

func DelCircleLink(head *CatNode,id int)*CatNode{
	temp:=head
	helper := head
	// empty
	if temp.next==nil{
		fmt.Println("can not delect empty link")
		return head
	}
	// one node
	if temp.next==head{
		if temp.no == id{
			temp.next=nil
			fmt.Println("one node delect")
		}
		return head
	}
	// more node
	for{
		if helper.next==head{
			break
		}
		helper = helper.next
	} // helper此时是尾部结点，temp是首结点，所以helper始终在temp前面
	flag := true
	for{
		if temp.next==head{//到最后一个 且最后一个还未比较
			break
		}
		if temp.no==id{
			if temp==head{
				// 删除头结点
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
		}
		temp.next = temp
		helper = helper.next
	}
	// 比较temp.next==head
	if flag{
		// 在上面无删除
		if temp.no==id{
			helper.next = temp.next
		}
	}

	return head
}


func main() {
	// 环形链表的头结点
	head := &CatNode{}
	// 创建一只猫
	Cat1 := &CatNode{
		no:1,
		name:"tom",
	}
	InsertCatNode(head, Cat1)
	head = DelCircleLink(head, 1)
	// DelCircleLink(head, 1)
	ListCircleLink(head)
}