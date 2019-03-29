package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

//通过AES测试OFB
func AesOFBEncryt(plaintxt []byte ,key []byte) []byte {
	block,_:=aes.NewCipher(key)

	ciphertxt:=make([]byte,aes.BlockSize+len(plaintxt))//切片分组

	/********************/
	fmt.Println("block:",block,"\n")
	fmt.Println("aes.BlockSize:",aes.BlockSize,"len(plaintxt):",len(plaintxt),"\n")
	fmt.Println("ciphertxt:",ciphertxt,"\n")
	/********************/
	iv:=ciphertxt[:block.BlockSize()]
	io.ReadFull(rand.Reader,iv)
	/******************/
	fmt.Println("iv:",iv)
	/********************/
	//设置加密模式
	stream:=cipher.NewOFB(block,iv)

	//其实是将加密后的密文存到ciphertxt[aes.BlockSize:]
	stream.XORKeyStream(ciphertxt[aes.BlockSize:],plaintxt)
	return ciphertxt

}

//解密
func AesOFBDecrypt(ciphertxt []byte,key []byte) []byte {
	block,_:=aes.NewCipher(key)
	iv:=ciphertxt[:aes.BlockSize]
	//存储解密后的信息
	plaintxt:=make([]byte,len(ciphertxt)-aes.BlockSize)
	//设置解密方式OFB
	stream:=cipher.NewOFB(block,iv)
	stream.XORKeyStream(plaintxt,ciphertxt[aes.BlockSize:])

	return plaintxt
}


func main () {
	ciphertxt:=AesOFBEncryt([]byte("hello world 123"),[]byte("1234567890123456"))
	fmt.Println(ciphertxt)

	fmt.Println(string(AesOFBDecrypt(ciphertxt,[]byte("1234567890123456"))))
}