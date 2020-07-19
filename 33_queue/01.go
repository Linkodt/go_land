package main
// 单向队列
import (
	"fmt"
	"errors"
	"os"
)

// 使用结构体管理队列
type Queue struct {
	maxSize int
	array [5]int // 数组->模拟队列
	front int  // 指向队列首,不包含第一个
	rear int // 指向队列尾巴
}

func (this *Queue) AddQueue(val int) (err error) {
	// 判断队列是否已经满了
	if this.rear == this.maxSize - 1{
		// 满了
		return errors.New("queue full")
	}

	this.rear++  // 后移
	this.array[this.rear] = val
	return 

}
// 显示队列，找到队首
func (this *Queue) ShowQueue(){
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("arrary[%d]=%d\t",i,this.array[i])
	}
	fmt.Println("")
}

// 取出
func (this *Queue) GetQueue()(val int,err error){
	// 判断是否为空
	if this.front == this.rear{
		// 队列空
		return -1,errors.New("queue empty")
	}
	this.front++
	val=this.array[this.front]
	return val,err
}

func main() {
	queue := &Queue{
		maxSize:5,
		front:-1,
		rear:-1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("err=",err.Error())
			}else{
				fmt.Println("ok!")
			}
		case "get":
			fmt.Println("get")
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			}else{
				fmt.Println("从队列中取出了,",val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}