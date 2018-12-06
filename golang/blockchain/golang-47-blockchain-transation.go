// 本文链接： https://dreamerjonson.com/2018/12/13/golang-47-blockchain-transation/
// golang[47]-区块链-比特币交易 2018-12-13
// 比特币的交易
/*
第一笔交易比较特殊，他是coinbase交易，矿工的收益。
金额每4年减少一半，从2009年一开始一个区块奖励50BTC、后来减少为了25个BTC、一直到18年 12.5BTC。总共有2100万BTC。

交易和我们的银行的转账非常不同，比特币中没有记录账户的信息、而是交易的信息。
go实现交易 demo
*/
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
)

const subsidy = 100

type Transation struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXInput struct {
	TXid      []byte
	Voutindex int
	Signature []byte
}

type TXOutput struct {
	value      int
	PubkeyHash []byte
}

//格式化打印交易完整信息
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
func (tx Transation) Serialize() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)

	err := enc.Encode(tx)

	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

//计算交易的hash值
func (tx *Transation) Hash() []byte {

	txcopy := *tx
	txcopy.ID = []byte{}

	hash := sha256.Sum256(txcopy.Serialize())

	return hash[:]
}

//根据金额与地址新建一个输出
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.PubkeyHash = []byte(address)
	return txo
}

//第一笔coinbase交易
func NewCoinbaseTX(to string) *Transation {
	txin := TXInput{[]byte{}, -1, nil}
	txout := NewTXOutput(subsidy, to)

	tx := Transation{nil, []TXInput{txin}, []TXOutput{*txout}}

	tx.ID = tx.Hash()

	return &tx
}

func main() {
	newTX := NewCoinbaseTX("jonson")
	fmt.Printf("%s", newTX.String())
}
