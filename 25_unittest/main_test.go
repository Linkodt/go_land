package main

import (
	_ "fmt"
	"testing" // go testing框架包
)

// 测试addUpper是否正确
func TestAddUpper(t *testing.T){
 res := addUpper(10)
 if res!=55{
 		// fmt.Println("错误，期望值=",55," 实际值=",res)
 		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v", 55,res)
	}
	t.Logf("AddUpper(10) 执行正确")
}