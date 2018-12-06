// 本文链接： https://dreamerjonson.com/2018/12/07/golang-37-blockchain-verifysign/
// golang[37]-区块链-验证数据签名 2018-12-07
// 验证数据签名

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

//生成私钥和公钥，生成的私钥为结构体ecdsa.PrivateKey的指针

//type PrivateKey struct {
//	PublicKey
//	D *big.Int
//}
func newKeyPair3() (ecdsa.PrivateKey, []byte) {

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

	//生成公钥要私钥
	privKey, pubkey := newKeyPair3()

	//生成某一串信息的哈希值，需要签名的数据
	hash := sha256.Sum256([]byte("跟着jonson老师实战区块链\n"))

	//根据私钥和信息的哈希值生成数字签名的r和s，r和s拼接在一起就是数字签名，在这里省略了拼接的步骤，欲查看，请看3.数字签名
	r, s, _ := ecdsa.Sign(rand.Reader, &privKey, hash[:])

	//fmt.Printf("%v\n", *r)
	//fmt.Printf("%v\n", *s)
	////生成secp256k1椭圆曲线
	curve := elliptic.P256()

	//公钥的长度
	keyLen := len(pubkey)

	//前一半为x轴坐标，后一半为y轴坐标
	x := big.Int{}
	y := big.Int{}
	x.SetBytes(pubkey[:(keyLen / 2)])
	y.SetBytes(pubkey[(keyLen / 2):])

	//rawPubKey为生成PublicKey结构体，作为下面ecdsa.Verify的参数
	//type PublicKey struct {
	//	elliptic.Curve
	//	X, Y * big.Int }

	//公钥
	rawPubKey := ecdsa.PublicKey{curve, &x, &y}

	//根据交易哈希、公钥、数字签名验证成功。ecdsa.Verify func Verify(pub *PublicKey, hash []byte, r *big.Int, s *big.Int) bool
	if ecdsa.Verify(&rawPubKey, hash[:], r, s) == false {
		fmt.Printf("%s\n", "验证失败")
	} else {
		fmt.Printf("%s\n", "验证成功")
	}

	//用其他的信息哈希——证明验证失败
	hash2 := sha256.Sum256([]byte("我要给你200愿\n"))

	if ecdsa.Verify(&rawPubKey, hash2[:], r, s) == false {
		fmt.Printf("%s\n", "验证失败")
	} else {
		fmt.Printf("%s\n", "验证成功")
	}
}
