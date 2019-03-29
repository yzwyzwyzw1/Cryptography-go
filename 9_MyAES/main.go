package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

//PKCS5Padding 要求分组长度只能为8

//PKCS7Padding 要求分组的长度可以[1-255]


func PKCS7Padding(org []byte,blockSize int) []byte {
	pad := blockSize-len(org)%blockSize
	padArr := bytes.Repeat([]byte{byte(pad)},pad)
	return append(org,padArr...)

}

//去掉补码
func PKCS7Unpadding(org []byte) []byte {
	l := len(org)
	pad :=org[l-1]
	return org[:l-int(pad)]
}

//通过CBC分组模式，完成AES的密码过程
//AES 也是对称加密，AES 是DES 的替代品
//AES 秘钥长度，要么16,或者 24, 或者32

func AesCBCEncrypt(org []byte,key []byte) []byte {

	//校验密钥
	block,_ := aes.NewCipher(key)

	org = PKCS7Padding(org,block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block,key)

	crypted := make([]byte,len(org))
	blockMode.CryptBlocks(crypted,org)

	return crypted

}


func AesCBCDecrypt(cipherTxt []byte,key []byte) []byte {
	//校验key
	block,_ := aes.NewCipher(key)
	//设置解密模式CDC
	blockMode := cipher.NewCBCDecrypter(block,key)

	//开始解密
	org := make([]byte,len(cipherTxt))
	blockMode.CryptBlocks(org,cipherTxt)

	//去除补码
	org = PKCS7Unpadding(org)

	return org
}

func main() {
	//测试pkcs7padding
	//pad:=PKCS7Padding([]byte("abc"),8)

	//fmt.Println(PKCS7Unpadding(pad))

	ciphertxt:=AesCBCEncrypt([]byte("hello 123"),[]byte("1234567890123456"))

	fmt.Println("解密后的结果",string(AesCBCDecrypt(ciphertxt,[]byte("1234567890123456"))))
}
