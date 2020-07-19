package main
// 环形队列
import (
	"fmt"
	"errors"
	"os"
)

// 使用结构体管理队列
type CircleQueue struct {
	maxSize int  // 4
	array [5]int // 数组->模拟队列
	head int  // 0 指向队列首
	tail int // 0 指向队列尾巴
}

// 入队列AddQueue(push) GetQueue(pop)
func (this *CircleQueue) Push(val int)(err error){
	if this.IsFull(){
		return errors.New("queue full")
	}
	// tail不含最后一个元素
	this.array[this.tail] = val // 把值给尾部
	this.tail = (this.tail + 1) % this.maxSize
	return 

}

// 出队列
func (this *CircleQueue) Pop() (val int,err error){
	if this.IsEmpty(){
		return 0,errors.New("queue empty")
	}
	// 取出,head是队首,包含队首
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val,nil
}

// 判断环形队列为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail + 1) % this.maxSize == this.head
}

// 判断环形队列是否空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	// !!!!!!!!!!!
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

// 显示队列
func (this *CircleQueue) ListQueue(){
	// 取出当前队列有多少个元素
	fmt.Println("环形队列情况如下：")
	size := this.Size()
	if size == 0{
		fmt.Println("队列空")
	}


	// 设计一个辅助变量
	tempHead := this.head
	for i:=0;i<size;i++{
		fmt.Printf("arr[%d]=%d\t", tempHead,this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println("")
}

func main() {
	queue := &CircleQueue{
		maxSize:5,
		head:0,
		tail:0,
	}
	var key string
	var val int
	for{
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示取出数据")
		fmt.Println("3.输入show显示数据")
		fmt.Println("4.输入exit退出")
		fmt.Scanln(&key)
		switch key{
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println("err=",err.Error())
			}else{
				fmt.Println("ok!")
			}
		case "get":
			fmt.Println("get")
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			}else{
				fmt.Println("从队列中取出了,",val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

