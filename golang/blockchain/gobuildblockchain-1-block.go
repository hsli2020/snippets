// 本文链接： https://dreamerjonson.com/2018/12/16/gobuildblockchain-1-block/ 
// go实现区块链[1]-block 2018-12-16

// 定义区块结构体
type Block struct{
	Version int32
	PrevBlockHash []byte
	Merkleroot []byte
	Hash []byte
	Time int32
	Bits int32
	Nonce int32
}

//序列化
func (b* Block) Serialize() []byte{

	var encoded bytes.Buffer
	enc:= gob.NewEncoder(&encoded)

	err:= enc.Encode(b)

	if err!=nil{
		log.Panic(err)
	}
	return encoded.Bytes()
}

//反序列化
func DeserializeBlock(d []byte) *Block{
	var block Block

	decode :=gob.NewDecoder(bytes.NewReader(d))
	err := decode.Decode(&block)
	if err!=nil{
		log.Panic(err)
	}
	return &block
}

//打印区块
func (b*Block)String(){
	fmt.Printf("version:%s\n",strconv.FormatInt(int64(b.Version),10))
	fmt.Printf("Prev.BlockHash:%x\n",b.PrevBlockHash)
	fmt.Printf("Prev.merkleroot:%x\n",b.Merkleroot)
	fmt.Printf("Prev.Hash:%x\n",b.Hash)
	fmt.Printf("Time:%s\n",strconv.FormatInt(int64(b.Time),10))
	fmt.Printf("Bits:%s\n",strconv.FormatInt(int64(b.Bits),10))
	fmt.Printf("nonce:%s\n",strconv.FormatInt(int64(b.Nonce),10))
}

//打印区块、测试序列化

func TestNewSerialize(){
	//初始化区块
	block := &Block{
		2,
		[]byte{},
		[]byte{},
		[]byte{},
		1418755780,
		404454260,
		0,
		[]*Transation{},
	}

	deBlock:=DeserializeBlock(block.Serialize())

	deBlock.String()
}

//添加交易

transition.go

package main

import (
	"fmt"
	"strings"
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

//挖矿奖励
const subsidy = 100

//交易
type Transation struct{
	ID []byte
	Vin []TXInput
	Vout []TXOutput

}

//输入
type TXInput struct {
	TXid []byte
	Voutindex int
	Signature []byte
}

//输出
type TXOutput struct {
	value int
	PubkeyHash []byte
}

//打印
func (tx Transation) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("--- Transaction %x:", tx.ID))

	for i, input := range tx.Vin {
		lines = append(lines, fmt.Sprintf("     Input %d:", i))
		lines = append(lines, fmt.Sprintf("       TXID:      %x", input.TXid))
		lines = append(lines, fmt.Sprintf("       Out:       %d", input.Voutindex))
		lines = append(lines, fmt.Sprintf("       Signature: %x", input.Signature))
	}

	for i, output := range tx.Vout {
		lines = append(lines, fmt.Sprintf("     Output %d:", i))
		lines = append(lines, fmt.Sprintf("       Value:  %d", output.value))
		lines = append(lines, fmt.Sprintf("       Script: %x", output.PubkeyHash))
	}

	return strings.Join(lines, "\n")
}

//序列化
func (tx Transation) Serialize() []byte{
	var encoded bytes.Buffer
	enc:= gob.NewEncoder(&encoded)

	err:= enc.Encode(tx)

	if err!=nil{
		log.Panic(err)
	}
	return encoded.Bytes()
}

//计算交易的hash值
func (tx *Transation) Hash() []byte{

	txcopy := *tx
	txcopy.ID = []byte{}

	hash:= sha256.Sum256(txcopy.Serialize())

	return hash[:]
}

//根据金额与地址新建一个输出
func NewTXOutput(value int,address string) * TXOutput{
	txo := &TXOutput{value,nil}
	txo.PubkeyHash = []byte(address)
	return txo
}

//第一笔coinbase交易
func NewCoinbaseTX(to string) *Transation{
	txin := TXInput{[]byte{},-1,nil}
	txout := NewTXOutput(subsidy,to)

	tx:= Transation{nil,[]TXInput{txin},[]TXOutput{*txout}}

	tx.ID = tx.Hash()

	return &tx
}

//工具类

utils.go

package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

//计算两个数的最小值
func min(a int,b int) int{

	if(a>b){
		return b
	}
	return a
}

//将类型转化为了字节数组,小端
func IntToHex(num int32) []byte{
	buff := new(bytes.Buffer)
	//binary.LittleEndian 小端模式
	err:= binary.Write(buff,binary.LittleEndian,num)

	if err!=nil{
		log.Panic(err)
	}

	return buff.Bytes()
}

//将类型转化为了字节数组，大端
func IntToHex2(num int32) []byte{
	buff := new(bytes.Buffer)
	//binary.LittleEndian 小端模式
	err:= binary.Write(buff,binary.BigEndian,num)

	if err!=nil{
		log.Panic(err)
	}

	return buff.Bytes()
}

//字节反转
func ReverseBytes(data []byte){
	for i,j :=0,len(data) - 1;i<j;i,j = i+1,j - 1{
		data[i],data[j] = data[j],data[i]
	}
}

// 修改区块

block.go

//增加交易
type Block struct{
	Version int32
	PrevBlockHash []byte
	Merkleroot []byte
	Hash []byte
	Time int32
	Bits int32
	Nonce int32
	Transations []*Transation

}

//根据前一个hash增加区块
func NewBlock(prevBlockHash []byte) * Block{

	block := &Block{
		2,
		prevBlockHash,
		[]byte{},
		[]byte{},
		int32(time.Now().Unix()),
		404454260,
		0,
		[]*Transation{},
	}

	pow := NewProofofWork(block)

	nonce,hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

//创世区块
func NewGensisBlock() * Block{
	block := &Block{
		2,
		[]byte{},
		[]byte{},
		[]byte{},
		int32(time.Now().Unix()),
		404454260,
		0,
		[]*Transation{},
	}

	pow:=NewProofofWork(block)

	nonce,hash:=pow.Run()

	block.Nonce = nonce
	block.Hash = hash

	//block.String()
	return block
}
