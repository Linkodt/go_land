package main

import (
	"fmt"
	"math"
)

func main() {
	p := 11
	q := 3
	N := 0
	F := 0
	e := 0
	d := 0
	c := 0
	m := 0
	N = p * q
	F = (p - 1) * (q - 1)
	for i := 2; i < N; i++ {
		if (F+1)%i == 0 {
			e = i
			break
		}
	}
	d = (F + 1) / e
	// 公钥(N,e), 私钥(N,d)
	// 公钥(33,3), 私钥(33,7)
	m = 31
	// 1. 计算c 加密
	c = int(math.Pow(float64(m), float64(e))) % N
	fmt.Println("c = ", c)

	// 2. 计算m  解密
	m = int(math.Pow(float64(c), float64(d))) % N
	fmt.Println("m = ", m)
}
