package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)


//椭圆曲线实现数字签名
//椭圆曲线基于DSA改进而成，数字签名的安全性比rsa高
//dsa专业做签名的，不做加密
//ecc专业做签名，比特币就是用ecc完成的签名
func main() {

	//通过椭圆曲线加密实现明文签名
	message := []byte("hello world")

	//生成私钥
	//elliptic.P256()设置生成私钥为256
	privatekey,_ := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)

	//创建公钥
	publickey := privatekey.PublicKey

	//hash散列明文
	digest := sha256.Sum256(message)

	//用私钥签名
	r,s,_:=ecdsa.Sign(rand.Reader,privatekey,digest[:])


	//------------保存签名结果------------------//
	param := privatekey.Curve.Params()

	//获取私钥长度字节
	curveOrderBytes := param.P.BitLen()/8

	//获得签名返回的字节
	rByte,sByte := r.Bytes(),s.Bytes()

	//创建数组
	signature := make([]byte,curveOrderBytes*2)
	copy(signature[:len(rByte)],rByte)
	copy(signature[len(sByte):],sByte)

	//假设signature 中就存放了完整的签名的结果

	//--------------------------------------//
	//假设通过TCP方式将signature传递给对方

	//验证签名
	digest = sha256.Sum256(message)
	//获取公钥的字节长度
	curveOrderBytes = publickey.Curve.Params().P.BitLen()/8

	//创建大整数类型保存rbyte,sbyte
	r,s = new(big.Int),new(big.Int)
	r.SetBytes(signature[:curveOrderBytes])
	s.SetBytes(signature[curveOrderBytes:])

	//开始认证
	e := ecdsa.Verify(&publickey,digest[:],r,s)
	if e == true {
		fmt.Println("验证签名")
	}

}