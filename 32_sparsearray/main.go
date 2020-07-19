package main
import(
	"fmt"
	"os"
)

type ValNode struct {
	row int
	col int
	val int
}

func main(){
	// 1.创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 1黑棋
	chessMap[2][3] = 2 // 2蓝子

	// 2.输出看看原始的数组
	for _, v := range chessMap{
		for _, v2 := range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	// 3.转成稀疏数组。想->算法
	// 思路
	// 1.遍历chessMap，发现元素的值不为0，创建一个node结构体
	// 2.将其放入对应切片
	var sparseArr []ValNode

	// 标准的一个稀疏数组头应该是二维数组的规模
	valNode := ValNode{
					row:11,
					col:11,
					val:0,
				}
	sparseArr = append(sparseArr, valNode)

	for i,v := range chessMap{
		for j,v2 :=range v{
			if v2 != 0{
				// 创建一个ValNode值结点
				valNode = ValNode{
					row:i,
					col:j,
					val:v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	// 输出
	for i,valNode := range sparseArr{
		fmt.Printf("%d:%d %d %d\n",i,valNode.row,valNode.col,valNode.val)
	}
	// 存盘稀疏数组  chessmap.data


	// 恢复原始数组
	// 1.创建原始数组
	var chessMap2 [11][11]int

	// 遍历稀疏数组[文件每一行]
	for i,valNode := range sparseArr{
		if  i >= 1 {  // 跳过第一行
		chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
}