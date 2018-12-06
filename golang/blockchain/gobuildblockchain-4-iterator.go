// 本文链接： https://dreamerjonson.com/2018/12/16/gobuildblockchain-4-iterator/ 

// go实现区块链[4]-遍历区块链与数据库持久化(下) 2018-12-16
// 遍历区块链

// blockchain.go完整代码

package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

const dbFile = "blockchain.db"
const blockBucket = "blocks"

type Blockchain struct{
	tip []byte //最近的一个区块的hash值
	db * bolt.DB
}


type BlockChainIterateor struct{
	currenthash []byte
	db * bolt.DB
}
func (bc * Blockchain) AddBlock(){
	var lasthash []byte

	err := bc.db.View(func(tx * bolt.Tx) error{
		b:= tx.Bucket([]byte(blockBucket))
		lasthash = b.Get([]byte("l"))
		return nil
	})
	if err!=nil{
		log.Panic(err)
	}
	newBlock := NewBlock(lasthash)


	bc.db.Update(func(tx *bolt.Tx) error {
			b:=tx.Bucket([]byte(blockBucket))
			err:= b.Put(newBlock.Hash,newBlock.Serialize())
		if err!=nil{
			log.Panic(err)
		}
			err = b.Put([]byte("l"),newBlock.Hash)

		if err!=nil{
			log.Panic(err)
		}
			bc.tip = newBlock.Hash
			return nil
	})
}

func NewBlockchain() * Blockchain{
	var tip []byte
	db,err := bolt.Open(dbFile,0600,nil)
	if err!=nil{
		log.Panic(err)
	}

	err = db.Update(func(tx * bolt.Tx) error{

		b:= tx.Bucket([]byte(blockBucket))

		if b==nil{

			fmt.Println("区块链不存在，创建一个新的区块链")

			genesis := NewGensisBlock()
			 b,err:=tx.CreateBucket([]byte(blockBucket))
			if err!=nil{
				log.Panic(err)
			}

			err = b.Put(genesis.Hash,genesis.Serialize())
			if err!=nil{
				log.Panic(err)
			}
			err =  b.Put([]byte("l"),genesis.Hash)
			tip = genesis.Hash

		}else{
			tip  =  b.Get([]byte("l"))
		}

		return nil
	})

	if err!=nil{
		log.Panic(err)
	}

	bc:=Blockchain{tip,db}
	return &bc
}

func (bc * Blockchain) iterator() * BlockChainIterateor{

	bci := &BlockChainIterateor{bc.tip,bc.db}

	return bci
}

func (i * BlockChainIterateor) Next() * Block{

	var block *Block

	err:= i.db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		deblock := b.Get(i.currenthash)
		block = DeserializeBlock(deblock)
		return nil
	})

	if err!=nil{
		log.Panic(err)
	}

	i.currenthash = block.PrevBlockHash
	return block
}

func (bc * Blockchain) printBlockchain(){
	bci:=bc.iterator()

	for{
		block:= bci.Next()
		block.String()
		fmt.Println()

		//fmt.Printf("长度：%d\n",len(block.PrevBlockHash))
		if len(block.PrevBlockHash)==0{
			break
		}

	}

}

// 测试

func TestBoltDB(){
	blockchain := NewBlockchain()
	blockchain.AddBlock()
	blockchain.AddBlock()
	blockchain.printBlockchain()
}

func main(){
	TestBoltDB()
}

// 第一次执行执行：

    go build .
    ./buildingBlockChain

区块链不存在，创建一个新的区块链
version:2
Prev.BlockHash:0000349e762f37b4f79f23c5270066cb2963610f5a6c999a846b781cec3152bc
Prev.merkleroot:
Prev.Hash:0000deb768a8e6c520081051d28756578c4c666bde404ff282d7a8e41a1e0107
Time:1544966755
Bits:404454260
nonce:13075

version:2
Prev.BlockHash:0000fdcb6bd475c8275ab47ac6d8d97ab2644ae33d574a914d36f9c1024099eb
Prev.merkleroot:
Prev.Hash:0000349e762f37b4f79f23c5270066cb2963610f5a6c999a846b781cec3152bc
Time:1544966755
Bits:404454260
nonce:142155

version:2
Prev.BlockHash:
Prev.merkleroot:
Prev.Hash:0000fdcb6bd475c8275ab47ac6d8d97ab2644ae33d574a914d36f9c1024099eb
Time:1544966755
Bits:404454260
nonce:105247

再次执行./buildingBlockChain

version:2
Prev.BlockHash:0000c336d1f0284faac173c1d68ca196b3f2e94684d12f201b2610aca39acc7b
Prev.merkleroot:
Prev.Hash:0000a00ab59ad06d7c5d29e9769171676705c270d9edb1e5bd4b39da41e0d40c
Time:1544966858
Bits:404454260
nonce:67063

version:2
Prev.BlockHash:0000deb768a8e6c520081051d28756578c4c666bde404ff282d7a8e41a1e0107
Prev.merkleroot:
Prev.Hash:0000c336d1f0284faac173c1d68ca196b3f2e94684d12f201b2610aca39acc7b
Time:1544966858
Bits:404454260
nonce:69856

version:2
Prev.BlockHash:0000349e762f37b4f79f23c5270066cb2963610f5a6c999a846b781cec3152bc
Prev.merkleroot:
Prev.Hash:0000deb768a8e6c520081051d28756578c4c666bde404ff282d7a8e41a1e0107
Time:1544966755
Bits:404454260
nonce:13075

version:2
Prev.BlockHash:0000fdcb6bd475c8275ab47ac6d8d97ab2644ae33d574a914d36f9c1024099eb
Prev.merkleroot:
Prev.Hash:0000349e762f37b4f79f23c5270066cb2963610f5a6c999a846b781cec3152bc
Time:1544966755
Bits:404454260
nonce:142155

version:2
Prev.BlockHash:
Prev.merkleroot:
Prev.Hash:0000fdcb6bd475c8275ab47ac6d8d97ab2644ae33d574a914d36f9c1024099eb
Time:1544966755
Bits:404454260
nonce:105247

