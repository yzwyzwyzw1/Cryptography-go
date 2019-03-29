package main

import (
	"awesomeProject/8_MyCS/CryptedDic"
	"fmt"
	"net"
)

func MyClient(cipherTxt []byte) {
	//创建准备连接的服务器
	netAddr,_ := net.ResolveTCPAddr("tcp4","127.0.0.1:1234")
	//连接服务器
	conn,_ := net.DialTCP("tcp",nil,netAddr)
	//发送数据
	_,err := conn.Write(cipherTxt)
	if err != nil {
		fmt.Println("发送数据失败")
	}
	fmt.Println("发送数据成功")
}

func main() {
	data := []byte("hello 123")

	//cipherTxt就是加密出的密文
	cipherTxt := CryptedDic.EnCrypt("12345678",data)

	//发送数据
	MyClient(cipherTxt)
}