// 本文链接： https://dreamerjonson.com/2018/12/07/golang-36-blockchain-signature/
// golang[36]-区块链-数据签名生成 2018-12-07
// 区块链-数据签名生成

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

//生成私钥和公钥，生成的私钥为结构体ecdsa.PrivateKey的指针

//type PrivateKey struct {
//	PublicKey
//	D *big.Int
//}
func newKeyPair2() (ecdsa.PrivateKey, []byte) {
	//生成secp256k1椭圆曲线
	curve := elliptic.P256()
	//产生的是一个结构体指针，结构体类型为ecdsa.PrivateKey
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	//x坐标与y坐标拼接在一起，生成公钥
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}

func main() {
	//调用函数生成私钥与公钥
	privKey, _ := newKeyPair2()

	//信息的哈希,签名什么样的数据
	hash := sha256.Sum256([]byte("hello world\n"))

	//根据私钥和信息的哈希进行数字签名，产生r和s
	r, s, err := ecdsa.Sign(rand.Reader, &privKey, hash[:])

	if err != nil {
		log.Panic(err)
	}

	//r和s拼接在一起实现了数字签名
	signature := append(r.Bytes(), s.Bytes()...)
	//打印数字签名的16进制显示
	fmt.Printf("%x\n", signature)

	fmt.Printf("%x\n", r.Bytes())
	fmt.Printf("%x\n", s.Bytes())

	//补充：如何把一个字符串转换为16进制数据
	//m := big.Int{}
	//n := big.Int{}
	//rr,_:=hex.DecodeString("7dccc0f58639584a3f0c879c3688d2f4a0137697cbf34245d075c764e36233d2")
	//ss,_:=hex.DecodeString("cf3713bf4369eb1c02e476cdbefb7f76a25b572f53fb71d4e4742fa11c827526")
	//
	//m.SetBytes(rr)
	//n.SetBytes(ss)
	//
	//fmt.Printf("%x\n", m.Bytes())
	//fmt.Printf("%x\n", n.Bytes())
}
