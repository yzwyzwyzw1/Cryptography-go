package main

import (
	"awesomeProject/8_MyCS/CryptedDic"
	"fmt"
	"net"
)

func MyServer() []byte {

	netListen,_ := net.Listen("tcp","127.0.0.1:1234")

	//延时关闭
	defer netListen.Close()
	for {
		conn,_ := netListen.Accept()

		//接收数据
		data := make([]byte,2048)
		for {
			n,_ := conn.Read(data)
			//返回接收到的密文
			return data[:n]
		}
	}
}

func main() {
	cipherTxt := MyServer()
	//解密密文
	orig := CryptedDic.Decrypt(cipherTxt,"12345678")
	fmt.Println("解密后的数据为",string(orig))
}