// 本文链接： https://dreamerjonson.com/2018/12/11/golang-42-blockchain-merkletree2/
// golang[42]-区块链-go实战比特币默克尔树 2018-12-11
//
// go实战比特币默克尔树

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

//默克尔树节点
type MerkleTree struct {
	RootNode *MerkleNode
}

//默克尔根节点
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

//生成默克尔树中的节点，如果是叶子节点，则Left，right为nil ，如果为非叶子节点，根据Left，right生成当前节点的hash
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mnode := MerkleNode{}

	if left == nil && right == nil {
		mnode.Data = data
	} else {
		prevhashes := append(left.Data, right.Data...)
		firsthash := sha256.Sum256(prevhashes)
		hash := sha256.Sum256(firsthash[:])
		mnode.Data = hash[:]
	}

	mnode.Left = left
	mnode.Right = right

	return &mnode
}

//构建默克尔树
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode
	//构建叶子节点。
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}
	//j代表的是某一层的第一个元素
	j := 0
	//第一层循环代表 nSize代表某一层的个数，每循环一次减半
	for nSize := len(data); nSize > 1; nSize = (nSize + 1) / 2 {
		//第二条循环i+=2代表两两拼接。 i2是为了当个数是基数的时候，拷贝最后的元素。
		for i := 0; i < nSize; i += 2 {
			i2 := min(i+1, nSize-1)

			node := NewMerkleNode(&nodes[j+i], &nodes[j+i2], nil)
			nodes = append(nodes, *node)
		}
		//j代表的是某一层的第一个元素
		j += nSize
	}

	mTree := MerkleTree{&(nodes[len(nodes)-1])}
	return &mTree
}

func ReverseBytes3(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func main() {
	//测试网站下的5个hash是否能够生成merkleRoot
	//https://www.blockchain.com/btc/block/00000000000090ff2791fe41d80509af6ffbd6c5b10294e29cdf1b603acab92c
	//传递hash
	data1, _ := hex.DecodeString("6b6a4236fb06fead0f1bd7fc4f4de123796eb51675fb55dc18c33fe12e33169d")
	data2, _ := hex.DecodeString("2af6b6f6bc6e613049637e32b1809dd767c72f912fef2b978992c6408483d77e")
	data3, _ := hex.DecodeString("6d76d15213c11fcbf4cc7e880f34c35dae43f8081ef30c6901f513ce41374583")
	data4, _ := hex.DecodeString("08c3b50053b010542dca85594af182f8fcf2f0d2bfe8a806e9494e4792222ad2")
	data5, _ := hex.DecodeString("612d035670b7b9dad50f987dfa000a5324ecb3e08745cfefa10a4cefc5544553")

	//大小段转换
	ReverseBytes3(data1)
	ReverseBytes3(data2)
	ReverseBytes3(data3)
	ReverseBytes3(data4)
	ReverseBytes3(data5)

	hehe := [][]byte{
		data1,
		data2,
		data3,
		data4,
		data5,
	}
	//生成默克尔树
	merleroot := NewMerkleTree(hehe)
	//反转
	ReverseBytes3(merleroot.RootNode.Data)
	fmt.Printf("%x", merleroot.RootNode.Data)

}

// 参考资料
// eth wiki:patricia-tree
//
// [csdn 默克尔树解释]https://blog.csdn.net/wo541075754/article/details/54632929
//
// https://github.com/ZtesoftCS/go-ethereum-code-analysis/blob/master/trie源码分析.md
