// main.go
package main

import (
    "fmt"
    "strconv"

    "github.com/nheingit/learnGo/blockchain"
)

func main() {
    chain := blockchain.InitBlockChain()

    chain.AddBlock("first block after genesis")
    chain.AddBlock("second block after genesis")
    chain.AddBlock("third block after genesis")

    for _, block := range chain.Blocks {
        fmt.Printf("Previous hash: %x\n", block.PrevHash)
        fmt.Printf("data: %s\n", block.Data)
        fmt.Printf("hash: %x\n", block.Hash)

        pow := blockchain.NewProofOfWork(block)
        fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()
    }
}

// blockchain/block.go
package blockchain

type BlockChain struct {
    Blocks []*Block
}

type Block struct {
    Hash     []byte
    Data     []byte
    PrevHash []byte
    Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
    block := &Block{[]byte{}, []byte(data), prevHash, 0} 
        // Don't forget to add the 0 at the end for the nonce!
    pow := NewProofOfWork(block)
    nonce, hash := pow.Run()

    block.Hash = hash[:]
    block.Nonce = nonce

    return block
}

func (chain *BlockChain) AddBlock(data string) {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    new := CreateBlock(data, prevBlock.Hash)
    chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// blockchain/proof.go
package blockchain

import (
    "bytes"
    "crypto/sha256"
    "encoding/binary"
    "fmt"
    "log"
    "math"
    "math/big"
)

const Difficulty = 12

type ProofOfWork struct {
    Block *Block
    Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
    target := big.NewInt(1)
    target.Lsh(target, uint(256-Difficulty))

    pow := &ProofOfWork{b, target}
    return pow
}

func ToHex(num int64) []byte {
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, num)
    if err != nil {
        log.Panic(err)
    }
    return buff.Bytes()
}

func (pow *ProofOfWork) InitNonce(nonce int) []byte {
    data := bytes.Join(
        [][]byte{
            pow.Block.PrevHash,
            pow.Block.Data,
            ToHex(int64(nonce)),
            ToHex(int64(Difficulty)),
        },
        []byte{},
    )
    return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
    var intHash big.Int
    var hash [32]byte

    nonce := 0
        // This is essentially an infinite loop due to how large
        // MaxInt64 is.
    for nonce < math.MaxInt64 {
        data := pow.InitNonce(nonce)
        hash = sha256.Sum256(data)

        fmt.Printf("\r%x", hash)
        intHash.SetBytes(hash[:])

        if intHash.Cmp(pow.Target) == -1 {
            break
        } else {
            nonce++
        }
    }
    fmt.Println()

    return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
    var intHash big.Int

    data := pow.InitNonce(pow.Block.Nonce)

    hash := sha256.Sum256(data)
    intHash.SetBytes(hash[:])

    return intHash.Cmp(pow.Target) == -1
}
