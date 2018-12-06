// 本文链接： https://dreamerjonson.com/2018/12/09/golang-39-blockchain-generateprivatekeywif/
// golang[39]-区块链-产生wif私钥 2018-12-09
// 压缩公钥
//
// 公钥一般来说是椭圆曲线上的x,y坐标拼接在一起的。压缩的公钥其实就是x的坐标。
// WIF 私钥产生

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

//base58编码
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)

	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod) // 对x取余数
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)

	for _, b := range input {

		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return result
}

//字节数组的反转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
func generatePrivateKey(hexprivatekey string, compressed bool) []byte {
	versionstr := ""
	//判断是否对应的是压缩的公钥，如果是，需要在后面加上0x01这个字节。同时任何的私钥，我们需要在前方0x80的字节
	if compressed {
		versionstr = "80" + hexprivatekey + "01"
	} else {
		versionstr = "80" + hexprivatekey
	}
	//字符串转化为16进制的字节
	privatekey, _ := hex.DecodeString(versionstr)
	//通过 double hash 计算checksum.checksum他是两次hash256以后的前4个字节。
	firsthash := sha256.Sum256(privatekey)

	secondhash := sha256.Sum256(firsthash[:])

	checksum := secondhash[:4]
	//拼接
	result := append(privatekey, checksum...)

	//最后进行base58的编码
	base58result := Base58Encode(result)
	return base58result
}

func main() {
	wifprivatekey := generatePrivateKey("18d3e15d48b2df76562fab783eac137aaeb611e6ff0a193e12ceef1354220ac7", true)
	fmt.Printf("%s", wifprivatekey)
}
