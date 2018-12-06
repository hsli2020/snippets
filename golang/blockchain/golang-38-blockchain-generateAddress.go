// 本文链接： https://dreamerjonson.com/2018/12/07/golang-38-blockchain-generateAddress/
// golang[38]-区块链- 生成比特币地址 2018-12-07
// 生成比特币地址

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ripemd160"
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

//产生比特币地址
func generateAddress(pubkey []byte) []byte {
	//1、计算pubkeuhash
	pubkeyHash256 := sha256.Sum256(pubkey)

	PIPEMD160Hasher := ripemd160.New()

	_, err := PIPEMD160Hasher.Write(pubkeyHash256[:])

	if err != nil {
		fmt.Println("error")
	}

	publicRIPEMD160 := PIPEMD160Hasher.Sum(nil)

	//2、计算checksum
	versionPayload := append([]byte{0x00}, publicRIPEMD160...)

	firstSHA := sha256.Sum256(versionPayload)
	secondSHA := sha256.Sum256(firstSHA[:])
	//checksum 是前面的4个字节
	checksum := secondSHA[:4]

	//3、base58编码
	fullPayload := append(versionPayload, checksum...)
	//返回地址
	address := Base58Encode(fullPayload)
	return address
}

func main() {
	//外部得到公钥
	publickpey, _ := hex.DecodeString("D4A6C78C0B13DBD8A07AAB17C7D79ED9CB2523B63EDAC4E7CACE93C6B66CEDC7918EE0E174E8B2B61468D0E6CAA099710EF72094ACBD70BDAE3D8E42C617ACC6")
	//fmt.Printf("%X",publickpey)
	//打印这个地址
	address := generateAddress(publickpey)

	fmt.Printf("%s", address)
}

// 参考资料
// Building Blockchain in Go. Part 5: Addresses
//
// 比特币公钥转地址工具
// 地址：比特币维基百科
