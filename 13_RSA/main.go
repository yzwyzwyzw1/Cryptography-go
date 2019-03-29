package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

//声明私钥
var priKey  = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDJ2uVcXVDYwvYdhks9fh3gcQVnAbL/i+1PVDxjAsTIohukCxia
3sxgWVVeCPL2mqiq6b0enRCf4haKRRBClQ0wBtCLC0gx8XWiibFXCD4/11P+kp7W
xRB2vkNGKR3mqoGZQg0DRhRPUZvz0vmKL69mPp3DYR5X/n7JZQbRE4HJHwIDAQAB
AoGAWJAT5cyDdjdD6HxNcrNsxFaSOjmCoaBxBEc/H2nNkQGfAwBjUT/Dh9pqHBHt
F8mPiz12XoInEx8NKKlYkv32hH5S+EnQjEz6odSNaSKIaTi8XI9ngNXsOx5XcA+F
Hei5SqzdT3v1dqShCztLijyJBiVPU7jBUNYARSMWLPTnQXECQQDoYq98Tltts8KW
gAvfl9bFmeJ834PnNZeBjI+0LbiWwvz5IyC0sRFX/kSK9TtLIYt+TjbhpdUIkZGS
rvcihFJ7AkEA3l38oq4SS1tSU5ooPcWyuS4tSbACGkwueEC1E8JaE+TmMnE6HBuD
ZpMytLwXH3/RL7F93EkLOExh1Q5LIDpkrQJAEJAQD90J2SzJvq5LqlkJHrZUFTBd
F1qTfNFG4MGteVlWbG9bgmVoZgwiWIEYA8fVSGBMuxJ1t5GgfYLo7FXgBwJAPN33
MER7p3rozw+MYlRCNqQoK1ptSOt/cNUte6ogY/+s7zPFFzqEvkgeghBWN/4RsX9h
FzI7Sv7FbL9L2xeAxQJBAK0tidabLqf3viEEcG3Us0mIBM2EurSErFMvShOgn4NL
vMk0MKhZ22s7X7aX8hhNR713jdrObymBdLvxPCAtkIk=
-----END RSA PRIVATE KEY-----`)


var pubKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJ2uVcXVDYwvYdhks9fh3gcQVn
AbL/i+1PVDxjAsTIohukCxia3sxgWVVeCPL2mqiq6b0enRCf4haKRRBClQ0wBtCL
C0gx8XWiibFXCD4/11P+kp7WxRB2vkNGKR3mqoGZQg0DRhRPUZvz0vmKL69mPp3D
YR5X/n7JZQbRE4HJHwIDAQAB
-----END PUBLIC KEY-----`)





func RSAEncrypt(origData []byte) []byte {
	//公钥加密
	block,_ := pem.Decode(pubKey)
	//解析公钥
	pubInterface,_ := x509.ParsePKIXPublicKey(block.Bytes)  //这函数别写错了
	//加载公钥
	pub := pubInterface.(*rsa.PublicKey)
	//pub := pubInterface.N//这个点不能丢
	//加密密文
	bits,_ := rsa.EncryptPKCS1v15(rand.Reader,pub,origData)
	//bits为加密密文
	return bits
}

//解密函数
func RSADecrypt(origData []byte) []byte {
	block,_:=pem.Decode(priKey)
	//解析私钥
	priv,_:=x509.ParsePKCS1PrivateKey(block.Bytes)
	//解密
	bts,_:=rsa.DecryptPKCS1v15(rand.Reader,priv,origData)
	//返回明文
	return  bts

}


func main() {
	cipherTxt := RSAEncrypt([]byte("hello world 123"))
	fmt.Println(cipherTxt)
	fmt.Println(base64.StdEncoding.EncodeToString(cipherTxt))

	fmt.Println("解密结果",string(RSADecrypt(cipherTxt)))
}
