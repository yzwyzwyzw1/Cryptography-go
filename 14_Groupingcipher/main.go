package main

import (
	"bytes"
	"fmt"
)

func PKCS8Padding(org []byte) []byte {
	//长度为8
	pad := 8-len(org)%8
	//补码的字节数组
	pading := bytes.Repeat([]byte{byte(pad)},pad)
	//补码
	return append(org,pading...)
}

func PKCS8Unpadding(org []byte) []byte {
	l := len(org)
	pad := int(org[l-1])
	return org[:l-pad]

}

//实现密文分组的过程
func gcrypt(org []byte) [][]byte {
	//因为利用了PKCS8实现的补码，所以就以8个字节为单位
	var g [][]byte
	var t []byte
	//遍历补码后的数据
	for i:=0;i<len(org);i++ {
		t = append(t,org[i]+1)
		if i%8==0 && i!=0 {
			//做分组
			g = append(g,t)
			t = make([]byte,0)
		}
	}
	return g
}


func main() {
	cipher := PKCS8Padding([]byte("helloworldhellochina"))
	fmt.Println("原文为",cipher,"\n")

	g := gcrypt(cipher)

	for index,value := range g {
		for _,val := range value {
			fmt.Printf("当前组%d,值为%d\n",index,val-1)
		}
	}
}