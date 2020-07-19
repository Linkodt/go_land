package main
import(
	"fmt"
	"errors"
	"strconv"
)

// 数组模拟一个栈
type Stack struct{
	MaxTop int // 表示我们栈最大的存储量
	Top int // 栈顶，栈顶固定，使用Top
	arr [20]int  // 数组模拟栈
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
// 判断是否空栈
func (this *Stack) Isempty()bool{
	if this.Top == -1{
		return true
	}else{
		return false
	}
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

func(this *Stack) Isopera(val int)bool{
	if val==42||val==43||val==45||val==47{
		return true
	}else{
		return false
	}
}

func(this *Stack) opera(num1,num2,oper int)int{
	res:=0
	switch oper{
	case 42:
		res = num2*num1
	case 43:
		res = num1+num2
	case 45:
		res = num2 - num1
	case 47:
		res = num2/num1
	default:
		fmt.Println("运算符出错")
	}
	return res
}
// 编写一个方法，返回某个运算符的优先级
// * / => 1, + - => 0
func(this *Stack) Priority(oper int)int{
	res:=0
	if oper == 42 || oper == 47{
		res = 1
	}else if oper==43 || oper==45{
		res = 0
	}
	return res
} 

func main(){
	num_stack := &Stack{
		MaxTop:20, //最多5个
		Top:-1,  // 默认空 -1
	}
	opra_stack := &Stack{
		MaxTop:20,
		Top:-1,
	}
	exp:= "3+2*6-2"
	// 定义一个index，帮助扫描exp
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	res := 0
	for{
		ch := exp[index:index+1]  // "3" 字符串
		temp := int([]byte(ch)[0])  // 字符对应的ASCII
		if opra_stack.Isopera(temp){ // 是否操作符
			if opra_stack.Isempty(){
				opra_stack.Push(temp)
			}else{
				if opra_stack.Priority(opra_stack.arr[opra_stack.Top])>=opra_stack.Priority(temp){
						num1,_ = num_stack.Pop()
						num2,_ = num_stack.Pop()
						oper,_ = opra_stack.Pop()
						res = opra_stack.opera(num1, num2, oper)
						num_stack.Push(res)
						opra_stack.Push(temp)
					}else{
						opra_stack.Push(temp)
					}
			}
		}else{  // 说明是数
			val,_ := strconv.ParseInt(ch, 10, 64)
			num_stack.Push(int(val))
		}
		if index + 1 ==len(exp){
			break
		}
		index++
	}
	for{
		if opra_stack.Top==-1{
			break
		}
			num1,_ = num_stack.Pop()
			num2,_ = num_stack.Pop()
			oper,_ = opra_stack.Pop()
			res = opra_stack.opera(num1, num2, oper)
			num_stack.Push(res)
	}

	// 如果我们的算法没有问题，表达式也是正确的，则结果就是numStack最后数
	res,_ = num_stack.Pop()
	fmt.Printf("表达式%s = %v",exp,res)
}