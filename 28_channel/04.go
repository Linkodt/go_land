package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main() {
	allchan := make(chan interface{}, 3)

	allchan<-10
	allchan<-"tom jack"
	cat := Cat{"小花猫",4}
	allchan <- cat

	// 先推出前两个
	<-allchan
	<-allchan

	newCat := (<-allchan).(Cat)
	fmt.Printf("newCat=%T newCat=%v\n", newCat,newCat)
	fmt.Printf("newCat.Name=%v", newCat.Name)
}