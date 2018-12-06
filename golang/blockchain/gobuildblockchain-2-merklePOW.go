// 本文链接： https://dreamerjonson.com/2018/12/16/gobuildblockchain-2-merklePOW/ 
// go实现区块链[2]-整合默克尔树+POW 2018-12-16

// 添加merkleRoot

merkleTree.go

package main

import "crypto/sha256"

//默克尔树节点
type MerkleTree struct{
	RootNode *MerkleNode
}

//默克尔根节点
type MerkleNode struct{
	Left *MerkleNode
	Right *MerkleNode
	Data []byte
}

//生成默克尔树中的节点，如果是叶子节点，则Left，right为nil ，如果为非叶子节点，根据Left，right生成当前节点的hash
func NewMerkleNode(left,right *MerkleNode,data []byte) *MerkleNode{
	mnode := MerkleNode{}

	if left ==nil && right==nil{
		mnode.Data = data
	}else{
		prevhashes := append(left.Data,right.Data...)
		firsthash:= sha256.Sum256(prevhashes)
		hash:=sha256.Sum256(firsthash[:])
		mnode.Data = hash[:]
	}

	mnode.Left = left
	mnode.Right = right

	return &mnode
}

//构建默克尔树
func NewMerkleTree(data [][]byte) *MerkleTree{
	var nodes []MerkleNode
	//构建叶子节点。
	for _,datum := range data{
		node:= NewMerkleNode(nil,nil,datum)
		nodes =  append(nodes,*node)
	}
	//j代表的是某一层的第一个元素
	j:=0
	//第一层循环代表 nSize代表某一层的个数，每循环一次减半
	for nSize :=len(data);nSize >1;nSize = (nSize+1)/2{
		//第二条循环i+=2代表两两拼接。 i2是为了当个数是基数的时候，拷贝最后的元素。
		for i:=0 ; i<nSize ;i+=2{
			i2 := min(i+1,nSize-1)

			node := NewMerkleNode(&nodes[j+i],&nodes[j+i2],nil)
			nodes = append(nodes,*node)
		}
		//j代表的是某一层的第一个元素
		j+=nSize

	}

	mTree := MerkleTree{&(nodes[len(nodes)-1])}
	return &mTree
}

// 根据交易创建merkleROOT

func (b*Block) createMerkelTreeRoot(transations []*Transation){
	var tranHash [][]byte

	for _,tx:= range transations{

		tranHash = append(tranHash,tx.Hash())
	}

	mTree := NewMerkleTree(tranHash)

	b.Merkleroot =  mTree.RootNode.Data
}

// 测试merkle

func TestCreateMerkleTreeRoot(){
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

	txin := TXInput{[]byte{},-1,nil}
	txout := NewTXOutput(subsidy,"first")
	tx := Transation{nil,[]TXInput{txin},[]TXOutput{*txout}}

	txin2 := TXInput{[]byte{},-1,nil}
	txout2 := NewTXOutput(subsidy,"second")
	tx2 := Transation{nil,[]TXInput{txin2},[]TXOutput{*txout2}}

	var Transations []*Transation

	Transations = append(Transations,&tx,&tx2)

	block.createMerkelTreeRoot(Transations)

	fmt.Printf("%x\n",block.Merkleroot)
}

// 增加挖矿逻辑

proofofwork.go:

package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
)

type  ProofOfWork struct{
	block * Block
	tartget * big.Int
}
const targetBits = 16


func NewProofofWork(b*Block) * ProofOfWork{

	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))
	pow := &ProofOfWork{b,target}
	return pow
}

func (pow * ProofOfWork) prepareData(nonce int32) []byte{

	data := bytes.Join(
		[][]byte{
			IntToHex(pow.block.Version),
			pow.block.PrevBlockHash,
			pow.block.Merkleroot,
			IntToHex(pow.block.Time),
			IntToHex(pow.block.Bits),
			IntToHex(nonce)},
		[]byte{},
	)
	return data
}

func (pow * ProofOfWork) Run() (int32,[]byte){

		var nonce int32
		var secondhash [32]byte
		nonce = 0
		var currenthash big.Int

		for  nonce < maxnonce{

			//序列化
			data:= pow.prepareData(nonce)
			//double hash
			fitstHash := sha256.Sum256(data)
			secondhash = sha256.Sum256(fitstHash[:])
		//	fmt.Printf("%x\n",secondhash)

			currenthash.SetBytes(secondhash[:])
			//比较
			if currenthash.Cmp(pow.tartget) == -1{
				break
			}else{
				nonce++
			}
		}

		return nonce,secondhash[:]
}

func (pow * ProofOfWork) Validate() bool{
	var hashInt big.Int

	data:=pow.prepareData(pow.block.Nonce)

	fitstHash := sha256.Sum256(data)
	secondhash := sha256.Sum256(fitstHash[:])
	hashInt.SetBytes(secondhash[:])
	isValid:= hashInt.Cmp(pow.tartget) == -1

	return isValid
}

// 测试：

func TestPow(){
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

	pow:=NewProofofWork(block)

	nonce,_:= pow.Run()

	block.Nonce = nonce

	fmt.Println("POW:",pow.Validate())
}
