package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// 生成密钥文件
func GenerateRSAKey(bits int) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
		return
	}
	// 保存私钥
	// 通过x509标准得到ras私钥 序列化为ASN.1的DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
		return
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	err = pem.Encode(privateFile, &privateBlock)
	if err != nil {
		panic(err)
		return
	}
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	// X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// pem格式编码
	// 创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	// 创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	// 保存到文件
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		panic(err)
	}

}

func GetRSAPrivateKey(path string) *rsa.PrivateKey {
	// 读取文件内容
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem编码
	block, _ := pem.Decode(buf)
	// X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return privateKey
}

// 读取RSA公钥
func GetRSAPublicKey(path string) *rsa.PublicKey {
	// 读取公钥内容
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem解码
	block, _ := pem.Decode(buf)
	// x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey
}

// 对消息的散列值进行数字签名
func GetSign(msg []byte, path string) []byte {
	// 取得私钥
	publicKey := GetRSAPublicKey(path)
	//SignPKCS1v15使用RSA PKCS#1 v1.5规定的RSASSA-PKCS1-V1_5-SIGN签名方案计算签名
	sign, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg)
	if err != nil {
		panic(err)
	}
	return sign
}

// 验证数字签名
func VerifySign(sign []byte, path string) ([]byte, error) {
	// 取得私钥
	privateKey := GetRSAPrivateKey(path)
	// 验证数字签名
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, sign)
}

func main() {
	// 生成密钥文件
	GenerateRSAKey(2048)
	// 加密数字
	input := []byte("www.baidu.com")
	sign := GetSign(input, "public.pem")
	a, _ := VerifySign(sign, "private.pem")
	fmt.Println(string(a))
}
