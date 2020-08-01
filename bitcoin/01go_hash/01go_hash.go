package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func Hash_md5(input []byte) { // 默认32位小写
	hash := md5.New()           // 此时hash是hash.Hash接口
	_, err := hash.Write(input) // 通过嵌入的匿名io.Writer接口的Write方法向hash中添加更多数据，永远不返回错误
	if err != nil {
		fmt.Println(err)
	}
	result := hash.Sum(nil) // 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	md52len := hash.Size()  // 返回Sum会返回的切片的长度
	block2md5 := hash.BlockSize()
	// 返回hash底层的块大小；Write方法可以接受任何大小的数据，
	// 但提供的数据是块大小的倍数时效率更高
	fmt.Printf("md5:%x\n LenOfmd5:%x   blocksizeOfmd5:%x\n", result, md52len, block2md5)
}

func Hash_sha1(input []byte) {
	hash := sha1.New()
	_, err := hash.Write(input)
	if err != nil {
		fmt.Println(err)
	}
	result := hash.Sum(nil)
	fmt.Printf("sha1:%x\n", result)
}

func Hash_sha256(input []byte) {
	hash := sha256.New()
	_, err := hash.Write(input)
	if err != nil {
		fmt.Println(err)
	}
	result := hash.Sum(nil)
	fmt.Printf("sha256:%x\n", result)
}

func Hash_sha512(input []byte) {
	hash := sha512.New()
	_, err := hash.Write(input)
	if err != nil {
		fmt.Println(err)
	}
	result := hash.Sum(nil)
	fmt.Printf("sha512:%x\n", result)
}

func Hash_ripemd160(input []byte) {
	hash := ripemd160.New()
	_, err := hash.Write(input)
	if err != nil {
		fmt.Println(err)
	}
	result := hash.Sum(nil)
	fmt.Printf("ripemd160:%x\n", result)
}

func main() {
	input := "123456"
	Hash_md5([]byte(input))
	Hash_sha1([]byte(input))
	Hash_sha256([]byte(input))
	Hash_sha512([]byte(input))
	Hash_ripemd160([]byte(input))
}
