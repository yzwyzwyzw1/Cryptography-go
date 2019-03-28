package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"io"
	"os"
)

//MD5 16个字节 128位
func MyMd5() {

	//第一种写法
	data := []byte("hello world")

	s := fmt.Sprintf("%x",md5.Sum(data))

	//密文为16进制的数字
	fmt.Println(s)

     //第二种写法
	data2 :=[]byte("hello world")
	m := md5.New()
	m.Write(data2)
	//字节数组转换成字符串
	s2 := hex.EncodeToString(m.Sum(nil))
	fmt.Println(s2)


}

//总结，使用sha256方式加密，占用32字节 = （32字节×8=256位），Sha256加密方法，通用在公链中
func MySha256() {

	//1
	data := []byte("hello world")
	s := fmt.Sprintf("%x",sha256.Sum256(data))
	fmt.Println(s)

	//2
	data2 :=[]byte("hello world")
	m :=sha256.New()
	m.Write(data2)
	s2 := hex.EncodeToString(m.Sum(nil))
	fmt.Println(s2)

	//第三种用法
	//可将文件中的数据进行sha256加密处理
	//通过文件流首先找到文件,将文件读入内存
	f,_:=os.Open("sha256Ftest")
	//创建sha256对象
	h:=sha256.New()
	//将f内存中的数据，拷贝到sha256中
	io.Copy(h,f)//io流拷贝
	//实现sha256计算过程
	s3:= hex.EncodeToString(h.Sum(nil))
	fmt.Println(s3)
}

//如果利用ripemd160加密，需要引入三方库
//引入三方法库的步骤
//1,进入gopath下，创建golang.org目录
//2,进入golang.org，创建x目录
//3,进入x目录，并在翻墙情况下,在github上下载三方库
//git clone https://github.com/golang/crypto.git


//以上的三个步骤可以通过一行命令在终端直接实现
//cd $GOPATH/src $ mkdir golang.org $ cd golang.org $ mkdir x $ cd x $ git clone https://github.com/golang/crypto.git


func MyRipemd160() {
	//只有一种写法
	hasher:=ripemd160.New()
	hasher.Write([]byte("hello world"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))


}

func main() {
	//MyMd5()
	MySha256()
}