package main
// 互斥锁
import (
	"fmt"
	"time"
	"sync"
)
// 计算1-200各个数的阶乘，放入map中
var(
	mymap = make(map[int]int,10)
	// 全局互斥锁
	// lock是一个全局的互斥锁
	// sunc 包名 同步
	// Mutex : 互斥
	lock sync.Mutex
)

func test(n int){
	// 计算加
	num:=0
	for i:=1;i<=n;i++{
		num += i
	}
	lock.Lock() // 锁
	mymap[n] = num
	lock.Unlock() // 解锁
}

func main() {
	for i:=1;i<=50;i++{
		go test(i)
	}
	time.Sleep(time.Second*10)
	lock.Lock()
	for i,v :=range mymap{
		fmt.Printf("map[%v]=%v\n",i,v)
	}
	lock.Unlock()
}