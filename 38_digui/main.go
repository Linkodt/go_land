package main

import (
	"fmt"
)

func Setway(Map *[8][7]int,i int,j int) bool {
	// 分析情况
	// myMap[6][5] == 2找到
	if Map[6][5] == 2{
		return true
	}else{
		// 继续找
		if Map[i][j] == 0{ // 可以探测

			// 假设这个点是可以通的,需要探测, 上下左右
			// 换成下右上左
			Map[i][j] = 2
			if Setway(Map, i + 1, j){
				// 通了
				return true
			}else if Setway(Map, i, j+1){
				return true
			}else if Setway(Map, i-1, j){
				return true
			}else if Setway(Map, i, j-1){
				return true
			}else{// 死路
				Map[i][j] = 3
				return false
			}
		}else{ // 1 or 3
			return false
		}
	}
}

func main() {
	// 模拟迷宫
	// 0是没有走过的
	// 1为墙
	// 2是通路
	// 3是走过的点但是走不通
	var Map [8][7]int = [8][7]int{
								{1,1,1,1,1,1,1},
								{1,0,0,0,0,0,1},
								{1,0,0,0,0,0,1},
								{1,1,1,0,0,0,1},
								{1,0,0,1,0,0,1},
								{1,0,0,0,0,0,1},
								{1,0,0,0,0,0,1},
								{1,1,1,1,1,1,1},
							}
	for i := 0;i<8;i++{
		fmt.Print("  ")
		for j:=0;j<7;j++{
			fmt.Print(Map[i][j]," ")
		}
		fmt.Println()
	}
	fmt.Println()
	Setway(&Map, 1, 1)

	for i := 0;i<8;i++{
		fmt.Print("  ")
		for j:=0;j<7;j++{
			fmt.Print(Map[i][j]," ")
		}
		fmt.Println()
	}

}