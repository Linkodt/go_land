package main
import(
	"fmt"
	"errors"
)

// 数组模拟一个栈
type Stack struct{
	MaxTop int // 表示我们栈最大的存储量
	Top int // 栈顶，栈顶固定，使用Top
	arr [5]int  // 数组模拟栈
}
// 入栈
func (this *Stack) Push(val int) (err error){
	// 判断栈是否满
	if this.Top ==this.MaxTop-1{
		fmt.Println("已经满了。。。。")
		return errors.New("this stack full")
	}
	this.Top++
	// 放入数组
	this.arr[this.Top] = val
	return 
}

// 出栈
func (this *Stack) Pop()(val int, err error){
	if this.Top==-1{
		fmt.Println("this stack is empty")
		return 0,errors.New("stack empty")
	}

	// 先取值，再this.Top--
	val =  this.arr[this.Top]
	this.Top--
	return val,nil
}


// 遍历栈
func(this *Stack) ListStack(){
	if this.Top == -1{
		fmt.Println("this stack is empty")
		return 
	}
	help := this.Top
	for ;help>=0;{
		fmt.Printf("数是:%d\n",this.arr[help])
		help--
	}
	return 
}

func main(){
	stack := &Stack{
		MaxTop:5, //最多5个
		Top:-1,  // 默认空 -1
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.ListStack()

	val,err :=stack.Pop()
	if err == nil {
		fmt.Printf("栈顶:%d\n", val)
	}else{
		fmt.Printf("err:%v\n", err)
	}
	stack.ListStack()
}