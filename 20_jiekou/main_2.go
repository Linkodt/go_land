package main

import (
	"fmt"
	"sort"
	"math/rand"
)

// 声明hero结构体
type Hero struct {
	Name string
	Age int
}

// 声明hero结构体切片类型
type HeroSlice []Hero

// 实现Interface接口
func (hs HeroSlice) Len() int{
	return len(hs)
}
// 决定你使用什么标准进行排序 按age从小到大排序
func (hs HeroSlice) Less(i,j int) bool{
	return hs[i].Age < hs[j].Age
	// return hs[i].Name < hs[j].Name
}
func (hs HeroSlice) Swap(i,j int){
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}


func  main() {
	// 先定义一个数组/切片
	var intslice = []int{0,-1,10,7,90}
	// 排序
	// 1. 冒泡排序 。。。
	// 2. 系统函数排序
	sort.Ints(intslice)
	fmt.Println(intslice)  // 有序

	// 编写对一个结构体切片进行排序
	var heroes HeroSlice
	for i:=0;i<10;i++{
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d", rand.Intn(100)),  // 返回0-100随机数
			Age : rand.Intn(100),
		}
		// 讲hero append 到切片
		heroes = append(heroes, hero)
	}
	for _, v := range heroes{
		fmt.Println(v)
	}
	 // 调用sort包下面的Sort方法
	 fmt.Print("排序后。。。。。。。。。。。。。。\n")
	 sort.Sort(heroes) // 因为完成了别人定义好的接口
	 for _, v := range heroes{
		fmt.Println(v)
	}
}