package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
	fmt.Println(client.BlockNumber(context.Background()))
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x6125B1E88eCae3d238Ceb8aC7749173D9dd6C111")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}

package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 1. Generate private key, used to sign transaction securely
	privateKey, err := crypto.GenerateKey() // 64 hex characters,
	if err != nil {
		log.Fatal(err)
	}

	// 1.1 Private key in readable format
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) 
	//strip off the 0x, e.g. b8caf1fcb42cad076ba79ab5c27d58f030a18a3b85423aab28513b14f9edfa01

	// 2. Generate public key from private key, used to verify transaction
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // convert to hex
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)           // convert to bytes
	fmt.Println("public key: ", hexutil.Encode(publicKeyBytes)[4:]) 
	//strip off the 0x and the first 2 characters 04

	// 3. Address
	// Keccak-256 hash of the public key, and then we take the last 40 characters (20 bytes) 
	// and prefix it with 0x
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)
}


package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//createKs()
	importKs()
}

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {
	// keystore file
	file := "./tmp/UTC--2022-09-08T06-57-12.754954000Z--bdfc113335aca302b5d463c55277797ed6e8670a"

	// generate a new keystore object
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	// read the keystore file
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("readfile error")
		log.Fatal(err)
	}

	password := "secret"
	// import the account AND you will need to set a "new" password which in this case is 
	// the same old password
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		fmt.Println("account error")
		log.Fatal(err)
	}

	// get the Key object
	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal(err)
	}

	// get private key from Key object
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:])

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		fmt.Println("remove error")
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 15495248

	blockNumber := big.NewInt(15495248)

	// full block
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 15495248
	fmt.Println(block.Time())                // 1662621329
	fmt.Println(block.Difficulty().Uint64()) // 12774904243062970
	fmt.Println(block.Hash().Hex())          
		// 0xacd7dd26f51fb92079d1486195b8d10dd3ea83b047f96922877fdb6226828260
	fmt.Println(len(block.Transactions()))   // 217

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 217
}
