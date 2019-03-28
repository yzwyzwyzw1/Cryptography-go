package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {

	var key = []byte("12345678")
	var data = []byte("hello world")
	var cipherTxt = DESEncrypt(data,key)

	fmt.Println("加密后的结果：",cipherTxt)


	var origData = DESDecrypt(cipherTxt,key)
	fmt.Println("解密后的结果：",string(origData))


}

//调用系统库中的DES加密
func DESEncrypt(origData []byte,key []byte) []byte {
	//DES加密中key长度必须为8
	//3DES加密中key的长度必须24

	//校验秘钥
	block,_ := des.NewCipher(key)

	//补码
	origData = PKCS5Padding(origData,block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block,key)
	//
	crypted := make([]byte,len(origData))
	blockMode.CryptBlocks(crypted,origData)

	return crypted
}

func DESDecrypt(crypted []byte,key []byte) []byte {

	//校验key的有效性
	block,_ := des.NewCipher(key)

	//通过CBC模式解密
	blockMode := cipher.NewCBCDecrypter(block,key)

	//实现解密
	origData := make([]byte,len(crypted))
	blockMode.CryptBlocks(origData,crypted)

	//去码
	origData = PKCS5UnPadding(origData)

	return origData

}

//PKCS5Unpadding 去码
func PKCS5UnPadding(cipherTxt []byte) []byte {
	var l = len(cipherTxt)
	var txt = int(cipherTxt[l-1])
	return cipherTxt[:l-txt]
}


//实现PKCS5Padding补码
func PKCS5Padding(cipherTxt[] byte,blockSize int) []byte {
	//计算准备添加的数字
	padding:=blockSize-len(cipherTxt)%blockSize

	//55555
	padTxt:=bytes.Repeat([]byte{byte(padding)},padding)

	//叠加两个数组
	var byteTxt =append(cipherTxt,padTxt...)

	return byteTxt

}

