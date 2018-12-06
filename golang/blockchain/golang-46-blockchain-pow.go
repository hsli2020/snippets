// 本文链接： https://dreamerjonson.com/2018/12/12/golang-46-blockchain-pow/
// golang[46]-区块链-比特币真实挖矿过程实现 2018-12-12
// 比特币真实挖矿过程实现

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
)

var (
	maxnonce int32 = math.MaxInt32
)

type Block struct {
	version       int32
	prevBlockHash []byte
	merkleroot    []byte
	hash          []byte
	time          int32
	bits          int32
	nonce         int32
}

//将类型转化为了字节数组
func IntToHex(num int32) []byte {
	buff := new(bytes.Buffer)
	//binary.LittleEndian 小端模式
	err := binary.Write(buff, binary.LittleEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

//将类型转化为了字节数组
func IntToHex2(num int32) []byte {
	buff := new(bytes.Buffer)
	//binary.LittleEndian 小端模式
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

//字节反转
func ReverseBytes4(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

//序列化
func (block *Block) serialize() []byte {

	result := bytes.Join(
		[][]byte{
			IntToHex(block.version),
			block.prevBlockHash,
			block.merkleroot,
			IntToHex(block.time),
			IntToHex(block.bits),
			IntToHex(block.nonce)},
		[]byte{},
	)

	return result
}

func main() {

	//前一个区块的hash
	prev, _ := hex.DecodeString("000000000000000016145aa12fa7e81a304c38aec3d7c5208f1d33b587f966a6")
	ReverseBytes4(prev)
	fmt.Printf("%x\n", prev)

	//默克尔根
	merkleroot, _ := hex.DecodeString("3a4f410269fcc4c7885770bc8841ce6781f15dd304ae5d2770fc93a21dbd70d7")
	ReverseBytes4(merkleroot)
	fmt.Printf("%x\n", merkleroot)

	//初始化区块
	block := &Block{
		2,
		prev,
		merkleroot,
		[]byte{},
		1418755780,
		404454260,
		0,
	}

	//目标hash
	//fmt.Printf("targethash:%x",CalculateTargetFast(IntToHex2(block.bits)))
	targetHash := CalculateTargetFast(IntToHex2(block.bits))

	//目标hash转换为bit.int
	var tartget big.Int
	tartget.SetBytes(targetHash)

	//当前hash
	var currenthash big.Int

	//一直计算到最大值，	block.nonce的值不断变化
	for block.nonce < maxnonce {

		//序列化，block.nonce的值不断变化带来序列化的变化
		data := block.serialize()
		//double hash
		fitstHash := sha256.Sum256(data)
		secondhash := sha256.Sum256(fitstHash[:])

		//反转
		ReverseBytes4(secondhash[:])
		fmt.Printf("nonce:%d,  currenthash:%x\n", block.nonce, secondhash)
		currenthash.SetBytes(secondhash[:])
		//比较
		if currenthash.Cmp(&tartget) == -1 {
			break
		} else {
			block.nonce++
		}
	}
}

//18   1B7B74

//计算困难度
func CalculateTargetFast(bits []byte) []byte {

	var result []byte
	//第一个字节  计算指数
	exponent := bits[:1]
	fmt.Printf("%x\n", exponent)

	//计算后面3个系数
	coeffient := bits[1:]
	fmt.Printf("%x\n", coeffient)

	//将字节，他的16进制为"18"  转化为了string "18"
	str := hex.EncodeToString(exponent) //"18"
	fmt.Printf("str=%s\n", str)
	//将字符串18转化为了10进制int64 24
	exp, _ := strconv.ParseInt(str, 16, 8)

	fmt.Printf("exp=%d\n", exp)
	//拼接，计算出目标hash
	result = append(bytes.Repeat([]byte{0x00}, 32-int(exp)), coeffient...)
	result = append(result, bytes.Repeat([]byte{0x00}, 32-len(result))...)

	return result
}
