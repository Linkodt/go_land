package main

import (
	"fmt"
)

type Emp struct{
	Id int
	Name string
	Next *Emp
}


// EmpLink
type EmpLink struct{
	Head *Emp

}
// insert,按id从小到大
func (this *EmpLink) Insert(emp *Emp){
	cur := this.Head
	var pre *Emp = nil // 保证在cur之前
	if this.Head == nil{
		this.Head = emp
		return 
	}
	// 不是空链表,给emp找到对应的位置并插入
	for{
		if cur != nil{
			if cur.Id > emp.Id{
				// 找到位置 
				break
			}
			pre = cur
			cur = cur.Next
		}else{
			break
		}
	}
	// 退出for
	if pre == nil{
		help := this.Head.Next
		this.Head = emp
		emp.Next = help
	}else{
		pre.Next = emp
		emp.Next = cur
	}


}

// 显示信息
func (this *EmpLink) ShowLink(no int){
	if this.Head == nil{
		fmt.Printf("链表%d为空\n", no)
		return 
	}
	fmt.Printf("链表id:%d", no)
	cur := this.Head
	for{
		if cur != nil{
			fmt.Printf("\t雇员id=%d 名字=%s\t->",cur.Id,cur.Name)
			cur = cur.Next
		}else{
			break
		}
	}
	fmt.Println("")
}



// hashtable,含有一个链表数组
type HashTable struct{
	LinkArr [7]EmpLink
}

// 编写insert雇员的方法
func(this *HashTable) Insert(emp *Emp){
	linkno := this.HashFun(emp.Id)
	this.LinkArr[linkno].Insert(emp)
}

// 编写一个散列方法
func (this *HashTable) HashFun(id int) int{
	return id%7
}


// 编写方法 显示hashtable所有的雇员
func (this *HashTable) ShowAll(){
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func main() {
	key := ""
	id := 0
	name := ""
	flag := true
	var hashtable HashTable
	for{
		fmt.Println("==========雇员菜单=========")
		fmt.Println("input 添加雇员")
		fmt.Println("show  显示雇员")
		fmt.Println("find  查找雇员")
		fmt.Println("exit  退出系统")
		fmt.Println("请输入你的选择:")
		fmt.Scanln(&key)

		switch key{
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp:=&Emp{
				Id:id,
				Name:name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "exit":
			flag = false
		case "find":
			fmt.Println("未开发功能")
		default:
			fmt.Println("输入错误")
		}
		if !flag{
			break
		}
	}


}