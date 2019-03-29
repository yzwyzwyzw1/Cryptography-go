package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

func main() {

	//创建私钥
	priv,_ := rsa.GenerateKey(rand.Reader,1024)


	fmt.Println("私钥为：",priv)
	//通过私钥创建公钥
	pub := priv.PublicKey

	//加密和解密
	org := []byte("hello china")

	//通过oaep函数实现公钥加密
	//EncryptOAEP的第一参数的作用为，将不同长度的明文，通过hash散列实现相同长度的散列值，此过程就是生成密文摘要过程

	cipherTxt,_ := rsa.EncryptOAEP(md5.New(),rand.Reader,&pub,org,nil)

	//打印密文
	fmt.Println(cipherTxt)
	fmt.Println(base64.StdEncoding.EncodeToString(cipherTxt))

	//解密
	plainTxt,_ := rsa.DecryptOAEP(md5.New(),rand.Reader,priv,cipherTxt,nil)
	//打印明文
	fmt.Println(plainTxt)
	fmt.Println(string(plainTxt))


}
