package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/crypto/sha3"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("https://rinkeby.infura.io")
    if err != nil { log.Fatal(err) }

    privateKey, err := crypto.HexToECDSA("fad9c8855b740a...cad6b3fe8d5817ac83d38b6a19")
    if err != nil { log.Fatal(err) }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil { log.Fatal(err) }

    value := big.NewInt(0) // in wei (0 eth)
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil { log.Fatal(err) }

    toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
    tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

    transferFnSignature := []byte("transfer(address,uint256)")
    hash := sha3.NewKeccak256()
    hash.Write(transferFnSignature)
    methodID := hash.Sum(nil)[:4]
    fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

    paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
    fmt.Println(hexutil.Encode(paddedAddress)) // 0x00000000000000000...001e72c3e4fa1806a51ac79d

    amount := new(big.Int)
    amount.SetString("1000000000000000000000", 10) // 1000 tokens
    paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
    fmt.Println(hexutil.Encode(paddedAmount)) // 0x000000000000000000...0000000635c9adc5dea00000

    var data []byte
    data = append(data, methodID...)
    data = append(data, paddedAddress...)
    data = append(data, paddedAmount...)

    gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
        To:   &toAddress,
        Data: data,
    })
    if err != nil { log.Fatal(err) }
    fmt.Println(gasLimit) // 23256

    tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil { log.Fatal(err)  }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil { log.Fatal(err) }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil { log.Fatal(err) }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}