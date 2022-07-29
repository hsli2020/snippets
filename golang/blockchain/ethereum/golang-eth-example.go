// Ethereum Development in Go using Go-Ethereum
// https://blog.logrocket.com/ethereum-development-using-go-ethereum/

// go get github.com/ethereum/go-ethereum

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
    "log"
)

var (
    ctx         = context.Background()
    url         = "Your Infura URL here"
    client, err = ethclient.DialContext(ctx, url)
)

func currentBlock() {
    block, err := client.BlockByNumber(ctx, nil)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(block.Number())
}

// Querying Ethereum wallet balances with Geth
address := common.HexToAddress("0x8335659d19e46e720e7894294630436501407c3e")
balance, err := client.BalanceAt(ctx, address, nil)
if err != nil {
	log.Print("There was an error", err)
}
fmt.Println("The balance at the current block number is", balance)

// Creating an Ethereum wallet with Go-Ethereum

func createWallet() (string, string) {
	getPrivateKey, err := crypto.GenerateKey()
    if err != nil {
        log.Println(err)
    }

	getPublicKey := crypto.FromECDSA(getPrivateKey)
    thePublicKey := hexutil.Encode(getPublicKey)

	thePublicAddress := crypto.PubkeyToAddress(getPrivateKey.PublicKey).Hex()
	return thePublicAddress, thePublicKey
}

// Making Ethereum transactions in Go using Go-Ethereum

	RecipientAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")

    privateKey, err := crypto.HexToECDSA("The Hexadecimal Private Key ")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Public Key Error")
    }

    SenderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, SenderAddress)
    if err != nil {
        log.Println(err)
    }

    amount := big.NewInt("amount In Wei")
    gasLimit := 3600
    gas, err := client.SuggestGasPrice(ctx)
    if err != nil {
        log.Println(err)
    }

    ChainID, err := client.NetworkID(ctx)
    if err != nil {
        log.Println(err)
    }

	transaction := types.NewTransaction(nonce, RecipientAddress, amount, uint64(gasLimit), gas, nil)
    signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(ChainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }
    err = client.SendTransaction(ctx, signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("transaction sent: %s", signedTx.Hash().Hex())

// Querying the number of transactions in a block

func getTransactionsPerBlock() {
    block, err := client.BlockByNumber(ctx, nil)
    figure, err := client.TransactionCount(ctx, block.Hash())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(figure)
}

// Querying details of transactions in a block

func QueryTransactions() {
    block, _ := client.BlockByNumber(ctx, nil)
    for _, transaction := range block.Transactions() {
        fmt.Println(transaction.Value().String())
        fmt.Println(transaction.Gas())
        fmt.Println(transaction.GasPrice().Uint64())
        fmt.Println(transaction.Nonce())
        fmt.Println(transaction.To().Hex())
    }
}

