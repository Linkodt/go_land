package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// 直接声明了 一般是从文件读取的
var send_pub []byte = []byte("-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCgktqIHI6rS+4vcYV20v7tzlA3\nSXPu9TWSdqk0mt9YdpmkrUHb31r51UIbw0C0WQlLNC4r2seWw7jaiL7zFgKSuhtY\no60MSmqpeEkL5EsCkW2MQ8rdcCsAYq0isG2XmKhoYMiMoN1Bpbb5jW5kPrTK4zwW\nAp7cfH18s/KymS88sQIDAQAB\n-----END PUBLIC KEY-----")
var send_pri []byte = []byte("-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCgktqIHI6rS+4vcYV20v7tzlA3SXPu9TWSdqk0mt9YdpmkrUHb\n31r51UIbw0C0WQlLNC4r2seWw7jaiL7zFgKSuhtYo60MSmqpeEkL5EsCkW2MQ8rd\ncCsAYq0isG2XmKhoYMiMoN1Bpbb5jW5kPrTK4zwWAp7cfH18s/KymS88sQIDAQAB\nAoGAD3x+UkqDyXpEHWQC0dYRY1zGoIBna5xPV4H1eXb206yzjn3tuD+LQgVa7/vF\nq6cbJftBOTZLUctsUnpGBUKCCp53ce6DXy7iclbWWuyD/EgOsSBm/gA45WFEAup9\n2kl+KV2Ac/DWogizdfgzc6CQNcqPXsUtgb5OxOTEDjcLBwkCQQCyntwVa/1qyX9i\nXvIv6TaDEsx8naQyi6+w4vaqdo4tsalYp2BeGga6EUH5UWKdgBcq3qZs5y/7DTek\noCZ3ziMJAkEA5iKUoDmRPXpx+x1Ca3TJJKrTLMG5ZtE+0PiGSvCPcPM9ByJvlDww\nwY41btgfujZnabkijCcb63pq5dHbai1uaQJAIw7ZBX9XkuEPloAqKsssPteuuPq7\n4ovWJPK3FUgFqeRH57WaTX3wOg7vEv8tStp8MZbXtWI+4Zh7hl1Ej5ku+QJAHsO2\nrKilctuAizpBG85T6VPwwQXwu/7y78qZYUFBW68YlHCVC/Lz7ZwOJpA7xY/qSSbX\nqHa987+8tJZWR55lKQJAMaW9E0dbGZWco/ZkmsrLaiz68SGBMvJ9tmL32z2kA/0N\ne+pHgqUBxiBhMm41vF/JVQ9zskRw6giDpT9E9jtRBg==\n-----END RSA PRIVATE KEY-----")
var recv_pub []byte = []byte("-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCE7iYx+jHrLc433htRIRRDZDgJ\n2DnSVB/mEIUFP4Bd+x/pnm0087SYgcBkfcHUxzDmPq6hkmhXfNwZ0HietzKWxFf/\n+PxQkMRHAQZhQi8KMlcFH1lSVnu5fFYnLUUMB/tF0/Y4/X6VMXFZ8l8NaI95+dDa\nHXJ5gQRZ3I0rx0bnXwIDAQAB\n-----END PUBLIC KEY-----")
var recv_pri []byte = []byte("-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQCE7iYx+jHrLc433htRIRRDZDgJ2DnSVB/mEIUFP4Bd+x/pnm00\n87SYgcBkfcHUxzDmPq6hkmhXfNwZ0HietzKWxFf/+PxQkMRHAQZhQi8KMlcFH1lS\nVnu5fFYnLUUMB/tF0/Y4/X6VMXFZ8l8NaI95+dDaHXJ5gQRZ3I0rx0bnXwIDAQAB\nAoGAJtj0J5nfGcQupmwN3qVEpzc0G0hqCXoVQmZkgXaf3JdFqb/nWVi5V/JfV97a\nlQTLWRZLcaN4Q80VA6RFg9i0+yWtXYzJIJZirQUGIPOmtmeR2b0iqLq3U1eiMU/S\n6f96tznM0/LlQokPTwaut89G5hliG4PHjb1qf3JC5kE1IW0CQQCXtUYkgWj5DEIC\nUJgNO5oqO80yuWEsakTAluLadPhUiJHhvFi1DI6AGjEdMdIox8Qs+lBVBASiRH/5\n206sjySLAkEA4FA2Sgq6lMScp3cyt6wRJYCXgD6i7rUPWYsxLr2UiofYh6KO9XUT\nGWL5aKo78AUfrabGVZD+u0kBbukPXgae/QJATnRshcx24PS53kCXoaBFJbovcB/2\nZOXAEGj3gEolcnH+H2Mr7NVFjHH/HD37lSNaDHiWuqCNVO1f0XLcyNht6QJBAIBS\nAqyatJ/0VbWLf+0ItwpgJMsNigHKfEx6TxlxBG+3RMxa1zNd1R45u3n/oKgvFi8t\n79n1R3P2UKaVFuBwIzUCQE9Nl9kMKNuR/ihtfG2vsZcAUgqnKDFoMJf+pPxlnPc9\n/x6V/OBlrot0YuAB8zJAk/Ea9snW9qhXx9AUdqdm9Tk=\n-----END RSA PRIVATE KEY-----")

/*
	总结： 要有一个通用的hash加密函数 传入字符串 返回密文
		  要有接收方跟发送方的私密钥
		  要有一个签名函数 传入hash后的摘要 并进行签名
		  要有一个解密签名的函数 得到摘要
*/

// hash加密函数 md5加密
func Hash_md5(input []byte) []byte {
	hash := md5.New()
	_, err := hash.Write(input)
	if err != nil {
		panic(err)
	}
	result := hash.Sum(nil)
	return result
}

// 私钥签名  用发送方的私钥加密字符串（签名）  传入[]byte
func RSASign(data []byte) (string, error) {
	// 解密私钥
	block, _ := pem.Decode(send_pri)
	if block == nil {
		return "", errors.New("privateKey error")
	}
	// 解析PKCS1格式的私钥
	PrivKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	myhash := crypto.MD5
	bytes, err := rsa.SignPKCS1v15(rand.Reader, PrivKey, myhash, data)
	if err != nil {
		return "", err
	}
	//return base64.StdEncoding.EncodeToString(bytes),nil
	return string(bytes), nil
}

// 接收方公钥加密字符串  这个可能出错
func RsaEncrypt(origData []byte) ([]byte, error) {
	// 解密接收方公钥
	block, _ := pem.Decode(recv_pub)
	if block == nil {
		return nil, errors.New("block error(recv_pub)")
	}
	//fmt.Println("block:",block)
	// 解析公钥
	pubinterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubinterface.(*rsa.PublicKey)

	//fmt.Println("pub:",pub)

	if err != nil {
		return nil, err
	}
	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

}

// 用发送方的公钥解密  这个digest是用hash函数弄得
func RSAVerify(data []byte, digest []byte) error {
	// 1. 解析公钥
	block, _ := pem.Decode(send_pub)
	if block == nil {
		return errors.New("send_pri error")
	}
	// 2、选择hash算法，对需要签名的数据进行hash运算
	myhash := crypto.MD5
	// 3、读取公钥文件，解析出公钥对象
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := publicKey.(*rsa.PublicKey)
	// 4、RSA验证数字签名（参数是公钥对象、哈希类型、签名文件的哈希串、签名后的字节）
	return rsa.VerifyPKCS1v15(pub, myhash, digest, data)
}

// 私钥解析密文  出来的应该是hello ketty
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	// 解密
	block, _ := pem.Decode(recv_pri)
	if block == nil {
		return nil, errors.New("privateKey error")
	}
	// 解析PKCS1格式的私钥
	PrivKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, PrivKey, ciphertext)
}

func main() {
	// 1. 输入字符串
	input := []byte("hello kitty")
	// 2. 加密字符串
	// 2.1 hash加密字符串（摘要）
	zhaiyao := Hash_md5(input) // zhaiyao([]byte)
	fmt.Println("摘要：", zhaiyao)
	// 2.2 用接收方的公钥加密字符串
	miwen, _ := RsaEncrypt(input) // 密文 []byte
	fmt.Println("密文：", miwen)
	// 2.2.2 用发送方的私钥加密摘要（签名）
	sign, _ := RSASign(zhaiyao) // 签名
	fmt.Println("签名：", []byte(sign))
	// 3 将两种方式加密好的密文都发送给接收方
	// 4.1 用接收方的私钥解密2.2的密文 得到字符串 []byte
	mingwen, _ := RsaDecrypt(miwen)
	fmt.Println("明文：", string(mingwen))
	// 4.1.2 将得到的字符串用同样的hash加密 得到摘要
	digest := Hash_md5(mingwen)
	fmt.Println("digest:", digest)
	// 4.2 用发送方的公钥解密发送方发过来的签名 得到摘要
	// 将两种方式的摘要进行比较 判断是否相同
	err := RSAVerify([]byte(sign), digest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("成功")
	}

}
