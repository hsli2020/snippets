// 本文链接： https://dreamerjonson.com/2018/12/09/golang-40-blockchain-verifyprivatekeywif/
// golang[40]-区块链-wif倒推出私钥 2018-12-09
// wif倒推出私钥

package main

import (
	"bytes"
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

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for _, b := range input {
		if b == '1' {
			zeroBytes++
		} else {
			break
		}
	}

	payload := input[zeroBytes:]

	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b) //反推出余数

		result.Mul(result, big.NewInt(58)) //之前的结果乘以58

		result.Add(result, big.NewInt(int64(charIndex))) //加上这个余数

	}

	decoded := result.Bytes()

	decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)
	return decoded
}

//检查checkWIF是否有效
func checkWIF(wifprivate string) bool {
	rawdata := []byte(wifprivate)
	//包含了80、私钥、checksum
	base58decodedata := Base58Decode(rawdata)

	fmt.Printf("base58decodedata：%x\n", base58decodedata)
	length := len(base58decodedata)

	if length < 37 {
		fmt.Printf("长度小于37，一定有问题")
		return false
	}

	private := base58decodedata[:(length - 4)]
	//得到检查码
	//fmt.Printf("private：%x\n",private)
	firstsha := sha256.Sum256(private)

	secondsha := sha256.Sum256(firstsha[:])

	checksum := secondsha[:4]
	//fmt.Printf("%x\n",checksum)
	//得到原始的检查码
	orignchecksum := base58decodedata[(length - 4):]
	//	fmt.Printf("%x\n",orignchecksum)

	//[]byte对比
	if bytes.Compare(checksum, orignchecksum) == 0 {
		return true
	}

	return false
}

//通过wif格式的私钥，得到原始的私钥。
func getPrivateKeyfromWIF(wifprivate string) []byte {
	if checkWIF(wifprivate) {
		rawdata := []byte(wifprivate)
		//包含了80、私钥、checksum
		base58decodedata := Base58Decode(rawdata)
		//私钥一共32个字节，排除了0x80
		return base58decodedata[1:33]
	}
	return []byte{}
}

func main() {
	//得到wif格式的私钥
	wifprivatekey := generatePrivateKey("18d3e15d48b2df76562fab783eac137aaeb611e6ff0a193e12ceef1354220ac7", false)
	fmt.Printf("%s\n", wifprivatekey)

	//【】byte转换为string
	str := fmt.Sprintf("%s", wifprivatekey)

	//得到原始的私钥
	privatekey := getPrivateKeyfromWIF(str)
	fmt.Printf("%x", privatekey)
}

// 参考资料
//
// https://en.bitcoin.it/wiki/Private_key
// https://en.bitcoin.it/wiki/Wallet_import_format
// http://gobittest.appspot.com/PrivateKey
