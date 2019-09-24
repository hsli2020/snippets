package blockchain  // https://github.com/EwanValentine/bloq

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Contract is a smart contract
type Contract struct {
	// Contract - example: "users:>:5&money:>:1000"
	Contract  string
	Fulfilled bool
	Callback  func(Block) error
}

// Block is an individual block
type Block struct {
	Index     int
	Timestamp string
	Hash      string
	PrevHash  string

	// Data to keep record of
	Data []byte
}

// Blockchain type
type Blockchain struct {
	Blocks    []Block
	Contracts []Contract
	mu        sync.Mutex
}

var bcServer chan []Block

// New blockchain
func New(genesis Block) *Blockchain {
	var blocks []Block
	blocks = append(blocks, genesis)
	return &Blockchain{
		Blocks: blocks,
		mu:     sync.Mutex{},
	}
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.Data) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// IsBlockValid validates a single block
func (bc *Blockchain) IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// GenerateBlock generates a new block
func (bc *Blockchain) GenerateBlock(oldBlock Block, data []byte) (Block, error) {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

// ReplaceChain replaces the blockchain with the new blocks
// if the new blocks are longer than the current, meaning that
// the new blocks are more up-to-date.
func (bc *Blockchain) ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(bc.Blocks) {
		bc.mu.Lock()
		bc.Blocks = newBlocks
		bc.mu.Unlock()
	}
}

// Append adds a new block to the blockchain
func (bc *Blockchain) Append(data []byte) (Block, error) {
	newBlock, err := bc.GenerateBlock(bc.Blocks[len(bc.Blocks)-1], data)
	if err != nil {
		return newBlock, err
	}

	if bc.IsBlockValid(newBlock, bc.Blocks[len(bc.Blocks)-1]) {
		newBlockchain := append(bc.Blocks, newBlock)
		bc.ReplaceChain(newBlockchain)

		// Does this async?
		// bc.checkContracts(newBlock)
		spew.Dump(bc.Blocks)
	}

	return newBlock, nil
}

// GetBlocks gets all blocks in the blockchain
func (bc *Blockchain) GetBlocks() []Block {
	return bc.Blocks
}

// AddContract adds a new smart contract to the block chain
func (bc *Blockchain) AddContract(contract string, callback func(Block) error) error {
	if err := validateContract(contract); err != nil {
		return err
	}
	bc.mu.Lock()
	bc.Contracts = append(bc.Contracts, Contract{
		Contract:  contract,
		Fulfilled: false,
		Callback:  callback,
	})
	bc.mu.Unlock()
	return nil
}

func validateContract(contract string) error {
	parts := strings.Split(contract, ":")
	if len(parts) != 3 {
		return errors.New("Invalid contract")
	}

	operator := parts[1]
	switch operator {
	case ">":
	case "<":
	case "=":
	case "!=":
	default:
		return errors.New("Invalid operator")
	}

	return nil
}

// checkContracts iterates through each unfulfilled contract
// against the current data to check if any of them have been
// fulfilled, if so, call the handler.
func (bc *Blockchain) checkContracts(block Block) error {
	bc.mu.Lock()
	var data map[string]interface{}
	if err := json.Unmarshal(block.Data, &data); err != nil {
		return err
	}
	// For each contract
	for key, val := range bc.Contracts {
		var rulesCount int
		rulesValidated := 0
		rules := parseContract(val.Contract)
		rulesCount = len(rules)
		if val.Fulfilled == true {
			continue
		}

		for _, rule := range rules {
			parts := parsePart(rule)
			field := parts[0]
			operator := parts[1]
			value := parts[2]
			log.Println("val:", data[field])
			switch operator {
			case ">":
				i, _ := strconv.Atoi(value)
				if int(data[field].(float64)) > i {
					rulesValidated++
				}
			case "<":
				i, _ := strconv.Atoi(value)
				if int(data[field].(float64)) < i {
					rulesValidated++
				}
			case "=":
				if data[field] == value {
					rulesValidated++
				}
			}
		}

		// If the amount of rules that are valid
		// if equal to all of the rules set, then
		// this smart contract has now been fulfilled
		if rulesValidated == rulesCount {
			val.Callback(block)
			val.Fulfilled = true
			bc.Contracts[key] = val
		}
	}
	bc.mu.Unlock()
	return nil
}

func parseContract(contract string) []string {
	return strings.Split(contract, "&")
}

func parsePart(part string) []string {
	return strings.Split(part, ":")
}
