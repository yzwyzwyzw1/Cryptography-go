package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	//output := PKCS5Padding([]byte("abc"),8)
	//fmt.Println("补码",output)
	//
	//fmt.Println("去码",PKCS5UnPadding(output))


	 var key = "123344234234"
	//
	 var data  = []byte("hello")

	//Encrypt(key,data)

	var cipherTxt = Encrypt(key,data)

	fmt.Println("加密后的结果",hex.EncodeToString(cipherTxt))

	var d = Decrypt(cipherTxt,key)
	fmt.Println("解密后的结果：",string(d))


}


func Encrypt(key string,data []byte) []byte {
	var sum = 0
	//加密算法：首先计算key的总和,利用key的总和与明文参与运算
	for i :=0;i<len(key);i++ {
		sum = sum + int(key[i])
	}

	var pad = PKCS5Padding(data,len(key))

	fmt.Println(pad)
	//通过加法运算，实现加密过程
	for i:=0;i<len(pad);i++{
		pad[i]= pad[i]+byte(sum)
	}

	return pad
}


func Decrypt(cipherTxt []byte,key string) []byte {

	var sum = 0
	for i:=0;i<len(key);i++ {
		sum += int(key[i])
	}

	for i:=0;i<len(cipherTxt);i++ {
		cipherTxt[i] = cipherTxt[i] - byte(sum)
	}

	var p = PKCS5UnPadding(cipherTxt)
	return p
}



//PKCS5Unpadding 去码
func PKCS5UnPadding(cipherTxt []byte) []byte {
	var l = len(cipherTxt)
	var txt =int(cipherTxt[l-1])
	return cipherTxt[:l-txt] //返回数组

}


func PKCS5Padding(cipherTxt[] byte,blockSize int) []byte {
	//计算准备添加的数字
	padding := blockSize-len(cipherTxt)%aes.BlockSize
	//55555
	padTxt := bytes.Repeat([]byte{byte(padding)},padding)

	var byteTxt = append(cipherTxt,padTxt...)

	return byteTxt
}

