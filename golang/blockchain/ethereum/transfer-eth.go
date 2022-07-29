package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("https://rinkeby.infura.io")
    if err != nil { log.Fatal(err) }

    // 私钥
    privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221...49cad6b3fe8d5817ac83d38b6a19")
    if err != nil { log.Fatal(err) }

    // 公钥
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    // 钱包地址
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil { log.Fatal(err) }

    // 金额
    value := big.NewInt(1000000000000000000) // in wei (1 eth)
    gasLimit := uint64(21000)                // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil { log.Fatal(err) }

    // 收款地址
    toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

    // 走哪个链
    chainID, err := client.NetworkID(context.Background())
    if err != nil { log.Fatal(err) }

    // 签名TX
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil { log.Fatal(err) }

    // 发送转账
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil { log.Fatal(err)  }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}