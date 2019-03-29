package main

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

)

func AesCFBEncrypt(plaintxt []byte,key []byte) []byte {

	block,_ := aes.NewCipher(key)
	cipherTxt := make([]byte,aes.BlockSize+len(plaintxt))
	iv := cipherTxt[:aes.BlockSize]


	io.ReadFull(rand.Reader,iv)
	//设置加密模式
	stream := cipher.NewCFBEncrypter(block,iv)

	//加密
	stream.XORKeyStream(cipherTxt[aes.BlockSize:],plaintxt)

	return cipherTxt
}

func AesCFBDecrypt(ciphertxt []byte,key []byte) []byte {
	block,_ := aes.NewCipher(key)
	iv := ciphertxt[:aes.BlockSize]
	ciphertxt = ciphertxt[aes.BlockSize:]
	//设置解密模式
	stream := cipher.NewCFBDecrypter(block,iv)
	var des = make([]byte,len(ciphertxt))
	//解密
	stream.XORKeyStream(des,ciphertxt)
	return des

	//block,_ := aes.NewCipher(key)
	//iv := ciphertxt[:aes.BlockSize]
	////存储解密后的信息
	//plaintxt :=make([]byte,len(ciphertxt)-aes.BlockSize)
	//stream := cipher.NewCTR(block,iv)
	//stream.XORKeyStream(plaintxt,ciphertxt[aes.BlockSize:])
	//return plaintxt





}

func main() {
	var ciphertxt = AesCFBEncrypt([]byte("hello world"),[]byte("1234567890123456"))

	//三种显示方式
	fmt.Println("加密数据",ciphertxt)
	fmt.Println(hex.EncodeToString(ciphertxt))
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertxt))


	//解密
	fmt.Println("解密数据",string((AesCFBDecrypt(ciphertxt,[]byte("1234567890123456")))))
}