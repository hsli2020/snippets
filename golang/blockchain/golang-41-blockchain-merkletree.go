// 本文链接： https://dreamerjonson.com/2018/12/10/golang-41-blockchain-merkletree/
// golang[41]-区块链-默克尔树构建 2018-12-10
// 大自然的启示

// Banyan树的启示：印度banyan树，最大的一颗可以长到1万平方米以上。其如此巨大的秘密就在于其枝干也会产生根，起到支撑，从而作为附属树干，继续生成分支。大自然给人太多启示……
// Merkle Tree

// Merkle Tree，通常也被称作Hash Tree，顾名思义，就是存储hash值的一棵树。Merkle树是一种数据结构，Merkle树的叶子是数据块(例如，文件或者文件的集合)的hash值。非叶节点是其对应子节点串联字符串的hash。
// Merkle树是使区块链发挥作用的基本组成部分。虽然理论上可以在没有Merkle树的情况下制作区块链，但只需创建直接包含每个事务的巨型块头，这样做会带来巨大的可扩展性挑战，可以说无可置疑地使用区块链的能力超出了所有范围，从长远来看，功能强大的电脑。感谢Merkle树，可以构建在所有计算机和大小笔记本电脑上运行的以太网节点，智能手机，甚至是物联网设备
// 比特币中默克尔树的构建过程：

// 对于网站中的交易：
// https://www.blockchain.com/btc/block/000000000001741120135274584b2a0da45b39c8cc78322a14f9004ae766a8e0
/*
第一笔hash：
16f0eb42cb4d9c2374b2cb1de4008162c06fdd8f1c18357f0c849eb423672f5f
大小端转换为：
5f2f6723b49e840c7f35181c8fdd6fc0628100e41dcbb274239c4dcb42ebf016

第二笔hash：
cce2f95fc282b3f2bc956f61d6924f73d658a1fdbc71027dd40b06c15822e061
大小端转换为：
61e02258c1060bd47d0271bcfda158d6734f92d6616f95bcf2b382c25ff9e2cc

将两个拼接在一起：
5f2f6723b49e840c7f35181c8fdd6fc0628100e41dcbb274239c4dcb42ebf01661e02258c1060bd47d0271bcfda158d6734f92d6616f95bcf2b382c25ff9e2cc

将上面拼接的字符串进行两次hash如下：

第一次hash结果：
9b2ec096d49fee8b310752082d63d8fe198386ae2172d90533d9186bb28df63d

将上面计算出的hash值再次进行hash：
525894ddd0891b36c5ff8658e2a978d615b35ce6dedb5cb83f2420dbcd40a0c7

大小端转换即为结果：
c7a040cddb20243fb85cdbdee65cb315d678a9e25886ffc5361b89d0dd945852

go语言实现上面的验证过程
*/
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ReverseBytes2(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func main() {

	//字符串hash转换为字节
	hash1, _ := hex.DecodeString("16f0eb42cb4d9c2374b2cb1de4008162c06fdd8f1c18357f0c849eb423672f5f")

	hash2, _ := hex.DecodeString("cce2f95fc282b3f2bc956f61d6924f73d658a1fdbc71027dd40b06c15822e061")

	//大小端的转换
	ReverseBytes2(hash1)

	ReverseBytes2(hash2)

	//拼接在一起
	rawdata := append(hash1, hash2...)
	//double hash256
	firsthash := sha256.Sum256(rawdata)
	secondhash := sha256.Sum256(firsthash[:])
	merkroot := secondhash[:]

	//反转，与浏览器当中的数据对比
	ReverseBytes2(merkroot)

	fmt.Printf("%x", merkroot)
}

// 参考资料
// eth wiki:patricia-tree
//
// [csdn 默克尔树解释]https://blog.csdn.net/wo541075754/article/details/54632929
//
// https://github.com/ZtesoftCS/go-ethereum-code-analysis/blob/master/trie源码分析.md
